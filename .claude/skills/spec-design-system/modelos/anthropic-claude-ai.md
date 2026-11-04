# Modelo de design system — Claude (claude.ai)

> Retrato de um design system que já existe, guardado na `/spec-design-system` para ser aplicado
> em qualquer projeto com `/spec-design-system modelo anthropic-claude-ai`. Descreve **o design
> system**, não um produto: nome, público e plataformas do projeto são perguntados quando o
> modelo é aplicado.

> **Aviso de origem:** a Anthropic não publica documentação desse design system. Todos os valores
> abaixo foram **lidos do CSS em produção de https://claude.ai/new** (folhas de estilo, variáveis
> CSS e estilos computados) em 2026-10-30, no tema `claude`, `data-color-version="v2"`. O prefixo
> interno dos tokens é `cds-` (Claude Design System). Regras de uso, "quando usar" e diretrizes
> não existem em fonte pública e ficam `[?]`.

**Notação:** `[?]` = a fonte não diz (ao aplicar, a skill decide pelo default consagrado ou
pergunta) · `[—]` = o design system não define isso.

---

## 1. Identificação → `docs/design-system.md`

- **Nome** (obrigatório): Claude Design System (interface do claude.ai) — prefixo de token `cds-`
- **Slug** (obrigatório): `anthropic-claude-ai`
- **Mantido por** (obrigatório): Anthropic
- **Link oficial** (obrigatório): `[—]` não há documentação pública; capturado de https://claude.ai/new
- **Versão capturada** (obrigatório): `data-color-version="v2"`, `data-theme="claude"` (build de produção de 2026-10-30)
- **Data da captura** (obrigatório): 2026-10-30
- **Licença** (obrigatório): proprietária, sem licença publicada. Fontes `anthropic-sans`,
  `anthropic-serif`, `anthropic-mono` e a fonte de ícones `Anthropicons-Variable` são proprietárias.
  Uso fora da organização dona permitido: **não** (reaproveitar só a estrutura e os valores
  numéricos; trocar fontes, ícones e a cor de marca)
- **Plataformas cobertas** (obrigatório): web (também usado no app desktop, que embute a web)
- **Pacote de tokens:** `[—]` não publicado — formato observado: variáveis CSS (`--cds-*`,
  `--bg-*`, `--text-*` …) + utilitários Tailwind
- **Implementações oficiais** (obrigatório):

| Stack | Pacote | Versão | Como importar |
|---|---|---|---|
| nenhuma | — | — | — |

- **Notas:** a base é Tailwind CSS com tema próprio (classes como `rounded-composer`,
  `shadow-focus`, `duration-base`, `bg-bg-000`, `text-text-300`). Há dois níveis de token:
  primitivas `--_<cor>-<tom>` / `--cds-hsl-<cor>-<tom>` e semânticas `--bg-*`, `--text-*`,
  `--border-*`, `--accent-*`, `--danger-*` etc.

## 2. Direção visual → `docs/design-system.md` (seção Direção)

- **Fonte:** inspeção visual e de CSS de https://claude.ai/new
- **Personalidade em 3 adjetivos** (obrigatório): calmo, editorial, discreto
- **Eixos** (obrigatório):
  - sóbrio `1 2 3 4 5` expressivo: **2** (neutros quase puros; cor só na marca `clay` e em estados)
  - denso `1 2 3 4 5` espaçado: **3**
  - geométrico `1 2 3 4 5` orgânico: **3** (cantos médios 8–14px; serifada nos títulos de destaque)
  - plano `1 2 3 4 5` com profundidade: **2** (sombras leves e difusas, bordas em alfa)
- **Densidade de informação:** compacta por padrão; modo **confortável** via
  `.cds-root[data-density="comfortable"]` aumenta controles, paddings e tamanhos de texto (ver seção 4)
- **O que o design system evita** (obrigatório): gradientes; cores saturadas fora de estados;
  bordas opacas (usa borda com alfa 10–20%); MAIÚSCULAS em botões; sombras duras
- **Tom do microcopy:** `[?]` (observado: conversacional e curto, ex.: saudação por período do dia)
- **Notas:** a cor de marca `clay` (#D97757) aparece só no logotipo/asterisco e em acentos de marca.

## 3. Cores → `docs/design/cores.md`

- **Fonte:** variáveis `--cds-hsl-*` em `:root` do CSS de claude.ai (valores convertidos de HSL para hex)
- **Modelo da paleta** (obrigatório): tons por cor, escala não contínua (0–900, apenas os tons usados)
- **Convenção de nome dos tokens** (obrigatório): primitiva `--cds-hsl-<cor>-<tom>` (exposta como
  `--_<cor>-<tom>`); semântica `--<papel>-<nível>` com nível `000`–`500` (`000` = mais forte/frente)
  e `900` = fundo suave do papel
- **Paleta base** (obrigatório — modo claro; as primitivas são as mesmas nos dois modos):

| Cor | Tons (token = valor) |
|---|---|
| cinza | `gray-0 = #FFFFFF` · `gray-20 = #F9F9F9` · `gray-40 = #F3F3F3` · `gray-50 = #EFEFEF` · `gray-80 = #E6E6E6` · `gray-100 = #E0E0E0` · `gray-200 = #C2C2C2` · `gray-350 = #959595` · `gray-450 = #797979` · `gray-650 = #444444` · `gray-700 = #383838` · `gray-750 = #2C2C2C` · `gray-800 = #202020` · `gray-810 = #1E1E1E` · `gray-850 = #151515` · `gray-860 = #131313` · `gray-870 = #111111` · `gray-890 = #0D0D0D` · `gray-900 = #0B0B0B` |
| vermelho | `red-100 = #FBD4D4` · `red-200 = #F7A7A7` · `red-350 = #EB6263` · `red-400 = #EC4647` · `red-450 = #DC3C3D` · `red-600 = #A42528` · `red-750 = #601314` |
| amarelo | `yellow-100 = #F9DCA4` · `yellow-200 = #FBB315` · `yellow-350 = #D28E05` · `yellow-400 = #BF8004` · `yellow-450 = #AE7303` · `yellow-600 = #794C02` · `yellow-750 = #442800` |
| verde | `green-100 = #CAECC8` · `green-200 = #90DA8E` · `green-400 = #0AAC23` · `green-450 = #10A122` · `green-600 = #006C10` · `green-750 = #123B13` |
| azul | `blue-100 = #CCE3FD` · `blue-200 = #9DC8F7` · `blue-350 = #519CED` · `blue-500 = #2370CD` · `blue-600 = #1757AC` · `blue-750 = #063167` |
| violeta | `violet-100 = #E4DEFF` · `violet-200 = #C6BDFA` · `violet-350 = #9887EF` · `violet-400 = #8873E9` · `violet-450 = #785EE8` · `violet-600 = #5437B8` · `violet-750 = #2E1D6B` |
| marca (clay) | `brand-clay = #D97757` · `brand-clay-emphasized = #C6613F` |

- **Primária ou acento:** `--accent-100` (azul `#2370CD` claro / `#519CED` escuro) — links,
  foco e ações de destaque; `--accent-brand` (`#D97757`) — identidade; `--accent-pro-*` (violeta) —
  recursos de plano pago
- **Cores de dados (gráficos):** `[—]`
- **Contraste alvo** (obrigatório): `[?]` (há tratamento de `prefers-contrast: more` e `forced-colors`)
- **Regras de uso:** `[?]`
- **Notas:** existe uma segunda escala neutra levemente quente usada nos componentes `cds`
  (`--cds-gray-10 = #FCFCFB`, `20 = #F9F9F7`, `50 = #F0EFEC`, `200 = #C3C2B7`, `300 = #A5A49A`,
  `400 = #898781`, `500 = #6D6B67`, `600 = #52514E`, `800 = #20201F`).

## 4. Tipografia → `docs/design/tipografia.md`

- **Fonte:** variáveis `--font-*` e `--cds-font-*`; estilos computados da página
- **Família de texto** (obrigatório): `anthropic-sans` — arquivo próprio (variável, pesos 300–800,
  com itálico) — licença: proprietária
- **Família de título:** `anthropic-serif` (variável, 300–800) nos títulos de destaque (ex.:
  saudação da página inicial); títulos de interface usam a sans
- **Família monoespaçada:** `anthropic-mono`
- **Fallback stack** (obrigatório): sans → `system-ui, "Segoe UI", Roboto, Helvetica, Arial,
  "PingFang SC", "PingFang TC", "Hiragino Sans", "Apple SD Gothic Neo", <fontes de scripts Noto/Kohinoor>, sans-serif`;
  serif → `"Anthropic Serif Fallback Georgia"`, `…Times`, `…DejaVu`, `…Noto` (fallbacks com métricas
  ajustadas); mono → `ui-monospace, monospace`
- **Unidade do tamanho de texto** (obrigatório): rem (multiplicado por `--cds-rem-scale`, padrão 1)
- **Estilos nomeados** (obrigatório):

| Estilo (token oficial) | Tamanho | Peso | Altura de linha | Espaçamento entre letras | Uso |
|---|---|---|---|---|---|
| `--cds-font-size-footnote` (compacto) | 12px (.75rem) | 400 | 15px | 0 | notas, metadados |
| `--cds-font-size-footnote` (confortável) | 13px | 400 | 17px | 0 | idem |
| `--cds-font-size-body` (compacto) | 13px (.8125rem) | 400 | 19px | 0 | texto de interface |
| `--cds-font-size-body` (confortável) | 14px | 400 | 20px | 0 | idem |
| `--cds-font-size-heading` (compacto) | 14px | `[?]` | 18px | 0 | título de componente |
| `--cds-font-size-heading` (confortável) | 15px | `[?]` | 20px | 0 | idem |
| `--cds-font-size-prose` (compacto / confortável) | 14px / 16px | 400 | `[?]` | 0 | texto longo (respostas) |
| `text-xs` | 12px | — | 16px | 0 | utilitário Tailwind |
| `text-sm` | 14px | — | 20px | 0 | itens de navegação, botões |
| `text-base` | 16px | — | 24px | 0 | corpo da página |
| `text-lg` / `text-xl` | 18px / 20px | — | 28px | 0 | subtítulos |
| `text-2xl` / `text-3xl` | 24px / 30px | — | 32px / 36px | 0 | títulos |
| saudação (serif) | ≈28.6px (fluido) | 360 (`wght` 320) | 1.2 | normal | título de destaque da tela inicial |

- **Números tabulares:** `[?]`
- **Notas:** existe suporte a fontes de acessibilidade (`OpenDyslexic`, `Atkinson Hyperlegible Next`)
  como preferência do usuário.

## 5. Espaçamento, raios e sombras → `docs/design/espacamento-raios-sombras.md`

- **Fonte:** variáveis `--cds-*` e utilitários Tailwind resolvidos no navegador
- **Escala de espaçamento** (obrigatório): base 4px (Tailwind)

| Token | Valor |
|---|---|
| `1` | 4px |
| `2` | 8px |
| `3` | 12px |
| `4` | 16px |
| `6` | 24px |
| `8` | 32px |
| `--cds-gap-xs` | 6px (compacto) · 8px (confortável) |
| `--cds-pad-md` | 8px (compacto) · 12px (confortável) |
| `--cds-h-control` | 24px (compacto) · 32px (confortável) — altura de controle |

- **Raios** (obrigatório):

| Token | Valor | Onde se usa |
|---|---|---|
| `rounded-sm` | 4px | detalhes pequenos |
| `rounded-chip` | 4.5px | chips, etiquetas |
| `rounded` / `rounded-md` | 6px | botões pequenos (24px) |
| `rounded-lg` | 8px | botões 28–32px, itens da navegação lateral |
| `rounded-xl` | 12px | cartões, menus |
| `--cds-radius-composer` | 12px (compacto) · 14px (confortável) | caixa de mensagem (composer) |
| `rounded-2xl` | 16px | diálogos |
| `rounded-3xl` | 24px | superfícies grandes |
| `rounded-full` | 9999px | avatar, pílulas |

- **Espessura de borda padrão:** 1px, cor `--cds-border` = `--cds-neutral-900` a 10% (`--cds-border-strong` = 20%);
  nas telas, `border-border-300` com opacidade 15%
- **Estratégia de elevação** (obrigatório): fundo diferente (`--bg-000`…`--bg-500`) + borda em alfa;
  sombra leve só em superfícies flutuantes
- **Sombras:**

| Token | Valor | Para quê |
|---|---|---|
| `shadow-sm` | `0 1px 2px rgb(11 11 11 / .06), 0 2px 8px rgb(0 0 0 / .24)` | botão elevado, chip |
| `shadow-md` | `0 2px 4px rgb(11 11 11 / .07), 0 6px 16px rgb(0 0 0 / .24)` | menus, popovers |
| `shadow-lg` | `0 4px 8px rgb(11 11 11 / .08), 0 12px 28px -2px rgb(0 0 0 / .24)` | diálogos |
| `shadow-composer-focus` | `0 4px 20px rgb(0 0 0 / .075), inset 0 0 0 1px rgb(255 255 255 / .08)` | composer (valor do modo escuro) |
| `shadow-focus` | `inset 0 0 0 1px <bg>, 0 0 0 1px #257DE0, 0 0 6px 1px rgb(23 87 172 / .6)` | anel de foco |

- **Notas:** valores de sombra lidos no modo escuro; a opacidade pode variar no claro `[?]`.

## 6. Ícones → `docs/design/icones.md`

- **Fonte:** fontes carregadas (`document.fonts`) e elementos da navegação
- **Biblioteca** (obrigatório): `Anthropicons-Variable` — fonte de ícones proprietária (glifos na área
  de uso privado Unicode, ex.: `U+E0C9`); alguns ícones em SVG inline de 20×20
- **Estilo** (obrigatório): outline — traço 1.25 (SVG); fonte variável com pesos 400–700
- **Tamanhos** (obrigatório): 16px · 20px — padrão: 20px
- **Cor:** herda a cor do texto (`currentColor`)
- **Ícones com significado fixo:** `[?]`

| Significado | Ícone |
|---|---|
| `[?]` | `[?]` |

- **Notas:** licença proprietária — ao aplicar, substituir por uma biblioteca outline de traço fino
  equivalente.

## 7. Layout → `docs/design/layout.md`

- **Fonte:** media queries do CSS de claude.ai
- **Breakpoints ou faixas de tela** (obrigatório):

| Token | Valor | O que muda |
|---|---|---|
| `sm` | 640px (40rem) | `[?]` |
| (sem nome) | 701px | navegação lateral passa de sobreposta a fixa `[?]` |
| `md` | 768px (48rem) | layout de tablet/desktop |
| `lg` | 1024px (64rem) | `[?]` |

- **Largura máxima do conteúdo:** `max-w-3xl = 768px` (coluna de conversa) `[?]`
- **Grid:** `[—]` (layout em flex, sem grid de colunas)
- **Padrões de navegação que o design system oferece:** navegação lateral recolhível (modo trilho,
  `data-sidebar-chrome="rail"`) com lista de recentes — no menor tamanho vira gaveta sobreposta
- **Barra de rolagem:** personalizada — largura `thin` · polegar `rgb(224 224 224 / .35)` no escuro,
  trilho transparente · `scrollbar-gutter: stable` · aparece sempre (fina)
- **Notas:** composer da página inicial com ≈540px de largura em janela de 1023px.

## 8. Tema → `docs/design/tema.md`

- **Fonte:** regras `[data-color-version="v2"][data-theme="claude"]` e `…[data-mode="dark"]`
- **Modos oferecidos** (obrigatório): claro, escuro (seguem o sistema por padrão); ajustes para
  `prefers-contrast: more` e `forced-colors`
- **Como o modo é ativado:** atributo `data-mode="light|dark"` no `<html>`, junto de
  `data-theme="claude"` e `data-color-version="v2"`
- **Tokens semânticos** (obrigatório):

| Token | Claro | Escuro | Uso |
|---|---|---|---|
| `--bg-000` | `#FFFFFF` | `#202020` | superfície elevada (cartão, composer) |
| `--bg-100` | `#F9F9F9` | `#151515` | fundo da página |
| `--bg-200` | `#F3F3F3` | `#111111` | fundo secundário |
| `--bg-300` | `#EFEFEF` | `#0D0D0D` | hover / área rebaixada |
| `--bg-400` / `--bg-500` | `#E6E6E6` | `#0B0B0B` | fundo mais profundo |
| `--text-000` / `--text-100` | `#131313` | `#F9F9F9` | texto primário |
| `--text-200` / `--text-300` | `#383838` | `#C2C2C2` | texto secundário |
| `--text-400` / `--text-500` | `#797979` | `#959595` | texto atenuado, placeholder |
| `--border-100`…`--border-400` | `#1E1E1E` | `#E0E0E0` | borda (usada com opacidade ≈15%) |
| `--accent-000` | `#1757AC` | `#9DC8F7` | acento forte (texto) |
| `--accent-100` / `--accent-200` | `#2370CD` | `#519CED` | acento (link, foco) |
| `--accent-900` | `#CCE3FD` | `#063167` | fundo de acento |
| `--accent-brand` / `--brand-200` | `#D97757` | `#D97757` | cor de marca |
| `--brand-000` | `#C6613F` | `#C6613F` | marca enfatizada |
| `--brand-100` | `#C6613F` | `#D97757` | marca (hover) |
| `--accent-pro-000` | `#5437B8` | `#C6BDFA` | plano pago, texto |
| `--accent-pro-100` | `#785EE8` | `#9887EF` | plano pago |
| `--accent-pro-900` | `#E4DEFF` | `#2E1D6B` | plano pago, fundo |
| `--success-000` | `#006C10` | `#90DA8E` | texto de sucesso |
| `--success-100` / `-200` | `#10A122` | `#0AAC23` | sucesso |
| `--success-900` | `#CAECC8` | `#123B13` | fundo de sucesso |
| `--warning-000` | `#794C02` | `#FBB315` | texto de aviso |
| `--warning-100` | `#AE7303` | `#D28E05` | aviso |
| `--warning-900` | `#F9DCA4` | `#442800` | fundo de aviso |
| `--danger-000` | `#A42528` | `#F7A7A7` | texto de erro |
| `--danger-100` | `#DC3C3D` | `#EB6263` | erro |
| `--danger-200` | `#DC3C3D` | `#EC4647` | erro (ícone/borda) |
| `--danger-900` | `#FBD4D4` | `#601314` | fundo de erro |
| `--pictogram-100`…`400` | `#E6E6E6` · `#E0E0E0` · `#FFFFFF` · `#F9F9F9` | `#444444` · `#383838` · `#2C2C2C` · `#202020` | ilustrações |

- **Notas:** no escuro, a página observada usa fundo `#191919`, composer `#303030` e texto
  `rgb(255 255 255 / .93)` (camada `cds` sobre os tokens acima).

## 9. Componentes → `docs/design/componentes.md`

- **Fonte:** observação da interface (sem catálogo público)
- **Catálogo** (obrigatório):

| Componente (nome oficial) | Variantes | Quando usar (pelo guia) |
|---|---|---|
| Botão (`[?]` nome) | alturas 24px (raio 6), 28px (raio 8), 32px (raio 8); ícone-só; ghost | `[?]` |
| Composer (caixa de mensagem) | página inicial, conversa | `[?]` |
| Item de navegação lateral | 32px, raio 8, texto 14px, ícone 20px | `[?]` |
| Chip | raio 4.5px | `[?]` |
| Menu suspenso / popover | `[?]` | `[?]` |
| Tooltip | `[?]` | `[?]` |
| Diálogo (modal) | `[?]` | `[?]` |
| Etiqueta ("Experimental") | acento azul | `[?]` |

- **Estados que todo interativo tem** (obrigatório): padrão · hover · foco · ativo · desabilitado ·
  selecionado observados; carregando e erro `[?]`
- **Notas:** botões sem caixa alta; ações de ícone dominam a interface.

### Campos de formulário → `docs/design/formularios.md`

- **Fonte:** composer e `--cds-h-control`
- **Tamanhos** (obrigatório):

| Tamanho (token) | Altura | Padding horizontal | Tamanho do texto |
|---|---|---|---|
| compacto (`--cds-h-control`) | 24px | 8px (`--cds-pad-md`) | 13px |
| confortável (`--cds-h-control`) | 32px | 12px (`--cds-pad-md`) | 14px |
| composer | automática | `[?]` | 15px / 21px |

- **Estados** (obrigatório):

| Estado | Borda | Fundo | Texto | Outro (sombra, contorno, ícone) |
|---|---|---|---|---|
| padrão | `--cds-border` (10%) | `--bg-000` | `--text-000` | sombra difusa leve |
| hover | `[?]` | `[?]` | `[—]` | `[?]` |
| foco | `[—]` | `[—]` | `[—]` | `shadow-composer-focus` / `shadow-focus` |
| erro | `[?]` | `[?]` | `--danger-000` `[?]` | `[?]` |
| desabilitado | `[?]` | `[?]` | `[?]` | `[?]` |
| somente leitura | `[?]` | `[?]` | `[?]` | `[?]` |

- **Anatomia:** rótulo `[?]` · obrigatório marcado `[?]` · texto de ajuda `[?]` · mensagem de erro `[?]`
  · contador de caracteres não
- **Placeholder:** cor `--text-500` (≈`#878787` no escuro) · regra de uso `[?]`
- **Quando o erro aparece:** `[?]`
- **Notas:**

## 10. Padrões de feedback → `docs/design/padroes-feedback.md`

- **Fonte:** `[?]` — não há guia público; nada foi inferido.

| Situação | Padrão | Componente |
|---|---|---|
| carregando a página | `[?]` | `[?]` |
| erro de requisição | `[?]` | `[?]` |

- **Notas:**

## 11. Acessibilidade visual → `docs/design/acessibilidade-visual.md`

- **Fonte:** media queries e `shadow-focus` do CSS
- **Nível alvo** (obrigatório): `[?]`
- **Contraste mínimo:** `[?]`
- **Foco visível** (obrigatório): anel `shadow-focus` — 1px interno da cor do fundo + 1px `#257DE0`
  + halo 6px `rgb(23 87 172 / .6)`; offset 0
- **Área mínima de toque ou clique** (obrigatório): `[?]` (menor controle observado: 24px)
- **Movimento reduzido:** 54 regras em `prefers-reduced-motion: reduce` desativam ou encurtam
  animações; também trata `prefers-reduced-transparency`, `prefers-contrast: more`, `forced-colors`
- **Notas:** fontes alternativas OpenDyslexic e Atkinson Hyperlegible disponíveis.

## 12. Movimento e transições → `docs/design/movimento-transicoes.md`

- **Fonte:** utilitários Tailwind resolvidos no navegador
- **Durações:** `duration-fast = 60ms` · `duration-snap = 120ms` · padrão Tailwind `150ms` ·
  `duration-base = 200ms` · `duration-slow = 450ms`
- **Easing:** `ease-out = cubic-bezier(0.165, 0.84, 0.44, 1)` · `ease-in-out = cubic-bezier(0.4, 0, 0.2, 1)`
- **Notas:**

## 13. Camadas e sobreposições → `docs/design/camadas-sobreposicoes.md`

- **Fonte:** utilitários `z-*` resolvidos no navegador

| Camada | z-index | Componentes |
|---|---|---|
| `z-header` | 20 | cabeçalho |
| `z-sidebar` | 30 | navegação lateral |
| `z-modal` | 40 | diálogo |
| `z-dropdown` / `z-popover` / `z-tooltip` | 50 | menus, popovers, tooltips |

- **Modal:** fundo escurecido `[?]` · fecha com Esc `[?]` · fecha com clique fora `[?]` · prende o foco `[?]`
- **Notas:**

## 14. Lacunas

- **O que o design system não define** e cada projeto terá de decidir: documentação pública,
  catálogo de componentes, regras de uso de cor, cores de gráfico, padrões de feedback, anatomia de
  formulário, grid, contraste-alvo e área mínima de toque.
- **Qualquer coisa que não cabia nos campos:** fontes e ícones são proprietários — ao aplicar o
  modelo, trocar `anthropic-sans/serif/mono` por equivalentes livres e `Anthropicons` por uma
  biblioteca outline; trocar `clay` pela cor de marca do projeto. Valores mudam sem aviso a cada
  deploy de claude.ai.
