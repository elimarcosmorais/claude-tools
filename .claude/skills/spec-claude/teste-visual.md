# Teste visual por stack

Consulta da Etapa 2, só em projeto com interface. Diz **como o agente confere uma mudança de
tela**, por tipo de projeto. A ferramenta que o repo **já usa** sempre ganha desta tabela: ela só
decide quando o repo não tem nenhuma, e aí vira a recomendação da pergunta da Etapa 3.

| Projeto | Sinal no repo | Ferramenta | Como o agente confere |
|---|---|---|---|
| Laravel com páginas no navegador (Blade, Livewire, Inertia) | `laravel/framework` no `composer.json` | Laravel Dusk (`laravel/dusk`) | teste em `tests/Browser/`, `php artisan dusk`; na falha, o print fica em `tests/Browser/screenshots/` e o agente o abre |
| Rails | `rails` no `Gemfile` | system tests (Capybara) | teste em `test/system/`, `bin/rails test:system`; `take_screenshot` no ponto a conferir |
| Web em JS/TS (React, Vue, Svelte, Angular, Next, Nuxt) ou backend web sem ferramenta própria (Django, Flask, Go, .NET) | `package.json` com o framework de front; ou templates HTML servidos pelo backend | Playwright (`@playwright/test`; em Python, `pytest-playwright`) | `npx playwright test`; `expect(page).toHaveScreenshot()` compara com a imagem aprovada (`--update-snapshots` só quando a mudança visual é a pedida) |
| Desktop Windows (WPF, WinForms, WinUI, MAUI, Delphi, VB6) | `.csproj` com `UseWPF`/`UseWindowsForms`/`UseWinUI`, `.dproj`, `.vbp` | Windows-MCP (`/spec-windows-mcp`) | o agente abre o app e confere com `Snapshot`/`Screenshot`, clicando e digitando como o usuário |
| Electron ou Tauri | `electron` ou `@tauri-apps/api` no `package.json` | Playwright (Electron) ou Windows-MCP (os dois, no Windows) | Electron: `_electron.launch()` do Playwright, preferido porque Electron é o alvo mais pesado do `Snapshot` (chega a pendurar o app); Tauri ou sem teste: Windows-MCP como no desktop |
| Flutter | `flutter` no `pubspec.yaml` | golden tests | `matchesGoldenFile` em `test/`, `flutter test`; `--update-goldens` só quando a mudança visual é a pedida |
| React Native ou Expo | `react-native` ou `expo` no `package.json` | Maestro | fluxo YAML em `.maestro/` com `assertVisible` e `takeScreenshot`, `maestro test .maestro/`; precisa de emulador ou aparelho ligado (iOS só no macOS) |

Repo já tem Cypress, Selenium, Detox, Storybook com teste visual (Chromatic, Loki) ou outra
ferramenta → use a dele, com o comando que estiver nos scripts; não proponha trocar.

## Ver a tela sem escrever teste

Teste automatizado prova e fica no repo; às vezes o agente só precisa **olhar** a tela durante o
trabalho (conferir um layout, reproduzir um bug). Web → Playwright MCP (`/spec-playwright-mcp`);
desktop Windows → Windows-MCP (`/spec-windows-mcp`). Se o projeto tem o servidor no `.mcp.json`,
a seção Teste visual diz quando usá-lo.

## Acessibilidade junto com o teste de tela

Projeto web com Playwright → o mesmo teste checa acessibilidade com `@axe-core/playwright`:
`await new AxeBuilder({ page }).withTags(['wcag2a', 'wcag2aa']).analyze()` e
`expect(results.violations).toEqual([])`. Com `docs/design/acessibilidade-visual.md`, use o nível
de contraste que ele define (AA ou AAA) nas tags. Só entra na seção Teste visual se o pacote já
está no repo ou se o usuário o escolher na Etapa 3; sem isso, fica de fora.

## O que vai para o `CLAUDE.md`

A seção **Teste visual** do template, em 2 a 4 bullets, só com fato que tem fonte:

- a ferramenta e o comando (do script do projeto; sem script, o comando padrão da ferramenta, cuja
  fonte é o pacote no manifest);
- onde o teste de tela nasce, se for teste automatizado;
- quando conferir: **toda mudança de tela** passa pela ferramenta antes de dizer que terminou; não
  conseguiu usar a ferramenta → pergunte no chat, sem pular a conferência nem passá-la ao
  usuário;
- o que não fazer, se houver: atualizar imagem de referência (`--update-snapshots`,
  `--update-goldens`) sem a mudança visual ter sido pedida.

Ferramenta recomendada que **ainda não está instalada** (o usuário a escolheu na Etapa 3): a
seção diz qual é, sem comando, e a instalação vai para a Etapa 6 como próximo passo. Instalar é
uma mudança no projeto, não trabalho desta skill.

O MCP de tela (ver a tela sem escrever teste) entra na Etapa 6 quando o `.mcp.json` não tem o
servidor ou quando ele está lá mas inativo (a sessão não tem as ferramentas dele): web sem `playwright` → `/spec-playwright-mcp`; desktop Windows sem `windows-mcp` →
`/spec-windows-mcp`. Nenhuma skill de MCP é chamada daqui: elas só rodam pelo usuário, e as
ferramentas só aparecem na sessão seguinte. Também não é preciso voltar a esta skill depois: a
skill de MCP acrescenta ela mesma o bullet na seção Teste visual.
