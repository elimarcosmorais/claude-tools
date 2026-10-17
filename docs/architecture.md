# Arquitetura

O `claude-tools` é um CLI de Windows, um executável único (`claude-tools.exe`) em Go com só a
biblioteca padrão ([ADR 001](adr/001-go-binario-unico.md)). Ele troca a conta do Claude Code
copiando os arquivos de login para dentro e para fora de `~/.claude/accounts/`. Não há rede,
banco nem serviço: tudo é leitura e escrita de arquivos na home do usuário.

## Camadas

```
cmd/claude-tools  →  internal/cli  →  internal/account  →  arquivos em ~ (home)
```

| Camada | Pasta | Faz | Não faz |
|---|---|---|---|
| Entrada | `cmd/claude-tools/main.go` | lê `os.UserHomeDir()` (o único lugar do código que o faz) e chama `os.Exit(cli.Run(os.Args[1:], home, os.Stdout, os.Stderr))` | nenhuma regra |
| CLI | `internal/cli` | `Run(args, home, stdout, stderr) int`: o parse dos argumentos, o texto de uso, as mensagens, os avisos de variáveis de ambiente e os códigos de saída | não toca em arquivo |
| Contas | `internal/account` | `Store{Home}`: os caminhos, a validação de nome, a escrita atômica e as operações sobre as contas | não escreve no terminal nem decide código de saída |

A home entra por parâmetro em `cli.Run` e em `account.Store`, e não é lida de novo em nenhum
outro ponto: é o que deixa os testes rodarem com `t.TempDir()` como home, sem tocar no
`~/.claude` real.

## Saída e códigos

- Sucesso e avisos vão para o stdout; erro e uso incorreto, para o stderr.
- Códigos: `0` sucesso · `1` erro de operação (`Erro: <detalhe>`) · `2` uso incorreto (comando
  desconhecido, número errado de argumentos, nome de conta recusado).
- O `internal/account` devolve erros; o `cli` os traduz: `*account.NameError` vira saída 2 com a
  mensagem dele; `*account.NoLoginError`, `*account.MissingAccountError` e
  `*account.ExistingAccountError` saem com a mensagem deles (a frase para o usuário, que o `cli`
  pode completar, como o `use` faz) e saída 1; o
  resto vira `Erro: <detalhe>` e saída 1.

## Build e instalação

O `Taskfile.yml` tem `build` (`bin/claude-tools.exe`), `test` (`go test ./...`), `lint`
(`go vet ./...` e `gofmt -l .` vazio) e `install`. A versão fica no arquivo `VERSION`.

`task install DIR=<pasta>` compila e roda `scripts/install.ps1 -Root <pasta>`, que:

1. copia `bin/claude-tools.exe` para `<pasta>\claude-tools-v<VERSION>\`;
2. recria a junction `<pasta>\current` apontando para essa pasta (remove
   só o link antigo, nunca a pasta para onde ele apontava; se `current` é uma pasta comum, para
   com erro em vez de apagá-la);
3. põe `<pasta>\current` no `PATH` do usuário, se ainda não estiver lá.

Não há pasta padrão: o repositório não guarda caminho de instalação. A pasta vem na variável
`DIR` porque o Task não deixa a tarefa ter flag própria (`task install --path=C:\x` falha com
`unknown flag: --path`) e lê argumento solto como nome de tarefa (`task install C:\x` falha com
`Task "C:\x" does not exist`). O nome não é `PATH` para não se confundir com a variável de
ambiente. O `Taskfile.yml` passa a pasta por `{{shellQuote .DIR}}`: pasta com espaço ou
terminada em `\` chega inteira ao `pwsh`. O que vier depois do `--` (`{{.CLI_ARGS}}`) vai para
o `install.ps1`, como o `-SkipPath`.

Sem pasta, a `precondition` do `install` para com a mensagem de uso **antes** do build — por
isso o build é o comando `task: build`, e não `deps`, que rodaria antes da `precondition`. O
`install.ps1` chamado direto sem `-Root` também falha (saída 1), sem `Mandatory` para o `pwsh`
não pedir a pasta no console.

Rodar de novo dá o mesmo resultado. O parâmetro `-SkipPath` não mexe no `PATH`; o teste
`scripts/install_test.go` usa `-Root` numa pasta temporária com `-SkipPath`.

O `PATH` é lido e gravado direto no registro (`HKCU\Environment`), sem expandir as variáveis e
mantendo o tipo do valor: `[Environment]::SetEnvironmentVariable` gravaria um `REG_SZ` já
expandido, e as entradas com `%USERPROFILE%` deixariam de acompanhar a variável. Depois da
gravação, o script avisa as janelas com `WM_SETTINGCHANGE`, para os terminais novos já virem com
o `PATH` atualizado.

Quando o pwsh roda com a saída redirecionada, as mensagens do script saem na página de código do
console, não em UTF-8; no terminal elas aparecem acentuadas.

## Módulos

| Módulo | Spec |
|---|---|
| Contas (armazenamento, nomes, `save`, `use`, `rename`, `remove`, `status`) | [specs/account.md](specs/account.md) |
