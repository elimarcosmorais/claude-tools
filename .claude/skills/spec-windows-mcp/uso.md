# Windows-MCP: operar a tela

Leia antes da primeira chamada `mcp__windows-mcp__*` de uma sessão. Configurar o servidor é outra
coisa: está na [SKILL.md](SKILL.md), que só o usuário roda (`/spec-windows-mcp`).

## Os parâmetros, medidos

Saíram do `tools/list` do próprio servidor (`windows-mcp 4.1.0`), nas duas fontes. Só `Shortcut`
(`shortcut`), `Type` (`text`) e `Wait` (`duration`) têm parâmetro obrigatório; todo o resto é
opcional.

| Ferramenta | Parâmetros (default entre parênteses) |
|---|---|
| `Snapshot` | `use_ui_tree` (`true`), `use_annotation` (`true`), `use_vision` (`false`), `use_dom` (`false`), `display`, `region`¹, `width_reference_line`, `height_reference_line` |
| `Screenshot` | `use_annotation` (`false`), `display`, `region`¹, `width_reference_line`, `height_reference_line` |
| `Click` | `loc` (`[x, y]`), `label`, `button` (`left`), `clicks` (`1`) |
| `Type` | **`text`**, `loc`, `label`, `clear` (`false`), `caret_position` (`idle`), `press_enter` (`false`) |
| `Scroll` | `loc`, `label`, `type` (`vertical`), `direction` (`down`), `wheel_times` (`1`) |
| `Move` | `loc`, `label`, `drag` (`false`), `from_loc`, `duration` |
| `App` | `mode` (`launch` \| `launch_executable` \| `resize` \| `switch`), `name`, `executable`, `args`, `cwd`, `window_loc`, `window_size` |
| `Wait` | **`duration`** (segundos) |
| `Shortcut` | **`shortcut`** (ex.: `ctrl+shift+p`) |

¹ `region` (`[left, top, right, bottom]`) **só existe na fonte A** (o fork pinado). Na `0.8.5` do
PyPI, `Snapshot` e `Screenshot` não a têm — medido.

**`mode` é parâmetro do `App`, e de mais ninguém.** Passá-lo no `Snapshot` dá erro de validação do
pydantic. O erro diz que *esse nome* não é dele, não que ele seja sem parâmetros: não generalize
um erro de nome para "a ferramenta não aceita argumento" — confira na descrição da própria
ferramenta, que o Claude Code mostra, ou nesta tabela.

## Uma chamada basta

`Screenshot` não é "a que tira print" e `Snapshot` "a que lê a árvore". O `Snapshot` faz as duas
coisas: com `use_vision: true` ele devolve a árvore **e** a imagem. O `Screenshot` é o caminho
rápido, que pula a travessia da árvore de propósito.

- **Conferir um layout:** um `Snapshot` com `use_vision: true` e, na fonte A, `region` na janela
  do app — em vez de `Screenshot` + `Snapshot`, e sem o desktop inteiro com o editor atrás.
- **`use_ui_tree: false` não devolve a lista de janelas.** Ele é o caminho rápido do `Screenshot`:
  pula a enumeração de janelas junto com a árvore, e a resposta vem com `Focused Window: No
  active window found` e `Opened Windows: No windows found` **mesmo com janelas abertas** — medido
  em sessão real, e o código diz o mesmo (`desktop/service.py`, "Fast path for Screenshot tool
  (use_ui_tree=False): skip window enumeration", igual na `0.8.5`, na `0.8.8` e no fork). Então:
  a lista de janelas só vem com a árvore ligada (o padrão), e para saber se um app subiu o mais
  barato é `Get-Process <nome>` no terminal. Use `use_ui_tree: false` só quando quiser a imagem
  ou os metadados da tela sem a travessia — e aí o `Screenshot` já faz isso.
- **Clicar pelo número:** com `use_annotation: true` (o padrão) a imagem vem com um número em
  cada elemento detectado, e `Click`, `Type`, `Scroll` e `Move` aceitam esse número em `label`,
  sem coordenada. Desligue a anotação (`use_annotation: false`) quando a imagem é para olho
  humano — é o caso do print que vai para o usuário.

## O que os testes mostraram

- **App com WebView2 (Wails, Electron, Tauri) expõe o HTML.** Os elementos da página aparecem na
  árvore com os rótulos acessíveis (`aria-label`, texto do link ou do botão). Dá para navegar e
  validar o layout sem pedir ao usuário.
- **`App` com `mode: "launch_executable"` não passa variáveis de ambiente.** App que precisa de
  uma variável para subir fecha logo após abrir. Embrulhar em `cmd /c "set X=...&& start ..."`
  pelo próprio `App` também não funcionou nos testes. Suba o app pelo terminal com a variável
  definida e use o MCP só para ver e operar. Confira se ele ficou de pé (`Get-Process <nome>`)
  antes de procurar a janela.
- **`App` com `mode: "switch"` casa o nome de forma aproximada.** Com o projeto aberto no VS Code,
  "Dev Tools" trouxe a janela "dev-tools - … - Visual Studio Code". Para focar a janela certa,
  clique no ícone dela na barra de tarefas (`Click` nas coordenadas que o `Snapshot` mostra).
- **Esperar a tela mudar.** Use `Wait` com alguns segundos e confira com um `Snapshot`. Siga sem
  `WaitFor`: o laço dele foi revisado no fork e respeita o `timeout` pedido, mas quem travava era
  a captura dentro dele, e `Wait` + `Snapshot` resolve com menos ferramenta liberada.
- **`Snapshot` no VS Code é a travessia mais pesada, e pode pendurar o editor.** Medido: com
  `editor.accessibilitySupport` em `auto` (o default) ou `off`, a árvore do editor vem
  praticamente vazia e a captura custa 0,05-1,40 s. Com a chave em **`on`**, a mesma captura leva
  23 s **parada numa única chamada UIA** (`TextPattern.GetVisibleRanges` sobre o texto aberto), e
  o Windows confirma o editor pendurado (`IsHungAppWindow`) por 12 s seguidos. Isso é
  independente do hook de input: acontece na `0.8.5` também.

  Duas consequências práticas: nenhum limite de **quantidade** de elementos corta esse caso (a
  chamada presa não conta elemento nenhum), e nenhuma correção do lado do servidor destrava o
  editor — o provider travado é dentro dele. A saída é não entrar no processo:
  `WINDOWS_MCP_EXCLUDE_PROCESSES=Code.exe` no `env`. Na `0.8.8` isso protege **só** o walk (o
  pré-passo de janelas entra no processo excluído de qualquer jeito); na fonte A vale em todo o
  caminho do `Snapshot`. A variável não existe na `0.8.5`.
- **Coordenadas.** Se a imagem vier reduzida, multiplique as coordenadas da imagem pela razão
  entre o tamanho real da tela e o da imagem antes de clicar.
- **Mostrar o print ao usuário.** O resultado do `Snapshot`/`Screenshot` diz onde a imagem foi
  salva (`tool-results\...png`). Mande esse arquivo para o usuário em vez de só descrevê-lo.

## O que ele não alcança: o diálogo do UAC

Ele nasce no **desktop seguro**, num processo de integridade alta (`consent.exe`), e o
`windows-mcp` roda em integridade média, no desktop normal: o `Snapshot` não o vê (é outro
desktop, não aparece no print) e o `Click` não chega nele. Baixar o `PromptOnSecureDesktop` para
`0` não resolve — o UIPI barra input sintético de um processo de integridade média para um de
integridade alta.

Então **não existe** "o agente clica em Sim". Comando que precisa de admin segue a regra 5 do
`behavior.md`: eleve só aquele comando, avise antes no chat, e trate "não autorizado" (recusa ou
expiração do diálogo, ~2 min) como passo não executado, sem prova.

## Só quando o controle de input está ligado

Vale apenas na `0.8.8`+ **sem** `--no-control`. Com `--no-control` (fonte A) ou na `0.8.5` o
subsistema de controle não está em pé, e nada desta seção acontece.

- **Disputa pelo mouse e teclado.** Enquanto o agente opera, o pacote assume o mouse e o teclado
  físicos e mostra um aviso na tela. Se o usuário mexer, o agente perde o controle e só o retoma
  depois de uma espera sem uso. Ação recusada nesse meio-tempo → consulte o `ControlStatus` (dono
  atual e segundos que faltam), espere com `Wait` e tente de novo. **Mouse em uso não é motivo
  para pular a conferência**, nem para passá-la ao usuário: depois de umas 3 esperas sem
  conseguir, pergunte no chat ("Preciso conferir a tela com o Windows-MCP. Solte o mouse e o
  teclado e me avise.") e, com a resposta, confira. O próprio projeto avisa que essa retomada é
  "best-effort, not a security boundary", e que o Windows pode remover um hook de baixo nível sem
  avisar o processo.
- **Tela congelada: `Ctrl+Alt+Shift+Backspace`.** Devolve o input na hora; mover bastante o mouse
  também devolve, e o controle volta ao agente depois de ~10 s sem uso. **Não ofereça esse atalho
  nas outras configurações:** sem o subsistema de controle ninguém o intercepta e ele chega no app
  em foco como tecla normal — num caso real fechou o VS Code em vez de salvar a sessão.
- **`ControlStatus` aparece na lista de ferramentas mesmo com `--no-control`** (medido, dez
  ferramentas na fonte A). Presença dele não quer dizer que o controle está ligado; quem diz é o
  `--no-control` nos `args` do `.mcp.json`.
