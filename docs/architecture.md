# Arquitetura

O `claude-tools` é um CLI em Go com só a biblioteca padrão, compilado num executável único por
sistema ([ADR 001](adr/001-go-binario-unico.md)): `claude-tools.exe` no Windows e `claude-tools`
no Linux, que roda dentro do WSL2 e cuida dos logins da home do Linux
([ADR 003](adr/003-windows-e-linux-do-mesmo-codigo.md)). Ele troca a conta do Claude Code
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

O `Taskfile.yml` é o mesmo no Windows e no Linux (o Task roda os comandos no interpretador `sh`
embutido nele, nos dois sistemas) e tem:

| Tarefa | Faz |
|---|---|
| `build` | `bin/claude-tools-<VERSION>` com `{{exeExt}}`: `.exe` no Windows, sem extensão no Linux |
| `build:all` | os anexos da release, por compilação cruzada (`GOOS`/`GOARCH`, `CGO_ENABLED=0`): `bin/claude-tools-<VERSION>-windows-amd64.exe` e `bin/claude-tools-<VERSION>-linux-amd64`, de qualquer um dos dois sistemas |
| `test` | `go test ./...` |
| `lint` | `go vet ./...` e `gofmt -l .` vazio |
| `install` | `build` e o instalador do sistema (abaixo) |

A versão fica no arquivo `VERSION`, lido pela variável `VERSION` do Taskfile com o `read` e o
`printf` do próprio interpretador, tirando o `\r`: o checkout do Windows (`core.autocrlf`) deixa o
arquivo em CRLF, e o WSL lê esse mesmo checkout por `/mnt/c`. Pela mesma razão o `.gitattributes`
fixa `*.sh` em LF: em CRLF, o `sh` do Linux leria o `\r` como parte dos comandos. O Task é uma
diretiva `tool` do `go.mod` e roda por `go tool task <tarefa>`, sem instalação na máquina
([ADR 002](adr/002-task-como-ferramenta-do-go-mod.md)).

### `install`: o que vale nos dois sistemas

Não há pasta padrão: o repositório não guarda caminho de instalação. A pasta vem na variável
`DIR` porque o Task não deixa a tarefa ter flag própria (`go tool task install --path=C:\x` falha com
`unknown flag: --path`) e lê argumento solto como nome de tarefa (`go tool task install C:\x` falha com
`Task "C:\x" does not exist`). O nome não é `PATH` para não se confundir com a variável de
ambiente. O `Taskfile.yml` passa a pasta por `{{shellQuote .DIR}}`: pasta com espaço ou
terminada em `\` chega inteira ao instalador. O que vier depois do `--` (`{{.CLI_ARGS}}`) vai para
o instalador do sistema, como o `-SkipPath` (Windows) ou o `--skip-path` (Linux). Cada sistema
roda o seu comando pelo `platforms` do Task.

Sem pasta, a `precondition` do `install` para com a mensagem de uso, com um exemplo para cada
sistema, **antes** do build — por isso o build é o comando `task: build` (dentro do Taskfile), e
não `deps`, que rodaria antes da `precondition`.

Nos dois, a instalação é uma pasta por versão (`<pasta>/claude-tools-v<VERSION>/`) e um link
`<pasta>/current` para a versão ativa, que é o que vai no `PATH`. Rodar de novo dá o mesmo
resultado; o link antigo é removido sem tocar na pasta para onde ele apontava, e um `current` que
é uma pasta comum faz o instalador parar com erro em vez de apagá-lo.

### Windows: `scripts/install.ps1`

`go tool task install DIR=<pasta>` compila e roda `scripts/install.ps1 -Root <pasta>`, que:

1. copia `bin/claude-tools-<VERSION>.exe` para `<pasta>\claude-tools-v<VERSION>\claude-tools.exe`;
2. recria a junction `<pasta>\current` apontando para essa pasta;
3. põe `<pasta>\current` no `PATH` do usuário, se ainda não estiver lá.

O `install.ps1` chamado direto sem `-Root` também falha (saída 1), sem `Mandatory` para o `pwsh`
não pedir a pasta no console. O parâmetro `-SkipPath` não mexe no `PATH`; o teste
`scripts/install_windows_test.go` usa `-Root` numa pasta temporária com `-SkipPath`.

O `PATH` é lido e gravado direto no registro (`HKCU\Environment`), sem expandir as variáveis e
mantendo o tipo do valor: `[Environment]::SetEnvironmentVariable` gravaria um `REG_SZ` já
expandido, e as entradas com `%USERPROFILE%` deixariam de acompanhar a variável. Depois da
gravação, o script avisa as janelas com `WM_SETTINGCHANGE`, para os terminais novos já virem com
o `PATH` atualizado.

Quando o pwsh roda com a saída redirecionada, as mensagens do script saem na página de código do
console, não em UTF-8; no terminal elas aparecem acentuadas.

### Linux (WSL2): `scripts/install.sh`

Dentro do WSL, com o Go do Linux, `go tool task install DIR=<pasta>` compila e roda
`sh scripts/install.sh --root <pasta>` (`sh` POSIX: o Ubuntu do WSL não tem `pwsh`), que:

1. confere tudo antes de gravar: o `bin/claude-tools-<VERSION>` (sem ele, pede o `go tool task
   build`), o `current` que não é link simbólico e, sem `--skip-path`, o `HOME` definido e uma
   pasta sem `"`, `$`, `` ` `` ou `\`, que mudariam o caminho dentro das aspas da linha do
   `~/.profile`;
2. copia o binário para `<pasta>/claude-tools-v<VERSION>/claude-tools` com modo `0755`, por um
   temporário e `mv` (gravar por cima de um binário em execução falharia com `Text file busy`);
3. recria o link simbólico `<pasta>/current` apontando para essa pasta (`rm` no link, `ln -s`);
4. acrescenta ao `$HOME/.profile`, se a linha ainda não estiver lá, `# claude-tools` e
   `export PATH="<pasta>/current:$PATH"` (com um `\n` antes, se o arquivo não terminava em um).

O `~/.profile` é lido pelos shells de login, que o terminal do WSL abre e que o servidor do VS
Code no WSL usa para montar o ambiente. O script acha o repositório pela pasta dele mesmo, tira o
`\r` do `VERSION` e sai com 1 e `Erro: <detalhe>` no stderr em qualquer falha (sem `--root`,
argumento desconhecido). `--skip-path` não mexe no `~/.profile`; o teste
`scripts/install_linux_test.go` (`//go:build linux`) roda o script com `HOME` e `--root` em pastas
temporárias.

## Módulos

| Módulo | Spec |
|---|---|
| Contas (armazenamento, nomes, `save`, `use`, `rename`, `remove`, `status`) | [specs/account.md](specs/account.md) |
