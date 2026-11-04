# ADRs

Decisões que valem para o projeto todo. Abra só as ADRs cujo "Abra quando" toca a tarefa; o
status que vale é o do arquivo.

| ADR | Abra quando |
|---|---|
| [001 — Go com biblioteca padrão, binário único](001-go-binario-unico.md) | for criar código, acrescentar uma dependência ou um módulo Go, mexer no build ou no empacotamento do executável |
| [002 — Task como ferramenta do `go.mod`](002-task-como-ferramenta-do-go-mod.md) | for rodar ou mudar o Task, o `Taskfile.yml`, ou acrescentar uma ferramenta de desenvolvimento ao `go.mod` |
| [003 — Windows e Linux (WSL2) a partir do mesmo código](003-windows-e-linux-do-mesmo-codigo.md) | for mexer em código, teste ou instalador que dependa do sistema (Windows, Linux, WSL), em permissões de arquivo ou nos anexos da release |
