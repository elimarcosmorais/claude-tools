---
status: aceito
---

# ADR 002 — Task como ferramenta do `go.mod`, chamado por `go tool task`

- **Data:** 2026-11-04
- **Origem:** pedido de usar o Task dentro do projeto, como o `dev-tools` faz
- **Substitui:** nenhuma
- **Substituída por:** nenhuma

## Contexto

Build, testes, lint e instalação rodam pelo Task (`Taskfile.yml`). O comando `task` vinha do
`PATH`, instalado à mão na máquina: quem clonava o repositório precisava instalá-lo antes de
qualquer coisa, e a versão dependia de quem montou o PC. O [ADR 001](001-go-binario-unico.md)
manda o código usar só a biblioteca padrão, e o `go.mod` não tinha `require`.

## Decisão

O Task (`github.com/go-task/task/v3/cmd/task`, v3.54.0) é uma diretiva `tool` do `go.mod`, e todo
comando o chama por `go tool task <tarefa>`. O código do `claude-tools` continua só com a
biblioteca padrão (ADR 001): a diretiva `tool` só serve ao desenvolvimento, e o `go build` do
`claude-tools.exe` não compila nem embute nada dela.

## Consequências

**Ganhos:**

- Da máquina, o desenvolvimento exige só o Go; a versão do Task é a mesma em todo clone.
- Atualizar o Task é um `go get -tool github.com/go-task/task/v3/cmd/task@<versão>`, revisado no diff.

**Custos:**

- O `go.mod` ganha as dependências `// indirect` do Task e o `go.sum` passa a existir.
- Todo comando fica mais longo (`go tool task build`), e o primeiro `go tool task` de uma
  versão compila o Task.
- O `go` do `go.mod` acompanha o maior entre o app e o Task (o Task v3.54.0 pede o 1.26.4; o
  projeto está no 1.27).

## Alternativas consideradas

- **Task instalado no `PATH`** — nenhuma mudança no repo. Rejeitada porque a versão fica sem
  controle e cada clone exige uma instalação antes do primeiro comando.
- **Dispensar o Task e usar scripts PowerShell** — sem dependência nova. Rejeitada porque obriga
  a reescrever `Taskfile.yml`, a precondição do `install` e as mensagens de uso.
