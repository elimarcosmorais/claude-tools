# claude-tools

> CLI para Windows e Linux (WSL2) que guarda vários logins do Claude Code com um nome e alterna entre eles com um comando.

## Sobre

O Claude Code guarda um login só por vez. Quem usa mais de uma conta (pessoal e trabalho, por
exemplo) precisa fazer `/login` de novo a cada troca. O `claude-tools` resolve isso: você salva o
login ativo com um nome e depois ativa qualquer conta salva com `claude-tools account use <conta>`.

A troca é só cópia de arquivos em `~/.claude` (`.credentials.json` e o campo `oauthAccount` de
`~/.claude.json`): sem rede, sem banco, sem serviço rodando. Antes de cada troca ele guarda um
backup do login ativo e devolve à conta de origem o token que o Claude Code renovou. As contas
ficam em `~/.claude/accounts/`.

- Como cada comando funciona, passo a passo: [docs/specs/account.md](docs/specs/account.md)
- Camadas, build e instalação: [docs/architecture.md](docs/architecture.md)
- Decisões de projeto: [docs/adr/README.md](docs/adr/README.md)

## Pré-requisitos

- Windows, ou Linux no WSL2
- [Go](https://go.dev) 1.27+ (no WSL, o Go do Linux)
- No Windows, PowerShell 7+ (`pwsh`), usado pelo instalador; no Linux, só o `sh`

No WSL2, o Claude Code (o `claude` no terminal do Linux ou a extensão do VS Code no modo remoto)
guarda o login na home do Linux, separado do Windows. Para trocar essas contas, instale o
`claude-tools` também dentro do WSL: cada ambiente tem as suas contas salvas.

## Build

```powershell
git clone https://github.com/elimarcosmorais/claude-tools.git
cd claude-tools
go tool task build
```

O [Task](https://taskfile.dev) não precisa estar instalado: ele é uma ferramenta do `go.mod` e roda
por `go tool task` (o Go o compila na primeira vez). O `go tool task build` compila o executável do
sistema em que roda, com a versão do arquivo `VERSION` no nome (o `bin/` fica fora do git):
`bin\claude-tools-<VERSION>.exe` no Windows e `bin/claude-tools-<VERSION>` no Linux. Dá para
usá-lo direto dali, sem instalar:

```powershell
bin\claude-tools-0.1.0.exe help
```

```sh
bin/claude-tools-0.1.0 help
```

Para gerar os executáveis dos dois sistemas de uma vez (os anexos da release), em qualquer um
deles:

```sh
go tool task build:all
# bin/claude-tools-<VERSION>-windows-amd64.exe e bin/claude-tools-<VERSION>-linux-amd64
```

## Instalação no Windows

```powershell
go tool task install DIR=C:\dev\programs\claude-tools
```

A pasta de instalação é sua escolha e vem na variável `DIR` (sem ela, o `go tool task install` para com
a mensagem de uso antes de compilar). Pasta com espaço vai entre aspas:
`go tool task install "DIR=D:\meus programas\claude-tools"`. O `go tool task install DIR=<pasta>` roda o
`build` (não precisa rodá-lo antes) e depois o `scripts/install.ps1`, que:

1. copia o executável para `<pasta>\claude-tools-v<VERSION>\`;
2. aponta a junction `<pasta>\current` para essa versão;
3. põe `<pasta>\current` no `PATH` do usuário, se ainda não estiver lá.

Rodar de novo dá o mesmo resultado. Para instalar sem mexer no `PATH`:

```powershell
go tool task install DIR=D:\tools\claude-tools -- -SkipPath
```

Abra um terminal novo depois da instalação para pegar o `PATH` atualizado.

## Instalação no Linux (WSL2)

No terminal do WSL, dentro do repositório (pode ser o mesmo clone do Windows, por `/mnt/c/...`):

```sh
go tool task install DIR=$HOME/programs/claude-tools
```

O `go tool task install DIR=<pasta>` roda o `build` e depois o `scripts/install.sh`, que:

1. copia o executável para `<pasta>/claude-tools-v<VERSION>/claude-tools`;
2. aponta o link simbólico `<pasta>/current` para essa versão;
3. acrescenta ao `~/.profile`, se ainda não estiver lá, a linha
   `export PATH="<pasta>/current:$PATH"` (com o comentário `# claude-tools` antes).

Rodar de novo dá o mesmo resultado. Para instalar sem mexer no `~/.profile`:

```sh
go tool task install DIR=$HOME/programs/claude-tools -- --skip-path
```

Abra um terminal novo do WSL depois da instalação para pegar o `PATH` atualizado.

## Uso

Faça login no Claude Code com a primeira conta e salve; depois, a segunda:

```console
> claude-tools account save pessoal
Conta 'pessoal' salva (eu@exemplo.com).

# no Claude Code: /login com a outra conta

> claude-tools account save trabalho
Conta 'trabalho' salva (eu@empresa.com).
```

Para trocar:

```console
> claude-tools account use pessoal
Conta ativa: pessoal (eu@exemplo.com)
Reabra o 'claude' no terminal e use 'Developer: Reload Window' no VS Code.

> claude-tools account status
Conta atual: pessoal
Email ativo no Claude Code: eu@exemplo.com
Contas salvas:
  pessoal: eu@exemplo.com (salva em 17/10/2026 09:12) <- ativa
  trabalho: eu@empresa.com (salva em 17/10/2026 09:15)
```

Se `ANTHROPIC_API_KEY` ou `ANTHROPIC_AUTH_TOKEN` estiverem definidas, o Claude Code usa a chave
da variável e ignora o login: `use` e `status` avisam, porque a troca não terá efeito.

## Comandos

| Comando                                      | Descrição                                                   |
| -------------------------------------------- | ----------------------------------------------------------- |
| `claude-tools account save <conta>`          | guarda o login atual do Claude Code como `<conta>`          |
| `claude-tools account use <conta>`           | ativa o login guardado em `<conta>`                         |
| `claude-tools account rename <conta> <nova>` | renomeia `<conta>` para `<nova>`                            |
| `claude-tools account remove <conta>`        | apaga o login guardado em `<conta>` (sem pedir confirmação) |
| `claude-tools account status`                | mostra a conta ativa e as contas salvas                     |
| `claude-tools help`                          | mostra o uso                                                |

Nome de conta: letras, números, `-` ou `_`, até 32 caracteres, gravado em minúsculas; `backup` é
reservado. Códigos de saída: `0` sucesso, `1` erro, `2` uso incorreto.

## Testes

```powershell
go tool task test   # go test ./...
go tool task lint   # go vet ./... e gofmt
```
