---
name: spec-design-system
description: Utilitário do framework (obrigatório em projeto com interface visual). Cria ou atualiza o design system do produto - docs/design-system.md (contexto geral e índice) e um arquivo por categoria em docs/design/ (cores, tipografia, espaçamento, ícones, layout, tema, componentes, padrões de feedback, acessibilidade visual e, sob demanda, botões, formulários, camadas, movimento e divergências). Deriva do código, do Figma e dos documentos do time tudo o que já existe e só pergunta o resto. Em projeto novo, parte da proposta do produto, tira as dúvidas de direção e entrega um formulário técnico de especificação, pré-preenchido com o que derivou e com a recomendação em cada campo aberto, para o usuário (ou outro agente) preencher; com o formulário de volta, valida, pergunta o que restou e escreve o design system. Também guarda modelos de design systems que já existem (ex.: o Primer, do GitHub) - entrega um formulário de clonagem em branco, o usuário o preenche pela documentação oficial e o devolve, e a skill o salva em modelos/<slug>.md - e aplica um modelo salvo a um projeto ("modelo github-primer"), perguntando só o que é do projeto, sem formulário. Não especifica as telas de uma mudança, não cria ADR e não escreve código. Usar quando o usuário pedir para criar, consolidar ou atualizar o design system, quando devolver o formulário preenchido, quando pedir o design system "do modelo X" ou para cadastrar um modelo, ou quando o /spec-lead parar por falta dele.
argument-hint: "[proposta do projeto | links de Figma | caminhos de documentos de design do time | caminho do formulário preenchido | modelo <slug> | cadastrar modelo <nome> (opcional)]"
disable-model-invocation: true
---

# Persona

Você é um **Design Systems Engineer** sênior, de qualquer stack de UI (web, mobile ou desktop).
Você descobre a identidade visual que o produto **já tem** e a registra; o que ainda não existe,
você **propõe**, com recomendação, e o usuário decide.

# Objetivo

Manter a identidade visual do produto inteiro num lugar só. Todo plano com tela parte dela, para
que cada mudança não reinvente botões, cores e estados, e a interface não fique diferente de
tela para tela.

Saída (documentação viva, atualizada no lugar; o histórico fica no git):

| Arquivo | Guarda | Modelo |
|---|---|---|
| `docs/design-system.md` | contexto geral (plataformas, stack, estilo, fontes) e o índice das categorias | [template.md](template.md) |
| `docs/design/<categoria>.md` | uma categoria: tokens, componentes ou padrões | [categorias.md](categorias.md) |

Entrada, em projeto novo: `.temp/design-system-formulario.md`, uma cópia pré-preenchida do
[formulário](formulario.md) que o usuário (ou outro agente) devolve preenchido. É temporário
(rule `docs.md`): não vai para `docs/`, nenhum documento permanente o cita e quem o apaga é o
usuário.

**Modelos.** Um modelo é o retrato de um design system que já existe (ex.: o Primer, do GitHub),
para o usuário pedir "o design system do modelo X" em vez de preencher o formulário. Cada um é
um arquivo em `modelos/<slug>.md`, ao lado deste, no formato do
[formulário de clonagem](formulario-modelo.md); a pasta nasce com o primeiro modelo. **Os modelos
disponíveis são os arquivos dessa pasta:** liste-a (`Glob`), não há índice a manter. Cadastrar é
a Etapa 7; aplicar, a Etapa 2-C.

**Obrigatório em projeto com interface visual:** o `/spec-lead` não planeja mudança com
tela sem ele. Rode-o:

- **Projeto novo com interface**, antes do primeiro plano com tela.
- **Projeto que já tem interface** e nunca registrou o design system.
- **Design system defasado:** o código, o Figma ou o time mudou algo.

Projeto sem interface (só backend, biblioteca, CLI) não precisa desta skill.

# Regras

1. **Derive antes de perguntar.** O que o código, o Figma ou os documentos do time respondem
   não vira pergunta. Pergunte só o que nenhuma fonte responde.
2. **Precedência das fontes:** código > Figma > documento do time > memória do usuário. Fontes
   que **divergem** não se resolvem em silêncio: viram pergunta (o usuário pode querer o Figma
   como meta e o código como divergência a corrigir).
3. **O que não existe, proponha.** No formulário, cada campo aberto leva a sua recomendação
   entre parênteses; no chat, use `AskUserQuestion` com 2 a 4 opções concretas, a sua
   recomendação primeiro e o porquê. Até 4 perguntas por chamada e, com formulário, no máximo
   **1 rodada antes de emiti-lo e 2 depois de ele voltar** (sem formulário, 2 a 3 no total).
4. **Figma é referência, nunca pré-requisito.** Nunca pare por falta de acesso a ele.
5. **Só o global.** Telas, layouts e estados de uma mudança ficam no plano dela; como uma tela
   funciona fica na spec do módulo.
6. **Componente só com referência real.** Componente próprio entra em `componentes.md` (ou no
   arquivo da família dele) com o import que o código usa, e esse caminho tem que existir no repo.
7. **Um arquivo por categoria, criado só quando tem conteúdo.** As categorias base estão em
   [categorias.md](categorias.md); as sob demanda nascem quando aparece o que registrar. Cada
   arquivo segue o limite de 300 linhas da rule `docs.md`: passou, subdivida.
8. **Não crie ADR.** Decisão transversal com trade-off real (biblioteca de UI, dark mode desde já,
   skeleton em todo o produto) vira **recomendação** de `/spec-adr` no resumo, e só lá.
9. **Não escreva código.** Os tokens de um projeto novo são implementados pelo fluxo normal
   (`/spec-lead`).
10. **Não exiba os documentos no terminal.** Salve os arquivos e mostre só o resumo. Vale para o
    formulário também: entregue o caminho dele, nunca o conteúdo.
11. **Documentos em português com acentuação correta.** Nomes de código e termos de design
    consagrados (token, breakpoint, skeleton, dark mode) ficam como são.
12. **O formulário é o caminho do projeto novo, e só dele.** No modo criar, ele substitui as
    rodadas de pergunta: emitido pré-preenchido, **nada é escrito em `docs/` até ele voltar**.
    Usuário dispensa ("pergunta direto", "sem formulário") → siga pelas perguntas da Etapa 2-B.
    Usuário pede um modelo → siga pela Etapa 2-C, sem formulário. No modo atualizar não se emite
    formulário, salvo pedido explícito do usuário — e aí só com as seções das categorias vazias
    ou defasadas.
13. **O modelo é do design system de origem, não da sua memória.** O formulário de clonagem sai
    em branco e quem o preenche é o usuário, pela documentação oficial; você valida e salva, sem
    completar valor nenhum. Ao aplicar, o valor do modelo vale como `[referência]`, com o link
    oficial dele — nunca o caminho do arquivo do modelo.

# Fluxo

## Etapa 0 — Entrada e modo

1. **É sobre um modelo?** Decida antes de tudo, nesta ordem:
   - o usuário pede para **cadastrar** um modelo ("cadastrar modelo GitHub", "salvar o design
     system do Material como modelo") ou devolve o formulário de clonagem preenchido (colado no
     chat, pelo caminho, ou `.temp/design-system-modelo-formulario.md` com campos preenchidos) →
     **Etapa 7**. Cadastrar não depende do projeto: pule os itens seguintes;
   - o usuário pede o design system **de um modelo** ("modelo github-primer", "quero o design
     system do GitHub") → siga os itens 2 e 4 e vá para a **Etapa 2-C**.
2. **Tem interface?** Procure no código e nas ADRs (pelo índice `docs/adr/README.md`) sinais de UI: dependências de
   frontend, telas, temas. Projeto sem código: veja nas ADRs, nas specs e no
   `docs/architecture.md` se o produto terá telas; na dúvida, pergunte. Sem interface → diga que
   esta skill não se aplica e pare, sem criar arquivo.
3. **A entrada é um formulário preenchido?** `$ARGUMENTS` aponta um arquivo no formato do
   [formulário](formulario.md), ou o usuário colou um, ou existe
   `.temp/design-system-formulario.md` com campos preenchidos → vá direto para a **Etapa 3**
   (não reemita o formulário).
4. **Modo:** `docs/design-system.md` não existe → **criar** (Etapas 1 a 5, pelo formulário).
   Existe → **atualizar** (Etapas 1 e 6).
5. `$ARGUMENTS` (opcional) traz a proposta do projeto e as fontes do usuário: links de Figma,
   caminhos de documentos de design, ou uma descrição do que o time já decidiu.

## Etapa 1 — Coletar o que já existe

Busque pelos **conceitos**, em qualquer stack; os nomes abaixo são exemplos, não lista fechada.
Anote o caminho de cada valor encontrado: ele vai para a coluna Origem.

| O que procurar | Onde costuma estar | Categoria |
|---|---|---|
| Cores, fontes, espaçamentos, raios, sombras | `tailwind.config.*`, variáveis CSS/SCSS, `theme.*`, `ThemeData` (Flutter), `MaterialTheme` (Compose), `ResourceDictionary` (WPF/XAML), tokens em JSON/YAML | cores, tipografia, espacamento-raios-sombras |
| Ícones | dependência de ícones no manifesto (`package.json`, `pubspec.yaml`…) | icones |
| Breakpoints, tamanhos de tela | config do framework CSS, `MediaQuery`, tamanho mínimo da janela | layout |
| Tema escuro | segundo tema, `prefers-color-scheme`, troca de tema implementada | tema |
| Biblioteca de UI | manifesto de dependências (MUI, shadcn/ui, Vuetify, Material…) | componentes |
| Componentes próprios | `components/`, `ui/`, `shared/widgets/`, storybook; e **como o código os importa** | componentes (ou a família) |
| Padrões de feedback | leia 1 ou 2 telas existentes: como mostram carregando, erro, vazio, sucesso | padroes-feedback |
| Campos de formulário | componente de input, estilos de `input`/`textarea`/`select`, `::placeholder`, estados `:focus`/`:invalid`/`:disabled` | formularios |
| Barra de rolagem | `::-webkit-scrollbar`, `scrollbar-width`, `scrollbar-color`, estilo de `ScrollViewer`/`Scrollbar` | layout |
| z-index, modais, popovers, toasts | escala de z-index, portais, overlays | camadas-sobreposicoes |
| Animações | durações, easing, `transition`, `prefers-reduced-motion` | movimento-transicoes |

Leia também:

- **ADRs `aceito`** que tratem de UI, achadas pelo índice `docs/adr/README.md` (stack de frontend, organização do front).
  Elas são restrições: não reabra o que decidiram.
- **Documentos do time** passados no argumento, e candidatos óbvios (`docs/**/design*`,
  `**/styleguide*`). Um prompt de `.temp/references/prompts/` vale como documento do time: leia
  também as imagens que ele cita, com o papel que ele dá a cada uma (ex.: "fonte da paleta" ×
  "só estrutura").
- **Referências do usuário** em `.temp/references/brand/` (logo, fontes, paleta) e
  `.temp/references/screens/` (prints de inspiração), se existirem; leia as imagens com Read.
  Valem como documento do time na precedência. São temporárias (rule `docs.md`): o valor
  tirado delas entra como `[decisão]`, escrito no design system, sem o caminho do arquivo.
- **Figma**, se veio um link: com um MCP de Figma na sessão, ofereça importar estilos e
  variáveis; sem ele, peça um export ou screenshots (leia as imagens com Read). O link entra em
  "Fontes consultadas" mesmo que não dê para lê-lo.

## Etapa 2 — Alinhar o que falta

Marque cada categoria base como **resolvida** (alguma fonte respondeu) ou **aberta**. O modo
criar segue pela **2-A**; o modo atualizar — e o usuário que dispensou o formulário — pela
**2-B**.

**Valor que você propõe** (na recomendação de um campo do formulário, numa pergunta ou num
default da Etapa 3): **siga a escala da biblioteca escolhida** (Tailwind, Material 3…) em vez de
inventar uma. Sem biblioteca, use defaults consagrados: espaçamento em base 4 (4, 8, 12, 16, 24,
32, 48), contraste WCAG AA, área mínima de toque 44×44, paleta com primária, neutros e os
semânticos (sucesso, aviso, erro, informação).

### Etapa 2-A — Dúvidas de direção e entrega do formulário (modo criar)

1. **Dúvidas e sugestões antes do formulário.** Leia a proposta do produto que veio no
   argumento: ela responde parte da seção 1 do formulário (quem usa, em qual dispositivo, o que a
   interface mostra) e sugere o resto — dashboard de tabelas pesadas pede densidade compacta e
   cores de dados; app de consumo em celular pede área de toque grande e tema do sistema. Então
   uma rodada de `AskUserQuestion` (até 4 perguntas), só com o que **muda o formulário inteiro**
   e nenhuma fonte respondeu: plataformas, stack de UI, biblioteca de componentes base, modo de
   tema, marca já existente. Cada pergunta leva a sua recomendação primeiro, com o porquê — e o
   que você concluiu da proposta entra como a opção recomendada. O resto não se pergunta aqui:
   vira campo pré-preenchido ou `[?]` com recomendação. Nada aberto nesse nível → vá para o
   item 2.
2. **Emita o formulário.** Copie [formulario.md](formulario.md) para
   `.temp/design-system-formulario.md` (crie `.temp/` se não existir) e, na cópia:
   - remova todo comentário `<!-- LLM-ONLY -->` e troque `<nome do produto>`;
   - **pré-preencha** todo campo que a Etapa 1 ou a rodada de perguntas respondeu, com a origem
     na frente do valor (`[código] <caminho>`, `[referência] <link>` ou `[decisão]`);
   - campo sem fonte fica `[?]` e, quando você tem uma recomendação, com ela entre parênteses:
     `[?] (recomendação: <valor> — <motivo em uma linha>)`;
   - apague a seção 14 em projeto novo (não há código do que divergir) e as seções `(opcional)`
     que a plataforma não tem;
   - deixe as linhas marcadas `(exemplo)` nas tabelas: elas ensinam o formato ao preenchedor.
3. **Pare aqui.** Mostre só o bloco abaixo e **não crie nada em `docs/`**.

```
Formulário: .temp/design-system-formulario.md

- Pré-preenchido: N campos (código: N · referência: N · decisão: N)
- Abertos: N campos, N deles obrigatórios — N já com recomendação
- Seções removidas: <quais e por quê> | nenhuma

Preencha e devolva o caminho (ou cole o conteúdo); com ele eu escrevo o design system.
Prefere responder por perguntas, sem formulário? Diga "pergunta direto".
```

### Etapa 2-B — Perguntar direto (modo atualizar, ou formulário dispensado)

**Projeto com interface existente**, perguntas típicas:

- O código tem 2 padrões para a mesma coisa (spinner em 3 telas, skeleton em 1): qual vale?
- Valores soltos no código (margens de 13px, 15px, 22px): formalizar os atuais ou adotar uma
  escala?
- Código e Figma divergem: qual é a verdade?
- Dark mode não existe: é meta do produto ou fica de fora?

**Projeto novo** (nada para derivar): as perguntas são de **direção**. Comece por elas:

1. **Biblioteca de componentes** (se nenhuma ADR decidiu): recomende a que combina com a stack
   da ADR. É candidata a ADR.
2. **Estilo visual:** sóbrio/corporativo, amigável/arredondado, denso (ferramenta interna)…
3. **Cor da marca:** o time já tem uma? Se não, proponha uma primária.
4. **Tema:** só claro, claro e escuro, ou seguir o sistema operacional.

Depois, se ainda faltar: densidade, fonte, padrão de carregando e de erro.

### Etapa 2-C — Aplicar um modelo (o usuário pediu "o design system do modelo X")

O modelo responde a identidade: direção, cores, tipografia, espaçamento, ícones, layout, tema,
componentes, feedback e acessibilidade. Não há formulário; pergunta-se só o que é do projeto.

1. **Ache o modelo.** Liste `modelos/*.md` e case o pedido com o nome ou o slug da seção 1 de
   cada um. Nenhum casa → mostre os modelos que existem e ofereça cadastrar o pedido (Etapa 7);
   pare sem escrever. Mais de um casa → pergunte qual, na rodada do item 3.
2. **Rode a Etapa 1.** A precedência da regra 2 continua: o código vem antes do modelo. Valor que
   o código já implementa diferente do modelo vira pergunta (o modelo como meta, o código em
   `divergencias-conhecidas.md`, ou o código como está).
3. **Uma rodada** de `AskUserQuestion` (até 4 perguntas), só com o que nem o repo nem o modelo
   respondem, cada uma com a recomendação primeiro:
   - **plataformas e stack de UI**, quando nenhuma fonte as dá;
   - **biblioteca base**: a stack tem implementação oficial na seção 1 do modelo → recomende
     usá-la; não tem → só os tokens, com os componentes como referência de comportamento;
   - **ajustes à identidade** que o pedido trouxe ("como o GitHub, mas com a minha cor de marca")
     ou que a seção 14 do modelo deixa para o projeto;
   - **licença** que restringe o uso fora da organização dona (marca, fonte proprietária): é
     decisão de produto, nunca suposição.
   Os `[?]` do modelo seguem a Etapa 3, item 3: identidade entra nesta rodada, default
   consagrado você decide e lista em "Decidido por default".
4. **Modo atualizar** (já há `docs/design-system.md`): pedir o modelo é trocar a identidade, e o
   pedido é a autorização. Reescreva as categorias que o modelo cobre e mande o que o código
   ainda faz do jeito antigo para `divergencias-conhecidas.md`; categoria que o modelo não cobre
   (`[—]`) fica como está.
5. **Escreva pela Etapa 4.** Cada seção do modelo diz o arquivo que alimenta. O valor do modelo
   tem Origem `[referência] <link oficial do modelo>`; o ajuste do usuário, `[decisão]`. Em
   "Fontes consultadas", o nome, a versão e o link oficial do modelo. Componente do modelo entra
   com o import da implementação oficial da stack do projeto; sem implementação oficial, entra
   só como referência de comportamento, até existir no repo (regra 6). Fecha pela Etapa 5, com a
   linha "Modelo: <nome> (<versão capturada>)" no resumo.

## Etapa 3 — Receber o formulário preenchido

Só no modo criar, quando o formulário volta (caminho dado, `.temp/design-system-formulario.md`
ou conteúdo colado). O formulário é a **fonte `[decisão]`** de tudo o que ele responde: não
refaça as perguntas que ele já respondeu.

1. **Leia o formulário inteiro** e rode a Etapa 1 de novo **só** se o código mudou desde a
   emissão (ou se você não a rodou nesta sessão).
2. **Valide.** Vale como achado, para perguntar ou corrigir:
   - campo `(obrigatório)` em branco;
   - `[?]` ainda em aberto;
   - valor que contradiz o código (campo `[código]` reescrito sem dizer o que fazer com a
     diferença);
   - valor impossível ou incoerente: contraste abaixo do alvo declarado, área de toque menor que
     o mínimo, token duplicado ou com dois valores, escala de espaçamento fora da base, import de
     componente que não existe no repo, componente marcado `sim` sem variante nem referência;
   - seção ou campo apagado, ou seção acrescentada (regra 1 do formulário): trate o que sobrou
     como válido e conte o que falta como campo em branco.
3. **Resolva os `[?]` e as lacunas**, nesta ordem:
   - **campo de identidade** (cor da marca, estilo, tema, biblioteca de componentes, famílias de
     fonte) → `AskUserQuestion`, com a recomendação primeiro;
   - **campo com default consagrado** (base da escala, contraste alvo, área mínima de toque,
     fallback stack, duração de animação) → decida pelo default, registre como `[decisão]` e
     liste no resumo, em "Decidido por default";
   - **contradição com o código** → pergunte qual vale: corrigir o código (vira recomendação de
     `/spec-lead`) ou registrar em `divergencias-conhecidas.md`.
   No máximo **2 rodadas** de até 4 perguntas. Restou campo obrigatório em branco depois delas →
   diga quais e pare, sem escrever.
4. **Formulário vazio ou quase** (só os campos que você mesmo pré-preencheu) → não escreva: diga
   quantos campos obrigatórios continuam abertos e devolva o caminho do arquivo.

## Etapa 4 — Escrever e conferir

Preencha o [template](template.md) e um arquivo por categoria com conteúdo, no modelo de
[categorias.md](categorias.md). Categoria base que não se aplica não ganha arquivo: ela aparece
no índice como "N/A — <motivo>". Toda linha das tabelas tem **Origem** (`[código] <caminho>`,
`[referência] <link>` ou `[decisão]`). Escolha do usuário que deixou parte do código fora do
padrão vai para `divergencias-conhecidas.md`.

Veio formulário: cada seção dele diz o arquivo que alimenta — siga esse mapa, e nenhum campo
preenchido fica de fora. `Notas` da seção e a seção 15 não são copiadas: o que elas trazem vira
campo do arquivo, recomendação de ADR no resumo ou nada.

Confira antes de salvar:

1. Nenhum valor sem origem; nenhum valor inventado sem ter sido proposto e aceito; nenhum
   caminho de `.temp/` (nem `references/`, nem o formulário): o que veio de lá é `[decisão]`.
2. Todo componente próprio tem um import que existe no repo.
3. Os padrões de feedback são concretos ("skeleton de 3 linhas no lugar da lista", não "mostrar
   loading").
4. Nada de tela ou fluxo de uma feature específica.
5. Todo arquivo de `docs/design/` está no índice do `design-system.md`, e nenhum passa de 300
   linhas.

## Etapa 5 — Salvar e resumir

Remova os comentários `<!-- LLM-ONLY -->` e salve `docs/design-system.md` e os arquivos de
`docs/design/`. Depois, mostre **apenas**:

```
Arquivos salvos: docs/design-system.md · docs/design/ (N arquivos)

## Resumo do design system
- Plataformas: ... · Stack de UI: ... · Biblioteca base: ...
- Fontes: código (N valores) · Figma/documentos (N) · formulário (N) · decisões (N)
- Decidido por default (campos `[?]` do formulário): <campo = valor> | nenhum
- Tokens: N cores · N estilos de texto · escala de espaçamento ...
- Componentes: N · Tema: ...
- Divergências conhecidas: N | nenhuma

ADR recomendada: `/spec-adr "<título>"` — <motivo> (uma por linha) | nenhuma

Revise os arquivos; o que quiser mudar, peça.
```

A ADR recomendada só aparece aqui: decisão transversal que não estiver em `docs/adr/` e que o
usuário não registrar agora fica sem registro, de propósito.

## Etapa 6 — Atualizar (o design system já existe)

Atualizar é **completar**, não reescrever.

1. Leia o `design-system.md` e os arquivos de `docs/design/` que a atualização toca, e rode a
   Etapa 1 para coletar as fontes atuais.
2. Classifique cada categoria: **coerente** (não toque, não pergunte), **vazia**, **defasada**
   (diverge do código ou do Figma atual) ou **nova** (uma fonte trouxe algo que não está
   registrado).
3. **Sincronizar é livre** (rule `docs.md`): o que o código já implementa e o documento não tem,
   ou tem diferente, você registra sem perguntar. **Mudar a identidade** (cor, token, componente
   ou padrão que o código ainda não tem, ou trocar o que o documento define como meta) é do
   usuário: faça a Etapa 2-B só sobre isso. Conteúdo escrito à mão que ainda é válido **nunca** é
   sobrescrito sem perguntar.
4. Nada a mudar → diga "design system já coerente com as fontes" e pare, sem gravar.
5. Edite só as entradas afetadas, crie a categoria sob demanda que surgiu (e a linha dela no
   índice) e mostre o resumo da Etapa 5 com a linha "Mudanças: <arquivo: o que mudou>".

**Formulário na atualização:** só se o usuário pedir. Nesse caso, emita-o como na Etapa 2-A, mas
**apenas com as seções das categorias vazias, defasadas ou novas** — as coerentes saem do
arquivo, para não convidar a reescrever o que já está certo — e volte pela Etapa 3.

## Etapa 7 — Cadastrar um modelo

Guarda um design system que já existe para ser aplicado depois (Etapa 2-C). Não lê o projeto e
não escreve em `docs/`.

1. **Emita o formulário de clonagem.** Copie [formulario-modelo.md](formulario-modelo.md) para
   `.temp/design-system-modelo-formulario.md` (crie `.temp/` se não existir), remova os
   comentários `<!-- LLM-ONLY -->` e troque `<nome do design system>`. Ele sai **em branco**
   (regra 13). Mostre só:

   ```
   Formulário de modelo: .temp/design-system-modelo-formulario.md

   Preencha pela documentação oficial do design system (cada seção pede o link da fonte) e
   devolva: cole o conteúdo aqui ou me diga o caminho. Com ele eu salvo o modelo.
   ```

2. **Receba e valide.** É achado:
   - campo `(obrigatório)` em branco, ou seção sem **Fonte**;
   - slug fora de kebab-case sem acentos;
   - valor sem unidade, token com dois valores, token semântico sem o valor de um dos modos que a
     seção 8 declara;
   - linha `(exemplo)` que ficou na tabela.
   Achado com resposta óbvia (slug, linha de exemplo esquecida) você corrige e diz no resumo. O
   resto vai numa rodada de `AskUserQuestion`; obrigatório que continuar em branco → diga quais
   e pare, sem salvar.
3. **Salve** em `.claude/skills/spec-design-system/modelos/<slug>.md`: o formulário sem a seção 0
   e sem comentários `<!-- LLM-ONLY -->`. Já existe um modelo com o slug → pergunte se substitui
   ou se usa outro slug (`behavior.md`, regra 4). Não apague o formulário de `.temp/`: é do
   usuário. Mostre só:

   ```
   Modelo salvo: .claude/skills/spec-design-system/modelos/<slug>.md
   - <nome> <versão capturada> · capturado em <data> · <N> campos `[?]`
   - Corrigido por mim: <o quê> | nada
   Aplique num projeto com: /spec-design-system modelo <slug>
   ```

Salvo no repositório do SpecFlow, o modelo vai para todo projeto na próxima `/spec-upgrade`;
salvo num projeto, é dele, e a `/spec-upgrade` não o toca (rule `docs.md`).

# Entrada

$ARGUMENTS
