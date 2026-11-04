# Modelo de design system — <nome do design system>

<!-- LLM-ONLY: este arquivo é o MODELO do formulário de clonagem; não vai para docs/. Ao emitir:
  copie-o para `.temp/design-system-modelo-formulario.md`, remova todo comentário LLM-ONLY e
  troque <nome do design system>. Entregue-o EM BRANCO: quem preenche é o usuário, a partir da
  documentação oficial do design system. Ao salvar (Etapa 7): grave em
  `modelos/<slug>.md` sem a seção 0 e sem comentários LLM-ONLY. -->

> Retrato de um design system que já existe (ex.: o Primer, do GitHub), guardado na
> `/spec-design-system` para ser aplicado em qualquer projeto com
> `/spec-design-system modelo <slug>`. Descreve **o design system**, não um produto: nome,
> público e plataformas do projeto são perguntados quando o modelo é aplicado.

## 0. Como preencher

**Notação**

| Marcador | Significa |
|---|---|
| `<valor>` | o valor, com unidade explícita e o **nome oficial do token** quando existe: `fgColor-default = #1F2328`, `16px`, `200ms` |
| `[?]` | a fonte não diz → ao aplicar, a skill decide pelo default consagrado ou pergunta |
| `[—]` | o design system não define isso |

**Regras**

1. **Copie da fonte oficial, não de memória.** Cada seção tem o campo **Fonte**: o link da página
   de onde saíram os valores. Valor sem fonte é `[?]`: o errado entra em todo projeto que usar o
   modelo.
2. **Não apague** campo nem seção. O que não couber vai para **Notas** da seção ou para a seção 14.
3. **Tabelas:** apague as linhas marcadas `(exemplo)` e acrescente quantas linhas reais precisar.
4. **Só o que o design system define para todos.** Tela de um produto específico do dono (o
   repositório do GitHub, a página de um app) não entra.
5. **Tema com mais de um modo:** cada token semântico leva o valor de cada modo que o design
   system oferece.

Cada seção aponta o arquivo de `docs/design/` que ela alimenta quando o modelo é aplicado.

---

## 1. Identificação → `docs/design-system.md`

- **Nome** (obrigatório): <ex.: Primer>
- **Slug** (obrigatório): <kebab-case, sem acentos; vira o nome do arquivo e o argumento — ex.:
  `github-primer`>
- **Mantido por** (obrigatório): <organização>
- **Link oficial** (obrigatório): <página inicial da documentação>
- **Versão capturada** (obrigatório): <versão do design system ou do pacote de tokens>
- **Data da captura** (obrigatório): AAAA-MM-DD
- **Licença** (obrigatório): <licença do código e dos tokens; fontes ou ícones com licença
  própria, à parte; uso fora da organização dona permitido <sim/não>>
- **Plataformas cobertas** (obrigatório): web · mobile iOS · mobile Android · desktop · outra
- **Pacote de tokens:** `<nome>` — formato <variáveis CSS · JSON · SCSS · outro> | `[—]`
- **Implementações oficiais** (obrigatório — uma linha por stack; nenhuma → só a linha `nenhuma`):

| Stack | Pacote | Versão | Como importar |
|---|---|---|---|
| (exemplo) React | `@primer/react` | 37.x | `import {Button} from '@primer/react'` |

- **Notas:**

## 2. Direção visual → `docs/design-system.md` (seção Direção)

- **Fonte:**
- **Personalidade em 3 adjetivos** (obrigatório):
- **Eixos** (obrigatório) — marque um ponto de 1 a 5 em cada:
  - sóbrio `1 2 3 4 5` expressivo:
  - denso `1 2 3 4 5` espaçado:
  - geométrico `1 2 3 4 5` orgânico:
  - plano `1 2 3 4 5` com profundidade:
- **Densidade de informação:** compacta · confortável · ampla — e onde muda:
- **O que o design system evita** (obrigatório): <ex.: gradiente, sombra forte, MAIÚSCULAS em botão>
- **Tom do microcopy:** <o guia de conteúdo dele, se tiver> | `[—]`
- **Notas:**

## 3. Cores → `docs/design/cores.md`

- **Fonte:**
- **Modelo da paleta** (obrigatório): tons por cor (ex.: 0–9) · cores chapadas
- **Convenção de nome dos tokens** (obrigatório): <ex.: base `base-color-<cor>-<tom>`,
  funcional `<papel>Color-<variação>`>
- **Paleta base** (obrigatório — uma linha por cor, no modo claro):

| Cor | Tons (token = valor) |
|---|---|
| (exemplo) azul | `blue-0 = #DDF4FF` … `blue-9 = #0A3069` |

- **Primária ou acento:** `<token>` — onde se usa:
- **Cores de dados (gráficos):** <tokens e ordem de uso> | `[—]`
- **Contraste alvo** (obrigatório): WCAG AA · AAA · outro
- **Regras de uso:** <o que o guia proíbe ou exige>
- **Notas:**

## 4. Tipografia → `docs/design/tipografia.md`

- **Fonte:**
- **Família de texto** (obrigatório): `<nome>` — de onde vem <Google Fonts · pacote · fonte do
  sistema · arquivo próprio> — licença da fonte:
- **Família de título:** `<nome>` | mesma do texto
- **Família monoespaçada:** `<nome>` | `[—]`
- **Fallback stack** (obrigatório):
- **Unidade do tamanho de texto** (obrigatório): px · rem · sp · pt
- **Estilos nomeados** (obrigatório):

| Estilo (token oficial) | Tamanho | Peso | Altura de linha | Espaçamento entre letras | Uso |
|---|---|---|---|---|---|
| (exemplo) `text-body-medium` | 14px | 400 | 20px | 0 | texto corrido |

- **Números tabulares:** onde o guia manda usar | `[—]`
- **Notas:**

## 5. Espaçamento, raios e sombras → `docs/design/espacamento-raios-sombras.md`

- **Fonte:**
- **Escala de espaçamento** (obrigatório):

| Token | Valor |
|---|---|
| (exemplo) `base-size-8` | 8px |

- **Raios** (obrigatório):

| Token | Valor | Onde se usa |
|---|---|---|
| (exemplo) `borderRadius-medium` | 6px | botão, campo, card |

- **Espessura de borda padrão:** `<token = valor>`
- **Estratégia de elevação** (obrigatório): sombra · borda em vez de sombra · fundo diferente ·
  nenhuma
- **Sombras:**

| Token | Valor | Para quê |
|---|---|---|

- **Notas:**

## 6. Ícones → `docs/design/icones.md`

- **Fonte:**
- **Biblioteca** (obrigatório): `<nome>` — pacote por stack:
- **Estilo** (obrigatório): outline · preenchido · duotone — espessura do traço:
- **Tamanhos** (obrigatório): <lista> — padrão:
- **Cor:** herda a cor do texto · token próprio `<qual>`
- **Ícones com significado fixo:**

| Significado | Ícone |
|---|---|
| (exemplo) excluir | `trash` |

- **Notas:**

## 7. Layout → `docs/design/layout.md`

- **Fonte:**
- **Breakpoints ou faixas de tela** (obrigatório):

| Token | Valor | O que muda |
|---|---|---|
| (exemplo) `breakpoint-medium` | 768px | |

- **Largura máxima do conteúdo:** `<token = valor>`
- **Grid:** colunas · espaço entre colunas · margem lateral
- **Padrões de navegação que o design system oferece:** <ex.: barra superior, navegação lateral,
  abas> — comportamento no menor tamanho:
- **Barra de rolagem:** nativa do sistema · personalizada — se personalizada: largura `<valor>` ·
  cor da barra e do trilho por modo `<tokens>` · aparece <sempre · só ao rolar>
- **Notas:**

## 8. Tema → `docs/design/tema.md`

- **Fonte:**
- **Modos oferecidos** (obrigatório): <ex.: claro, escuro, escuro atenuado, alto contraste,
  daltonismo>
- **Como o modo é ativado:** <ex.: atributo `data-color-mode`, classe, prop do provider>
- **Tokens semânticos** (obrigatório — um por papel; uma coluna por modo que o design system
  oferece, acrescente as que faltarem):

| Token | Claro | Escuro | Uso |
|---|---|---|---|
| (exemplo) `bgColor-default` | `#FFFFFF` | `#0D1117` | fundo da página |
| (exemplo) `fgColor-default` | `#1F2328` | `#E6EDF3` | texto primário |
| (exemplo) `bgColor-success-emphasis` | | | fundo de sucesso |
| (exemplo) `fgColor-danger` | | | texto de erro |

- **Notas:**

## 9. Componentes → `docs/design/componentes.md`

- **Fonte:**
- **Catálogo** (obrigatório — os componentes que o design system documenta; o import vem da
  seção 1):

| Componente (nome oficial) | Variantes | Quando usar (pelo guia) |
|---|---|---|
| (exemplo) `Button` | primary, secondary, invisible, danger | primary: a ação principal, uma por tela |

- **Estados que todo interativo tem** (obrigatório): padrão · hover · foco · ativo ·
  desabilitado · carregando · erro · selecionado — marque os que o design system define:
- **Notas:**

### Campos de formulário → `docs/design/formularios.md`

- **Fonte:**
- **Tamanhos** (obrigatório):

| Tamanho (token) | Altura | Padding horizontal | Tamanho do texto |
|---|---|---|---|
| (exemplo) `medium` | 32px | 12px | 14px |

- **Estados** (obrigatório — token ou valor em cada célula; `[—]` onde o estado não muda aquilo):

| Estado | Borda | Fundo | Texto | Outro (sombra, contorno, ícone) |
|---|---|---|---|---|
| padrão | | | | |
| hover | | | | |
| foco | | | | |
| erro | | | | |
| desabilitado | | | | |
| somente leitura | | | | |

- **Anatomia:** rótulo <acima · ao lado · flutuante> · obrigatório marcado <asterisco · "(opcional)"
  nos opcionais> · texto de ajuda <onde> · mensagem de erro <onde, com ícone?> · contador de
  caracteres <sim/não>
- **Placeholder:** cor `<token>` · regra de uso <ex.: nunca no lugar do rótulo>
- **Quando o erro aparece:** ao sair do campo · ao enviar · ao digitar
- **Notas:**

## 10. Padrões de feedback → `docs/design/padroes-feedback.md`

- **Fonte:**

Só as situações que o design system orienta; apague as outras. Padrão **concreto**: o quê,
onde e com qual saída.

| Situação | Padrão | Componente |
|---|---|---|
| carregando a página | | |
| carregando uma ação (botão) | | |
| carregando lista ou tabela | | |
| erro de requisição | | |
| erro de validação de campo | | |
| lista vazia | | |
| sucesso de uma ação | | |
| confirmação de ação destrutiva | | |

- **Notas:**

## 11. Acessibilidade visual → `docs/design/acessibilidade-visual.md`

- **Fonte:**
- **Nível alvo** (obrigatório): WCAG 2.2 A · AA · AAA
- **Contraste mínimo:** texto normal · texto grande · elemento de interface e borda
- **Foco visível** (obrigatório): <cor (token), espessura, offset>
- **Área mínima de toque ou clique** (obrigatório): `<valor>`
- **Movimento reduzido:** <o que o design system muda>
- **Notas:**

## 12. Movimento e transições → `docs/design/movimento-transicoes.md`

- **Fonte:** | `[—]` se o design system não define movimento
- **Durações:** `<token = valor>` …
- **Easing:** `<token = função>` …
- **Notas:**

## 13. Camadas e sobreposições → `docs/design/camadas-sobreposicoes.md`

- **Fonte:** | `[—]` se o design system não define camadas

| Camada | z-index | Componentes |
|---|---|---|

- **Modal:** fundo escurecido `<token>` · fecha com Esc · fecha com clique fora · prende o foco
- **Notas:**

## 14. Lacunas

- **O que o design system não define** e cada projeto terá de decidir:
- **Qualquer coisa que não cabia nos campos:**
