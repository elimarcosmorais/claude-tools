"""Connects to the Playwright MCP server declared in .mcp.json, lists its tools and
proves the browser launches by opening a page and taking a screenshot.

Usage: python -I verificar.py <path to .mcp.json> [server name]
Exit code 0, the tool list and the page result on success; non-zero and the reason on failure.
"""

import json
import os
import queue
import shutil
import subprocess
import sys
import threading

TIMEOUT_SECONDS = 180  # the first run may download the package
TEST_PAGE = "data:text/html,<title>verificar</title><h1>Playwright MCP ok</h1>"


def main() -> int:
    if len(sys.argv) < 2:
        print("usage: verificar.py <.mcp.json> [server]", file=sys.stderr)
        return 2
    config_path = sys.argv[1]
    server_name = sys.argv[2] if len(sys.argv) > 2 else "playwright"

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

    def call_tool(request_id: int, name: str, arguments: dict) -> str:
        send({
            "jsonrpc": "2.0", "id": request_id, "method": "tools/call",
            "params": {"name": name, "arguments": arguments},
        })
        reply = wait_for(request_id)
        if "error" in reply:
            raise RuntimeError(f"{name} failed: {reply['error']}")
        result = reply["result"]
        text = "\n".join(c.get("text", "") for c in result.get("content", []) if c.get("type") == "text")
        if result.get("isError"):
            raise RuntimeError(f"{name} failed: {text}")
        return text

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

        navigated = call_tool(3, "browser_navigate", {"url": TEST_PAGE})
        title = next((l.strip() for l in navigated.splitlines() if "Page Title" in l), "")
        print(f"navigate: {title or 'ok'}")
        shot = call_tool(4, "browser_take_screenshot", {})
        saved = next((l.strip() for l in shot.splitlines() if ".png" in l), "ok")
        print(f"screenshot: {saved}")
        call_tool(5, "browser_close", {})
        return 0
    except (TimeoutError, RuntimeError) as err:
        print(str(err), file=sys.stderr)
        return 1
    finally:
        proc.kill()


if __name__ == "__main__":
    sys.exit(main())
