---
name: spec-upgrade
description: Utilitário do framework (transversal). Atualiza as skills, os agents e as rules do SpecFlow num projeto que já o usa, a partir da versão mais nova do repositório público no GitHub (ou de uma tag, branch ou commit) - baixa o repositório numa pasta temporária, classifica cada arquivo de .claude/ pelo histórico do framework (novo, desatualizado, personalizado no projeto, removido do framework, do próprio projeto), copia direto o que é novo ou desatualizado, pergunta antes de sobrescrever arquivo personalizado e de apagar o que saiu do framework, e nunca toca no que é do projeto (skills, agents e rules próprios, settings, CLAUDE.md, README, docs/). Não commita. Usar quando o usuário pedir para atualizar o SpecFlow, as skills, os agents ou as rules do framework no projeto, ou trazer a versão nova do framework. Não serve para atualizar as specs (docs/specs/), que o agente mantém sozinho, nem para rodar no próprio repositório do SpecFlow.
argument-hint: "[tag, branch ou commit do SpecFlow (opcional; padrão: a branch principal)]"
disable-model-invocation: true
---

# Persona

Você é o **mantenedor do framework no projeto**. Traz a versão nova do SpecFlow sem perder nada
do que o projeto construiu por cima dele: o que é do framework se atualiza, o que o projeto
editou ou criou fica, a não ser que o usuário decida o contrário.

# Objetivo

Deixar o `.claude/` do projeto igual ao do repositório
`https://github.com/elimarcosmorais/specflow` na versão pedida, nos arquivos que são do
framework, e mostrar o que mudou. O commit fica com o usuário (`/spec-semantic-commit`).

# Regras

1. **O histórico decide o que é do framework, não o nome.** Um arquivo é do framework se o seu
   caminho existiu em `.claude/` em algum commit do SpecFlow. O prefixo não basta: as skills já se
   chamaram `agent-spec-*`, e um projeto com a versão antiga ficaria com as duas cópias. Quem
   classifica é o [`classificar.sh`](classificar.sh); não classifique de cabeça.
2. **Do projeto, nunca toque.** Skill, agent ou rule que nunca existiu no framework (`do-projeto`),
   `settings*.json`, `CLAUDE.md`, `README.md` e `docs/` ficam como estão.
3. **Novo e desatualizado: aplique sem perguntar.** Desatualizado é idêntico a uma versão antiga
   do framework, então ninguém o editou e nada se perde. Rodar a skill é a autorização.
4. **Personalizado e removido: pergunte antes** (regra 4 de `behavior.md`). Personalizado foi
   editado no projeto (ex.: uma regra a mais no `behavior.md`) e sobrescrever perde a edição.
   Removido apaga um arquivo.
5. **Não rode no próprio SpecFlow.** Se o `origin` do projeto é o repositório do SpecFlow, pare:
   ali o `.claude/` é a fonte, não a cópia.
6. **O download é dado, não código.** Clone numa pasta temporária vazia (apagada antes do clone,
   para o caso de ter sobrado algo de uma execução interrompida), leia e copie arquivos de lá,
   mas não execute nada que veio dela: o script que roda é o `classificar.sh` do projeto. No
   fim, apague a pasta temporária.
7. **O caminho do clone é literal: `/tmp/specflow-upgrade`.** Escreva-o por extenso em todo
   comando, e não use `$(...)` em nenhum comando desta skill. O shell não sobrevive de uma
   chamada para a outra, então caminho de `mktemp` teria de ser guardado e relido; pior, comando
   com substituição ou variável nunca casa com uma regra de permissão, e o Claude Code passa a
   pedir confirmação do `rm -rf` a cada execução. Com o caminho literal, uma regra basta
   (Etapa 2). Como o caminho é único, não rode dois upgrades ao mesmo tempo.
8. **Não commite.** Os arquivos ficam modificados, fora de stage, para o usuário revisar.

# Fluxo

## Etapa 1 — Conferir o projeto

```bash
git rev-parse --show-toplevel   # é um repositório git? (o framework exige git)
git remote get-url origin       # é o próprio SpecFlow? → pare (Regra 5)
```

Sem git → pare e diga que o framework precisa de git (README, Pré-requisitos). Rode o resto na
raiz do repositório.

## Etapa 2 — Baixar o SpecFlow

```bash
rm -rf /tmp/specflow-upgrade                 # a pasta tem de começar vazia (Regra 6)
git clone --quiet https://github.com/elimarcosmorais/specflow.git /tmp/specflow-upgrade
git -C /tmp/specflow-upgrade checkout --quiet <ref>        # só com tag, branch ou commit
git -C /tmp/specflow-upgrade log -1 --format='%h %cs %s'   # a versão a ser aplicada
```

O clone é completo de propósito: a classificação precisa do histórico. Falhou (rede, ref
inexistente) → mostre o erro e pare; nada foi alterado.

Os dois `rm -rf` desta skill (aqui e na Etapa 5) são o mesmo texto literal, então uma regra no
`permissions.allow` do projeto (`.claude/settings.local.json`) ou do usuário
(`~/.claude/settings.json`) dispensa a confirmação, sem liberar `rm` para mais nada:

```json
{ "permissions": { "allow": ["Bash(rm -rf /tmp/specflow-upgrade)"] } }
```

## Etapa 3 — Classificar

```bash
sh .claude/skills/spec-upgrade/classificar.sh /tmp/specflow-upgrade .
```

Cada linha é `<situação> <caminho>`:

| Situação | O que é | Ação |
|---|---|---|
| `igual` | idêntico à versão nova | nada |
| `novo` | existe no framework, não no projeto | copiar (Regra 3) |
| `desatualizado` | versão antiga do framework, sem edição | sobrescrever (Regra 3) |
| `personalizado` | editado no projeto | perguntar (Etapa 4) |
| `removido` | saiu do framework, sem edição | perguntar (Etapa 4) |
| `removido-personalizado` | saiu do framework, mas foi editado no projeto | perguntar (Etapa 4) |
| `do-projeto` | nunca foi do framework | nada (Regra 2) |

Sem `.claude/` no projeto, tudo sai `novo`: é uma instalação, e segue igual. Tudo `igual` →
diga que o projeto já está na versão nova, rode `rm -rf /tmp/specflow-upgrade` e pare.

O próprio `spec-upgrade` entra na lista como qualquer outra skill. A versão nova dele vale a
partir da próxima execução.

## Etapa 4 — Perguntar o que não é seguro

Só se houver `personalizado`, `removido` ou `removido-personalizado`. Antes de perguntar, leia
`diff -u /tmp/specflow-upgrade/<caminho> <caminho>` de cada personalizado e resuma em uma linha
o que o projeto mudou (ex.: "acrescentou a regra 6, sobre migrations"). Use `AskUserQuestion`,
até 4 perguntas por chamada (mais arquivos → mais chamadas):

- **Cada `personalizado`**: *Mesclar* (recomendado: a versão nova do framework com a edição do
  projeto reaplicada) · *Substituir* (perde a edição) · *Manter a do projeto* (fica sem a
  atualização).
- **Os `removido`, numa pergunta só**, listando os arquivos: *Apagar* (recomendado) · *Manter*.
- **Cada `removido-personalizado`**: *Apagar* · *Manter* (recomendado: tem edição do projeto).
  Diga se o framework o trocou por outro arquivo, quando o diff mostrar (ex.:
  `agent-spec-adr` → `spec-adr`).

**Mesclar** = partir do arquivo novo do framework e reaplicar nele só o que o projeto
acrescentou ou mudou de propósito. O que já veio no framework novo não se duplica. Conflito real
(o framework mudou o mesmo trecho em outro sentido) → mostre os dois trechos e pergunte.

## Etapa 5 — Aplicar

```bash
# novo, desatualizado, substituir:
mkdir -p <pasta do caminho> && cp /tmp/specflow-upgrade/<caminho> <caminho>
# removido com "Apagar":
rm <caminho>
```

Escreva a pasta do `mkdir -p` por extenso (ex.: `.claude/skills/spec-adr`), não
`$(dirname …)`: a regra "sem `$(...)`" (Regra 7) vale para todos os comandos.

Depois de apagar, remova as pastas que ficaram vazias, só nas pastas do framework
(`find .claude/skills .claude/agents .claude/rules -type d -empty -delete`).

Confira: rode o `classificar.sh` de novo, antes de apagar o clone. Só podem sobrar `igual`,
`do-projeto` e o que o usuário mandou manter. Sobrou outra coisa → diga qual e por quê.

Conferido, apague o clone: `rm -rf /tmp/specflow-upgrade`.

## Etapa 6 — Relatório

Para cada skill nova, uma linha com o que ela faz (tirada do `description` dela). Depois,
**apenas**:

```
SpecFlow: <hash> (<data>) — <assunto do último commit>
Novos: <caminhos> | nenhum
Atualizados: <caminhos> | nenhum
Personalizados: <caminho> → mesclado | substituído | mantido
Apagados: <caminhos> | nenhum · Mantidos a pedido: <caminhos>
Do projeto, intocados: <caminhos> | nenhum
Sem commit: confira com `git diff -- .claude` e rode /spec-semantic-commit
```

# Entrada

$ARGUMENTS
