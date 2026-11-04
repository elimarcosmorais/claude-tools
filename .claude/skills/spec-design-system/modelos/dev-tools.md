# Modelo de design system — Dev Tools

> Retrato de um design system que já existe (ex.: o Primer, do GitHub), guardado na
> `/spec-design-system` para ser aplicado em qualquer projeto com
> `/spec-design-system modelo <slug>`. Descreve **o design system**, não um produto: nome,
> público e plataformas do projeto são perguntados quando o modelo é aplicado.

> **Sobre as fontes deste preenchimento:** o design system do Dev Tools não tem site público; a
> fonte oficial é o próprio repositório (`docs/design-system.md`, `docs/design/*.md` e o código em
> `frontend/`). Cada **Fonte** abaixo é um caminho relativo à raiz do repo, no commit `737e08a`.

**Notação:** `[?]` = a fonte não diz (ao aplicar, a skill decide pelo default consagrado ou
pergunta) · `[—]` = o design system não define isso.

---

## 1. Identificação → `docs/design-system.md`

- **Nome** (obrigatório): Dev Tools
- **Slug** (obrigatório): `dev-tools`
- **Mantido por** (obrigatório): Elimarcos Morais (projeto pessoal, repositório `dev-tools`)
- **Link oficial** (obrigatório): `docs/design-system.md` (no repositório; sem página pública)
- **Versão capturada** (obrigatório): app 0.2.0 (`build/config.yml`), commit `737e08a`
- **Data da captura** (obrigatório): 2026-10-30
- **Licença** (obrigatório): `[?]` — o repo não tem arquivo `LICENSE`. Partes de terceiros:
  shadcn/ui, Radix, Tailwind CSS, sonner e tw-animate-css em MIT; lucide-react em ISC; Simple
  Icons em CC0-1.0 (as marcas continuam dos donos); Segoe UI Variable, Segoe UI, Cascadia Mono e
  Consolas são fontes do Windows, não redistribuídas. Uso fora da organização dona permitido: `[?]`
- **Plataformas cobertas** (obrigatório): desktop (Windows, janela WebView2 do Wails v3, sem
  moldura)
- **Pacote de tokens:** `frontend/src/styles/tokens.css` — formato variáveis CSS (`--<token>`, valor
  claro em `:root`, escuro em `.dark`), registradas como utilitários do Tailwind v4 por
  `@theme inline` (`--color-<token>`)
- **Implementações oficiais** (obrigatório — uma linha por stack; nenhuma → só a linha `nenhuma`):

| Stack | Pacote | Versão | Como importar |
|---|---|---|---|
| React 19 + Tailwind CSS v4 | componentes shadcn/ui copiados para `frontend/src/components/ui` (estilo `radix-vega`, CLI `shadcn`) | shadcn 4.21.4 · radix-ui 1.7.0 · tailwindcss 4.3.3 | `import { Button } from "@/components/ui/button"` |
| React (toasts) | `sonner`, embrulhado em `@/components/ui/sonner` | 2.0.8 | `import { toast } from "@/components/ui/sonner"` (o ESLint reprova importar do `sonner` direto) |
| React (ícones) | `lucide-react` · `simple-icons` | 1.53.0 · 16.34.0 | `import { Power } from "lucide-react"` |
| React (editor) | CodeMirror 6, embrulhado em `@/components/editor` | 6.x | `import { CodeEditor } from "@/components/editor/CodeEditor"` |

- **Notas:** os componentes não vêm de um pacote: o código do shadcn mora no repo e foi ajustado
  aos tokens (ADR 001). Ao trazer um componente novo com `npx shadcn add`, os variantes
  `data-open:`/`data-checked:` viram `data-[state=open]:`/`data-[state=checked]:`. O nome do token
  vira a variável `--<token>` e o utilitário `bg-<token>`/`text-<token>`; `text-primary` e
  `text-secondary` não são utilitários: usam-se `text-foreground` e `text-muted-foreground`.

## 2. Direção visual → `docs/design-system.md` (seção Direção)

- **Fonte:** `docs/design-system.md` (Direção), `docs/design/layout.md`
- **Personalidade em 3 adjetivos** (obrigatório): sóbrio, técnico, nativo (de app desktop, no
  espírito do FlyEnv)
- **Eixos** (obrigatório) — marque um ponto de 1 a 5 em cada:
  - sóbrio `1 2 3 4 5` expressivo: 1
  - denso `1 2 3 4 5` espaçado: 3
  - geométrico `1 2 3 4 5` orgânico: 2
  - plano `1 2 3 4 5` com profundidade: 1
- **Densidade de informação:** confortável — itens de 44 px, texto base de 16 px; não muda com o
  tamanho da janela (só a sidebar recolhe abaixo de 1150 px)
- **O que o design system evita** (obrigatório): sombra em card e superfície fixa (separação só
  por borda); cor solta fora dos tokens (o lint reprova hexadecimal fora de `tokens.css`);
  gradiente (só no logo); barra de rolagem padrão da web; rolagem horizontal (texto longo quebra
  a linha); uma cor com dois papéis; toast que espera ser fechado (sem botão fechar)
- **Tom do microcopy:** `[—]` (todo texto vem de `lang/pt-br.json` e `lang/en.json`, mas não há
  guia de conteúdo)
- **Notas:** o tema escuro é o principal e a referência visual; o claro deriva dele mantendo a
  hierarquia. Cada cor tem um papel só: âmbar (`highlight`) = foco pelo teclado; verde
  (`running`) = serviço rodando; azul (`primary`, e `link` em texto e ícone) = ação principal,
  aba ativa, navegação ativa dos botões do topo, link e campo em foco; vermelho, amarelo, verde e
  azul semânticos = feedback.

## 3. Cores → `docs/design/cores.md`

- **Fonte:** `docs/design/cores.md`, `frontend/src/styles/tokens.css`
- **Modelo da paleta** (obrigatório): cores chapadas — não há escala de tons; cada token
  funcional tem um valor por modo. Os semânticos e os avatares vêm da escala do Tailwind (tom 400
  no escuro, 700 no claro).
- **Convenção de nome dos tokens** (obrigatório): só tokens funcionais, em kebab-case pelo papel
  (`<superfície>-bg`, `text-<nível>`, `<estado>`), como variável CSS `--<token>`; no Tailwind,
  `--color-<token>`. As variáveis do shadcn (`--background`, `--primary`…) apontam para eles.
- **Paleta base** (obrigatório — uma linha por cor, no modo claro):

| Cor | Tons (token = valor) |
|---|---|
| neutros (superfícies) | `titlebar-bg = #E4E7EF` · `sidebar-bg = #EEF1F7` · `content-bg = #F7F8FB` · `card-bg = #FFFFFF` · `item-selected-bg = #DCE2EE` · `field-group-bg = #EEF1F7` · `icon-button-active-bg = #DCE0EA` · `tooltip-bg = #FFFFFF` · `log-bg = #FAFAFA` · `mail-frame-bg = #FFFFFF` |
| neutros (borda e controle) | `border = #C9CEDA` · `tooltip-border = #C9CEDA` · `toggle-off = #9CA3B5` · `toggle-knob = #FFFFFF` |
| texto | `text-primary = #1A2138` · `text-secondary = #5B6278` |
| azul (ação) | `primary = #1F6FD1` · `primary-foreground = #FFFFFF` · `link = #1F6FD1` |
| âmbar (foco) | `highlight = #B26A00` |
| verde (rodando) | `running = #007A45` |
| vermelho (fechar janela) | `titlebar-close-hover = #C42B1C` |
| semânticas | `success = #15803D` · `warning = #B45309` · `error = #B91C1C` · `info = #1D4ED8` |
| sobreposição | `overlay = #0F1428` a 35% (`#0F142859`) |
| avatares (iguais nos dois modos) | `avatar-1 = #1D4ED8` · `avatar-2 = #0F766E` · `avatar-3 = #6D28D9` · `avatar-4 = #A16207` · `avatar-5 = #BE123C` · `avatar-6 = #15803D` |
| realce de log (apelidos) | `log-timestamp = error` · `log-number = success` · `log-text = text-primary` |

- **Primária ou acento:** `primary = #1F6FD1` (mesmo valor nos dois modos) — fundo do botão
  principal e da aba ativa, borda do campo em foco, barra do `Progress`, trilho do `Switch`
  `primary`. Em texto e ícone sobre superfície usa-se `link` (`#409EFF` no escuro). O acento de
  foco é `highlight`.
- **Cores de dados (gráficos):** `[—]`
- **Contraste alvo** (obrigatório): WCAG AA
- **Regras de uso:** nenhuma cor solta, só tokens (o `go tool task lint` reprova hexadecimal fora
  de `tokens.css`); cada cor tem um papel só; `primary` não vai em texto sobre superfície escura
  (fica abaixo de 3:1), vai `link`; avatar novo só entra se mantiver o AA com o texto branco;
  `highlight` não se chama `accent` porque o shadcn já usa `--accent` para fundo de hover
  discreto.
- **Notas:** os valores do escuro foram amostrados da tela de referência; o claro deriva deles. Os
  valores do escuro estão na seção 8.

## 4. Tipografia → `docs/design/tipografia.md`

- **Fonte:** `docs/design/tipografia.md`
- **Família de texto** (obrigatório): `Segoe UI Variable` — fonte do sistema (Windows 11; no 10,
  `Segoe UI`) — licença da fonte: Microsoft, do Windows, não empacotada
- **Família de título:** mesma do texto
- **Família monoespaçada:** `Cascadia Mono` (Windows 11), com `Consolas` no Windows 10 — fontes do
  sistema
- **Fallback stack** (obrigatório): texto `"Segoe UI Variable", "Segoe UI", sans-serif` ·
  mono `"Cascadia Mono", Consolas, monospace`
- **Unidade do tamanho de texto** (obrigatório): px (pela escala do Tailwind, em rem com base 16 px)
- **Estilos nomeados** (obrigatório):

| Estilo (token oficial) | Tamanho | Peso | Altura de linha | Espaçamento entre letras | Uso |
|---|---|---|---|---|---|
| `page-title` (`text-2xl font-semibold`) | 24px | 600 | 32px | `[—]` | título do cabeçalho de cada tela |
| `section-title` (`text-lg font-semibold`) | 18px | 600 | 28px | `[—]` | título de bloco dentro do card; título de painel e diálogo |
| `body` (`text-base`) | 16px | 400 | 24px | `[—]` | texto padrão, itens da sidebar, valores de campo |
| `small` (`text-sm`) | 14px | 400 | 20px | `[—]` | título dos grupos da sidebar, tooltip, texto auxiliar |
| `log` (`font-mono text-sm leading-6`) | 14px | 400 | 24px | `[—]` | visualizador de log e editor de código |

- **Números tabulares:** `[—]`
- **Notas:** os estilos são nomes do design system, não tokens no CSS: na tela usam-se as classes
  do Tailwind da coluna Estilo. O `Button` usa `text-sm font-medium` (14 px / 500), e o `Input` e
  o `Textarea` caem para `text-sm` a partir de 768 px de largura (`md:text-sm`), ou seja, 14 px em
  toda janela do app.

## 5. Espaçamento, raios e sombras → `docs/design/espacamento-raios-sombras.md`

- **Fonte:** `docs/design/espacamento-raios-sombras.md`, `frontend/src/styles/tokens.css`
- **Escala de espaçamento** (obrigatório):

| Token | Valor |
|---|---|
| `1` (Tailwind) | 4px |
| `2` | 8px |
| `3` | 12px |
| `4` | 16px |
| `6` | 24px |
| `8` | 32px |
| `12` | 48px |

- **Raios** (obrigatório):

| Token | Valor | Onde se usa |
|---|---|---|
| `radius-sm` | 4px | card, select, botão, campo, menu, diálogo, toast, abas |
| `radius-md` | 6px | tooltip |
| `radius-lg` | 8px | item selecionado da sidebar, bloco interno de um card |
| `radius-full` | 9999px | botão de ícone (círculo), toggle, polegar da barra de rolagem |

- **Espessura de borda padrão:** 1px, cor `border`
- **Estratégia de elevação** (obrigatório): borda em vez de sombra (sombra só em sobreposições)
- **Sombras:**

| Token | Valor | Para quê |
|---|---|---|
| sem sombra | — | card e superfícies fixas: separação só pela `border` |
| `shadow-xs` (Tailwind) | padrão do Tailwind | botão `outline`, `Input`, `SelectTrigger`, `Textarea` (herdado do shadcn) |
| `shadow-md` (Tailwind) | padrão do Tailwind | sobreposições: lista do select, menu, tooltip, painel lateral, diálogo, painel da bandeja |

- **Notas:** escala de espaçamento base 4 (a do Tailwind); valor fora dela não entra. Os botões
  `xs`/`sm`/`icon-xs`/`icon-sm` do shadcn usam `rounded-[min(var(--radius-md),8–10px)]`, ou seja,
  6 px.

## 6. Ícones → `docs/design/icones.md`

- **Fonte:** `docs/design/icones.md`
- **Biblioteca** (obrigatório): `lucide` para os genéricos — pacote React `lucide-react`; `Simple
  Icons` para as marcas de tecnologia — pacote `simple-icons`; marca que não existe no Simple
  Icons vira SVG próprio no repo, monocromático
- **Estilo** (obrigatório): outline — espessura do traço: a padrão do lucide (2px) `[?]`
- **Tamanhos** (obrigatório): 16px (dentro de botão e select, padrão do shadcn), 20px (item da
  sidebar, botão de ícone, linha da bandeja), 26px (cabeçalho de tela), 48px (estado vazio) —
  padrão: 20px
- **Cor:** herda a cor do texto (`currentColor`); muda só pelos tokens de estado (`link`,
  `running`, semânticos)
- **Ícones com significado fixo:**

| Significado | Ícone |
|---|---|
| iniciar/parar todos os serviços | `Power` |
| log do aplicativo | `Clock` |
| configurações | `Settings` |
| expandir/recolher a sidebar | `Menu` |
| projetos web | `Globe` |
| recarregar | `RefreshCw` |
| abrir pasta no Explorer | `FolderOpen` |
| fechar painel ou diálogo | `X` |
| abrir select | `ChevronDown` |
| ação em andamento | `LoaderCircle` girando |
| operação que pede administrador | `ShieldAlert` |
| aviso que não bloqueia | `CircleAlert` |
| versão nova disponível | `CircleArrowUp` |
| sair | `LogOut` |

- **Notas:** cada tecnologia (Apache, Go, Java, NodeJS, PHP, Python, MySQL, PostgreSQL, Redis,
  Android, Windows) e cada conector (GitHub) usa o próprio ícone de marca. O logo do app
  (quadrado com gradiente de `#3BB2FC` a `#3420FC` e `</>` branco) é do produto, não do design
  system: as cores dele não entram na interface.

## 7. Layout → `docs/design/layout.md`

- **Fonte:** `docs/design/layout.md`
- **Breakpoints ou faixas de tela** (obrigatório):

| Token | Valor | O que muda |
|---|---|---|
| janela mínima | 1024×640 | a janela não encolhe abaixo disso |
| janela inicial | 1170×737 (pixels lógicos) | tamanho com que a janela abre |
| estreita | largura < 1150px | sidebar recolhida, botão expandir/recolher desabilitado |
| larga | largura ≥ 1150px | sidebar inicia expandida; o botão volta a funcionar |

- **Largura máxima do conteúdo:** `[—]` (o card ocupa toda a área de conteúdo, com 16 px de margem)
- **Grid:** `[—]`
- **Padrões de navegação que o design system oferece:** barra de título própria (32 px,
  `titlebar-bg`, botões da janela de 46×32 px); sidebar à esquerda (304 px expandida, 64 px
  recolhida, itens de 44 px, grupos com título `small` e divisor, botões de ícone redondos no
  topo); cabeçalho de tela de 64 px (ícone de 26 px + `page-title`) com abas à direita
  (`PageTabs`, que navegam entre rotas, ou `Tabs`, que trocam painéis); card único que rola;
  painel lateral pela direita (75% da janela, mínimo 560 px; variante estreita 37,5%/280 px;
  larga 95%) — comportamento no menor tamanho: abaixo de 1150 px a sidebar fica recolhida, só
  ícones com tooltip à direita
- **Barra de rolagem:** personalizada — largura `6px`, arredondada, sem setas nem trilho · polegar
  `toggle-off`, `text-secondary` no hover do polegar (os dois por modo, seção 8); trilho
  transparente · aparece só com o mouse sobre a área (utilitário `scrollbar-hover`, que reserva o
  espaço com `scrollbar-gutter: stable` para o conteúdo não se mexer)
- **Notas:** só rolagem vertical: texto longo quebra a linha; a horizontal fica só onde a quebra
  não cabe (tabela larga). Não se usa `scrollbar-width` nem `scrollbar-color`: com eles o
  Chromium ignora os `::-webkit-scrollbar` e volta à barra padrão. Na sidebar, só os grupos
  rolam, por baixo dos botões do topo.

## 8. Tema → `docs/design/tema.md`

- **Fonte:** `docs/design/tema.md`, `docs/design/cores.md`, `frontend/src/styles/tokens.css`
- **Modos oferecidos** (obrigatório): escuro (principal), claro e sistema (segue o Windows)
- **Como o modo é ativado:** classe `dark` no elemento raiz (convenção do shadcn), posta pelo
  `ThemeProvider`, que também troca o `color-scheme`; em Sistema, segue `prefers-color-scheme`.
  Sem recarregar a página.
- **Tokens semânticos** (obrigatório — um por papel; uma coluna por modo que o design system
  oferece, acrescente as que faltarem):

| Token | Claro | Escuro | Uso |
|---|---|---|---|
| `titlebar-bg` | `#E4E7EF` | `#000000` | barra de título da janela |
| `titlebar-close-hover` | `#C42B1C` | `#C42B1C` | fundo do botão fechar da janela no hover |
| `sidebar-bg` | `#EEF1F7` | `#222940` | fundo da sidebar |
| `content-bg` | `#F7F8FB` | `#192037` | fundo da área de conteúdo |
| `card-bg` | `#FFFFFF` | `#222940` | card das telas, campo, menu, diálogo |
| `border` | `#C9CEDA` | `#464956` | borda do card, divisor, borda de campo |
| `item-selected-bg` | `#DCE2EE` | `#131931` | item selecionado da sidebar; `--muted`/`--secondary` do shadcn |
| `field-group-bg` | `#EEF1F7` | `#131931` | bloco que agrupa os campos de um formulário |
| `text-primary` | `#1A2138` | `#FCFAF8` | texto e ícone primários |
| `text-secondary` | `#5B6278` | `#A9AEC1` | texto auxiliar, estado vazio, placeholder |
| `icon-button-active-bg` | `#DCE0EA` | `#484B5A` | círculo do botão de ícone no hover/ativo; `--accent` do shadcn |
| `tooltip-bg` | `#FFFFFF` | `#1D1E1F` | fundo do tooltip |
| `tooltip-border` | `#C9CEDA` | `#414243` | borda do tooltip |
| `log-bg` | `#FAFAFA` | `#1B1B1B` | fundo do visualizador de log e do editor |
| `overlay` | `#0F1428` a 35% | `#000000` a 50% | fundo atrás de painel lateral e diálogo |
| `highlight` | `#B26A00` | `#E69F26` | anel de foco (`--ring`); realce de seção e chave no editor |
| `primary` | `#1F6FD1` | `#1F6FD1` | botão principal, aba ativa, campo em foco |
| `primary-foreground` | `#FFFFFF` | `#FFFFFF` | texto e ícone sobre `primary` |
| `link` | `#1F6FD1` | `#409EFF` | link, valor clicável, ícone ativo do topo |
| `running` | `#007A45` | `#01CC74` | serviço rodando: trilho ligado, ícone do item |
| `toggle-off` | `#9CA3B5` | `#505158` | trilho desligado; polegar da barra de rolagem |
| `toggle-knob` | `#FFFFFF` | `#FFFFFF` | bolinha do toggle |
| `success` | `#15803D` | `#4ADE80` | feedback de sucesso |
| `warning` | `#B45309` | `#FBBF24` | aviso |
| `error` | `#B91C1C` | `#F87171` | erro; ação destrutiva (`--destructive`) |
| `info` | `#1D4ED8` | `#60A5FA` | informação |
| `mail-frame-bg` | `#FFFFFF` | `#FFFFFF` | fundo de conteúdo HTML externo (como o destinatário o veria) |

- **Notas:** ligação com o shadcn (as variáveis dele apontam para os tokens, nunca para valor
  solto): `--background` = `content-bg`; `--foreground`, `--card-foreground`,
  `--popover-foreground`, `--secondary-foreground`, `--accent-foreground`, `--sidebar-foreground`
  e `--sidebar-accent-foreground` = `text-primary`; `--card` e `--popover` = `card-bg`;
  `--primary`/`--primary-foreground` = `primary`/`primary-foreground`; `--secondary`, `--muted` e
  `--sidebar-accent` = `item-selected-bg`; `--muted-foreground` = `text-secondary`; `--accent` =
  `icon-button-active-bg`; `--destructive` = `error`; `--border`, `--input` e `--sidebar-border`
  = `border`; `--ring` e `--sidebar-ring` = `highlight`; `--sidebar` = `sidebar-bg`;
  `--sidebar-primary` = `primary`. A barra de título acompanha o tema.

## 9. Componentes → `docs/design/componentes.md`

- **Fonte:** `docs/design/componentes.md`, `frontend/src/components/ui/*.tsx`
- **Catálogo** (obrigatório — os componentes que o design system documenta; o import vem da
  seção 1):

| Componente (nome oficial) | Variantes | Quando usar (pelo guia) |
|---|---|---|
| `Button` | `default` (`primary`), `outline`, `ghost`, `destructive` (também `secondary` e `link` no código); tamanhos `xs` 24px, `sm` 32px, `default` 36px, `lg` 40px, `icon-xs` 24px, `icon-sm` 32px, `icon` 36px, `icon-lg` 40px | ação; `icon` + `rounded-full` nos botões de ícone do topo; `ghost` `icon-sm` com `aria-label` nos botões de ícone de barra e de linha |
| `Tooltip` | — | nome de botão de ícone e de item com a sidebar recolhida; abre à direita |
| `Select` | tamanhos `default` 36px, `sm` 32px | escolha única numa lista fechada; o gatilho tem o visual do `Input` |
| `Switch` | `running` (padrão), `primary` | `running`: ligar/parar um serviço; `primary`: preferência. Desligado sempre `toggle-off` |
| `Checkbox` | — | opção que entra ou sai de um conjunto; marcada em `primary`; sempre com rótulo |
| `Input` | — | texto de uma linha; sempre com `aria-label` ou rótulo |
| `Textarea` | — | texto de várias linhas, ou valor do servidor que pode trazer quebra de linha; cresce com o texto |
| `Skeleton` | — | conteúdo carregando (depois de 300 ms) |
| `Table` | — | dados em linhas e colunas; hover em `muted` a 50% |
| `Progress` | — | andamento com porcentagem; trilho `muted`, barra `primary`; texto da fase ao lado |
| `DropdownMenu` | item `default`, `destructive` | menu de ações de uma linha (⋮); "Excluir" é `destructive` e o último |
| `AlertDialog` | ações com as variantes do `Button` | confirmar ação destrutiva ou operação que pede administrador; Cancelar com o foco |
| `Dialog` | — | formulário curto no meio da janela, com Cancelar `outline` e Confirmar `submit` |
| `Sheet` | padrão (75%, mín. 560px), estreita (37,5%, mín. 280px), larga (95%) | painel lateral pela direita, modal, sobre toda a interface |
| `Tabs` | — | alternar painéis da mesma página; ativa em `primary` |
| `PageTabs` | — | abas que navegam entre rotas, no cabeçalho da tela; mesmo visual do `Tabs` |
| `Toaster` (sonner) | `success`, `info`, `warning`, `error` | resultado de ação sem efeito visível; embaixo à direita, sem botão fechar, com barra de tempo |
| `ToolbarButton` | `active` | botão de ícone da barra de baixo de um painel, em grupo com borda |
| `EmptyState` | — | estado vazio dentro de um card: ícone de 48px, título, descrição e ação opcional |
| `ContentLayout` | `actions`, `tabs` | cabeçalho (ícone + título) e card de toda tela |
| `ComingSoon` | — | tooltip "Em breve" num controle ainda sem ação |
| Faixa de aviso (`ConfigWarningNotice`) | — | aviso que não bloqueia, no alto do card, com `CircleAlert` em `warning` |
| `CodeEditor` | `language` `httpd`, `ini`, `hosts`, `redis`, `yaml`, `json`, `apache-error`, `app-log`, `plain`; `minimap` | editar arquivo de texto (CodeMirror 6, fonte `log`, fundo `log-bg`, sempre quebra a linha) |
| `LogViewer` / `LogView` | `highlight` `apache-error`, `app-log`, `plain`; `follow`, `subscribe` | ler um log, só leitura; o `LogView` traz a barra Abrir/Atualizar/Limpar |

- **Estados que todo interativo tem** (obrigatório): padrão · hover · foco · ativo ·
  desabilitado · carregando · erro · selecionado — marque os que o design system define:
  padrão ✓ · hover ✓ · foco ✓ (anel `highlight`; campo só pela borda `primary`) · ativo ✓ (botão
  desce 1px; aba e navegação ativas em `primary`/`link`) · desabilitado ✓ (opacidade 50%, sem
  ponteiro) · carregando ✓ (`LoaderCircle` girando no lugar do ícone; spinner na bolinha do
  toggle) · erro ✓ (`aria-invalid`: borda e anel `destructive`) · selecionado ✓
  (`item-selected-bg`)
- **Notas:** botão, link e item de menu mostram o cursor `pointer` (regra no `@layer base`, porque
  o Tailwind v4 o tirou do botão); desabilitado não. Menu, `Dialog`, `AlertDialog` e `Sheet` só
  devolvem o foco ao gatilho quando o usuário navega pelo teclado (`lib/focusReturn.ts`).
  Componente reutilizável novo entra no catálogo antes de ser usado por uma segunda feature.

### Campos de formulário → `docs/design/formularios.md`

- **Fonte:** `frontend/src/components/ui/{input,select,textarea}.tsx`, `docs/design/componentes.md`,
  `docs/design/acessibilidade-visual.md`
- **Tamanhos** (obrigatório):

| Tamanho (token) | Altura | Padding horizontal | Tamanho do texto |
|---|---|---|---|
| `Input` (único) | 36px (`h-9`) | 10px (`px-2.5`) | 14px (`md:text-sm`; 16px abaixo de 768px) |
| `SelectTrigger` `default` | 36px | 10px à esquerda, 8px à direita | 16px (`text-base`) |
| `SelectTrigger` `sm` | 32px | 10px à esquerda, 8px à direita | 16px |
| `Textarea` | mínimo 64px (`min-h-16`; `min-h-9 py-1.5` para a altura de um `Input`), cresce com o texto | 10px | 14px |

- **Estados** (obrigatório — token ou valor em cada célula; `[—]` onde o estado não muda aquilo):

| Estado | Borda | Fundo | Texto | Outro (sombra, contorno, ícone) |
|---|---|---|---|---|
| padrão | `border` 1px | `card-bg` (no escuro, `input` a 30%) | `text-primary` | `shadow-xs` |
| hover | `[—]` | `[—]` | `[—]` | `[—]` |
| foco | `primary` 1px | `[—]` | `[—]` | sem anel (exceção ao foco `highlight`) |
| erro | `error` (`destructive`; a 50% no escuro) | `[—]` | `[—]` | anel de 3px em `error` a 20% (40% no escuro); mensagem abaixo |
| desabilitado | `[—]` | `[—]` | `[—]` | opacidade 50%, cursor `not-allowed` |
| somente leitura | `[?]` | `[?]` | `[?]` | `[?]` (só o `CodeEditor` define `readOnly`: mantém foco, seleção e cópia) |

- **Anatomia:** rótulo `[?]` (o guia exige só `aria-label` ou `<label htmlFor>`; nas telas, o
  rótulo visível fica acima, em `text-sm`) · obrigatório marcado `[—]` · texto de ajuda `[—]` ·
  mensagem de erro `[?]` (nas telas, abaixo do campo, `text-sm` em `error`, `role="alert"` ligado
  por `aria-describedby`, sem ícone) · contador de caracteres não
- **Placeholder:** cor `text-secondary` (`text-muted-foreground`) · regra de uso `[—]`
- **Quando o erro aparece:** `[?]` (nas telas, ao enviar, e some ao editar o campo)
- **Notas:** campos de um formulário ficam num bloco `field-group-bg`, `radius-sm`, `p-4`, com
  16px entre campos. Valor lido do servidor que pode ter quebra de linha vai no `Textarea`: o
  `Input` a tira.

## 10. Padrões de feedback → `docs/design/padroes-feedback.md`

- **Fonte:** `docs/design/padroes-feedback.md`, `docs/design/componentes.md`

| Situação | Padrão | Componente |
|---|---|---|
| carregando a página | `Skeleton` no lugar do conteúdo do card, só se a carga passar de 300 ms | `Skeleton` |
| carregando uma ação (botão) | o controle fica desabilitado e o ícone dá lugar ao `LoaderCircle` girando; no toggle, trilho desabilitado e spinner na bolinha | `Button`, `Switch` |
| erro de requisição | ao carregar: dentro do card, ícone, mensagem e botão "Tentar novamente"; numa ação: toast de erro embaixo à direita, some após 8 s (pausa com o mouse em cima), com o botão "Ver log" que abre o log do aplicativo | `EmptyState` + `Button`; `Toaster` |
| lista vazia | centralizado no card: ícone de 48 px em `text-secondary`, título curto, uma linha de explicação e a ação principal, se houver | `EmptyState` |
| sucesso de uma ação | ação com efeito visível (toggle mudou, select aplicou) não tem toast; sem efeito visível, toast de sucesso por 4 s | `Toaster` |
| aviso | risco que não bloqueia: toast de aviso por 4 s quando ele aparece **e** uma faixa no alto do card enquanto ele existir | `Toaster`, faixa de aviso |
| confirmação de ação destrutiva | diálogo com título, a consequência, botão `destructive` e Cancelar com o foco inicial | `AlertDialog` |
| operação que pede administrador | o mesmo diálogo, com `ShieldAlert`, dizendo **antes** do UAC o que será alterado no Windows | `AlertDialog` |
| descartar texto não salvo | "Descartar as alterações?", Cancelar com o foco e "Descartar" `destructive` | `DiscardDialog` |

- **Notas:** nenhum toast espera ser fechado (não há botão fechar); o que precisa ficar na tela é
  faixa, nunca toast. Toast: ícone em anel de 32 px na cor do tipo, barra de 3 px na base que
  encolhe no tempo do toast; pausa com o mouse em cima, com o botão pressionado ou com a janela
  minimizada. A feature que foge destes padrões diz por quê no plano. Erro de validação de campo:
  ver a seção 9 (Campos).

## 11. Acessibilidade visual → `docs/design/acessibilidade-visual.md`

- **Fonte:** `docs/design/acessibilidade-visual.md`
- **Nível alvo** (obrigatório): WCAG 2.2 AA
- **Contraste mínimo:** texto normal 4,5:1 · texto grande `[?]` · elemento de interface e borda
  3:1 (ícone e componente)
- **Foco visível** (obrigatório): anel `highlight` (`--ring`), 2px pelo guia, sem offset `[?]` (no
  código, o `Button` usa `ring-3` em `ring` a 50% e borda `ring`; o `Switch`, `ring-2`). Exceções:
  campo de formulário só pela borda `primary` 1px; `CodeEditor`/`LogViewer` sem anel (cursor e
  linha ativa mostram o foco)
- **Área mínima de toque ou clique** (obrigatório): `44px` de altura (item da sidebar; botões de
  ícone do topo com 44×44)
- **Movimento reduzido:** com `prefers-reduced-motion`, painel lateral, diálogos, menus e o painel
  da bandeja aparecem sem deslizar; a transição do toggle some; o spinner continua (é estado, não
  decoração) — exceto nos botões de recarregar, onde ele para (`motion-reduce:animate-none`)
- **Notas:** pares medidos: `text-primary` passa de 13:1 em todas as superfícies,
  `text-secondary` de 5:1, `primary-foreground` sobre `primary` dá 4,9:1 (por isso o botão é
  `#1F6FD1` e não `#409EFF`, que dá 2,8:1). Exceções aceitas: trilho `toggle-off` com 1,8:1 contra
  a sidebar no escuro (o estado se lê pela posição da bolinha branca); borda azul do campo em foco
  no escuro com 2,5 a 3,5:1. Ícone e item de menu clicáveis mostram cursor `pointer`.

## 12. Movimento e transições → `docs/design/movimento-transicoes.md`

- **Fonte:** `frontend/src/components/ui/{dialog,alert-dialog,sheet,dropdown-menu,select}.tsx`,
  `docs/design/componentes.md` (não há página de movimento no guia)
- **Durações:** sem tokens; valores do código: `duration-100` = 100ms (diálogo, menu, select,
  fundo do painel) · `duration-150` = 150ms (painel da bandeja) · `duration-200` = 200ms (painel
  lateral deslizando) · toast: 4s (8s o de erro)
- **Easing:** sem tokens; `ease-in-out` no painel lateral; o resto usa o padrão do
  `tw-animate-css` (`animate-in`/`animate-out` com `fade`, `zoom-in-95` e deslize de 2) `[?]`
- **Notas:** entradas e saídas vêm do `tw-animate-css` pelo `data-state` do Radix. Tudo some com
  `prefers-reduced-motion` (seção 11).

## 13. Camadas e sobreposições → `docs/design/camadas-sobreposicoes.md`

- **Fonte:** `frontend/src/components/ui/*.tsx`, `docs/design/componentes.md` (não há escala de
  camadas no guia)

| Camada | z-index | Componentes |
|---|---|---|
| conteúdo fixo dentro da rolagem | 10 | cabeçalho fixo de tabela, botões de rolar do select |
| sobreposições | 50 | `Dialog`, `AlertDialog`, `Sheet` (fundo e conteúdo), `Tooltip`, `DropdownMenu`, lista do `Select` |
| toasts | 999999999 (padrão do sonner) | `Toaster` |

- **Modal:** fundo escurecido `overlay` · fecha com Esc (sim) · fecha com clique fora (sim no
  `Dialog` e no `Sheet`; o `AlertDialog` não) · prende o foco (sim; o resto fica `aria-hidden` e
  sem `pointer-events`)
- **Notas:** com um painel modal aberto, o `Toaster` volta a receber cliques por
  `styles/toast.css`. Quem fecha um painel com texto não salvo pergunta antes, no
  `onOpenChange(false)`.

## 14. Lacunas

- **O que o design system não define** e cada projeto terá de decidir: guia de microcopy; licença
  do conjunto; grid e largura máxima de conteúdo (o modelo é app desktop de janela única);
  breakpoints de web (só há faixas de largura de janela); cores de dados para gráficos; números
  tabulares; espaçamento entre letras; tokens de duração e easing (há só valores no código);
  escala de z-index (há só valores no código); estados de hover e somente leitura do campo;
  anatomia do campo (rótulo visível, marca de obrigatório, texto de ajuda) e quando o erro
  aparece; contraste de texto grande; offset do anel de foco; tipografia e ícones para web fora
  do Windows (as fontes são do sistema e não existem em macOS/Linux).
- **Qualquer coisa que não cabia nos campos:** o design system nasce de um app desktop sem
  moldura: barra de título própria (32 px, botões da janela de 46×32 px, fechar em
  `titlebar-close-hover`), painel de bandeja (300 px, `sidebar-bg`, `shadow-md`, linhas de 48 px)
  e janela mínima de 1024×640. Ao aplicar em web, essas partes não se transportam. A identidade é
  de tema escuro azul-marinho primeiro; o claro é derivado. Diferenças entre o guia e o código
  registradas acima: anel de foco 2px no guia e 3px a 50% no `Button`; corpo de texto 16px no
  guia e 14px no `Input`.
