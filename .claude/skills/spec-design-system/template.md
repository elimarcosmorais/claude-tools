# Design System — <nome do produto>

> Identidade visual do produto: tokens, componentes e padrões que **toda feature com interface
> segue**. Este arquivo dá o contexto geral e aponta para as categorias em `docs/design/`.
> Gerado e atualizado por `/spec-design-system`; sincronizar com o código é livre, mudar a
> identidade é decisão do usuário (rule `docs.md`).
>
> **Origem de cada valor:** `[código]` já implementado (com o caminho) · `[referência]` Figma ou
> documento do time (com o link) · `[decisão]` decidido com o usuário (proposto e aceito, ou
> tirado do material que ele trouxe).

- **Plataformas:** web | mobile | desktop
- **Stack de UI:** <!-- LLM-ONLY: detectada no código ou na ADR. Ex.: React + Tailwind, Flutter (Material 3), WPF. Não invente. -->
- **Biblioteca de componentes base:** <!-- LLM-ONLY: ex.: shadcn/ui, MUI, Material 3, nenhuma (componentes próprios) -->
- **Fontes consultadas:** <!-- LLM-ONLY: caminhos do código, links de Figma, documentos do time. Nunca .temp/references/: é temporário (rule docs.md); o que veio de lá é [decisão]. -->
- **ADRs relacionadas:** `docs/adr/NNN-<titulo>.md` | nenhuma

## Direção

<!-- LLM-ONLY: 3 a 6 linhas que orientam quem cria uma tela nova e não acha o caso exato nas
  categorias: o estilo visual (sóbrio, amigável, denso), a densidade, o tema, a cor da marca e
  o que o produto evita. Só o que alguma fonte ou o usuário disse. -->

## Categorias

<!-- LLM-ONLY: uma linha por arquivo de docs/design/. Categoria base que não se aplica fica com
  "N/A — <motivo>" no lugar do link. Arquivo novo entra aqui; apagado sai. -->

| Arquivo | Abra quando |
|---|---|
| [cores](design/cores.md) | for usar ou criar uma cor, um fundo ou um estado (sucesso, erro…) |
| [tipografia](design/tipografia.md) | for escrever título, texto ou rótulo |
| [espaçamento, raios e sombras](design/espacamento-raios-sombras.md) | for definir margem, padding, borda arredondada ou elevação |
| [ícones](design/icones.md) | for usar um ícone |
| [layout](design/layout.md) | for montar a estrutura de uma tela ou tratar um tamanho de tela |
| [tema](design/tema.md) | for mexer em claro/escuro ou nos tokens semânticos |
| [componentes](design/componentes.md) | **antes de criar qualquer componente**: o que já existe e como importar |
| [padrões de feedback](design/padroes-feedback.md) | for mostrar carregando, erro, vazio, sucesso ou confirmar ação destrutiva |
| [acessibilidade visual](design/acessibilidade-visual.md) | for conferir contraste, foco, área de toque ou movimento reduzido |
