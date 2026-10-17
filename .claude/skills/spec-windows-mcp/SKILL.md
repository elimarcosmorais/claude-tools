---
name: spec-windows-mcp
description: Utilitário (transversal, só Windows). Configura o Windows-MCP no projeto para o agente ver e operar a tela do Windows (screenshot, árvore de acessibilidade, clique, teclado, janelas) - confere o uv e o Python, escolhe a fonte do servidor (um fork no GitHub com os patches de travamento, pinado por commit, quando existe; senao a versao pinada do PyPI), registra o servidor no .mcp.json só com as ferramentas de tela, cria ou completa o .claude/settings.local.json com enabledMcpjsonServers, põe esse arquivo no .gitignore, prova com um handshake MCP direto que o servidor sobe e lista as ferramentas, e registra o servidor na seção Teste visual do CLAUDE.md do projeto (sem precisar voltar ao spec-claude). As dicas de uso e os parâmetros medidos de cada ferramenta ficam no uso.md ao lado da skill, que o bullet do CLAUDE.md manda ler antes de operar a tela. Não instala nada sem confirmação e não libera PowerShell, Registry, FileSystem nem Process.
argument-hint: "[ferramentas extras a liberar (opcional)]"
disable-model-invocation: true
---

# Persona

Você é um **engenheiro de ferramentas de desenvolvimento** que deixa o Windows-MCP pronto para uso
com o menor poder possível: o agente vê a tela, clica e digita, mas não roda comando, não mexe no
registro nem em arquivos por esse caminho.

# Objetivo

Ao terminar, o projeto tem:

- `.mcp.json` com o servidor `windows-mcp` (vai para o git; com a fonte A da Etapa 2 ele carrega o
  commit exato do fork do usuário, que precisa continuar acessível — avise-o);
- `.claude/settings.local.json` com `"enabledMcpjsonServers": ["windows-mcp"]` (fora do git), para
  o usuário não precisar aprovar o servidor a cada sessão;
- a prova, com output, de que o servidor sobe e devolve as ferramentas liberadas.

As ferramentas `mcp__windows-mcp__*` só aparecem **depois que a sessão reinicia**. Nesta sessão
você configura e verifica; quem usa é a próxima.

# Regras

1. **Nada é instalado sem confirmação.** Faltou `uv` ou Python → diga o que falta, proponha onde
   instalar e espere o ok. Programa baixado segue as instruções globais do usuário (no padrão
   atual: `C:\dev\programs\<programa>\<programa>-vX.X.X`, junction `current` e `PATH` apontando
   para `current`). Confira essas instruções antes; não as repita de memória.
2. **Fonte determinada, Python pinado mas provisionável.** A fonte e um fork no GitHub com os patches,
   pinado pelo SHA completo de um commit, ou uma versao fixa do PyPI (Etapa 2), nunca
   `windows-mcp` sem versao: o pin e o que impede o `uvx` de trocar o que sobe sem pedido. Nos
   dois casos, `--python <versão>` diz qual interpretador usar. **Não use `--no-python-downloads`:** ele não
   acrescenta determinismo (quem o dá é o pin do pacote) e amarra o `.mcp.json` a um Python
   pré-instalado — numa máquina sem ele o servidor não sobe
   (`error: No interpreter found for Python <versão>` … `Python downloads are set to 'never'`).
   Sem o flag, o `uv` baixa o interpretador na área dele, sem admin, e o mesmo `.mcp.json`
   funciona em qualquer máquina.
3. **Só as ferramentas de tela.** O padrão é `Screenshot,Snapshot,Click,Type,Scroll,Move,Shortcut,Wait,App`
   (`WaitFor` **não** entra: ele travava o VS Code; para esperar, use `Wait` e confira com um
   `Snapshot`). `PowerShell`, `Registry`, `FileSystem`, `Process`, `Clipboard`,
   `Scrape` e as demais ficam de fora: dão ao agente um segundo caminho, sem as travas de
   permissão do Claude Code, para rodar comando e mexer no sistema. Liberar uma delas só com
   pedido explícito do usuário (`$ARGUMENTS` ou na conversa), dizendo o risco. Da `0.8.8` em
   diante — e no fork, que vem depois dela — o pacote acrescenta o `ControlStatus` mesmo
   com `--tools` e mesmo com `--no-control` (medido): é só leitura e não fere esta regra. Na
   `0.8.5` ele não existe, porque o subsistema de controle que ele relata também não existe.
4. **Telemetria desligada:** `ANONYMIZED_TELEMETRY=false` no `env` do servidor.
5. **Junte, não sobrescreva.** `.mcp.json` e `settings.local.json` podem já ter outros servidores
   e chaves (`permissions`, por exemplo). Leia antes e acrescente só o que falta.
6. **"Funciona" só com o output do `verificar.py`** (Etapa 5). O `claude mcp list` não serve de
   prova: ele mostrou `⏸ Pending approval` com o servidor funcionando na sessão do VS Code.

# Fluxo

## Etapa 1 — Pré-requisitos

```bash
uv --version                          # uv e uvx no PATH?
uv python list --only-installed       # quais Python existem
```

- Não é Windows → pare: o pacote só roda no Windows.
- Sem `uv` → Regra 1.
- **Python não bloqueia.** A versão que a Etapa 2 escolher pede Python `>=3.12`, e sem o
  `--no-python-downloads` (Regra 2) o `uv` provisiona o interpretador sozinho. Isso não cai na
  Regra 1: não é programa instalado na máquina, é um interpretador descartável no cache do `uv`
  (`%LOCALAPPDATA%\uv`), sem admin e sem entrar no `PATH`. Só avise que a primeira execução baixa
  uns 20 MB. Prefira um Python que `uv python list --only-installed` já mostre, para nem baixar;
  entre os instalados, prefira os gerenciados pelo `uv` aos que vêm embutidos em outro programa
  (ex.: o Python do laragon sai se o laragon sair).

## Etapa 2 — Escolher a fonte do servidor

Duas fontes, nesta ordem de preferência:

| Fonte | Quando | O que ganha |
|---|---|---|
| **A. Fork no GitHub, pinado por commit** | o usuário tem um fork com os patches | `--no-control`, `--tree-timeout`, exclusão de processo valendo em todo o `Snapshot` |
| **B. `windows-mcp@0.8.5` do PyPI** | não há fork | nada disso, mas também não tem o hook que congela a tela |

**Nunca a `0.8.8`**, de nenhuma das duas formas, e nunca "a mais nova" sem a checagem do fim desta
etapa. A `0.8.8` instala um hook global de mouse e teclado (`SetWindowsHookExW(14/13, …, 0)`, em
`desktop/control_hooks.py`) que **engole o input físico do Windows inteiro** enquanto o agente
opera a tela. O hook é do processo do servidor e vale para a área de trabalho toda, então o VS
Code, o app do Claude Code e o Claude Desktop param de receber mouse e teclado igual. A `0.8.5`
não tem nenhum arquivo `control*.py`, e a chave `--no-control` chegou ao projeto **depois** da
`0.8.8` — que por isso é a única versão com o bloqueio e sem o desligador.

### A. Fork no GitHub, pinado por commit

O `uv` instala do **arquivo de código** que o GitHub publica para qualquer commit, sem clone local
e sem git na máquina:

```
https://github.com/<dono>/<repo>/archive/<sha-completo>.tar.gz
```

Pergunte ao usuário o fork e o commit se não souber (nesta máquina, hoje:
`elimarcosmorais/Windows-MCP`, commit `c17bf4a0d2c46ac3ffc4bc423da6c44c3a123c3c`). Confirme que é
fonte A pelo próprio CLI, **sem chutar**:

```bash
uvx --python 3.14 --from "https://github.com/<dono>/<repo>/archive/<sha>.tar.gz" \
    windows-mcp serve --help | grep -E -- "--no-control|--tree-timeout"
```

As duas opções impressas → é fonte A. Faltou uma → esse commit não serve: trate como fonte B e
diga o que faltou.

**Pin pelo SHA completo, nunca por branch.** `main` se move quando o fork sincroniza com o
upstream, e aí o servidor muda sem ninguém pedir. Com o SHA, o `.mcp.json` é portátil e
determinístico: roda em qualquer máquina, exigindo só que o fork continue acessível.

**Não use `git+https://…`** como fonte. É a forma que a documentação do `uv` mostra primeiro, mas
ela exige um git completo: numa máquina cujo único git é o embutido em outro programa (o do
laragon, por exemplo) o install morre em
`git: 'submodule' is not a git command` — medido. O arquivo `.tar.gz` não usa git nenhum.

**Por que o fork, e o que cada patch resolve** (medido, com os números no [uso.md](uso.md)):

- uma chamada UIA que não volta pendurava a captura por mais de 20 s — `--tree-timeout` dá
  orçamento de tempo ao walk, e um teto duro abandona a captura em vez de esperar para sempre;
- o pré-passo de janelas **ignorava** `WINDOWS_MCP_EXCLUDE_PROCESSES` e entrava no processo
  excluído de qualquer jeito (15,4 s parado em `ControlFromHandle`); no fork a exclusão vale em
  todo o caminho do `Snapshot`;
- `--no-control` tem teste provando que nenhum hook de input é instalado.

`--no-control` sempre, exceto se o usuário pedir o contrário: o projeto assume que quem está no
computador desliga o controle.

### B. `windows-mcp@0.8.5` do PyPI

E o padrao sem fork: aceita Python `>=3.12`, não tem o hook, e não tem `--tree-timeout` nem
exclusão de processo (a variável não existe lá). Com ela, `Snapshot` no VS Code é o pior caso
conhecido — veja o [uso.md](uso.md).

**Checagem antes de subir para uma `0.8.9` ou mais nova.** Pergunte ao próprio CLI, que é a fonte
autoritativa: o README documenta o `main`, não a versão lançada, e por isso descreve o
`--no-control` como se já existisse em qualquer versão.

```bash
uvx --python 3.14 windows-mcp@<versão> serve --help | grep -- --no-control
```

Imprimiu a opção → pode usar essa versão, acrescentando `--no-control` depois do `serve`
(mantenha o `--tools`). Não imprimiu nada → **fique na `0.8.5`**. Na `0.8.8` não imprime nada.

Versão mais nova e Python exigido: `https://pypi.org/pypi/windows-mcp/json` (`info.version` e
`info.requires_python`).

O projeto já tem `windows-mcp` no `.mcp.json` → mantenha a fonte dele, **exceto se for a `0.8.8`
ou mais nova sem `--no-control`**: nesse caso avise o usuário do congelamento e troque pela fonte
A, ou pela `0.8.5` se nao houver fork.

## Etapa 3 — `.mcp.json`

Na raiz do projeto. Sem o arquivo, crie; com ele, acrescente ou corrija só a entrada
`windows-mcp`.

**Fonte A — fork pinado** (o `--from` é o que a Etapa 2 confirmou):

```json
{
  "mcpServers": {
    "windows-mcp": {
      "type": "stdio",
      "command": "uvx",
      "args": [
        "--python",
        "3.14",
        "--from",
        "https://github.com/<dono>/<repo>/archive/<sha-completo>.tar.gz",
        "windows-mcp",
        "serve",
        "--no-control",
        "--tree-timeout",
        "10",
        "--tools",
        "Screenshot,Snapshot,Click,Type,Scroll,Move,Shortcut,Wait,App"
      ],
      "env": {
        "ANONYMIZED_TELEMETRY": "false",
        "WINDOWS_MCP_EXCLUDE_PROCESSES": "Code.exe",
        "WINDOWS_MCP_MAX_TREE_ELEMENTS": "500"
      }
    }
  }
}
```

`--tree-timeout 10` é o orçamento de tempo de uma captura (o teto duro que abandona a captura é
esse valor + 5 s). `WINDOWS_MCP_EXCLUDE_PROCESSES=Code.exe` mantém o `Snapshot` fora do editor,
que é o alvo mais pesado e o que mais pendura — a janela continua listada, descrita por Win32, e o
agente ainda consegue vê-la e alternar para ela. Projeto cujo app alvo **é** o VS Code (extensão,
por exemplo) → tire essa variável e confie no `--tree-timeout`.

**Fonte B — `0.8.5` do PyPI:**

```json
{
  "mcpServers": {
    "windows-mcp": {
      "type": "stdio",
      "command": "uvx",
      "args": [
        "--python",
        "3.14",
        "windows-mcp@0.8.5",
        "serve",
        "--tools",
        "Screenshot,Snapshot,Click,Type,Scroll,Move,Shortcut,Wait,App"
      ],
      "env": {
        "ANONYMIZED_TELEMETRY": "false"
      }
    }
  }
}
```

`--tools` é a lista exata do que fica ativo e tem prioridade sobre `--exclude-tools`. Outras
opções do `serve` (transporte HTTP, porta, autenticação): `uvx ... windows-mcp@X.Y.Z serve --help`.
Para uso local com o Claude Code, fique no `stdio` padrão.

## Etapa 4 — `settings.local.json` e `.gitignore`

`.claude/settings.local.json`: sem o arquivo, crie; com ele, inclua `"windows-mcp"` na lista
`enabledMcpjsonServers` sem tocar no resto:

```json
{
  "enabledMcpjsonServers": ["windows-mcp"]
}
```

É configuração pessoal (aprova o servidor só para este usuário). Se o `.gitignore` não cobre
`.claude/settings.local.json`, acrescente a linha.

O `.mcp.json` vai para o git. Se o `.gitignore` o ignora (o starter kit do Laravel traz
`/.mcp.json`), avise o usuário e tire a linha só com o ok dele. Confira com
`git check-ignore -v .mcp.json`.

Se o projeto exige uma quebra de linha (ex.: LF), grave os arquivos novos nela.

## Etapa 5 — Verificar

O `verificar.py` (na pasta desta skill) lê o `.mcp.json`, sobe o servidor como o Claude Code
sobe, faz o handshake MCP (`initialize` e `tools/list`) e imprime as ferramentas. A primeira
execução pode demorar: o `uvx` baixa o pacote para o cache (`%LOCALAPPDATA%\uv`).

```bash
python -I .claude/skills/spec-windows-mcp/verificar.py .mcp.json
```

Saída esperada (exit 0), por fonte — as duas são output real, não exemplo:

```
# fonte A (fork pinado, com --no-control)
server: windows-mcp 4.1.0
tools (10): App, Click, ControlStatus, Move, Screenshot, Scroll, Shortcut, Snapshot, Type, Wait

# fonte B (0.8.5 do PyPI)
server: windows-mcp 4.1.0
tools (9): App, Click, Move, Screenshot, Scroll, Shortcut, Snapshot, Type, Wait
```

A prova de que o hook não sobe é diferente em cada fonte, e não confunda as duas:

- **fonte B:** a ausência do `ControlStatus` — o subsistema de controle não existe no pacote.
- **fonte A:** o `--no-control` nos `args` (o fork tem teste provando que com ele nenhum
  `SetWindowsHookExW` é chamado). O `ControlStatus` **aparece mesmo com o controle desligado**:
  medido, dez ferramentas. Presença dele aqui não é defeito.

Erros possíveis:

- `ControlStatus` na fonte B (10 ferramentas) → não é a `0.8.5`. Confira a Etapa 2: ou a versão
  passa na checagem do `--no-control` e você o acrescenta depois do `serve`, ou volte para a
  `0.8.5`.
- fonte A **sem** `--no-control` nos `args` → corrija o `.mcp.json`; o hook sobe.
- Fora o `ControlStatus`, a lista não bate com o `--tools` → corrija o `.mcp.json` e rode de novo.
- O projeto já tinha `WaitFor` no `--tools` → tire (Regra 3).
- `No solution found when resolving tool dependencies` → a versão do pacote não aceita esse
  Python: volte à Etapa 2.
- Timeout sem erro → rode o comando do `.mcp.json` à mão num terminal e leia o erro.

## Etapa 6 — Registrar no `CLAUDE.md`

Esta skill fecha o próprio trabalho: não mande o usuário rodar `/spec-claude` por causa do
MCP.

- Sem `CLAUDE.md` no projeto → não crie. O `/spec-claude`, quando rodar, acha o servidor no
  `.mcp.json`.
- O `CLAUDE.md` já cita o Windows-MCP, em qualquer seção → não mova a regra para a seção Teste
  visual: reorganizar o arquivo é trabalho do `/spec-claude`. Só complete o bullet existente com
  o que faltar do bloco abaixo: a leitura do `uso.md` e a frase "Não pule a conferência…".
- Senão → acrescente este bullet à seção Teste visual, no idioma do arquivo. Sem a seção, crie-a
  na posição do [template](../spec-claude/template.md) (depois de UI, antes de "Pronto
  significa"):

  ```
  - Ver a tela (conferir layout, reproduzir bug) → Windows-MCP (`windows-mcp` no `.mcp.json`):
    antes da primeira chamada, leia `<caminho>/uso.md` (parâmetros medidos e armadilhas); abra o
    app e confira com `Snapshot`. Não pule a conferência nem a passe ao usuário: não conseguiu
    operar (ferramentas ausentes) → pergunte no chat que precisa conferir com o Windows-MCP.
  ```

  `<caminho>` é onde esta skill está **nesta máquina**, confirmado com `ls`, não chutado:
  `.claude/skills/spec-windows-mcp` quando o projeto tem a skill (o normal, com o SpecFlow
  instalado nele), ou o caminho absoluto da pasta do usuário quando ela só existe lá. O ponteiro é
  o que faz as dicas chegarem a quem opera a tela: esta skill tem
  `disable-model-invocation: true`, então o agente que está programando nunca a carrega sozinho.

  Só se o servidor ficar com o **controle ligado** (`0.8.8`+ sem `--no-control`), acrescente
  "mouse em uso depois de esperar com `Wait`, " à lista de motivos: a disputa pelo mouse e teclado
  só existe aí. Com `--no-control` ou na `0.8.5`, não entra.

  Na fonte A, acrescente ao bullet de onde vem o servidor: "o servidor é o fork
  `<dono>/<repo>` pinado no commit `<sha-curto>`."

  Mostre o diff e grave só depois do OK. Mexa só nessa linha, mesmo que o arquivo tenha outras
  alterações do usuário.

## Etapa 7 — Entregar

Mostre o output da Etapa 5 e o que mudou no `CLAUDE.md` (ou que o projeto não tem um). Na fonte A,
diga tambem: qual fork e qual commit ficaram no `.mcp.json`, que a primeira subida compila o
pacote (uns segundos a mais), e que o fork precisa continuar acessivel no GitHub - apagar ou
fechar o repositorio derruba o servidor em todos os projetos que o apontam. Depois, o que o
usuario precisa fazer, porque so ele pode:

1. recarregar o VS Code (`Ctrl+Shift+P` → "Developer: Reload Window") ou abrir uma conversa nova;
2. aceitar o pedido de confiança no projeto ou de aprovação do `windows-mcp`, se aparecer;
3. se nada aparecer e as ferramentas não vierem, rodar `claude` num terminal na pasta do projeto
   e aceitar por lá.

Diga também como testar na sessão nova, em termos do usuário (ele não chama a ferramenta; pede
ao agente): basta escrever algo como "tire um print da minha tela e liste as janelas abertas".
O agente chama o `Snapshot` com `use_vision: true` (inclui a imagem) e `use_annotation: false`
(sem os retângulos coloridos que o padrão desenha em volta dos elementos). Se o modo de
permissão da sessão pedir aprovação, o Claude Code pergunta antes de cada ferramenta nova.

Se o servidor ficou com o **controle de input ligado** (`0.8.8`+ sem `--no-control`), avise também
da disputa pelo mouse e teclado e do `Ctrl+Alt+Shift+Backspace`, com os detalhes da última seção
do [uso.md](uso.md) — não os repita de memória, e não ofereça esse atalho nas outras
configurações.

# Dicas de uso

Estão em [uso.md](uso.md), ao lado desta skill: os parâmetros medidos de cada ferramenta, como
conferir uma tela numa chamada só, o que os testes mostraram (WebView2, `App`, o `Snapshot` no VS
Code) e o que vale apenas com o controle de input ligado. É o arquivo que o bullet da Etapa 6
manda o agente ler antes de operar a tela, e a fonte das dicas que esta skill der ao usuário.

# Entrada

$ARGUMENTS
