# Modelo de design system — Primer

> Retrato de um design system que já existe (ex.: o Primer, do GitHub), guardado na
> `/spec-design-system` para ser aplicado em qualquer projeto com
> `/spec-design-system modelo <slug>`. Descreve **o design system**, não um produto: nome,
> público e plataformas do projeto são perguntados quando o modelo é aplicado.

> **Como este formulário foi preenchido (2026-10-30).** Navegação em primer.style (páginas de Getting
> started, Primitives, Componentes, UI Patterns, Acessibilidade e Octicons) e leitura dos arquivos
> publicados dos pacotes `@primer/primitives@11.10.0` e `@primer/react@38.43.0` (CSS dos tokens e dos
> componentes, via jsDelivr). Todos os valores de cor abaixo são do **modo claro** salvo quando a
> coluna diz outro modo. Escopo: **Primer Product UI** (o Primer Brand é outro conjunto, para marketing,
> e não foi capturado).

**Notação:** `[?]` = a fonte não diz (ao aplicar, a skill decide pelo default consagrado ou
pergunta) · `[—]` = o design system não define isso.

---

## 1. Identificação → `docs/design-system.md`

- **Nome** (obrigatório): Primer (Product UI)
- **Slug** (obrigatório): `github-primer`
- **Mantido por** (obrigatório): GitHub
- **Link oficial** (obrigatório): https://primer.style/
- **Versão capturada** (obrigatório): `@primer/primitives` 11.10.0 (tokens) · `@primer/react` 38.43.0 · `@primer/octicons-react` 19.42.0 (versões `latest` do npm na data da captura)
- **Data da captura** (obrigatório): 2026-10-30
- **Licença** (obrigatório): código e tokens **MIT** (campo `license` dos pacotes `@primer/primitives`, `@primer/react`, `@primer/css`, `@primer/view-components` e `@primer/octicons*` no npm). Fonte Mona Sans: **SIL Open Font License 1.1** (arquivo LICENSE em github.com/github/mona-sans), à parte. Logotipos da GitHub dentro do conjunto Octicons: política de uso própria `[?]` — não verificada. Uso fora da organização dona permitido: **sim** (MIT)
- **Plataformas cobertas** (obrigatório): **web** (React, Rails/ViewComponents e CSS). Mobile iOS `[—]` · Mobile Android `[—]` · Desktop `[—]` (nada na documentação do Product UI)
- **Pacote de tokens:** `@primer/primitives` — formato: **variáveis CSS** (`dist/css/base/*`, `dist/css/functional/*`, um arquivo por tema em `functional/themes/`) e variáveis do Figma. Outros formatos de distribuição (JSON, SCSS, JS) `[?]` — não conferidos
- **Implementações oficiais** (obrigatório):

| Stack | Pacote | Versão | Como importar |
|---|---|---|---|
| React | `@primer/react` | 38.43.0 | `import {Button} from '@primer/react'` |
| Ícones (React) | `@primer/octicons-react` | 19.42.0 | `[?]` (a página de código dos Octicons não foi lida) |
| Rails | `@primer/view-components` | 0.53.5 | ViewComponents; guia em https://primer.style/product/getting-started/rails/ (nome da gem `[?]`) |
| CSS | `@primer/css` | 22.3.3 | `[?]` |

- **Notas:** o Primer tem três frentes: **Product UI** (esta captura), **Brand UI** (`@primer/react-brand` 0.80.1, marketing) e fundações compartilhadas (Acessibilidade, Octicons, Primitives). Status por componente (Experimental · Ready · Deprecated) em https://primer.style/product/getting-started/component-status/: só "Ready" tem suporte de longo prazo e migração em mudança incompatível. Fonte: https://primer.style/ e https://primer.style/product/getting-started/

## 2. Direção visual → `docs/design-system.md` (seção Direção)

- **Fonte:** https://primer.style/product/getting-started/ (A cohesive experience · Inclusive and accessible design · Design for efficiency · The Zen of GitHub) · https://primer.style/product/getting-started/foundations/typography/
- **Personalidade em 3 adjetivos** (obrigatório): **coeso** · **inclusivo/acessível** · **eficiente** (a própria página fala em experiência "fast and compact", "calm", "logical")
- **Eixos** (obrigatório) — a fonte não dá notas; os valores abaixo são **leitura sugerida a partir do texto e dos tokens, não oficial** `[?]`:
  - sóbrio `1 2 3 4 5` expressivo: **2** `[?]` (uso restrito de cor como ênfase: "Refrain from utilizing color as a primary method of emphasis")
  - denso `1 2 3 4 5` espaçado: **2** `[?]` (experiência "fast and compact"; controles de 32px)
  - geométrico `1 2 3 4 5` orgânico: **2** `[?]` (raio padrão 6px)
  - plano `1 2 3 4 5` com profundidade: **2** `[?]` (borda de 1px + sombras sutis; sombra forte só em sobreposição)
- **Densidade de informação:** **compacta** por padrão. Muda por: tamanho do controle (`small` 28px · `medium` 32px · `large` 40px · `xlarge` 48px), variantes de padding `condensed · normal · spacious` nos controles e no `Stack`, e ponteiro grosso (alvo mínimo sobe de 16px para 44px).
- **O que o design system evita** (obrigatório): (a) mais de um botão primário num grupo de botões e quase nunca mais de um por página; (b) cor como único meio de dar ênfase ou de mostrar estado; (c) botões desabilitados ("discouraged") — para formulário, validar no envio; (d) placeholder no lugar do rótulo; (e) mais de um `Banner` ao mesmo tempo na página, e estilo de título ou vários tamanhos de texto dentro do Banner; (f) animação decorativa ("design" motion é a que mais causa dano sem agregar valor); (g) botão `large` em excesso; (h) adicionar o que dilui o resto ("Anything added dilutes everything else").
- **Tom do microcopy:** princípios do "Zen of GitHub": "Mind your words, they are important" · "Speak like a human". Regras de formulário: rótulo descritivo e conciso (até ~3 palavras), em *sentence case* (ex.: "Repository name"). Guia completo de conteúdo em https://primer.style/product/getting-started/foundations/content/ `[?]` (página não lida)
- **Notas:** os 14 aforismos do Zen incluem: "Responsive is better than fast", "Approachable is better than simple", "Encourage flow", "Non-blocking is better than blocking", "Favor focus over features", "Avoid administrative distraction", "Design for failure".

## 3. Cores → `docs/design/cores.md`

- **Fonte:** https://primer.style/product/primitives/color/ · https://primer.style/product/getting-started/foundations/color-usage/ · `@primer/primitives@11.10.0/dist/css/functional/themes/light.css` e `dark.css`
- **Modelo da paleta** (obrigatório): **tokens funcionais por papel**, sobre escalas base. A documentação pública de Color só lista os funcionais; a página Color usage descreve escalas base (uma neutra com 11 passos, invertida entre claro e escuro, e escalas por cor), mas os valores brutos de cada passo **não** estão nas páginas lidas `[?]`.
- **Convenção de nome dos tokens** (obrigatório): funcional `<fgColor|bgColor|borderColor>-<papel>[-emphasis|-muted]` (ex.: `fgColor-default`, `bgColor-accent-muted`, `borderColor-danger-emphasis`); de componente `button-<variante>-<propriedade>-<estado>` e `control-<propriedade>-<estado>`; dados `data-<cor>-color-<emphasis|muted>`; base `base-size-*`, `base-text-*`, `base-duration-*`.
- **Paleta base** (obrigatório — uma linha por papel, modo claro; escalas numeradas `[?]`, ver acima):

| Papel | fgColor | bgColor-emphasis | bgColor-muted | borderColor-emphasis | borderColor-muted |
|---|---|---|---|---|---|
| neutral | `fgColor-neutral = #59636e` | `#59636e` | `#818b981f` | `#59636e` | `#d1d9e0b3` |
| accent | `fgColor-accent = #0969da` | `#0969da` | `#ddf4ff` | `#0969da` | `#54aeff66` |
| success | `fgColor-success = #1a7f37` | `#1f883d` | `#dafbe1` | `#1a7f37` | `#4ac26b66` |
| attention | `fgColor-attention = #9a6700` | `#9a6700` | `#fff8c5` | `#9a6700` | `#d4a72c66` |
| severe | `fgColor-severe = #bc4c00` | `#bc4c00` | `#fff1e5` | `#bc4c00` | `#fb8f4466` |
| danger | `fgColor-danger = #d1242f` | `#cf222e` | `#ffebe9` | `#cf222e` | `#ff818266` |
| done | `fgColor-done = #8250df` | `#8250df` | `#fbefff` | `#8250df` | `#c297ff66` |
| sponsors | `fgColor-sponsors = #bf3989` | `#bf3989` | `#ffeff7` | `#bf3989` | `#ff80c866` |

  Aliases de estado: `open` = success (`fgColor-open = #1a7f37`) · `closed` = danger (`fgColor-closed = #d1242f`) · `draft` = neutral (`fgColor-draft = #59636e`) · `upsell` = done (`fgColor-upsell = #8250df`). Neutros: `bgColor-default = #ffffff` · `bgColor-muted = #f6f8fa` · `bgColor-inset = #f6f8fa` · `bgColor-emphasis = #25292e` · `bgColor-inverse = #25292e` · `bgColor-disabled = #eff2f5` · `bgColor-black = #1f2328` · `bgColor-white = #ffffff` · `fgColor-default = #1f2328` · `fgColor-muted = #59636e` · `fgColor-disabled = #818b98` · `fgColor-onEmphasis = #ffffff` · `fgColor-onInverse = #ffffff` · `borderColor-default = #d1d9e0` · `borderColor-muted = #d1d9e0b3` · `borderColor-emphasis = #818b98` · `borderColor-disabled = #818b981a` · `borderColor-translucent = #1f232826` · `borderColor-transparent = #ffffff00`.
- **Primária ou acento:** `fgColor-accent` / `bgColor-accent-emphasis = #0969da` — usado em links (`fgColor-link` = accent), anel de foco (`focus-outlineColor`), controles marcados (`control-checked-bgColor-rest = #0969da`), borda de campo em foco e seleção (`selection-bgColor = #0969da33`). Atenção: o **botão primário é verde** (`button-primary-bgColor-rest = bgColor-success-emphasis = #1f883d`), não azul.
- **Cores de dados (gráficos):** 17 cores, cada uma com `--data-<cor>-color-emphasis` / `-muted`: auburn `#9d615c`/`#f2e9e9` · blue `#006edb`/`#d1f0ff` · brown `#856d4c`/`#eeeae2` · coral `#d43511`/`#ffe5db` · gray `#808fa3`/`#e8ecf2` · green `#30a147`/`#caf7ca` · lemon `#866e04`/`#f7eea1` · lime `#527a29`/`#e3f2b5` · olive `#64762d`/`#f0f0ad` · orange `#eb670f`/`#ffe7d1` · pine `#167e53`/`#bff8db` · pink `#ce2c85`/`#ffe5f1` · plum `#a830e8`/`#f8e5ff` · purple `#894ceb`/`#f1e5ff` · red `#df0c24`/`#ffe2e0` · teal `#179b9b`/`#c7f5ef` · yellow `#b88700`/`#ffec9e`. Ordem de uso `[?]` (página https://primer.style/product/ui-patterns/data-visualization/ não lida).
- **Contraste alvo** (obrigatório): **WCAG AA** (níveis A e AA). Exigido nos temas padrão Day (claro) e Night (escuro); os demais subtemas não precisam cumprir (o "Dark dimmed" é de baixo contraste de propósito).
- **Regras de uso:** não depender só de cor (mostrar estado também por texto ou ícone); texto normal 4,5:1 e texto grande 3:1; borda de campo contra o fundo onde ele está 3:1; estado selecionado 3:1; validação com ícone sozinho 3:1; ícone funcional (sem texto) 3:1; desabilitado isento; cálculo de contraste de texto e borda contra `bgColor-muted` (vale para `default` e `muted`); cores de ícone por significado: `fgColor-accent` = informação, `fgColor-success` = sucesso, `fgColor-danger` = erro, `fgColor-attention` = aviso, `fgColor-onEmphasis` sobre fundo escuro ou colorido; em gráficos, ligar rótulos às séries com linhas ou padrões.
- **Notas:** a página Color usage diz "nine themes", mas o pacote 11.10.0 publica **14** arquivos de tema (ver seção 8). Valores de todos os tokens, em todos os temas, estão nos arquivos `functional/themes/*.css`. Focus: `--focus-outlineColor = #0969da`. Overlay: `--overlay-bgColor = #ffffff` · `--overlay-borderColor = #d1d9e080` · `--overlay-backdrop-bgColor = #c8d1da66`. Controles: ver seção 9. Botões: ver tokens `button-*` na página de Color.

## 4. Tipografia → `docs/design/tipografia.md`

- **Fonte:** https://primer.style/product/primitives/typography/ · https://primer.style/product/getting-started/foundations/typography/ · `@primer/primitives@11.10.0/dist/css/functional/typography/typography.css`
- **Família de texto** (obrigatório): `Mona Sans VF` (fonte variável) — de onde vem: fonte do próprio GitHub, hospedagem/pacote `[?]` (a documentação só cita o nome na pilha) — licença da fonte: SIL OFL 1.1
- **Família de título:** mesma do texto (`--fontStack-sansSerifDisplay` tem o mesmo valor de `--fontStack-sansSerif`)
- **Família monoespaçada:** `ui-monospace` (fontes do sistema) — `--fontStack-monospace`
- **Fallback stack** (obrigatório): sans: `'Mona Sans VF', -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans Backtick Fix', 'Noto Sans', Helvetica, Arial, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji'` · mono: `ui-monospace, SFMono-Regular, SF Mono, Menlo, Consolas, Liberation Mono, monospace`
- **Unidade do tamanho de texto** (obrigatório): **rem** (para respeitar o zoom do navegador); altura de linha **sem unidade**, alinhada à grade de 4px
- **Estilos nomeados** (obrigatório) — pixels = rem × 16; altura de linha em px = tamanho × fator:

| Estilo (token oficial) | Tamanho | Peso | Altura de linha | Espaçamento entre letras | Uso |
|---|---|---|---|---|---|
| `text-display` | 2.5rem = 40px | 500 | 1.375 (55px) | `[—]` | `[?]` |
| `text-title-large` | 2rem = 32px | 600 | 1.5 (48px) | `[—]` | `[?]` |
| `text-title-medium` | 1.25rem = 20px | 600 | 1.625 (32,5px) | `[—]` | `[?]` |
| `text-title-small` | 1rem = 16px | 600 | 1.5 (24px) | `[—]` | `[?]` |
| `text-subtitle` | 1.25rem = 20px | 400 | 1.625 (32,5px) | `[—]` | `[?]` |
| `text-body-large` | 1rem = 16px | 400 | 1.5 (24px) | `[—]` | `[?]` |
| `text-body-medium` | 0.875rem = 14px | 400 | 1.5 (21px) | `[—]` | texto padrão dos controles (ex.: TextInput) |
| `text-body-small` | 0.75rem = 12px | 400 | 1.625 (19,5px) | `[—]` | `[?]` |
| `text-caption` | 0.75rem = 12px | 400 | 1.25 (15px) | `[—]` | `[?]` |
| `text-codeBlock` | 0.8125rem = 13px | 400 | 1.5 (19,5px) | `[—]` | bloco de código (mono) |
| `text-codeInline` | 0.9285em | 400 | `[—]` | `[—]` | código em linha (mono) |

  Família por estilo: display, subtitle, title-large e title-medium usam `sansSerifDisplay`; title-small, body e caption usam `sansSerif`; código usa `monospace`. Atalhos `font:` prontos: `--text-<estilo>-shorthand[-<tamanho>]` (ex.: `font: var(--text-display-shorthand)`). Escalas base: tamanhos `base-text-size-xs/sm/md/lg/xl/2xl = 12/14/16/20/32/40px` · pesos `light 300 · normal 400 · medium 500 · semibold 600` · alturas de linha `tight 1.25 · snug 1.375 · normal 1.5 · relaxed 1.625 · loose 1.75`.
- **Números tabulares:** `[?]` — não citado nas páginas lidas
- **Notas:** boas práticas do guia: não usar cor como principal método de ênfase; usar os tokens de altura de linha para alinhar; manter linhas com cerca de 80 caracteres ou menos; preferir os tokens "shorthand".

## 5. Espaçamento, raios e sombras → `docs/design/espacamento-raios-sombras.md`

- **Fonte:** https://primer.style/product/primitives/size/ · https://primer.style/product/primitives/color/ (seção Shadow) · `@primer/primitives@11.10.0/dist/css/functional/{size,spacing}/*.css`
- **Escala de espaçamento** (obrigatório):

| Token | Valor |
|---|---|
| `base-size-2` | 2px (0.125rem) |
| `base-size-4` | 4px |
| `base-size-6` | 6px |
| `base-size-8` | 8px |
| `base-size-12` | 12px |
| `base-size-16` | 16px |
| `base-size-20` | 20px |
| `base-size-24` | 24px |
| `base-size-28` | 28px |
| `base-size-32` | 32px |
| `base-size-36` | 36px |
| `base-size-40` | 40px |
| `base-size-44` | 44px |
| `base-size-48` | 48px |
| `base-size-64` | 64px |
| `base-size-80` | 80px |
| `base-size-96` | 96px |
| `base-size-112` | 112px |
| `base-size-128` | 128px |
| `base-size-negative-<n>` | −n px, para n = 2, 4, 6, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48 |
| `space-xxs / xs / sm / md / lg / xl` | 2 / 4 / 8 / 12 / 16 / 24px |
| `stack-gap-condensed / normal / spacious` (e `stack-padding-*`) | 8 / 16 / 24px |

- **Raios** (obrigatório):

| Token | Valor | Onde se usa |
|---|---|---|
| `borderRadius-small` | 3px (0.1875rem) | `[?]` |
| `borderRadius-medium` = `borderRadius-default` | 6px (0.375rem) | campo de texto (TextInput), sobreposições (`overlay-borderRadius`); padrão geral |
| `borderRadius-large` | 12px (0.75rem) | `[?]` |
| `borderRadius-full` | 9999px | `[?]` (pílulas e círculos, pelo nome) |

- **Espessura de borda padrão:** `borderWidth-default = borderWidth-thin = 1px (0.0625rem)`; também `borderWidth-thick = 2px` e `borderWidth-thicker = 4px`. Bordas por sombra interna: `boxShadow-thin / thick / thicker` = `inset 0 0 0 <espessura>`.
- **Estratégia de elevação** (obrigatório): **borda de 1px + sombra sutil**; a sombra forte só em elementos flutuantes. Campos usam `shadow-inset`; sobreposições usam `shadow-floating-*` e borda de 1px na própria sombra.
- **Sombras** (modo claro; o modo escuro tem valores próprios em `dark.css`):

| Token | Valor | Para quê |
|---|---|---|
| `shadow-inset` | `inset 0 1px 0 0 #1f23280a` | campo de texto (confirmado no CSS do TextInput) |
| `shadow-resting-xsmall` | `0 1px 1px 0 #1f23280d` | elevação "em repouso" (pelo nome) `[?]` |
| `shadow-resting-small` | `0 1px 1px 0 #1f23280a, 0 1px 2px 0 #1f232808` | em repouso `[?]` |
| `shadow-resting-medium` | `0 1px 1px 0 #25292e1a, 0 3px 6px 0 #25292e1f` | em repouso `[?]` |
| `shadow-floating-small` | `0 0 0 1px #d1d9e040, 0 6px 12px -3px #25292e0a, 0 6px 18px 0 #25292e1f` | flutuante `[?]` |
| `shadow-floating-medium` | `0 0 0 1px #d1d9e000, 0 8px 16px -4px #25292e14, 0 4px 32px -4px #25292e14, 0 24px 48px -12px #25292e14, 0 48px 96px -24px #25292e14` | flutuante `[?]` |
| `shadow-floating-large` | `0 0 0 1px #d1d9e000, 0 40px 80px 0 #25292e3d` | flutuante `[?]` |
| `shadow-floating-xlarge` | `0 0 0 1px #d1d9e000, 0 56px 112px 0 #25292e52` | flutuante `[?]` |
| `button-default-shadow-resting` | `0 1px 0 0 #1f23280a` | botão padrão |

- **Notas:** exemplos de tamanhos de controle e de sobreposição na seção 9 e 13. Qual sombra cada componente usa só está nos componentes (`[?]` onde não foi lido).

## 6. Ícones → `docs/design/icones.md`

- **Fonte:** https://primer.style/octicons/ · https://primer.style/octicons/design-guidelines/ · https://primer.style/octicons/usage-guidelines/
- **Biblioteca** (obrigatório): **Octicons** — pacotes: React `@primer/octicons-react` 19.42.0 · SVG `@primer/octicons` 19.42.0 · Rails (via `@primer/view-components`) `[?]`
- **Estilo** (obrigatório): **outline**, com variantes preenchidas para alguns ícones (sufixo `-fill`, ex.: `alert-fill-16`) — traço de **1,5px** em 16px e em 24px
- **Tamanhos** (obrigatório): **16px** e **24px** (cada ícone tem as duas versões); 12px só em caso especial em que 16px não serve — padrão: 16px `[?]` (a página de uso não nomeia o padrão)
- **Cor:** token funcional por significado (ver ícones fixos abaixo); ícone de controle `control-iconColor-rest = #59636e`; `fgColor-onEmphasis` sobre fundo escuro ou colorido (ex.: botões). Herdar a cor do texto `[?]`.
- **Ícones com significado fixo:**

| Significado | Ícone |
|---|---|
| informação importante (`fgColor-accent`) | `info` |
| sucesso, passou (`fgColor-success`) | `check` |
| erro, perigo (`fgColor-danger`) | `x` |
| aviso (`fgColor-attention`) | `alert` |
| commit (`fgColor-muted`) | `git-commit` |
| issue aberta ou reaberta (`fgColor-success`) | `issue-opened` · `issue-reopened` |
| issue fechada/concluída (`fgColor-done`) | `issue-closed` |
| issue rascunho (`fgColor-muted`) | `issue-draft` |

- **Notas:** regras de uso: ícone complementa texto, não o substitui; o significado precisa ficar claro sem texto; ícones com significado predefinido não mudam de significado nem de cor, salvo sobre fundo colorido; evitar ícone para conceito complexo ou abstrato; seguir os tamanhos predefinidos; testar com usuários. Ícone funcional (sem rótulo) precisa de contraste 3:1; ícone com texto ao lado é decorativo.

## 7. Layout → `docs/design/layout.md`

- **Fonte:** https://primer.style/product/primitives/size/ · https://primer.style/product/getting-started/foundations/layout/ · https://primer.style/product/components/page-layout/ · CSS do `PageLayout` em `@primer/react@38.43.0`
- **Breakpoints ou faixas de tela** (obrigatório):

| Token | Valor | O que muda |
|---|---|---|
| `breakpoint-xsmall` | 320px (20rem) | `[?]` |
| `breakpoint-small` | 544px (34rem) | `[?]` |
| `breakpoint-medium` | 768px (48rem) | abaixo: layout de **1 coluna** ("narrow"); a partir daqui: "regular" (2 colunas); largura de pane do `PageLayout` começa a aumentar |
| `breakpoint-large` | 1012px (63.25rem) | `PageLayout` passa o espaçamento para 24px; panes maiores |
| `breakpoint-xlarge` | 1280px (80rem) | `[?]` |
| `breakpoint-xxlarge` | 1400px (87.5rem) | faixa "wide" (`viewportRange-wide`): terceira coluna opcional |

  Faixas prontas: `viewportRange-narrow = (max-width: calc(48rem - 0.02px))` · `regular = (min-width: 48rem)` · `wide = (min-width: 87.5rem)` · `portrait` · `landscape` · `narrowLandscape`. O guia diz que os breakpoints não são mais a base de CSS mobile-first com `min-width`; as faixas de viewport são o alvo de projeto.
- **Largura máxima do conteúdo:** `PageLayout` `containerWidth`: **medium 768px · large 1012px · xlarge 1280px · full 100%** (valores no CSS do componente, sem token `[—]`)
- **Grid:** sem grade de colunas numerada `[—]`. A estrutura é a do `PageLayout`: Header, Content, Pane (barra lateral) e Footer. Largura do pane (a partir de 768px): small 240px, medium 256px, large 256px; a partir de 1012px: small 256px, medium 296px, large 320px. Espaço entre áreas e margem lateral: 24px a partir de 1012px (`--spacing-normal`); valor abaixo disso `[?]`.
- **Padrões de navegação que o design system oferece:** `UnderlineNav` (abas horizontais), `UnderlinePanels`, `NavList`, `ActionList`, `ActionMenu`, `Breadcrumbs`, `Pagination`, `TreeView`, `PageHeader`, `SegmentedControl`, `ActionBar` (com menu de estouro quando faltar espaço) — comportamento no menor tamanho: o guia manda quebrar layouts de várias colunas em várias telas no mobile; o comportamento de cada componente `[?]`. Padrões de navegação em https://primer.style/product/ui-patterns/navigation/ `[?]` (não lida).
- **Barra de rolagem:** nativa do sistema `[?]` — nada sobre barra personalizada nas páginas lidas
- **Notas:** `Stack` oferece fluxos horizontal e vertical responsivos (`stack-gap-*`, `stack-padding-*` = 8/16/24px). Princípios: ordem natural de leitura, foco no conteúdo, "design fully-functional responsive experiences", sem perda de função em tela pequena.

## 8. Tema → `docs/design/tema.md`

- **Fonte:** https://primer.style/product/getting-started/foundations/color-usage/ · https://primer.style/accessibility/design-guidance/color-considerations/ · `@primer/primitives@11.10.0/dist/css/functional/themes/*.css` · atributos do `<html>` do próprio primer.style
- **Modos oferecidos** (obrigatório): 2 modos (claro e escuro) e **14 temas** nos arquivos do pacote: `light` · `light-high-contrast` · `light-colorblind` · `light-colorblind-high-contrast` · `light-tritanopia` · `light-tritanopia-high-contrast` · `dark` · `dark-dimmed` · `dark-dimmed-high-contrast` · `dark-high-contrast` · `dark-colorblind` · `dark-colorblind-high-contrast` · `dark-tritanopia` · `dark-tritanopia-high-contrast`. A documentação fala em "nine themes" — divergência, ver Notas.
- **Como o modo é ativado:** atributos no `<html>`: `data-color-mode="light | dark | auto"` + `data-light-theme="<tema claro>"` + `data-dark-theme="<tema escuro>"` (no primer.style: `data-color-mode="light"`, `data-light-theme="light"`, `data-dark-theme="dark"`). O GitHub segue o modo do sistema quando o usuário não está logado. Via React `[?]` (página de React não lida).
- **Tokens semânticos** (obrigatório — claro · escuro · escuro atenuado · claro alto contraste · escuro alto contraste):

| Token | Claro | Escuro | Escuro atenuado | Claro alto contraste | Escuro alto contraste | Uso |
|---|---|---|---|---|---|---|
| `bgColor-default` | `#ffffff` | `#0d1117` | `#212830` | `#ffffff` | `#010409` | fundo da página |
| `bgColor-muted` | `#f6f8fa` | `#151b23` | `#262c36` | `#e6eaef` | `#151b23` | fundo secundário |
| `bgColor-inset` | `#f6f8fa` | `#010409` | `[?]` | `[?]` | `[?]` | fundo rebaixado |
| `bgColor-emphasis` | `#25292e` | `#3d444d` | `[?]` | `[?]` | `[?]` | fundo de ênfase neutra |
| `bgColor-disabled` | `#eff2f5` | `#212830` | `[?]` | `[?]` | `[?]` | fundo de desabilitado |
| `fgColor-default` | `#1f2328` | `#f0f6fc` | `#d1d7e0` | `#010409` | `#ffffff` | texto primário |
| `fgColor-muted` | `#59636e` | `#9198a1` | `#9198a1` | `#454c54` | `#b7bdc8` | texto secundário |
| `fgColor-disabled` | `#818b98` | `#656c76` | `[?]` | `[?]` | `[?]` | texto desabilitado |
| `fgColor-onEmphasis` | `#ffffff` | `#ffffff` | `[?]` | `[?]` | `[?]` | texto sobre fundo de ênfase |
| `borderColor-default` | `#d1d9e0` | `#3d444d` | `#3d444d` | `#454c54` | `#b7bdc8` | borda padrão |
| `borderColor-muted` | `#d1d9e0b3` | `#3d444db3` | `[?]` | `[?]` | `[?]` | divisor suave |
| `borderColor-emphasis` | `#818b98` | `#656c76` | `[?]` | `[?]` | `[?]` | borda de ênfase |
| `fgColor-accent` (= `fgColor-link`) | `#0969da` | `#4493f8` | `#478be6` | `#023b95` | `#74b9ff` | link, ação, informação |
| `bgColor-accent-emphasis` | `#0969da` | `#1f6feb` | `#316dca` | `#0349b4` | `#194fb1` | fundo de ênfase accent |
| `bgColor-accent-muted` | `#ddf4ff` | `#388bfd1a` | `[?]` | `[?]` | `[?]` | fundo de informação |
| `fgColor-success` | `#1a7f37` | `#3fb950` | `#57ab5a` | `#04591f` | `#2bd853` | texto/ícone de sucesso |
| `bgColor-success-emphasis` | `#1f883d` | `#238636` | `#347d39` | `#055d20` | `#006222` | fundo de sucesso; botão primário |
| `bgColor-success-muted` | `#dafbe1` | `#2ea04326` | `[?]` | `[?]` | `[?]` | fundo suave de sucesso |
| `fgColor-danger` | `#d1242f` | `#f85149` | `#e5534b` | `#960d1e` | `#ff9492` | texto de erro |
| `bgColor-danger-emphasis` | `#cf222e` | `#da3633` | `#c93c37` | `#a0111f` | `#ad0116` | fundo de erro |
| `bgColor-danger-muted` | `#ffebe9` | `#f851491a` | `[?]` | `[?]` | `[?]` | fundo suave de erro |
| `fgColor-attention` | `#9a6700` | `#d29922` | `#c69026` | `#603700` | `#f0b72f` | texto de aviso |
| `bgColor-attention-emphasis` | `#9a6700` | `#9e6a03` | `[?]` | `[?]` | `[?]` | fundo de aviso |
| `bgColor-attention-muted` | `#fff8c5` | `#bb800926` | `[?]` | `[?]` | `[?]` | fundo suave de aviso |
| `focus-outlineColor` (= `borderColor-accent-emphasis`) | `#0969da` | `#1f6feb` | `[?]` | `[?]` | `[?]` | anel de foco |
| `control-bgColor-rest` | `#f6f8fa` | `#212830` | `[?]` | `[?]` | `[?]` | fundo de controle em repouso |
| `overlay-backdrop-bgColor` | `#c8d1da66` | `#21283066` | `[?]` | `[?]` | `[?]` | fundo escurecido de modal |

- **Notas:** a escala neutra é invertida entre claro e escuro, o que permite os mesmos tokens funcionais nos dois. Os temas padrão Day e Night precisam cumprir WCAG; subtemas (alto contraste, daltonismo, "dark dimmed") são opção do usuário e não precisam. O GitHub **não** ativa tema de alto contraste sozinho quando o sistema está em modo de alto contraste. Não sobrescrever o modo de cores forçadas do sistema (`forced-colors`). Valores dos temas de daltonismo (colorblind, tritanopia) e de `light-high-contrast` para os demais tokens: arquivos `functional/themes/*.css` (não transcritos aqui). A documentação diz "nine themes" e o pacote traz 14 arquivos `[?]` — o motivo da diferença não consta nas páginas lidas.

## 9. Componentes → `docs/design/componentes.md`

- **Fonte:** https://primer.style/product/components/ (lista) · páginas de cada componente (Button, FormControl, TextInput, Banner, Blankslate, ConfirmationDialog, PageLayout lidas) · CSS de `@primer/react@38.43.0`
- **Catálogo** (obrigatório — importar de `@primer/react`; o nome do pacote e a regra de import estão na seção 1; "Quando usar" vem do texto do guia, `[?]` onde a página não foi lida além da descrição curta):

| Componente (nome oficial) | Variantes | Quando usar (pelo guia) |
|---|---|---|
| `Button` | `primary`, `default`, `invisible`, `danger`, `link`; tamanhos `small`, `medium` (padrão), `large`; `block`, `loading`, `inactive`, `leadingVisual`, `trailingVisual`, `trailingAction`, `count` | `primary`: ação de maior prioridade, nunca mais de um num grupo e raramente mais de um por página; `default`: ações gerais ou secundárias ao lado do primário; `invisible`: UI mínima ou componentes compostos (ex.: ActionList); `danger`: ação destrutiva, normalmente com diálogo de confirmação; `large` com parcimônia; evitar `disabled` (usar `inactive` quando um erro de sistema impede a ação) |
| `IconButton` | `[?]` | botão com ícone no lugar do texto; precisa de nome acessível |
| `ButtonGroup` | `[?]` | série de botões |
| `ActionBar` | `[?]` | IconButtons alinhados, com divisores e menu de estouro quando falta espaço |
| `ActionList` | `[?]` | lista vertical de ações ou opções, com ícone, descrição e informação lateral |
| `ActionMenu` | `[?]` | ações e seleções rápidas (ActionList + Overlay) |
| `AnchoredOverlay` | `[?]` | abre um Overlay relativo a uma âncora |
| `Autocomplete` | `[?]` | filtrar opções e escolher um ou mais valores de um campo |
| `Avatar` / `AvatarStack` | `[?]` | imagem de usuário ou organização / vários avatares em linha |
| `Banner` | `critical`, `warning`, `success`, `info`, `upsell`; dispensável (`onDismiss`); título pode ficar oculto só visualmente | destacar informação importante ou dar feedback após ação; usar com parcimônia; mensagem curta; um Banner por página por vez; `critical` = erro crítico ou ação destrutiva (ex.: erro de validação, confirmação de exclusão); `warning` = possíveis problemas (ex.: conectividade) |
| `Blankslate` | com gráfico (várias versões), `PrimaryAction`, `SecondaryAction`; borda opcional | explicar por que o conteúdo está ausente e como seguir |
| `BranchName` | `[?]` | nome de branch |
| `Breadcrumbs` | `[?]` | página atual dentro da hierarquia do site |
| `Card` | `[?]` | contêiner que agrupa conteúdo relacionado |
| `Checkbox` / `CheckboxGroup` | `[?]` | seleção única ou múltipla / conjunto de caixas |
| `Radio` / `RadioGroup` | `[?]` | uma escolha entre poucas opções excludentes |
| `CircleBadge` | `[?]` | liga logotipos de serviços de terceiros (ex.: marketplace) |
| `ConfirmationDialog` | rótulos de confirmar e cancelar personalizáveis | confirmar ação importante ou destrutiva |
| `CounterLabel` | `[?]` | contagem em itens de navegação e botões |
| `DataTable` | `[?]` | dados 2D (linha = item, coluna = dado); tem estado de carregamento próprio |
| `Details` | `[?]` | melhora o elemento nativo `<details>` |
| `Dialog` | `[?]` | superfície flutuante para conteúdo transitório |
| `FormControl` | `layout` vertical (padrão) ou horizontal (checkbox e radio); `required`, `disabled`; partes `Label`, `Caption`, `Validation` (`error` · `success`), `LeadingVisual` | campo rotulado com validação e ajuda; garante os atributos ARIA |
| `Heading` | `[?]` | estrutura hierárquica do conteúdo |
| `InlineMessage` | `[?]` | informar resultado de uma ação dentro do conteúdo |
| `KeybindingHint` | `[?]` | indica atalho de teclado disponível |
| `Label` / `LabelGroup` | `[?]` | metadado contextual / margens de grupos de Labels |
| `Link` | `[?]` | estilo de hiperlink |
| `NavList` | `[?]` | links de navegação em lista vertical |
| `Overlay` | tamanhos `xsmall`–`xlarge` (ver seção 13) | padrões de superfícies flutuantes (diálogos e menus) |
| `PageHeader` | `[?]` | título de nível superior da UI |
| `PageLayout` | áreas Header, Content, Pane, Footer; `containerWidth`; `divider`; `padding` | estrutura da página |
| `Pagination` | `[?]` | links horizontais para conteúdo paginado |
| `Popover` | `[?]` | chamar atenção para um elemento da interface |
| `ProgressBar` | `[?]` | quão completo está algo, ou partes de um todo |
| `RelativeTime` | `[?]` | tempo de forma clara, concisa e acessível |
| `SegmentedControl` | `[?]` | escolher uma entre opções próximas, aplicando de imediato |
| `Select` | `[?]` | uma opção de uma lista predefinida (dropdown) |
| `SelectPanel` | modal, ancorado, tela cheia | navegar e escolher um ou mais itens com filtro; ancorado aplica ao fechar; modal e tela cheia aplicam ao confirmar |
| `SkeletonAvatar` / `SkeletonBox` / `SkeletonText` | `[?]` | placeholders para avatar, elementos não textuais e texto em carregamento |
| `Spinner` | `[?]` | indicador de carregamento indeterminado (duração desconhecida) |
| `SplitPageLayout` | `[?]` | duas colunas: conteúdo + barra lateral |
| `Stack` | `direction` horizontal ou vertical; `gap` condensed, normal, spacious | fluxos horizontal e vertical responsivos |
| `StateLabel` | `[?]` | status de issue ou pull request |
| `Text` | `[?]` | abstração dos estilos tipográficos do Primer |
| `Textarea` | `[?]` | valor de várias linhas |
| `TextInput` | tamanhos `small` · `medium` (padrão) · `large`; `block`; `monospace`; `contrast`; com visual à esquerda e à direita e ação final | valor de uma linha de texto |
| `TextInputWithTokens` | `[?]` | campo cujo valor é uma lista |
| `Timeline` | `[?]` | itens em linha do tempo vertical |
| `ToggleSwitch` | `[?]` | liga/desliga uma configuração de imediato |
| `Token` | `[?]` | representação compacta de um objeto (metadados) |
| `Tooltip` | `[?]` | contexto extra em hover e foco de teclado |
| `TreeView` | `[?]` | lista hierárquica expansível; tem estado de carregamento próprio |
| `Truncate` | `[?]` | encurta texto com reticências |
| `UnderlineNav` / `UnderlinePanels` | `[?]` | navegação em abas / painéis em abas |

- **Estados que todo interativo tem** (obrigatório — marque os que o design system define): ✔ padrão (`-rest`) · ✔ hover (`-hover`) · ✔ foco (anel `focus-outline`) · ✔ ativo (`-active`) · ✔ desabilitado (`-disabled`) · ✔ carregando (`Button loading`) · ✔ erro (`control-borderColor-danger`, `FormControl.Validation error`) · ✔ selecionado (`-selected`, `control-checked-*`). Extra: ✔ inativo (`Button inactive`, `button-inactive-bgColor = #e6eaef`, `button-inactive-fgColor = #59636e`).
- **Notas:** tokens de botão e controle (claro) — `button-default`: fundo `rest #f6f8fa` · `hover #eff2f5` · `active #e6eaef` · `disabled #eff2f5`; borda `#d1d9e0`; texto `rest #25292e` · `disabled #818b98`. `button-primary`: fundo `rest #1f883d` · `hover #1c8139` · `active #197935` · `disabled #95d8a6`; texto `#ffffff` (`disabled #ffffffcc`); borda `#1f232826`. `button-danger`: fundo `rest #f6f8fa` · `hover #cf222e` · `active #a40e26`; texto `rest #d1242f` → `hover #ffffff`. `button-invisible`: fundo `rest #ffffff00` · `hover #818b981a` · `active #818b9826`; texto `#25292e`. `button-outline`: texto `rest #0969da` → `hover #ffffff` sobre `#0969da`. `control-bgColor`: `rest #f6f8fa` · `hover #eff2f5` · `active #e6eaef` · `disabled #eff2f5`. `control-checked-bgColor`: `rest #0969da` · `hover #0860ca` · `active #0757ba` · `disabled #818b98`. Tamanhos de controle: `control-xsmall-size = 24px · control-small-size = 28px · control-medium-size = 32px · control-large-size = 40px · control-xlarge-size = 48px`; padding horizontal `control-<tamanho>-paddingInline-condensed = 8px` e `-normal = 12px` (xsmall: 4/8/12px); padding vertical `control-small-paddingBlock = 4px · control-medium-paddingBlock = 6px · control-large-paddingBlock = 10px · control-xlarge-paddingBlock = 14px`; espaço interno `control-<tamanho>-gap = 8px` (small: 4px). O catálogo vem da página de lista; "Quando usar" por componente só foi lido onde a coluna dá detalhes.

### Campos de formulário → `docs/design/formularios.md`

- **Fonte:** https://primer.style/product/ui-patterns/forms/ · https://primer.style/product/components/form-control/ · https://primer.style/product/components/text-input/guidelines/ · CSS do TextInput (`TextInputWrapper`) em `@primer/react@38.43.0`
- **Tamanhos** (obrigatório):

| Tamanho (token) | Altura | Padding horizontal | Tamanho do texto |
|---|---|---|---|
| `small` | 28px (`min-height` = `base-size-28`) | 8px (CSS do campo; token `control-small-paddingInline-normal` = 12px) | 12px (`text-body-size-small`) |
| `medium` (padrão) | 32px (`min-height` = `base-size-32`) | `[?]` (token `control-medium-paddingInline-normal` = 12px) | 14px (`text-body-size-medium`) |
| `large` | 40px (`height` = `base-size-40`) | 8px (CSS do campo; token `control-large-paddingInline-normal` = 12px) | 14px (a regra `data-size=large` não muda o texto); 20px na variante antiga `variant=large` (`text-title-size-medium`) |

- **Estados** (obrigatório — modo claro; `[—]` onde o estado não muda aquilo):

| Estado | Borda | Fundo | Texto | Outro (sombra, contorno, ícone) |
|---|---|---|---|---|
| padrão | `control-borderColor-rest = borderColor-default = #d1d9e0`, 1px | `bgColor-default = #ffffff` | `fgColor-default = #1f2328` | `shadow-inset = inset 0 1px 0 0 #1f23280a`; raio `borderRadius-medium = 6px`; cursor texto |
| hover | `[?]` | `[?]` | `[?]` | `[?]` — o trecho do CSS lido não tem regra de hover |
| foco | `borderColor-accent-emphasis = #0969da` | `[—]` | `[—]` | contorno `2px solid borderColor-accent-emphasis`, `outline-offset: -1px` (via `:focus-within`) |
| erro | `borderColor-danger-emphasis = #cf222e` | `[—]` | `[—]` | com foco: borda e contorno de 2px em `control-borderColor-danger = #cf222e` (offset −1px); ícone junto à mensagem `[?]` |
| desabilitado | `control-borderColor-disabled = #818b981a` | `control-bgColor-disabled = #eff2f5` | `fgColor-disabled = #818b98` | `box-shadow: none`; cursor `not-allowed` |
| somente leitura | `[?]` | `[?]` | `[?]` | `[?]` |

  Sucesso (`Validation variant="success"`): borda `bgColor-success-emphasis = #1f883d`.
- **Anatomia:** rótulo **acima** (layout vertical padrão; horizontal só para checkbox e radio) · obrigatório marcado com **asterisco `*`** (`requiredText`; "deve aparecer em todo campo obrigatório, mesmo que todos os campos do formulário sejam obrigatórios"; campo marcado como obrigatório também precisa ser `required` no código; checkbox ou radio individual não pode ser obrigatório) · marcação "(opcional)" nos opcionais `[—]` · texto de ajuda: `FormControl.Caption` (curto, sem repetir o rótulo; se a mensagem de erro repetir a ajuda, esconder a ajuda) · mensagem de erro: `FormControl.Validation` (uma só mensagem por campo; para grupo de checkbox/radio, uma mensagem para o grupo) · contador de caracteres `[?]`. Regras do rótulo: descritivo e conciso (até ~3 palavras), em *sentence case*; nunca só visual (precisa existir para leitor de tela; pode ficar visualmente oculto).
- **Placeholder:** cor `fgColor-muted = #59636e` (`control-fgColor-placeholder`) · regra: só como exemplo do valor esperado, **nunca** no lugar do rótulo nem para informação crítica (some ao digitar, contraste costuma ser baixo, leitor de tela não o lê como rótulo)
- **Quando o erro aparece:** **ao enviar** (comportamento padrão da web). Depois de uma falha de envio, pode passar para validação em linha; um campo válido só é revalidado depois que o usuário o altera e tira o foco; não validar antes de o usuário terminar o campo.
- **Notas:** botão de enviar desabilitado até preencher é desencorajado. Largura do campo deve acompanhar o tamanho esperado do valor; para texto livre ou opção de uma lista, considerar `Autocomplete`. Contraste: texto 4,5:1, borda de campo 3:1 (ver seção 11).

## 10. Padrões de feedback → `docs/design/padroes-feedback.md`

- **Fonte:** https://primer.style/product/ui-patterns/loading/ · https://primer.style/product/ui-patterns/forms/ · https://primer.style/product/ui-patterns/empty-states/ · https://primer.style/product/ui-patterns/degraded-experiences/ · https://primer.style/product/components/banner/guidelines/ · https://primer.style/product/components/button/ · https://primer.style/product/components/confirmation-dialog/

| Situação | Padrão | Componente |
|---|---|---|
| carregando a página | por tempo de espera: **<1s** não mostrar indicador; **1–3s** indicador indeterminado; **3–10s** indicador determinado se possível; **>10s** determinado e, se der, tratar como tarefa em segundo plano sem bloquear o resto. Para áreas grandes, esqueleto do conteúdo no lugar do conteúdo | `Spinner`, `SkeletonText`, `SkeletonAvatar`, `SkeletonBox` |
| carregando uma ação (botão) | prop `loading`: spinner dentro do botão; marca `aria-disabled` e **preserva o foco**; com ícone, o spinner substitui só o ícone e o rótulo continua | `Button loading` (+ `loadingAnnouncement` para leitor de tela) |
| carregando lista ou tabela | usar o estado de carregamento que o componente já traz antes de criar o seu; senão esqueleto | `DataTable`, `SelectPanel`, `TreeView` (estado próprio); `Skeleton*` |
| erro de requisição | queda de serviço opcional: renderizar página degradada, sem esconder o problema; **Banner no topo da página, acima da navegação global**, variante `warning` por padrão; avisar também no contexto do elemento afetado; se a experiência principal não puder ser entregue, mostrar página de erro. Botão que não pode agir por falha de sistema: `inactive` com explicação | `Banner` (`warning`/`critical`), `InlineMessage`, `Button inactive` |
| erro de validação de campo | uma mensagem por campo, abaixo do input; mostrar ao enviar; tirar a ajuda se ficar redundante com o erro | `FormControl.Validation variant="error"` |
| lista vazia | "Blankslate": gráfico opcional, título, descrição e ação principal/secundária para o próximo passo; explicar por que está vazio; se o gráfico transmitir significado, repetir em texto | `Blankslate` |
| sucesso de uma ação | feedback depois de ação do usuário, curto e direto, em Banner ou mensagem em linha; detalhe da variante `success` `[?]` | `Banner success`, `InlineMessage` |
| confirmação de ação destrutiva | botão `danger` abre diálogo de confirmação com confirmar e cancelar | `Button variant="danger"` + `ConfirmationDialog` |

- **Notas:** Banner: usar com parcimônia, mensagem concisa, um por página por vez, título quando fizer sentido. Salvamento (https://primer.style/product/ui-patterns/saving/), notificações (…/notification-messaging/) e divulgação progressiva (…/progressive-disclosure/) não foram lidos `[?]`.

## 11. Acessibilidade visual → `docs/design/acessibilidade-visual.md`

- **Fonte:** https://primer.style/accessibility/design-guidance/color-considerations/ · https://primer.style/accessibility/design-guidance/motion-and-animation/ · https://primer.style/accessibility/design-guidance/focus-management · `@primer/primitives@11.10.0` (tokens de foco e alvo)
- **Nível alvo** (obrigatório): **WCAG 2.2 AA** (o GitHub é obrigado a cumprir os níveis A e AA; critérios citados: 1.4.1, 1.4.3, 1.4.11, 2.4.7, 2.2.2, 2.3.1)
- **Contraste mínimo:** texto normal **4,5:1** · texto grande **3:1** · elemento de interface e borda **3:1** (borda de campo contra o fundo onde está; estado selecionado; ícone funcional). Desabilitado isento. APCA é só proposta, não vale como conformidade.
- **Foco visível** (obrigatório): `focus-outline = 2px solid focus-outlineColor` (`focus-outline-width = 0.125rem = 2px`); cor `focus-outlineColor = borderColor-accent-emphasis` (`#0969da` claro, `#1f6feb` escuro); offset `focus-outline-offset = −0.125rem = −2px` (para o contorno ficar por dentro); campos de texto usam offset de −1px. Todo interativo deve ser alcançável e ativável por teclado.
- **Área mínima de toque ou clique** (obrigatório): `control-minTarget-coarse = 2.75rem = 44px` (ponteiro grosso, toque) · `control-minTarget-fine = 1rem = 16px` (mouse)
- **Movimento reduzido:** envolver transições e animações CSS em `@media (prefers-reduced-motion: no-preference)`; em JS, checar `matchMedia('(prefers-reduced-motion: reduce)')` e não iniciar animação; animação com mais de 5s precisa de botão de pausar/parar ou não pode iniciar sozinha; animação pequena de atenção para após 5s com estado final estático; nunca piscar mais de 3 vezes por segundo, nem em área grande, nem em vermelho saturado; animação que informa precisa de alternativa em texto.
- **Notas:** não usar só cor para mostrar estado; foco e contraste de campos e botões em https://primer.style/accessibility/design-guidance/color-considerations/. Texto precisa poder ser redimensionado sem cortar a interface (https://primer.style/accessibility/design-guidance/text-resize-and-spacing, não lida em detalhe `[?]`). Tokens de foco e alvo estão no pacote, não nas páginas de componente.

## 12. Movimento e transições → `docs/design/movimento-transicoes.md`

- **Fonte:** `@primer/primitives@11.10.0/dist/css/base/motion/motion.css` e `functional/motion/motion.css`. Página de documentação própria de motion em primer.style `[?]` (não encontrada nas páginas lidas); orientação de uso em https://primer.style/accessibility/design-guidance/motion-and-animation
- **Durações:** funcionais: `motion-duration-micro = 100ms` · `motion-duration-short = 200ms` · `motion-duration-medium = 300ms` · `motion-duration-long = 500ms`. Base: `base-duration-0 = 0ms` · `50ms` · `100ms` · `200ms` · `300ms` · `400ms` · `500ms` · `600ms` · `700ms` · `800ms` · `900ms` · `1000ms` (`base-duration-<n>`).
- **Easing:** `motion-easing-enter = base-easing-easeOut = cubic-bezier(0.3, 0.8, 0.6, 1)` · `motion-easing-exit = base-easing-easeIn = cubic-bezier(0.7, 0.1, 0.75, 0.9)` · `motion-easing-hover = base-easing-ease = cubic-bezier(0.25, 0.1, 0.25, 1)` · `motion-easing-move = base-easing-easeInOut = cubic-bezier(0.6, 0, 0.2, 1)` · `motion-easing-linear = base-easing-linear = cubic-bezier(0, 0, 1, 1)`. Combinações prontas: `motion-transition-enter = 300ms + enter` · `motion-transition-exit = 200ms + exit` · `motion-transition-hover = 100ms + hover` · `motion-transition-stateChange = 200ms + move`.
- **Notas:** respeitar `prefers-reduced-motion` (seção 11). Movimento "de design" (decorativo) é o mais arriscado e deve ser raro.

## 13. Camadas e sobreposições → `docs/design/camadas-sobreposicoes.md`

- **Fonte:** `@primer/primitives@11.10.0/dist/css/functional/size/z-index.css` · https://primer.style/product/primitives/size/ (seção Overlay) · https://primer.style/product/primitives/color/ (seção Overlay)

| Camada | z-index | Componentes |
|---|---|---|
| `zIndex-behind` | -1 | `[?]` |
| `zIndex-default` | 0 | conteúdo normal |
| `zIndex-sticky` | 100 | elementos fixos ao rolar (pelo nome) `[?]` |
| `zIndex-dropdown` | 200 | menus suspensos (pelo nome) `[?]` |
| `zIndex-overlay` | 300 | `Overlay` (pelo nome) `[?]` |
| `zIndex-modal` | 400 | `Dialog`, `ConfirmationDialog` (pelo nome) `[?]` |
| `zIndex-popover` | 500 | `Popover`, `Tooltip` (pelo nome) `[?]` |
| `zIndex-skipLink` | 600 | link "pular para o conteúdo" (pelo nome) |

- **Modal:** fundo escurecido `overlay-backdrop-bgColor` (`#c8d1da66` claro · `#21283066` escuro) · fecha com Esc `[?]` · fecha com clique fora `[?]` · prende o foco `[?]` (páginas do Dialog e de acessibilidade do Dialog não foram lidas)
- **Notas:** tokens de `Overlay`: `overlay-bgColor = #ffffff` (claro) · `overlay-borderColor = #d1d9e080` · `overlay-borderRadius = 6px` · `overlay-offset = 4px` · largura `xsmall 192px · small 320px · medium 480px · large 640px · xlarge 960px` · altura `small 256px · medium 320px · large 432px · xlarge 600px` · padding `condensed 8px · normal 16px` · padding vertical `condensed 4px · normal 12px`.

## 14. Lacunas

- **O que o design system não define** e cada projeto terá de decidir: plataformas mobile nativas (iOS, Android) e desktop; barra de rolagem personalizada; espaçamento entre letras (nenhum token); números tabulares; grade de colunas numerada e margens laterais abaixo de 1012px; documentação de movimento própria; ordem de uso das cores de dados.
- **Qualquer coisa que não cabia nos campos:**
  - **Divergência de contagem de temas:** a página Color usage fala em "nine themes", o pacote `@primer/primitives` 11.10.0 publica 14 arquivos de tema.
  - **Paleta base:** as escalas numéricas por cor (tons) existem, mas os valores de cada passo não estão nas páginas lidas; só os tokens funcionais foram copiados.
  - **Fonte Mona Sans:** a pilha tipográfica dos tokens já começa em `Mona Sans VF`; como a fonte é carregada e hospedada em outro projeto não está documentado nas páginas lidas.
  - **Botão primário verde:** o Primer usa verde de sucesso para a ação primária, e azul (accent) para links, foco e seleção.
  - **Páginas relevantes não lidas** (daí vários `[?]`): Content, Responsive, Data visualization, Navigation, Saving, Notification messaging, Progressive disclosure, página de React e Rails (como importar), página de código dos Octicons, páginas Dialog, Overlay, SelectPanel, Tooltip e a maior parte das páginas de guidelines por componente.
  - **Primer Brand UI** (`@primer/react-brand`, marketing) é outro conjunto de tokens e componentes; não faz parte desta captura.
  - **Hover de campo, estado somente leitura e ícone de erro:** não aparecem no trecho do CSS e das páginas lidos; confirmar no Storybook (https://primer.style/product/components/text-input/).
