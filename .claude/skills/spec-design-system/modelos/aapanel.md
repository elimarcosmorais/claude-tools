# Modelo de design system — aaPanel

> Retrato de um design system que já existe (ex.: o Primer, do GitHub), guardado na
> `/spec-design-system` para ser aplicado em qualquer projeto com
> `/spec-design-system modelo <slug>`. Descreve **o design system**, não um produto: nome,
> público e plataformas do projeto são perguntados quando o modelo é aplicado.

> **Origem desta captura:** o aaPanel não publica um site de documentação de design system.
> Todos os valores abaixo foram lidos do próprio painel em execução (aaPanel 8.0.7, "New UI",
> modo escuro ativo), a partir de: (a) as 285 variáveis CSS declaradas em `:root` e
> `:root[theme-mode="dark"]` de `/static/vite/css/app.css`; (b) o objeto `themeOverrides` passado
> ao `NConfigProvider` do Naive UI; (c) estilos computados de componentes na tela.
> Onde aparece **Fonte: painel — …**, é o arquivo ou objeto de onde o valor saiu.
> O `themeOverrides` capturado é o do modo escuro; os valores do modo claro vêm das variáveis CSS.

**Notação:** `[?]` = a fonte não diz (ao aplicar, a skill decide pelo default consagrado ou
pergunta) · `[—]` = o design system não define isso.

---

## 1. Identificação → `docs/design-system.md`

- **Nome** (obrigatório): aaPanel UI (interface "New UI" do aaPanel)
- **Slug** (obrigatório): `aapanel`
- **Mantido por** (obrigatório): aaPanel (equipe do BT/宝塔 Panel)
- **Link oficial** (obrigatório): https://www.aapanel.com — não há documentação pública de design
  system; a referência é o próprio painel
- **Versão capturada** (obrigatório): aaPanel 8.0.7 (New UI); versão do Naive UI embutido `[?]`
- **Data da captura** (obrigatório): 2026-10-30
- **Licença** (obrigatório): `[?]` — o aaPanel é distribuído gratuitamente, mas não publica licença
  para reuso dos tokens/estilos. Fontes: Inter (SIL OFL 1.1); PingFang SC e Microsoft YaHei são
  fontes de sistema proprietárias (Apple/Microsoft). Ícones: coleções Iconify, cada uma com licença
  própria (ex.: Lucide ISC, MDI Apache 2.0, Carbon Apache 2.0, Ant Design MIT, Font Awesome CC BY
  4.0). Uso fora da organização dona permitido: `[?]`
- **Plataformas cobertas** (obrigatório): web (desktop; responsivo limitado)
- **Pacote de tokens:** nenhum publicado — tokens existem como variáveis CSS em `app.css`
  (`--color-*`, `--input-*`, `--modal-*` etc.) + `themeOverrides` do Naive UI em JS (`[—]` como pacote)
- **Implementações oficiais** (obrigatório — uma linha por stack; nenhuma → só a linha `nenhuma`):

| Stack | Pacote | Versão | Como importar |
|---|---|---|---|
| nenhuma | — | — | — |

- **Notas:** Pilha observada no painel: Vue 3 + Vite, biblioteca de componentes **Naive UI**
  (classes `n-*`, `NConfigProvider` com `theme = dark` e `themeOverrides`), utilitários atômicos
  estilo Tailwind/UnoCSS (`flex`, `items-center`, `h-28px`, `mb-16px`, `dark:bg-gray-700`) e ícones
  via UnoCSS preset-icons/Iconify (`i-<coleção>:<ícone>`). Telas legadas ainda carregam Bootstrap
  3.3.5, layui (`layui-layer`), CodeMirror e Ace. Componentes próprios com prefixo `bt-`
  (ex.: `BtDialogConfig`, `bt-tabs-modal`).

## 2. Direção visual → `docs/design-system.md` (seção Direção)

- **Fonte:** painel — observação das telas Início e Site + `app.css`
- **Personalidade em 3 adjetivos** (obrigatório): utilitário, técnico, compacto
- **Eixos** (obrigatório) — marque um ponto de 1 a 5 em cada:
  - sóbrio `1 2 3 4 5` expressivo: **2**
  - denso `1 2 3 4 5` espaçado: **2**
  - geométrico `1 2 3 4 5` orgânico: **2**
  - plano `1 2 3 4 5` com profundidade: **2**
- **Densidade de informação:** compacta — texto base 12px, controles de 32px (26px no tamanho
  small), tabela com padding de célula 6–8px; cards de métricas na home usam números maiores
  (15–18px)
- **O que o design system evita** (obrigatório): sombras fortes (cards com `0 0 8px rgba(0,0,0,.06)`
  ou só mudança de fundo), cor fora do verde de marca para ação primária, MAIÚSCULAS em botão,
  gradientes no modo escuro (só existem em alguns fundos do modo claro)
- **Tom do microcopy:** `[—]` (não há guia; o painel é traduzido para vários idiomas, com tom
  direto/instrucional — "Por favor, insira…", "Não há dados")
- **Notas:** Verde `#20a53a` é a assinatura visual (marca, ação primária, sucesso, item ativo).

## 3. Cores → `docs/design/cores.md`

- **Fonte:** painel — `/static/vite/css/app.css` (`:root`) e `themeOverrides.common`
- **Modelo da paleta** (obrigatório): cores chapadas (não há escala de tons por cor)
- **Convenção de nome dos tokens** (obrigatório): base/semântica `--color-<papel>[-n]`
  (`--color-primary`, `--color-text-1`, `--color-bg-2`); por componente `--<componente>-<parte>-<propriedade>`
  (`--input-text-color`, `--modal-header-bg`, `--scrollbar-thumb-bg-color`)
- **Paleta base** (obrigatório — uma linha por cor, no modo claro):

| Cor | Tons (token = valor) |
|---|---|
| verde (primária/sucesso) | `--color-primary = #20a53a` · `--color-primary-1 = #e4f4e7` · `--primary-button-color-hover = #1d9534` · `--primary-button-color-pressed = #1a8a30` |
| laranja (aviso) | `--color-warning = #ffae45` |
| vermelho (erro) | `--color-error = #e73a33` |
| âmbar (Pro/pago) | `--color-pro = #ff8f00` · `--pay-color-bg = #fff5e1` · `--pay-color-border = #ecb566` |
| neutros de texto | `--color-text-1 = #131313` · `--color-text-5 = #333` · `--color-text-base = #3a424d` · `--color-text-4 = #666` · `--color-text-3 = #999` |
| neutros de fundo | `--color-bg-2 = #fff` · `--color-bg-4 = #f7f7f7` · `--color-bg-1 = #f2f5f9` · `--color-bg-3 = #f2f5f9` |
| borda | `--color-border = #dcdfe6` |

- **Primária ou acento:** `--color-primary = #20a53a` — botão primário, item de menu ativo, links,
  switch/checkbox/radio marcados, borda de foco, cursor de texto, contadores de status
- **Cores de dados (gráficos):** `[?]` — gráficos usam o verde primário nos anéis de uso (CPU/RAM);
  tooltip de gráfico: `--chart-tooltip-bg-color = #fff | #181818`; sem ordem de série definida
- **Contraste alvo** (obrigatório): `[?]` (não declarado)
- **Regras de uso:** `[—]` (não há guia escrito)
- **Notas:** No modo escuro, vários estados "ativos" deixam de ser verdes e passam a neutros
  (ex.: `--router-menu-active-text`, `--radio-active-dot-color`, `--check-color-checked`,
  `--pagination-item-active-color` → `var(--color-text-base)`).

## 4. Tipografia → `docs/design/tipografia.md`

- **Fonte:** painel — estilo computado do `body` e `themeOverrides.common`
- **Família de texto** (obrigatório): `PingFang SC` → `Inter` → `Microsoft YaHei` → `Segoe UI` —
  fontes do sistema (Inter, se não instalada, cai no fallback) — licença: PingFang/YaHei
  proprietárias de sistema; Inter SIL OFL 1.1
- **Família de título:** mesma do texto
- **Família monoespaçada:** `[?]` (editores Ace/CodeMirror usam a monoespaçada própria deles)
- **Fallback stack** (obrigatório): `"PingFang SC", Inter, "Microsoft YaHei", "Segoe UI", sans-serif`
- **Unidade do tamanho de texto** (obrigatório): px
- **Estilos nomeados** (obrigatório):

| Estilo (token oficial) | Tamanho | Peso | Altura de linha | Espaçamento entre letras | Uso |
|---|---|---|---|---|---|
| `common.fontSize` / `--el-font-size-base` | 12px | 400 | normal | 0 | texto base, botões, campos, tabela, rótulos de formulário |
| `Alert.fontSize` / `Breadcrumb.fontSize` | 13px / 14px | 400 | normal | 0 | alerta / breadcrumb |
| item de menu lateral (sem token) | 14px | 400 | 24.5px | 0 | navegação lateral |
| `Dialog.titleFontSize` / `Drawer.titleFontSize` / `Tabs.tabFontSizeLarge` / `Message.fontSize` | 14px | 500 (título) | normal | 0 | título de diálogo/gaveta, abas grandes, mensagens |
| `Card.titleFontSizeMedium` | 16px | 500 | 22px | 0 | título de card |
| título de seção (sem token) | 18px | 500 | 26px | 0 | ex.: "Visão Geral" na home |
| `DataTable` cabeçalho | 12px | 500 | normal | 0 | cabeçalho de tabela |

- **Números tabulares:** `[—]`
- **Notas:** `common.lineHeight = normal`; `Input.lineHeightTextarea = 22px`. Escala pequena e
  plana (12/13/14/16/18), sem tokens de título nomeados.

## 5. Espaçamento, raios e sombras → `docs/design/espacamento-raios-sombras.md`

- **Fonte:** painel — `themeOverrides` e estilos computados
- **Escala de espaçamento** (obrigatório): `[?]` — não há escala nomeada; os valores recorrentes são:

| Token | Valor |
|---|---|
| (sem token — utilitário) | 2px, 4px, 8px, 10px, 12px, 16px, 20px |
| `Button.paddingMedium` | 0 12px |
| `Button.paddingSmall` | 0 10px |
| `Input.paddingMedium` | 0 8px |
| `DataTable.thPaddingMedium` / `tdPaddingMedium` | 8px / 6px 8px |
| `Alert.padding` | 10px 12px |
| `Message.padding` | 16px 20px 16px 15px |
| `Form.labelPaddingHorizontal` | 0 20px 0 0 |
| espaço entre blocos da página (utilitário `mb-16px`) | 16px |

- **Raios** (obrigatório):

| Token | Valor | Onde se usa |
|---|---|---|
| `common.borderRadius` | 10px | card, painel, diálogo, popover, campo de texto, select, tabela, paginação |
| `Button.borderRadiusMedium` | 6px | botão |
| checkbox (`--n-border-radius`) | 2px | checkbox |
| item de menu lateral | 0 (fundo ativo com cantos ~6px no contêiner) | menu |

- **Espessura de borda padrão:** `1px solid var(--color-border)` (`#dcdfe6` claro · `#434343` escuro)
- **Estratégia de elevação** (obrigatório): fundo diferente (bg-1 página → bg-2 card) + sombra leve
  em sobreposições
- **Sombras:**

| Token | Valor | Para quê |
|---|---|---|
| (painel/card da home) | `0 0 8px 0 rgba(0,0,0,.06)` | cards de conteúdo |
| `--modal-boxshadow` (claro) | `0 6px 16px -9px rgba(0,0,0,.08), 0 9px 28px 0 rgba(0,0,0,.05), 0 12px 48px 16px rgba(0,0,0,.03)` | modal |
| `--modal-boxshadow` / `Modal.boxShadow` (escuro) | `0 0 3px 1px rgba(255,255,255,.4)` | modal no escuro (contorno claro) |
| popover (Naive) | `0 3px 6px -4px rgba(0,0,0,.24), 0 6px 12px 0 rgba(0,0,0,.16), 0 9px 18px 8px rgba(0,0,0,.1)` | popover/dropdown |
| `--select-box-shadow` | `0 0 8px 0 rgba(32,165,58,.4)` (claro) · `0 0 8px 0 rgba(48,53,49,.4)` (escuro) | select em foco |
| input foco (`--n-box-shadow-focus`) | `0 0 8px 0 rgba(32,165,58,.3)` | campo em foco |
| `--bt-tabs-modal-left-shadow` | `2px 0 3px #e4e3e3` · `2px 0 3px #444` | coluna de abas laterais em modal |

- **Notas:**

## 6. Ícones → `docs/design/icones.md`

- **Fonte:** painel — classes `i-*` no DOM e `/static/vite/iconfont/iconfont.css`
- **Biblioteca** (obrigatório): Iconify via UnoCSS preset-icons (classes `i-<coleção>:<nome>`),
  misturando coleções: `lucide`, `mdi`, `ant-design`, `carbon`, `fa`, `ri`, `uiw`, `famicons`,
  `material-symbols-light`, `dashicons`, `mage`, mais coleções próprias `common`, `home`, `system`,
  `theme`. Legado: iconfont (`iconfont.css`). — pacote por stack: `@iconify-json/<coleção>` + UnoCSS `[?]`
- **Estilo** (obrigatório): misto — predominantemente outline (traço fino ~1.5px) no menu lateral e
  no cabeçalho; preenchido em alguns ícones de status
- **Tamanhos** (obrigatório): 14px (`Upload.itemIconSize`), 16px, 18px (`Button` `--n-icon-size`),
  20px (`Alert.iconSize`), 30px (`Message.iconSize`), 52px (`Empty.iconSizeLarge`) — padrão: 18px
  em botão/menu `[?]`
- **Cor:** herda a cor do texto (`currentColor`); ícones de cards da home usam `--color-primary`
- **Ícones com significado fixo:**

| Significado | Ícone |
|---|---|
| buscar | `i-fa:search` |
| recolher menu | `i-ant-design:menu-fold-outlined` |
| abrir dropdown | `i-lucide:chevron-down` / `i-mdi:chevron-down` |
| ajuda | `i-uiw:question-circle-o` |
| mais ações | `i-famicons:ellipsis-vertical` |
| arrastar | `i-material-symbols-light:drag-pan` |
| site | `i-dashicons:admin-site-alt3` |
| banco de dados | `i-carbon:data-base` |
| segurança | `i-carbon:security` |
| idioma / modo escuro | `i-common:lang` / `i-common:mode`, `i-theme:dark` |

- **Notas:** Sem biblioteca única: cada tela escolhe de coleções diferentes.

## 7. Layout → `docs/design/layout.md`

- **Fonte:** painel — media queries em `app.css`/vendor e estilos computados
- **Breakpoints ou faixas de tela** (obrigatório):

| Token | Valor | O que muda |
|---|---|---|
| (sem token) `max-width: 640px` | 640px | ajustes pontuais |
| (sem token) `min-width: 768px` / `max-width: 767px` | 768px | breakpoint mais usado (Bootstrap 3 `sm` + utilitário `md`) |
| (sem token) `min-width: 992px` | 992px | Bootstrap 3 `md` (telas legadas) |
| (sem token) `min-width: 1200px` | 1200px | Bootstrap 3 `lg` |
| (sem token) `min-width: 1280px` | 1280px | utilitário `xl` |
| (sem token) `min-width: 1536px` | 1536px | utilitário `2xl` |

- **Largura máxima do conteúdo:** `[—]` (conteúdo ocupa toda a largura à direita do menu)
- **Grid:** `[—]` — layout em flex; Bootstrap 3 (12 colunas, gutter 30px) só em telas legadas
- **Padrões de navegação que o design system oferece:** menu lateral fixo (`n-layout-sider`,
  200px, itens de 38px com ícone + rótulo, busca `Ctrl+K` no topo, botão de recolher no rodapé);
  barra superior por página (conta, SO, selo PRO, ações rápidas, idioma, modo, versão); abas
  horizontais no topo das páginas (ex.: Projeto PHP / Node.js / Proxy…) — comportamento no menor
  tamanho: menu recolhível para só ícones `[?]`
- **Barra de rolagem:** personalizada — largura `10px` (`::-webkit-scrollbar`) · polegar
  `--scrollbar-thumb-bg-color = #999 | rgba(255,255,255,.2)` · trilho
  `--scrollbar-track-bg-color = #ededed | rgba(255,255,255,0)` · aparece sempre (WebKit); nas áreas
  Naive (`n-scrollbar`) aparece só ao passar o mouse/rolar
- **Notas:** Fundo da página no modo claro usa imagem (`--layout-bg = url(/static/vite/images/bg2.png)`);
  no escuro, `none`.

## 8. Tema → `docs/design/tema.md`

- **Fonte:** painel — `app.css` (`:root` × `:root[theme-mode="dark"]`) e `NConfigProvider`
- **Modos oferecidos** (obrigatório): claro, escuro
- **Como o modo é ativado:** atributo `theme-mode="dark"` + classe `dark` no `<html>`; no Naive UI,
  prop `theme` do `NConfigProvider` (`darkTheme`) com `themeOverrides` do modo
- **Tokens semânticos** (obrigatório — um por papel; uma coluna por modo que o design system
  oferece, acrescente as que faltarem):

| Token | Claro | Escuro | Uso |
|---|---|---|---|
| `--color-bg-1` | `#f2f5f9` | `#18191c` | fundo da página / modal no escuro |
| `--color-bg-2` | `#fff` | `#202020` | superfície (card, tabela, campo) |
| `--color-bg-3` | `#f2f5f9` | `#2e2e2e` | superfície secundária / hover |
| `--color-bg-4` | `#f7f7f7` | `#2e2e2e` | fundo sutil |
| `--color-text-1` | `#131313` | `#d8dce2` | texto de maior ênfase |
| `--color-text-base` / `--color-text-2` | `#3a424d` | `#c7c7c7` | texto primário |
| `--color-text-5` | `#333` | `#e0e0e0` | texto forte |
| `--color-text-4` | `#666` | `#aaa` | texto secundário |
| `--color-text-3` | `#999` | `#919191` | texto terciário |
| `--color-text-desc` | `#999` | `#777` | descrição/ajuda |
| `--color-border` | `#dcdfe6` | `#434343` | borda padrão |
| `--border-hover-focus-color` | `var(--color-primary)` | `var(--color-text-base)` | borda em hover/foco |
| `--color-primary` | `#20a53a` | `#20a53a` | ação primária/marca |
| `--color-success` | `#20a53a` | `#20a53a` | sucesso |
| `--color-warning` | `#ffae45` | `#e67e22` | aviso |
| `--color-error` | `#e73a33` | `#f16575` | erro |
| `--color-pro` | `#ff8f00` | `#feaa04` | recurso Pro/pago |
| `--primary-button-color-hover` | `#1d9534` | `#267544` | hover do botão primário (Naive no escuro: `rgb(29,149,52)`) |
| `--primary-button-color-pressed` | `#1a8a30` | `#20633a` | botão primário pressionado |
| `--color-sider` | `rgba(242,245,249,.6)` | `transparent` | fundo do menu lateral (Naive `Layout.siderColor = #202020`) |
| `--color-sider-text` | `#3a424d` | `#a1a1aa` | texto do menu |
| `--color-sider-active` | `#20a53a` | `#353535` | fundo do item ativo |
| `--color-sider-active-text` | `#fff` | `#fff` | texto do item ativo |
| `--color-sider-hover` | `rgba(32,165,58,.1)` | `#353535` | hover do item |
| `--color-table-th` | `#f6f6f6` | `#232323` | cabeçalho de tabela |
| `--color-table-td-hover` | `#f0f9f7` | `#353535` | linha em hover |
| `--color-modal` | `#fff` | `var(--color-bg-1)` | corpo do modal |
| `--modal-header-bg` | `#f6f8f8` | `var(--color-bg-2)` | cabeçalho/rodapé do modal |
| `--popover-color` | `#fff` | `var(--color-bg-1)` | popover |
| `--color-message` | `#fff` | `#48484e` | toast/mensagem |
| `--alert-default-bg` | `#fafafc` | `#202020` | alerta neutro |
| `--alert-warning-bg` | `#fdf6ec` | `#4a3c1a` | alerta de aviso |
| `--alert-error-bg` | `#fde6e6` | `#4a1f1f` | alerta de erro |
| `--tooltip-color-text` | `#666` | `var(--color-text-base)` | texto do tooltip (fundo Naive escuro `#202020`) |
| `--progress-rail-color` | `#f2f5f9` | `#444` | trilho de progresso |

- **Notas:** Há ainda ~150 tokens por tela (`--home-*`, `--site-*`, `--docker-*`, `--setting-*`…)
  que não entram aqui por serem específicos do produto.

## 9. Componentes → `docs/design/componentes.md`

- **Fonte:** painel — componentes Naive UI encontrados no DOM + `themeOverrides`
- **Catálogo** (obrigatório — os componentes que o design system documenta; o import vem da
  seção 1):

| Componente (nome oficial) | Variantes | Quando usar (pelo guia) |
|---|---|---|
| `Button` (`n-button`) | primary (verde cheio), default (contorno, hover verde 15%), text, desabilitado; tamanhos small 26px / medium 32px | primary: ação principal da barra de ações (ex.: "Adicionar item"); default: ações secundárias `[?]` |
| `Input` (`n-input`) | com ícone de sufixo (busca), textarea | `[—]` |
| `Select` (`n-select`) | — | `[—]` |
| `Checkbox` / `Radio` / `Switch` | switch small (trilho 28×16) | `[—]` |
| `DataTable` (`n-data-table`) | com seleção, ações em lote + paginação no rodapé | listagens |
| `Pagination` | item 28px | rodapé de tabela |
| `Tabs` (`n-tabs`) | line (abas de topo de página), card em modal (`bt-tabs-modal`, com coluna lateral) | `[—]` |
| `Card` (`n-card`) | — | blocos de conteúdo |
| `Dialog` / `Modal` (`BtDialogConfig`, layui-layer legado) | confirmação, formulário | `[—]` |
| `Drawer` | — | `[—]` |
| `Message` | success, warning, error | retorno de ação |
| `Alert` | default, success, warning, error | avisos em linha |
| `Tag`, `Badge`, `Avatar` | — | status/contadores |
| `Popover`, `Tooltip`, `Dropdown` | — | ajuda e menus contextuais |
| `Progress` | circular (uso de CPU/RAM/disco), linear | métricas |
| `Collapse`, `Divider`, `Ellipsis`, `Empty`, `Spin`, `Scrollbar`, `Menu`, `Layout` | — | `[—]` |
| `Upload`, `DatePicker`, `TimePicker`, `Cascader`, `Breadcrumb`, `Log`, `Code` | — | `[—]` |

- **Estados que todo interativo tem** (obrigatório): padrão · hover · foco · ativo ·
  desabilitado · carregando · erro · selecionado — marque os que o design system define:
  padrão ✔ · hover ✔ · foco ✔ · ativo/pressionado ✔ · desabilitado ✔ (opacidade 0.38 no botão) ·
  carregando ✔ (Naive `loading`/`Spin`) · erro ✔ (campos) · selecionado ✔ (menu, abas, tabela)
- **Notas:** Botão tem efeito "wave" do Naive (`--n-ripple-duration = .6s`, cor do ripple = primária).

### Campos de formulário → `docs/design/formularios.md`

- **Fonte:** painel — `themeOverrides.Input`, `.Select`, `.Form` e estilo computado de `.n-input`
- **Tamanhos** (obrigatório):

| Tamanho (token) | Altura | Padding horizontal | Tamanho do texto |
|---|---|---|---|
| `small` | 26px `[?]` (herda `Button.heightSmall`) | 8px | 12px |
| `medium` (`common.heightMedium`) | 32px | 8px (`Input.paddingMedium = 0 8px`) | 12px |

- **Estados** (obrigatório — token ou valor em cada célula; `[—]` onde o estado não muda aquilo):

| Estado | Borda | Fundo | Texto | Outro (sombra, contorno, ícone) |
|---|---|---|---|---|
| padrão | `1px solid --color-border` (`#dcdfe6` / `#434343`) | `--color-bg-2` (`#fff` / `#202020`) | `--input-text-color` (`#333` / `#c7c7c7`) | raio 10px |
| hover | `1px solid #20a53a` | `[—]` | `[—]` | `[—]` |
| foco | `1px solid #20a53a` | `--input-focus-bg-color` (= bg-2) | `[—]` | `0 0 8px 0 rgba(32,165,58,.3)`; cursor `#20a53a` |
| erro | `1px solid --color-error` (`#e73a33` / `#f16575`) | `[—]` | `[—]` | foco: `0 0 8px 0 rgba(241,101,117,.3)` |
| desabilitado | `--input-disabled-border-color` (`1px solid #d9d9d9` / `1px solid #434343`) | `--input-disabled-color` (`#f5f5f5` / `#18191c`) | `#b8b8b8` (escuro) | `[—]` |
| somente leitura | `[?]` | `[?]` | `[?]` | `[?]` |

- **Anatomia:** rótulo ao lado (à esquerda, 12px, `labelPaddingHorizontal = 0 20px 0 0`); acima em
  alguns formulários (`labelFontSizeTopMedium = 12px`) · obrigatório marcado `[?]` · texto de ajuda
  abaixo do campo `[?]` · mensagem de erro abaixo do campo (feedback 12px, altura 20px, padding
  `2px 0 0`), sem ícone · contador de caracteres `[?]`
- **Placeholder:** cor `#999` claro `[?]` / `rgba(255,255,255,.38)` escuro (select: `#999`) · regra de
  uso: `[—]`
- **Quando o erro aparece:** `[?]` (Naive UI valida por regra — `blur`/`input` — configurado por formulário)
- **Notas:** `--input-group-label-color = #f5f5f5 | #333` para rótulos agrupados (addon).

## 10. Padrões de feedback → `docs/design/padroes-feedback.md`

- **Fonte:** painel — comportamento observado e `themeOverrides` (Message, Empty, Spin, Alert, Dialog)

| Situação | Padrão | Componente |
|---|---|---|
| carregando uma ação (botão) | spinner dentro do botão, botão bloqueado `[?]` | `Button loading` / `Spin` (small 16px) |
| carregando lista ou tabela | spinner sobre a tabela `[?]` | `DataTable loading` / `Spin` |
| erro de requisição | mensagem flutuante no topo `[?]` | `Message` (error) |
| erro de validação de campo | borda vermelha + texto 12px abaixo do campo | `Form` feedback |
| lista vazia | ícone 52px + "Não há dados" centralizado na tabela | `Empty` |
| sucesso de uma ação | mensagem flutuante (ícone 30px, texto 14px) | `Message` (success) |
| confirmação de ação destrutiva | diálogo de confirmação; algumas exigem digitar texto/cálculo (`--confirm-calc-bg`, `--data-base-del-input-bg`) | `Dialog` / modal `bt` |

- **Notas:** "carregando a página" não tem padrão identificado (removido).

## 11. Acessibilidade visual → `docs/design/acessibilidade-visual.md`

- **Fonte:** painel — `app.css`
- **Nível alvo** (obrigatório): `[?]` (não declarado)
- **Contraste mínimo:** `[?]` — nota: `#999` (`--color-text-3`) sobre `#fff` ≈ 2.8:1 (abaixo de AA)
- **Foco visível** (obrigatório): borda `1px solid #20a53a` + sombra `0 0 8px rgba(32,165,58,.3)` em
  campos; botão: borda `1px solid rgb(29,149,52)`; sem anel de foco dedicado (offset `[—]`)
- **Área mínima de toque ou clique** (obrigatório): `[?]` — menor controle observado: checkbox 16px;
  botão small 26px; medium 32px
- **Movimento reduzido:** há 2 regras `@media (prefers-reduced-motion: reduce)` (conteúdo `[?]`)
- **Notas:**

## 12. Movimento e transições → `docs/design/movimento-transicoes.md`

- **Fonte:** painel — `--n-bezier` dos componentes Naive e regras `transition` mais frequentes
- **Durações:** sem tokens; recorrentes: `.15s`, `.2s`, `.25s`, `.3s`, `500ms` (mais usada em CSS
  próprio), ripple do botão `.6s`
- **Easing:** `--n-bezier = cubic-bezier(.4, 0, .2, 1)` (padrão Naive); `ease` em CSS próprio
- **Notas:**

## 13. Camadas e sobreposições → `docs/design/camadas-sobreposicoes.md`

- **Fonte:** painel — z-index computado de popover/modal Naive

| Camada | z-index | Componentes |
|---|---|---|
| sobreposições Naive (popover, modal, dropdown) | 2000+ (incremental; observado 2001) | `Popover`, `Modal`, `Dialog`, `Dropdown`, `Tooltip` |
| layer legado | `[?]` (layui-layer gerencia o próprio) | `layui-layer` |

- **Modal:** fundo escurecido `[?]` (padrão Naive) · fecha com Esc ✔ (Naive) · fecha com clique fora
  `[?]` · prende o foco `[?]`
- **Notas:** Modal: raio 10px; cabeçalho `--modal-header-bg` com borda `--modal-header-bottom-border`;
  título 14px; rodapé de ações `--modal-action-bg` com borda superior `--modal-action-top-border`.

## 14. Lacunas

- **O que o design system não define** e cada projeto terá de decidir: documentação pública e
  licença de reuso; escala de espaçamento nomeada; escala tipográfica nomeada; paleta de dados para
  gráficos; nível de contraste alvo; área mínima de toque; tokens de duração/easing; z-index
  nomeados; guia de microcopy; biblioteca de ícones única; comportamento mobile.
- **Qualquer coisa que não cabia nos campos:** O painel convive com três gerações de UI (Bootstrap
  3 + layui, CSS próprio com variáveis, Naive UI + utilitários). Há um seletor "Classic UI / New UI"
  no cabeçalho; esta captura é da New UI. Tokens por tela (`--home-*`, `--site-*`, `--docker-*`,
  `--setting-*`, `--pay-*`) foram omitidos por serem do produto.
