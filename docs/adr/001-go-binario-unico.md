---
status: aceito
---

# ADR 001 — Go com biblioteca padrão, binário único

- **Data:** 2026-10-17
- **Origem:** plano `2026-10-17-troca-de-conta` (Decisões)
- **Substitui:** nenhuma
- **Substituída por:** nenhuma

## Contexto

O `claude-tools` é um CLI de Windows que vai para o `PATH` do usuário e precisa rodar em
qualquer terminal sem pedir a instalação de um runtime. A ferramenta que ele substitui era um
script Python, que exigia o Python instalado. O CLI é pequeno (cinco comandos que leem e gravam
arquivos JSON) e precisa abrir rápido, porque roda a cada troca de conta.

## Decisão

Escrever o `claude-tools` em Go (1.27, em `C:\dev\programs\go\current`) usando só a biblioteca
padrão, compilado num executável único `claude-tools.exe` (`go build ./cmd/claude-tools`).

## Consequências

**Ganhos:**

- Um `.exe` só, sem runtime nem dependência para instalar: copiar o arquivo basta.
- Abre rápido e tem tamanho pequeno.
- Teste, formatação e análise estática já vêm com o toolchain (`go test`, `gofmt`, `go vet`).
- Sem módulos de terceiros, não há dependência para atualizar nem auditar.

**Custos:**

- Sem biblioteca de CLI: o parse dos argumentos e o texto de uso são escritos à mão.
- O `encoding/json` exige cuidado para regravar JSON de terceiros sem alterá-lo
  (`json.RawMessage` para preservar valores e `SetEscapeHTML(false)` para não escapar `<`, `>`
  e `&`).
- Quem for mexer no código precisa do Go instalado.

**Neutras:**

- O build e a instalação ficam num `Taskfile.yml` (`task build`, `test`, `lint`, `install`).

## Alternativas consideradas

- **Python + PyInstaller** — empacotar o script Python existente num `.exe`. Rejeitada porque
  gera um executável grande e lento para abrir, que se descompacta a cada execução.
- **.NET com AOT nativo** — C# compilado para código nativo. Rejeitada porque o toolchain é
  bem mais pesado do que um CLI de cinco comandos pede.
- **Continuar com o script Python e os `.bat`** — Rejeitada porque exige o Python instalado,
  que é o problema a resolver.
