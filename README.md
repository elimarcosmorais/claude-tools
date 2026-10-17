# claude-tools

> CLI para Windows que guarda vários logins do Claude Code com um nome e alterna entre eles com um comando.

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

- Windows
- [Go](https://go.dev) 1.27+
- [Task](https://taskfile.dev)
- PowerShell 7+ (`pwsh`), usado pelo instalador

## Build

```powershell
git clone https://github.com/elimarcosmorais/claude-tools.git
cd claude-tools
task build
```

O `task build` compila o executável em `bin\claude-tools.exe` (o `bin\` fica fora do git). Dá
para usá-lo direto dali, sem instalar:

```powershell
bin\claude-tools.exe help
```

## Instalação

```powershell
task install DIR=C:\dev\programs\claude-tools
```

A pasta de instalação é sua escolha e vem na variável `DIR` (sem ela, o `task install` para com
a mensagem de uso antes de compilar). Pasta com espaço vai entre aspas:
`task install "DIR=D:\meus programas\claude-tools"`. O `task install DIR=<pasta>` roda o
`task build` (não precisa rodá-lo antes) e depois o `scripts/install.ps1`, que:

1. copia o executável para `<pasta>\claude-tools-v<VERSION>\`;
2. aponta a junction `<pasta>\current` para essa versão;
3. põe `<pasta>\current` no `PATH` do usuário, se ainda não estiver lá.

Rodar de novo dá o mesmo resultado. Para instalar sem mexer no `PATH`:

```powershell
task install DIR=D:\tools\claude-tools -- -SkipPath
```

Abra um terminal novo depois da instalação para pegar o `PATH` atualizado.

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

| Comando | Descrição |
|---|---|
| `claude-tools account save <conta>` | guarda o login atual do Claude Code como `<conta>` |
| `claude-tools account use <conta>` | ativa o login guardado em `<conta>` |
| `claude-tools account rename <conta> <nova>` | renomeia `<conta>` para `<nova>` |
| `claude-tools account remove <conta>` | apaga o login guardado em `<conta>` (sem pedir confirmação) |
| `claude-tools account status` | mostra a conta ativa e as contas salvas |
| `claude-tools help` | mostra o uso |

Nome de conta: letras, números, `-` ou `_`, até 32 caracteres, gravado em minúsculas; `backup` é
reservado. Códigos de saída: `0` sucesso, `1` erro, `2` uso incorreto.

## Testes

```powershell
task test   # go test ./...
task lint   # go vet ./... e gofmt
```
