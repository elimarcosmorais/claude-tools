---
status: aceito
---

# ADR 003 — Windows e Linux (WSL2) a partir do mesmo código

- **Data:** 2026-11-04
- **Origem:** plano `2026-11-04-suporte-linux-wsl2` (Decisões)
- **Substitui:** nenhuma
- **Substituída por:** nenhuma

## Contexto

O Claude Code dentro do WSL2 (o `claude` no terminal do Linux ou a extensão do VS Code no modo
remoto) guarda o login na home do Linux, separado da home do Windows. O `claude-tools` só
compilava, testava e instalava no Windows. O código de contas (`internal/`) já usa só
`path/filepath` e `os`, mas os testes dependiam de `icacls` e de `pwsh`, e o rename atômico e as
permissões só valem num sistema de arquivos local do próprio sistema.

## Decisão

Um código só, compilado num binário por sistema (`claude-tools.exe` no Windows, `claude-tools`
no Linux, que roda dentro do WSL e mexe na home do Linux), com um instalador por sistema
(`scripts/install.ps1` e `scripts/install.sh`). O que depende do sistema fica em arquivos com
restrição de build (`_windows_test.go`, `//go:build !windows`), nunca em teste pulado.

## Consequências

**Ganhos:**

- A lógica de contas e as mensagens são as mesmas nos dois sistemas; a spec vale para os dois.
- Cada binário grava no sistema de arquivos local dele: rename atômico do `writeFileAtomic`,
  dono certo e permissões `0600`/`0700` nos arquivos de login do Linux.

**Custos:**

- Dois instaladores para manter em paralelo, com testes de cada um rodando só no próprio
  sistema.
- A suíte inteira só é provada rodando nos dois sistemas (no Windows e no WSL).
- As contas de cada ambiente ficam separadas: salvar no Windows não aparece no WSL.

**Neutras:**

- A release passa a ter um anexo por sistema.

## Alternativas consideradas

- **O `.exe` do Windows mexendo na home do WSL por `\\wsl$`** — um binário só. Rejeitada porque
  o 9P não garante o rename atômico do `writeFileAtomic` e grava com dono e permissões errados.
- **Projeto separado para o Linux** — liberdade total no lado Linux. Rejeitada porque duplica a
  lógica de contas, que já é portável.
- **Compartilhar as contas entre Windows e WSL** — salvar uma vez só. Rejeitada porque o token
  de renovação é rotacionado, e usar a conta de um lado derrubaria o login do outro.
