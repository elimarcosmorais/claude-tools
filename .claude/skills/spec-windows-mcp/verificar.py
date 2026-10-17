"""Connects to the windows-mcp server declared in .mcp.json and lists its tools.

Usage: python -I verificar.py <path to .mcp.json> [server name]
Exit code 0 and the tool list on success; non-zero and the reason on failure.
"""

import json
import os
import queue
import shutil
import subprocess
import sys
import threading

TIMEOUT_SECONDS = 180  # the first run may download the package


def main() -> int:
    if len(sys.argv) < 2:
        print("usage: verificar.py <.mcp.json> [server]", file=sys.stderr)
        return 2
    config_path = sys.argv[1]
    server_name = sys.argv[2] if len(sys.argv) > 2 else "windows-mcp"

    with open(config_path, encoding="utf-8") as f:
        server = json.load(f)["mcpServers"][server_name]

    executable = shutil.which(server["command"])
    if executable is None:
        print(f"command not found in PATH: {server['command']}", file=sys.stderr)
        return 1

    env = {**os.environ, **server.get("env", {})}
    proc = subprocess.Popen(
        [executable, *server.get("args", [])],
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.DEVNULL,
        env=env,
        text=True,
        encoding="utf-8",
    )

    lines: queue.Queue = queue.Queue()
    threading.Thread(target=lambda: [lines.put(l) for l in proc.stdout], daemon=True).start()

    def send(message: dict) -> None:
        proc.stdin.write(json.dumps(message) + "\n")
        proc.stdin.flush()

    def wait_for(request_id: int) -> dict:
        while True:
            try:
                line = lines.get(timeout=TIMEOUT_SECONDS)
            except queue.Empty:
                raise TimeoutError(f"no response to request {request_id} in {TIMEOUT_SECONDS}s")
            try:
                message = json.loads(line)
            except json.JSONDecodeError:
                continue
            if message.get("id") == request_id:
                return message

    try:
        send({
            "jsonrpc": "2.0", "id": 1, "method": "initialize",
            "params": {
                "protocolVersion": "2024-11-05",
                "capabilities": {},
                "clientInfo": {"name": "verificar", "version": "1"},
            },
        })
        init = wait_for(1)
        if "error" in init:
            print(f"initialize failed: {init['error']}", file=sys.stderr)
            return 1
        send({"jsonrpc": "2.0", "method": "notifications/initialized"})
        send({"jsonrpc": "2.0", "id": 2, "method": "tools/list"})
        listed = wait_for(2)
        if "error" in listed:
            print(f"tools/list failed: {listed['error']}", file=sys.stderr)
            return 1
        server_info = init["result"].get("serverInfo", {})
        names = sorted(tool["name"] for tool in listed["result"]["tools"])
        print(f"server: {server_info.get('name')} {server_info.get('version', '')}".rstrip())
        print(f"tools ({len(names)}): {', '.join(names)}")
        return 0
    except TimeoutError as err:
        print(str(err), file=sys.stderr)
        return 1
    finally:
        proc.kill()


if __name__ == "__main__":
    sys.exit(main())
