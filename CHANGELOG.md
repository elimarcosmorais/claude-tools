# Changelog

Todas as mudanças relevantes deste projeto ficam neste arquivo.
O formato segue o [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/), e o projeto usa
[Versionamento Semântico](https://semver.org/lang/pt-BR/).

## [0.1.0] - 2026-10-17

### Incompatível

- **install:** torna obrigatória a pasta de instalação (ea70b10)
  `task install` sem `DIR` e `install.ps1` sem `-Root` não instalam mais em
  `C:\dev\programs\claude-tools`; passe a pasta com `task install DIR=C:\dev\programs\claude-tools`.

### Adicionado

- **account:** implementa o CLI de troca de contas do Claude Code (ac54bde)

[0.1.0]: https://github.com/elimarcosmorais/claude-tools/releases/tag/v0.1.0
