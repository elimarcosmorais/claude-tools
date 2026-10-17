# CLAUDE.md

## Visão do projeto
CLI de Windows (`claude-tools.exe`) que guarda logins do Claude Code com um nome e troca entre
eles copiando arquivos em `~/.claude`; sem rede, banco nem serviço. Uso pessoal, instalado em
`C:\dev\programs\claude-tools`.

## Stack
- Go 1.27 (`go.mod`), só a biblioteca padrão: um executável único ([ADR 001](docs/adr/001-go-binario-unico.md)).
- Build, testes e instalação pelo [Task](https://taskfile.dev) (`Taskfile.yml`); versão no arquivo `VERSION`.
- NÃO usar: módulo de terceiros (o `go.mod` não tem `require`); mudar isso exige ADR nova.

## Comandos
- Build: `task build` (gera `bin/claude-tools.exe`)
- Testes: `task test` (`go test ./...`)
- Lint: `task lint` (`go vet ./...` e falha se `gofmt -l .` listar algum arquivo)
- Instalar: `task install` (build + `scripts/install.ps1`: cópia versionada, junction `current` e `PATH` do usuário)

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
- Teste → `<arquivo>_test.go` ao lado do código; o do instalador em `scripts/install_test.go`.

## Arquitetura
- Camadas `cmd/claude-tools → internal/cli → internal/account → arquivos na home`: ver
  [docs/architecture.md](docs/architecture.md).
- Como cada comando funciona: [docs/specs/README.md](docs/specs/README.md). Decisões:
  [docs/adr/README.md](docs/adr/README.md).

## Pronto significa
- `task test` e `task lint` verdes, e `task build` compilando.
