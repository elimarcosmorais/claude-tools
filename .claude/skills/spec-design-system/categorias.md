# Categorias do design system

Modelo de cada arquivo de `docs/design/`. Todo arquivo começa com `# <Categoria>` e uma linha
dizendo o que ele cobre; depois, só as tabelas e campos da categoria. Toda linha de tabela tem
**Origem** (`[código] <caminho>`, `[referência] <link>` ou `[decisão]`).

<!-- LLM-ONLY: este arquivo é modelo, não vai para o projeto. Base = criada sempre que se
  aplica ao produto. Sob demanda = criada só quando aparece o que registrar. Limite de 300
  linhas por arquivo (rule docs.md): passou, subdivida (ex.: componentes.md vira índice de
  componentes/formularios.md, componentes/navegacao.md). -->

## Base

### `cores.md`

<!-- LLM-ONLY: se o código já declara o token (variável CSS, chave do tema), use o nome dele.
  Mínimo num projeto novo: primária, neutros (fundo, superfície, borda, texto, texto secundário)
  e os semânticos (sucesso, aviso, erro, informação). -->

| Token | Valor | Uso | Origem |
|-------|-------|-----|--------|

### `tipografia.md`

- **Fonte(s):** <!-- LLM-ONLY: família e de onde vem (Google Fonts, do sistema, pacote) -->

| Estilo | Tamanho / Peso / Altura de linha | Uso | Origem |
|--------|----------------------------------|-----|--------|

### `espacamento-raios-sombras.md`

- **Escala de espaçamento:** <!-- LLM-ONLY: ex.: base 4 (4, 8, 12, 16, 24, 32, 48) -->

| Token | Valor | Uso | Origem |
|-------|-------|-----|--------|

### `icones.md`

- **Biblioteca:** <!-- LLM-ONLY: ex.: lucide, Material Symbols, SF Symbols -->
- **Tamanhos:**
- **Ícones com significado fixo:** <!-- LLM-ONLY: ex.: lixeira = excluir. Só os que o produto já usa assim. -->

### `layout.md`

<!-- LLM-ONLY: as faixas de tamanho de tela da plataforma. Web: breakpoints. Mobile: celular,
  tablet e orientação. Desktop: tamanho mínimo da janela e o que muda ao redimensionar. -->

| Faixa | Valor | O que muda |
|-------|-------|------------|

- **Densidade:** <!-- LLM-ONLY: ex.: confortável; compacta só em tabelas -->
- **Estrutura das telas:** <!-- LLM-ONLY: ex.: barra lateral + conteúdo com largura máxima de 1200px -->
- **Barra de rolagem:** <!-- LLM-ONLY: nativa do sistema, ou personalizada com largura e as cores da barra e do trilho por tema -->


### `tema.md`

- **Modos:** claro | escuro | claro e escuro | segue o sistema operacional
- **Como o tema é aplicado:** <!-- LLM-ONLY: ex.: tokens semânticos (fundo, texto) sobre a paleta, trocados por classe/tema -->

| Token semântico | Claro | Escuro | Origem |
|-----------------|-------|--------|--------|

### `componentes.md`

<!-- LLM-ONLY: só componentes REUTILIZÁVEIS entre features. A Referência é o import real que o
  código usa (`@/components/ui/button`, `import { Button } from '@acme/ui'`): é por ela que o
  executor reusa o componente em vez de criar outro. Componente próprio sem referência que
  exista no repo não entra. Projeto novo, sem componentes ainda: liste os da biblioteca base que
  serão o padrão. Uma família que cresceu (botões, formulários) ganha arquivo próprio e fica aqui
  só como link. -->

| Componente | Referência (import) | Variantes | Quando usar |
|------------|---------------------|-----------|-------------|

Componente reutilizável novo entra aqui antes de ser usado por uma segunda feature.

### `padroes-feedback.md`

<!-- LLM-ONLY: comportamento CONCRETO (o quê, onde, com qual ação), nunca "mostrar loading".
  É o padrão que toda tela herda; a feature que fugir dele diz por quê no plano. -->

| Situação | Padrão | Componente | Origem |
|----------|--------|------------|--------|
| Carregando |      |            |        |
| Erro     |        |            |        |
| Vazio    |        |            |        |
| Sucesso  |        |            |        |
| Ação destrutiva |  |           |        |

### `acessibilidade-visual.md`

- **Contraste mínimo:** <!-- LLM-ONLY: ex.: WCAG AA (4,5:1 em texto normal) -->
- **Foco visível:**
- **Área mínima de toque/clique:** <!-- LLM-ONLY: ex.: 44×44 -->
- **Movimento reduzido:**

## Sob demanda

### `botoes.md` (ou outra família de componentes)

Nasce quando a família não cabe numa linha de `componentes.md`: variantes, tamanhos, estados
(padrão, hover, foco, desabilitado, carregando) e quando usar cada uma.

| Variante | Referência (import) | Quando usar | Origem |
|----------|---------------------|-------------|--------|

### `formularios.md`

Nasce quando o produto tem campo de texto: tamanhos, estados e anatomia dos campos, que toda
tela com formulário segue.

| Tamanho | Altura / Padding horizontal / Texto | Origem |
|---------|-------------------------------------|--------|

| Estado | Borda / Fundo / Texto / Outro | Origem |
|--------|-------------------------------|--------|

- **Anatomia:** <!-- LLM-ONLY: posição do rótulo, marcação de obrigatório, texto de ajuda, mensagem de erro, contador -->
- **Placeholder:** <!-- LLM-ONLY: cor e regra de uso -->
- **Quando o erro aparece:**

### `camadas-sobreposicoes.md`

Escala de z-index e o comportamento de modal, drawer, popover, tooltip e toast: o que fica por
cima de quê, fundo escurecido, fechar com Esc ou clique fora.

| Camada | z-index | Componentes | Origem |
|--------|---------|-------------|--------|

### `movimento-transicoes.md`

| Token | Duração / Easing | Uso | Origem |
|-------|------------------|-----|--------|

- **Movimento reduzido:** <!-- LLM-ONLY: o que muda com prefers-reduced-motion (ou o equivalente da plataforma) -->

### `divergencias-conhecidas.md`

Onde o código ainda não segue o design system (ex.: 2 telas usam spinner, mas o padrão é
skeleton). Item corrigido sai da lista; lista vazia → apague o arquivo e a linha do índice.

- <onde> — <o que diverge> — <o que vale: o design system>
