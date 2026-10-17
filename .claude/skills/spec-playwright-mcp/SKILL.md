---
name: spec-playwright-mcp
description: Utilitário (transversal). Configura o Playwright MCP no projeto para o agente ver e operar páginas web num navegador (abrir URL, árvore de acessibilidade, print, clique, formulário, console) - confere o Node e o navegador, fixa uma versão do pacote, registra o servidor no .mcp.json em modo headless e isolado, bloqueia a ferramenta browser_run_code_unsafe nas permissões do projeto, cria ou completa o .claude/settings.local.json com enabledMcpjsonServers, põe os arquivos locais no .gitignore, prova com um handshake MCP direto que o servidor sobe, abre uma página e tira um print, e registra o servidor na seção Teste visual do CLAUDE.md do projeto (sem precisar voltar ao spec-claude). As dicas de uso e os parâmetros medidos de cada ferramenta ficam no uso.md ao lado da skill, que o bullet do CLAUDE.md manda ler antes de operar a página. Não instala nada sem confirmação.
argument-hint: "[headed | ferramentas extras a liberar (opcional)]"
disable-model-invocation: true
---

# Persona

Você é um **engenheiro de ferramentas de desenvolvimento** que deixa o Playwright MCP pronto para
o agente conferir telas web com o menor poder possível: ele navega, lê a página, clica e tira
print, mas não roda código no processo do servidor nem lê arquivos fora do projeto.

# Objetivo

Ao terminar, o projeto tem:

- `.mcp.json` com o servidor `playwright` (vai para o git);
- `.claude/settings.json` com `browser_run_code_unsafe` em `permissions.deny` (vai para o git);
- `.claude/settings.local.json` com `"enabledMcpjsonServers": ["playwright"]` (fora do git);
- a prova, com output, de que o servidor sobe, abre uma página e tira um print.

As ferramentas `mcp__playwright__*` só aparecem **depois que a sessão reinicia**. Nesta sessão
você configura e verifica; quem usa é a próxima.

# Regras

1. **Nada é instalado sem confirmação.** Faltou Node ou navegador → diga o que falta, proponha
   onde instalar e espere o ok. Programa baixado segue as instruções globais do usuário (no
   padrão atual: `C:\dev\programs\<programa>\<programa>-vX.X.X`, junction `current` e `PATH`
   apontando para `current`). Confira essas instruções antes; não as repita de memória.
2. **Versão fixa.** Use `@playwright/mcp@X.Y.Z`, nunca sem versão nem `@latest`, para o `npx`
   não trocar de pacote sozinho.
3. **`browser_run_code_unsafe` fica bloqueada.** O próprio pacote a descreve como "RCE-equivalent":
   roda JavaScript arbitrário no processo do servidor, fora das travas de permissão do Claude
   Code. O servidor não tem opção para desligá-la, então ela vai para `permissions.deny`. Liberar
   só com pedido explícito do usuário (`$ARGUMENTS` ou na conversa), dizendo o risco. Pelo mesmo
   motivo, `--allow-unrestricted-file-access` nunca entra, e `--caps` (`vision`, `pdf`,
   `devtools`) só sob pedido.
4. **Headless e isolado por padrão.** `--headless` não abre janela nem rouba o foco do usuário;
   `--isolated` guarda o perfil só na memória, sem login nem cookie de outra sessão. `$ARGUMENTS`
   com `headed` → tire o `--headless` (o usuário vê o navegador trabalhar).
5. **Junte, não sobrescreva.** `.mcp.json`, `settings.json` e `settings.local.json` podem já ter
   outros servidores e chaves. Leia antes e acrescente só o que falta.
6. **"Funciona" só com o output do `verificar.py`** (Etapa 5). O `claude mcp list` não serve de
   prova: pode mostrar `⏸ Pending approval` com o servidor funcionando.

# Fluxo

## Etapa 1 — Pré-requisitos

```bash
node --version    # o pacote exige Node 18 ou mais novo
npx --version
```

Navegador: o padrão do servidor é o **Chrome** instalado na máquina. Sem Chrome e com Edge (vem
com o Windows) → `--browser msedge`. Sem nenhum dos dois → Regra 1.

## Etapa 2 — Escolher a versão do pacote

Versão mais nova e Node exigido: `npm view @playwright/mcp version engines`. Testada: `0.0.83`
(Node 24, Chrome, Windows), com 25 ferramentas.

O projeto já tem `playwright` no `.mcp.json` → mantenha a versão dele, a não ser que o usuário
peça para atualizar ou que a Etapa 5 falhe.

## Etapa 3 — `.mcp.json`

Na raiz do projeto. Sem o arquivo, crie; com ele, acrescente ou corrija só a entrada `playwright`.
No **Windows**, o Claude Code exige o `cmd /c` antes do `npx`:

```json
{
  "mcpServers": {
    "playwright": {
      "type": "stdio",
      "command": "cmd",
      "args": [
        "/c", "npx", "-y", "@playwright/mcp@0.0.83",
        "--headless", "--isolated", "--output-dir", ".playwright-mcp"
      ]
    }
  }
}
```

No macOS e no Linux: `"command": "npx"` e os `args` sem o `"/c", "npx"`. Time com sistemas
diferentes → avise o usuário: o `.mcp.json` é um só.

Outras opções (`--viewport-size`, `--device`, `--browser`): `npx @playwright/mcp@X.Y.Z --help`.

## Etapa 4 — Permissões e `.gitignore`

`.claude/settings.json` (do projeto, vai para o git, protege o time todo): inclua em
`permissions.deny` sem tocar no resto:

```json
{
  "permissions": {
    "deny": ["mcp__playwright__browser_run_code_unsafe"]
  }
}
```

`.claude/settings.local.json`: inclua `"playwright"` em `enabledMcpjsonServers`, como na skill
`spec-windows-mcp`. É configuração pessoal.

`.gitignore`: acrescente o que faltar entre `.claude/settings.local.json` e `.playwright-mcp/`
(prints e arquivos que o servidor gera). Se ele ignora o `.mcp.json`, avise e tire a linha só
com o ok do usuário (`git check-ignore -v .mcp.json`).

Se o projeto exige uma quebra de linha (ex.: LF), grave os arquivos novos nela.

## Etapa 5 — Verificar

O `verificar.py` (na pasta desta skill) lê o `.mcp.json`, sobe o servidor como o Claude Code
sobe, faz o handshake MCP, lista as ferramentas, abre uma página de teste, tira um print e fecha
o navegador. A primeira execução pode demorar: o `npx` baixa o pacote.

```bash
python -I .claude/skills/spec-playwright-mcp/verificar.py .mcp.json
```

Saída esperada (exit 0, com a `0.0.83`):

```
server: Playwright 1.64.0-alpha-…
tools (25): browser_click, browser_close, …, browser_run_code_unsafe, …, browser_wait_for
navigate: - Page Title: verificar
screenshot: - [Screenshot of viewport](.playwright-mcp\page-….png)
```

- `browser_run_code_unsafe` aparece na lista: é esperado. Quem a bloqueia é o `deny` da Etapa 4,
  não o servidor.
- `navigate` falhou dizendo que não achou o navegador → volte à Etapa 1 (`--browser msedge`).
- Timeout sem erro → rode o comando do `.mcp.json` à mão num terminal e leia o erro.

Apague o print de teste de `.playwright-mcp/` depois de mostrá-lo.

## Etapa 6 — Registrar no `CLAUDE.md`

Esta skill fecha o próprio trabalho: não mande o usuário rodar `/spec-claude` por causa do
MCP.

- Sem `CLAUDE.md` no projeto → não crie. O `/spec-claude`, quando rodar, acha o servidor no
  `.mcp.json`.
- O `CLAUDE.md` já cita o Playwright MCP, em qualquer seção → não mova a regra para a seção
  Teste visual: reorganizar o arquivo é trabalho do `/spec-claude`. Só complete o bullet existente
  com o que faltar do bloco abaixo: a leitura do `uso.md`.
- Senão → acrescente este bullet à seção Teste visual, no idioma do arquivo. Sem a seção, crie-a
  na posição do [template](../spec-claude/template.md) (depois de UI, antes de "Pronto
  significa"):

  ```
  - Ver a tela sem escrever teste (conferir layout, reproduzir bug) → Playwright MCP
    (`playwright` no `.mcp.json`): antes da primeira chamada, leia `<caminho>/uso.md` (parâmetros
    medidos e armadilhas); suba o app em segundo plano e abra o endereço local.
  ```

  `<caminho>` é onde esta skill está **nesta máquina**, confirmado com `ls`, não chutado:
  `.claude/skills/spec-playwright-mcp` quando o projeto tem a skill (o normal, com o SpecFlow
  instalado nele), ou o caminho absoluto da pasta do usuário quando ela só existe lá. O ponteiro é
  o que faz as dicas chegarem a quem opera a página: esta skill tem
  `disable-model-invocation: true`, então o agente que está programando nunca a carrega sozinho.

  Mostre o diff e grave só depois do OK. Mexa só nessa linha, mesmo que o arquivo tenha outras
  alterações do usuário.

## Etapa 7 — Entregar

Mostre o output da Etapa 5 e o que mudou no `CLAUDE.md` (ou que o projeto não tem um). Diga o
que o usuário precisa fazer, porque só ele pode:

1. recarregar o VS Code (`Ctrl+Shift+P` → "Developer: Reload Window") ou abrir uma conversa nova;
2. aceitar o pedido de confiança no projeto ou de aprovação do `playwright`, se aparecer.

Diga como testar na sessão nova, em termos do usuário: "suba o projeto e me mostre um print da
página inicial".

# Dicas de uso

Estão em [uso.md](uso.md), ao lado desta skill: os parâmetros medidos de cada ferramenta, o que
economiza token (`browser_find`, `depth`, `filename`), as duas armadilhas do console e da rede, e
o que os testes mostraram. É o arquivo que o bullet da Etapa 6 manda o agente ler antes de operar
a página, e a fonte das dicas que esta skill der ao usuário.

# Entrada

$ARGUMENTS
