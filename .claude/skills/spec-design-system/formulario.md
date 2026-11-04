# Formulário de especificação do design system — <nome do produto>

<!-- LLM-ONLY: este arquivo é o MODELO entregue ao usuário; não vai para docs/. Ao emitir:
  copie-o para `.temp/design-system-formulario.md`, remova todo comentário LLM-ONLY, troque
  <nome do produto> e pré-preencha cada campo que alguma fonte respondeu (Etapa 1), com a
  origem na frente do valor. Campo sem fonte fica `[?]`, com a sua recomendação entre
  parênteses quando você tem uma. Nunca entregue o modelo cru. -->

> Especificação de entrada do design system. Preenchido por quem conhece o produto — pessoa ou
> agente — e devolvido a `/spec-design-system`, que o valida, pergunta o que restar e escreve
> `docs/design-system.md` e os arquivos de `docs/design/`.
>
> Campo já preenchido veio de uma fonte real: `[código] <caminho>` (o que o código tem hoje),
> `[referência] <link>` (Figma ou documento do time) ou `[decisão]` (acertado com o usuário).
> O que você escrever entra como `[decisão]`.

## 0. Como preencher

**Notação** — use exatamente estes marcadores:

| Marcador | Significa |
|---|---|
| `<valor>` | o valor, com unidade explícita: `#0F62FE`, `16px`, `1rem`, `200ms`, `1.5` |
| `[?]` | não sei ou não decidido → o agente decide e explica a recomendação |
| `[—]` | não se aplica a este produto |
| `[=]` | confirmo o valor que já está no campo (derivado do código ou da referência) |

**Regras**

1. **Não apague** campo, seção ou coluna, e **não acrescente** seção. O que não couber em um
   campo vai para **Notas** da seção; o que não couber em seção nenhuma vai para a seção 15.
2. **Não invente.** `[?]` é mais barato que um valor errado: o errado entra no documento como
   verdade do produto, e toda tela nova o segue.
3. **Campo com `[código]`** é o que está implementado. Confirme com `[=]`, ou escreva o valor
   novo e diga no mesmo campo o que fazer com a diferença: `corrigir o código` ou
   `manter, é divergência conhecida`.
4. **Tabelas:** apague as linhas marcadas `(exemplo)` e deixe só as reais. Linha sem valor é
   linha apagada.
5. **Só o global.** Tela, fluxo e estado de uma feature específica não entram aqui: ficam no
   plano da mudança. Campo que só tem a resposta "depende da tela" recebe `[—]`.
6. **Decisão transversal com trade-off real** (biblioteca de UI, dark mode desde já, skeleton em
   todo o produto) vai também na seção 15, como candidata a ADR.
7. **Obrigatório em branco bloqueia:** o agente volta a perguntar antes de escrever qualquer
   arquivo. Os campos `(obrigatório)` são o mínimo para o design system existir.

Cada seção aponta o arquivo que ela alimenta. Seção `(opcional)` só é preenchida quando o produto
tem o assunto; `(só projeto existente)`, só quando já há interface implementada.

---

## 1. Produto e contexto → `docs/design-system.md`

- **Nome do produto** (obrigatório):
- **O que é, em uma frase** (obrigatório):
- **Domínio / setor:**
- **Usuário principal e contexto de uso** (obrigatório): <quem, em qual dispositivo, por quanto
  tempo, em que ambiente — ex.: operador interno, desktop 1920×1080, 8h/dia, sala com luz forte>
- **Plataformas-alvo** (obrigatório): web · mobile iOS · mobile Android · desktop Windows ·
  desktop macOS · desktop Linux · TV · outra
- **Resolução ou tamanho dominante:**
- **Stack de UI** (obrigatório): <framework + linguagem + camada de estilo — ex.: React 19 +
  TypeScript + Tailwind 4>
- **Biblioteca de componentes base** (obrigatório): <nome e versão, ou `nenhuma (componentes
  próprios)`>
- **Maturidade** (obrigatório): projeto novo · já tem interface implementada · redesign de
  produto existente
- **Idioma(s) da interface e direção do texto:** <ex.: pt-BR, LTR>
- **Restrições externas:** <norma de acessibilidade exigida por contrato ou lei, diretriz de
  plataforma (HIG, Material), identidade corporativa já definida, limite de peso ou de
  performance, navegador ou SO mínimo suportado>
- **Produtos de referência:** <produto — o que imitar nele>
- **Anti-referência:** <produto ou padrão — o que não repetir>
- **ADRs aceitas que restringem a UI:** <NNN — o que decidem> | nenhuma
- **Notas:**

## 2. Direção visual → `docs/design-system.md` (seção Direção)

- **Personalidade em 3 adjetivos** (obrigatório): <ex.: sóbrio, preciso, silencioso>
- **Eixos** (obrigatório) — marque um ponto de 1 a 5 em cada:
  - sóbrio `1 2 3 4 5` expressivo:
  - denso `1 2 3 4 5` espaçado:
  - geométrico `1 2 3 4 5` orgânico:
  - plano `1 2 3 4 5` com profundidade:
- **Densidade de informação** (obrigatório): compacta · confortável · ampla — e onde ela muda
  (ex.: ampla em formulário, compacta em tabela):
- **Marca já existente:** logo <sim/não, formatos> · cor da marca <valor> · fontes da marca
  <nomes> · manual da marca <sim/não>
- **O que o produto evita** (obrigatório): <ex.: gradiente, ilustração, ícone colorido,
  MAIÚSCULAS em botão, sombra forte>
- **Tom do microcopy:** formal · neutro · direto · próximo — e a pessoa usada (você · nós ·
  impessoal):
- **Notas:**

## 3. Cores → `docs/design/cores.md`

- **Modelo da paleta** (obrigatório): tons por cor (ex.: 50–900) · cores chapadas · herdada da
  biblioteca base
- **Primária** (obrigatório): `<valor>` — onde se usa:
- **Primária: onde NÃO pode entrar:** <ex.: fundo de página, texto corrido>
- **Secundária ou acento:** `<valor>` — onde se usa:
- **Neutros** (obrigatório):

| Papel | Valor | Uso |
|---|---|---|
| fundo da página | | |
| superfície (card, painel) | | |
| superfície elevada | | |
| borda e divisória | | |
| texto primário | | |
| texto secundário | | |
| texto desabilitado | | |

- **Semânticos** (obrigatório):

| Papel | Valor | Variante de fundo | Uso |
|---|---|---|---|
| sucesso | | | |
| aviso | | | |
| erro | | | |
| informação | | | |

- **Cores de dados (gráficos):** precisa <sim/não> · quantas séries · ordem de uso ·
  distinguíveis em escala de cinza <sim/não>
- **Contraste alvo** (obrigatório): WCAG AA (4,5:1) · AAA (7:1) · outro
- **Regras de uso:** <cores que não podem conviver, par proibido, uso de cor em estado>
- **Notas:**

## 4. Tipografia → `docs/design/tipografia.md`

- **Família de texto** (obrigatório): `<nome>` — de onde vem <Google Fonts · pacote · fonte do
  sistema · arquivo próprio> — pesos usados:
- **Família de título:** `<nome>` | mesma do texto
- **Família monoespaçada:** `<nome>` | `[—]`
- **Fallback stack** (obrigatório): <ex.: system-ui, -apple-system, Segoe UI, sans-serif>
- **Escala** (obrigatório): base `<valor>` · razão `<valor>`, ou a lista explícita de tamanhos:
- **Unidade do tamanho de texto** (obrigatório): px · rem · sp · pt
- **Estilos nomeados** (obrigatório — apague os que o produto não tem):

| Estilo | Tamanho | Peso | Altura de linha | Espaçamento entre letras | Uso |
|---|---|---|---|---|---|
| display | | | | | |
| título 1 | | | | | |
| título 2 | | | | | |
| título 3 | | | | | |
| corpo | | | | | |
| corpo pequeno | | | | | |
| rótulo | | | | | |
| legenda ou auxiliar | | | | | |
| código | | | | | |

- **Largura máxima de linha de leitura:** <ex.: 72 caracteres>
- **MAIÚSCULAS em botão ou rótulo:** sim · não
- **Números tabulares em tabela e em valor monetário:** sim · não
- **Notas:**

## 5. Espaçamento, raios e sombras → `docs/design/espacamento-raios-sombras.md`

- **Base da escala de espaçamento** (obrigatório): 4 · 8 · outra — valores:
- **Unidade** (obrigatório): px · rem · dp · pt
- **Raios** (obrigatório):

| Token | Valor | Onde se usa |
|---|---|---|
| nenhum | `0` | |
| pequeno | | |
| médio | | |
| grande | | |
| pill ou círculo | | |

- **Espessura de borda padrão:** `<valor>`
- **Estratégia de elevação** (obrigatório): sombra · borda em vez de sombra · fundo diferente ·
  nenhuma
- **Níveis de sombra:** <nível — valor — para quê>
- **Padding padrão por elemento:** botão `<valor>` · card `<valor>` · célula de tabela
  `<valor>` · container da página `<valor>` (o do campo fica em Campos de formulário, seção 9)
- **Espaço vertical entre blocos:** <ex.: 24px entre seções, 8px entre rótulo e campo>
- **Notas:**

## 6. Ícones → `docs/design/icones.md`

- **Biblioteca** (obrigatório): `<nome e versão>` · SVG próprio · ambos (com a regra de quando
  usar cada)
- **Estilo** (obrigatório): outline · preenchido · duotone — espessura do traço:
- **Tamanhos** (obrigatório): <lista> — padrão:
- **Cor:** herda a cor do texto · token próprio `<qual>`
- **Alinhamento com o texto:** <ex.: ícone de 16px alinhado ao baseline do corpo>
- **Ícones com significado fixo** (obrigatório — só os que o produto usa sempre assim):

| Significado | Ícone |
|---|---|
| excluir | |
| editar | |
| adicionar | |
| buscar | |
| fechar | |
| mais opções | |

- **Ícone sem rótulo visível exige rótulo acessível:** sim · não
- **Notas:**

## 7. Layout → `docs/design/layout.md`

- **Faixas de tamanho de tela** (obrigatório — breakpoints na web; celular, tablet e orientação
  em mobile; janela mínima no desktop):

| Faixa | Valor | O que muda |
|---|---|---|
| (exemplo) celular | < 640px | navegação vira drawer; tabela vira lista de cards |

- **Largura máxima do conteúdo:** `<valor>`
- **Grid:** colunas `<n>` · espaço entre colunas `<valor>` · margem lateral `<valor>`
- **Estrutura de navegação** (obrigatório): barra lateral · barra superior · abas · drawer ·
  barra inferior — o que ela contém e como se comporta no menor tamanho:
- **Regiões fixas:** <header, toolbar, footer — quais ficam fixas na rolagem>
- **Rolagem:** página inteira · só a área de conteúdo
- **Tamanho mínimo de janela (desktop) ou safe area (mobile):**
- **Zoom ou escala do SO suportada:** <ex.: até 200% sem quebra>
- **Barra de rolagem:** nativa do sistema · personalizada — se personalizada: largura `<valor>` ·
  cor da barra e do trilho por tema `<tokens>` · aparece <sempre · só ao rolar>
- **Notas:**

## 8. Tema → `docs/design/tema.md`

- **Modos** (obrigatório): só claro · só escuro · claro e escuro · segue o sistema operacional
- **Modo padrão quando não há preferência:**
- **Estratégia** (obrigatório): tokens semânticos sobre a paleta · dois conjuntos de valores ·
  inversão calculada
- **Nomenclatura dos tokens semânticos** (obrigatório): <padrão de nome, ex.:
  `--color-<papel>-<variação>`; papéis: bg, surface, fg, fg-muted, border, accent>
- **Tokens semânticos** (obrigatório — um por papel que o produto usa):

| Token | Claro | Escuro | Uso |
|---|---|---|---|
| (exemplo) `--color-bg-default` | | | fundo da página |

- **Onde a escolha do usuário é guardada:** <ex.: localStorage, perfil no servidor> | `[—]`
- **Alto contraste previsto:** sim · não
- **Notas:**

## 9. Componentes → `docs/design/componentes.md`

- **Uso da biblioteca base** (obrigatório): usar como vem · tematizar via tokens · copiar para o
  repo (estilo shadcn) · só como referência
- **Caminho e convenção de import** (obrigatório): <ex.: `@/components/ui/<nome>`, PascalCase,
  um componente por arquivo>
- **Inventário** (obrigatório) — `Precisa?` = sim (nesta leva) · depois · não. Apague as linhas
  `não` ao devolver:

| Componente | Precisa? | Variantes | Observação |
|---|---|---|---|
| botão | | | |
| link | | | |
| campo de texto | | | |
| área de texto | | | |
| select | | | |
| combobox ou autocomplete | | | |
| checkbox | | | |
| radio | | | |
| switch | | | |
| slider | | | |
| seletor de data | | | |
| upload de arquivo | | | |
| formulário (layout, rótulo, erro, ajuda) | | | |
| tabela de dados | | | |
| lista | | | |
| card | | | |
| badge ou tag | | | |
| avatar | | | |
| breadcrumb | | | |
| abas | | | |
| paginação | | | |
| menu ou dropdown | | | |
| tooltip | | | |
| popover | | | |
| modal | | | |
| drawer | | | |
| toast | | | |
| alerta inline | | | |
| banner de página | | | |
| barra de progresso | | | |
| spinner | | | |
| skeleton | | | |
| estado vazio | | | |
| stepper | | | |
| acordeão | | | |
| busca | | | |
| filtros | | | |
| gráfico | | | |

- **Variantes de botão** (obrigatório): <quantas e quando usar cada — ex.: primário (ação
  principal, uma por tela), secundário, texto, destrutivo>
- **Estados que todo interativo precisa ter** (obrigatório): padrão · hover · foco · ativo ·
  desabilitado · carregando · erro · selecionado — marque os que o produto usa:
- **O que NÃO deve virar componente próprio:** <ex.: layout de uma tela, wrapper de uma só
  chamada>
- **Notas:**

### Campos de formulário → `docs/design/formularios.md`

Preencha quando o inventário marca `campo de texto` como `sim`; senão, `[—]` no bloco inteiro.

- **Tamanhos** (obrigatório):

| Tamanho | Altura | Padding horizontal | Tamanho do texto |
|---|---|---|---|
| (exemplo) médio | 40px | 12px | 16px |

- **Estados** (obrigatório — token ou valor em cada célula; `[—]` onde o estado não muda aquilo):

| Estado | Borda | Fundo | Texto | Outro (sombra, contorno, ícone) |
|---|---|---|---|---|
| padrão | | | | |
| hover | | | | |
| foco | | | | |
| erro | | | | |
| desabilitado | | | | |
| somente leitura | | | | |

- **Anatomia** (obrigatório): rótulo <acima · ao lado · flutuante> · obrigatório marcado
  <asterisco · "(opcional)" nos opcionais> · texto de ajuda <onde> · mensagem de erro <onde, com
  ícone?> · contador de caracteres <sim/não>
- **Placeholder:** cor `<token>` · regra de uso <ex.: nunca no lugar do rótulo>
- **Quando o erro aparece:** ao sair do campo · ao enviar · ao digitar
- **Notas:**

## 10. Padrões de feedback → `docs/design/padroes-feedback.md`

Padrão **concreto**, nunca "mostrar loading": diga o quê, onde e com qual saída. Apague a linha
da situação que não existe no produto.

| Situação | Padrão | Componente |
|---|---|---|
| (exemplo) carregando a lista | skeleton de 5 linhas no lugar da tabela, por no mínimo 300ms | `Skeleton` |
| carregando a página | | |
| carregando uma ação (botão) | | |
| carregando lista ou tabela | | |
| atualizando dado já na tela | | |
| erro de requisição | | |
| erro de validação de campo | | |
| erro de permissão | | |
| não encontrado (404) | | |
| lista vazia no primeiro uso | | |
| busca ou filtro sem resultado | | |
| sucesso de uma ação | | |
| confirmação de ação destrutiva | | |
| operação longa (acima de 5s) | | |
| sem conexão | | |
| sessão expirada | | |
| salvamento automático | | |

- **Toast:** posição `<onde>` · duração `<valor>` · empilha <sim/não> · fecha sozinho <sim/não>
- **Mensagem de erro** (obrigatório): linguagem <técnica · amigável> · mostra código ou id de
  correlação <sim/não> · sempre oferece a próxima ação <sim/não>
- **Notas:**

## 11. Acessibilidade visual → `docs/design/acessibilidade-visual.md`

- **Nível alvo** (obrigatório): WCAG 2.2 A · AA · AAA · nenhum formal
- **Contraste mínimo** (obrigatório): texto normal `<valor>` · texto grande `<valor>` ·
  elemento de interface e borda `<valor>`
- **Foco visível** (obrigatório): <cor, espessura, offset> · só na navegação por teclado
  <sim/não>
- **Área mínima de toque ou clique** (obrigatório): `<valor>` (ex.: 44×44)
- **Cor nunca é o único portador de significado** — como se reforça: <ícone, texto, padrão de
  preenchimento>
- **Movimento reduzido** (obrigatório): <o que muda com `prefers-reduced-motion` ou o
  equivalente da plataforma>
- **Tamanho mínimo de fonte:** `<valor>`
- **Navegação por teclado:** ordem de foco <regra> · skip link <sim/não> · atalhos <quais>
- **Leitor de tela:** rótulo em campo e em ícone <regra> · região viva para o feedback <sim/não>
- **Auditoria:** ferramenta `<qual>` · quando roda <ex.: no CI, em cada PR> | `[—]`
- **Notas:**

## 12. Movimento e transições (opcional) → `docs/design/movimento-transicoes.md`

- **O produto usa animação:** sim · não — onde:
- **Durações nomeadas:** rápida `<valor>` · média `<valor>` · lenta `<valor>`
- **Easing:** <função por tipo de movimento: entrada, saída, movimento na tela>
- **O que nunca anima:** <ex.: valor numérico, conteúdo de tabela>
- **Com movimento reduzido:** <o que acontece: corta a animação, mantém só a opacidade>
- **Notas:**

## 13. Camadas e sobreposições (opcional) → `docs/design/camadas-sobreposicoes.md`

| Camada | z-index | Componentes |
|---|---|---|
| (exemplo) conteúdo | 0 | |
| (exemplo) toast | 400 | |

- **Modal:** fundo escurecido `<valor>` · fecha com Esc <sim/não> · fecha com clique fora
  <sim/não> · prende o foco <sim/não> · trava a rolagem do fundo <sim/não>
- **Mais de um modal ao mesmo tempo:** permitido · proibido
- **Quem fica por cima:** <ex.: toast acima de modal>
- **Notas:**

## 14. Divergências conhecidas (só projeto existente) → `docs/design/divergencias-conhecidas.md`

Onde o código de hoje não segue o que este formulário define. `Decisão` = `corrigir o código` ou
`manter como divergência conhecida`.

| Onde (caminho) | O que diverge | O que vale | Decisão |
|---|---|---|---|

## 15. Lacunas e decisões delegadas

- **Campos deixados em `[?]` de propósito** (o agente decide e recomenda): <campo — por quê>
- **Candidatas a ADR** (decisão transversal, cara de mudar, com alternativa real): <título —
  o que está em jogo>
- **Perguntas para o agente:** <o que você quer que ele proponha antes de escrever>
- **Fora do escopo deste design system:** <o que o time decidiu não padronizar agora>
- **Qualquer coisa que não cabia nos campos:**
