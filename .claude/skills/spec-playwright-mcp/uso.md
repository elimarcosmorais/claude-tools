# Playwright MCP: operar uma página

Leia antes da primeira chamada `mcp__playwright__*` de uma sessão. Configurar o servidor é outra
coisa: está na [SKILL.md](SKILL.md), que só o usuário roda (`/spec-playwright-mcp`).

## Os parâmetros, medidos

Saíram do `tools/list` do próprio servidor, na `@playwright/mcp@0.0.83` (25 ferramentas,
`Playwright 1.64.0-alpha`) — a versão que a skill fixa. **Obrigatório em negrito**, default entre
parênteses. Se o projeto fixa outra versão, confirme o parâmetro na descrição da própria
ferramenta, que o Claude Code mostra, antes de usá-lo.

| Ferramenta | Parâmetros |
|---|---|
| `browser_navigate` | **`url`** |
| `browser_snapshot` | `target`, `depth`, `boxes`, `filename` |
| `browser_find` | `text` \| `regex`, `filename` |
| `browser_take_screenshot` | `element`, `target`, `fullPage`, `type`, **`scale`** (`css`), `filename` |
| `browser_click` | `element`, **`target`**, `doubleClick`, `button` (`left`), `modifiers` |
| `browser_type` | `element`, **`target`**, **`text`**, `submit`, `slowly` |
| `browser_fill_form` | **`fields`** |
| `browser_select_option` | `element`, **`target`**, **`values`** |
| `browser_hover` | `element`, **`target`** |
| `browser_press_key` | **`key`** |
| `browser_wait_for` | `time` (s, máximo 30), `text`, `textGone` |
| `browser_console_messages` | **`level`** (`info`), `all`, `filename` |
| `browser_network_requests` | **`static`** (`false`), `filter` (regexp), `filename` |
| `browser_network_request` | **`index`**, `part`, `filename` |
| `browser_resize` | **`width`**, **`height`** |
| `browser_tabs` | **`action`**, `index`, `url` |
| `browser_evaluate` | `element`, `target`, **`function`**, `filename` |
| `browser_handle_dialog` | **`accept`**, `promptText` |
| `browser_drag` | `startElement`, **`startTarget`**, `endElement`, **`endTarget`** |
| `browser_drop` | `element`, **`target`**, `paths`, `data` |
| `browser_file_upload` | `paths` |
| `browser_emulate_media` | `colorScheme`, `reducedMotion`, `forcedColors`, `contrast`, `media` |
| `browser_navigate_back`, `browser_close` | sem parâmetro |
| `browser_run_code_unsafe` | **bloqueada** no `deny` do projeto — não chame |

**`target` é a referência do snapshot** (ou um seletor único) e é o que de fato aponta o
elemento; `element` é só a descrição legível que o servidor usa no pedido de permissão. Agir sem
um snapshot antes não funciona: a referência vem dele.

## O que economiza token

- **`browser_find` em vez de um snapshot inteiro.** Procura texto ou regexp na árvore e devolve
  só os nós que casam, com algumas linhas de contexto e o caminho até a raiz. É o caminho padrão
  quando você já sabe o que está procurando ("o botão Salvar aparece?").
- **`depth` e `target` no `browser_snapshot`.** `depth` corta a profundidade da árvore; `target`
  captura a subárvore de um elemento só, em vez da página toda.
- **`filename` quase em toda ferramenta.** Grava o resultado num arquivo (resolvido a partir da
  raiz do projeto) em vez de devolvê-lo na resposta. Vale para snapshot, print, console e rede.
- **`boxes: true`** acrescenta a caixa de cada elemento (`[box=x,y,width,height]`, em pixels CSS
  relativos à viewport) quando a posição importa — layout quebrado, elemento sobreposto.

## Duas armadilhas medidas

- **O console vem filtrado por padrão.** `browser_console_messages` tem `level` em `info` (cada
  nível inclui os mais graves) e `all` em `false`, o que mostra só as mensagens **desde a última
  navegação**. Erro que aconteceu antes dela não aparece: use `all: true`, e `level: "error"` para
  ir direto ao que importa.
- **A rede esconde os estáticos.** `browser_network_requests` tem `static` em `false`, então
  imagem, fonte e script que deram certo ficam fora da lista. Procure 4xx e 5xx aí mesmo; para o
  detalhe de uma requisição, `browser_network_request` com o número dela. `filter` aceita regexp
  (`"/api/.*user"`).

## O que os testes mostraram

- **O app precisa estar no ar.** O MCP só abre URL. Suba o servidor de dev com o comando do
  `CLAUDE.md` em segundo plano e navegue para o endereço local dele.
- **`browser_snapshot` x `browser_take_screenshot`.** O snapshot é a árvore de acessibilidade e é
  com ele que se age (é dele que sai o `target`). O print só mostra — a descrição da própria
  ferramenta diz que não dá para agir a partir dele. Para conferir o visual: `fullPage: true`
  pega a página rolável inteira, e `element` + `target` recortam um elemento.
- **Confira console e rede antes de dizer que a tela está certa.** Tela bonita com erro de
  JavaScript ou um 500 numa chamada de API não é tela certa.
- **Responsivo.** `browser_resize` (largura e altura obrigatórias) muda o tamanho da janela para
  conferir no celular sem reconfigurar o servidor.
- **Login.** Com `--isolated`, cada sessão começa deslogada. Use o usuário de teste que o projeto
  já documenta (seed, `README`); sem um, pergunte. Nunca digite uma senha real do usuário.
- **Mostrar o print ao usuário.** O resultado do `browser_take_screenshot` diz onde a imagem foi
  salva (`.playwright-mcp/…png`). Mande esse arquivo em vez de só descrevê-lo.
- **`browser_run_code_unsafe` não é alternativa.** O `deny` do projeto a bloqueia porque ela roda
  JavaScript arbitrário no processo do servidor, fora das travas do Claude Code. Precisou avaliar
  algo na página → `browser_evaluate`, que roda no contexto da página.
