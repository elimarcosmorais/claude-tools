# CLAUDE.md

## Visão do projeto
CLI de Windows e Linux/WSL2 (`claude-tools.exe` e `claude-tools`, do mesmo código:
[ADR 003](docs/adr/003-windows-e-linux-do-mesmo-codigo.md)) que guarda logins do Claude Code com
um nome e troca entre eles copiando arquivos em `~/.claude`; sem rede, banco nem serviço. Uso
pessoal, instalado na pasta passada ao `go tool task install` (sem pasta padrão).

## Stack
- Go 1.27 (`go.mod`), só a biblioteca padrão no código: um executável único ([ADR 001](docs/adr/001-go-binario-unico.md)).
- Build, testes e instalação pelo [Task](https://taskfile.dev) (`Taskfile.yml`), que é diretiva `tool` do `go.mod` e roda por `go tool task`, sem instalar nada na máquina ([ADR 002](docs/adr/002-task-como-ferramenta-do-go-mod.md)); versão no arquivo `VERSION`.
- NÃO usar: módulo de terceiros no código do `claude-tools` (o `go.mod` só tem as dependências `// indirect` da diretiva `tool` do Task); mudar isso exige ADR nova.

## Comandos
- Build: `go tool task build` (gera `bin/claude-tools-<VERSION>.exe` no Windows e `bin/claude-tools-<VERSION>` no Linux, com a versão do arquivo `VERSION`)
- Release: versão em `VERSION` · anexos: `go tool task build:all` → `bin/claude-tools-<versão>-windows-amd64.exe` e `bin/claude-tools-<versão>-linux-amd64`
- Testes: `go tool task test` (`go test ./...`)
- Lint: `go tool task lint` (`go vet ./...` e falha se `gofmt -l .` listar algum arquivo)
- Instalar: `go tool task install DIR=<pasta>` (build + instalador do sistema: no Windows, `scripts/install.ps1 -Root <pasta>` — cópia versionada, junction `current` e `PATH` do usuário; no Linux, `scripts/install.sh --root <pasta>` — cópia versionada, link simbólico `current` e linha `export PATH` no `~/.profile`)

## Convenções
- A home do usuário só é lida em `cmd/claude-tools/main.go`; `cli.Run` e `account.Store` a
  recebem por parâmetro. NÃO chame `os.UserHomeDir()` em outro lugar: é o que mantém os testes
  longe do `~/.claude` real.
- Toda escrita em `~/.claude.json`, `~/.claude/.credentials.json` ou `~/.claude/accounts/` passa
  por `writeFileAtomic` (`internal/account/store.go`), nunca `os.WriteFile` direto.
- Testes que gravam arquivo usam o helper `tempDir(t)`, não `t.TempDir()`: o antivírus do Windows
  segura os arquivos na limpeza e o `t.TempDir()` reprova o teste.
- `internal/account` devolve erros tipados e não escreve no terminal; mensagens, stdout/stderr e
  códigos de saída (0 ok · 1 erro · 2 uso) são do `internal/cli`.
- Mensagens ao usuário e comentários em português.

## Onde coisas novas vão
- Comando novo do CLI → `internal/cli/cli.go` (parse e mensagens) + operação no pacote de domínio.
- Operação sobre contas → método de `account.Store` em `internal/account/store.go`.
- Teste → `<arquivo>_test.go` ao lado do código; os dos instaladores em `scripts/install_windows_test.go` e `scripts/install_linux_test.go`.

## Arquitetura
- Camadas `cmd/claude-tools → internal/cli → internal/account → arquivos na home`: ver
  [docs/architecture.md](docs/architecture.md).
- Como cada comando funciona: [docs/specs/README.md](docs/specs/README.md). Decisões:
  [docs/adr/README.md](docs/adr/README.md).

## Pronto significa
- `go tool task test` e `go tool task lint` verdes, e `go tool task build` compilando.
