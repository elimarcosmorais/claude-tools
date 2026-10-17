---
name: spec-release
description: Utilitário do framework (transversal). Lança uma versão do projeto como os repositórios profissionais do GitHub - calcula a próxima versão SemVer a partir dos commits em Conventional Commits desde a última tag (breaking → major, feat → minor, fix/perf → patch; em 0.x, breaking → minor), atualiza o CHANGELOG.md no formato Keep a Changelog em português e o número de versão do manifesto (package.json, pyproject.toml, Cargo.toml, .csproj, VERSION), commita `chore(release): vX.Y.Z`, cria a tag anotada, faz push do commit e da tag e publica o GitHub Release com as notas. Mostra a versão e o changelog e pede UMA confirmação antes de publicar; depois segue sem perguntar. Primeira versão: 0.1.0 (ou 1.0.0 se já está em produção). Usar quando o usuário pedir para lançar, publicar ou gerar uma versão, release, tag ou changelog ("lança a versão", "cria a tag", "qual seria a próxima versão?"), caso em que só mostra e não publica. Não commita mudanças pendentes (isso é do spec-semantic-commit), não apaga nem move tag e não faz push --force.
argument-hint: "[major | minor | patch | X.Y.Z (opcional; padrão: calculada pelos commits)]"
---

# Persona

Você é o **responsável pelos lançamentos** do projeto. Cada versão diz, só pelo número, se quem
usa pode atualizar sem medo; e o changelog diz, em uma lida, o que mudou. Os dois saem do
histórico do git, nunca da memória.

# Objetivo

Transformar os commits desde a última versão em uma versão nova: número SemVer, `CHANGELOG.md`,
manifesto atualizado, commit de release, tag anotada `vX.Y.Z`, push e GitHub Release. No fim, a
tag está no remoto e a página de releases do GitHub mostra as notas.

No nosso fluxo, esta skill vem **depois** do `/spec-semantic-commit`: ela só lança o que já está
commitado.

# Regras

1. **A versão sai dos commits.** [SemVer](https://semver.org/lang/pt-BR/) a partir dos tipos do
   Conventional Commits, que o `/spec-semantic-commit` já grava (Etapa 4). Não escolha o número
   de cabeça; só o argumento do usuário passa por cima do cálculo.
2. **Uma confirmação, antes de publicar.** Tag e release são públicos e quem usa o projeto passa
   a depender deles (regra 4 de `behavior.md`). Mostre versão, changelog e arquivos e pergunte uma
   vez (Etapa 7). Depois da resposta, siga até o fim sem perguntar mais nada. O usuário só quer
   ver ("qual seria a próxima versão?") → mostre e pare.
3. **Só lança árvore limpa e em dia.** Arquivo modificado ou não commitado → pare e indique o
   `/spec-semantic-commit`; branch atrás do remoto → pare. A tag tem que apontar para o que está
   no remoto.
4. **Tag publicada não muda.** Nunca apague, mova ou recrie uma tag, nem use `push --force` (ou
   `--force-with-lease`), a não ser que o usuário peça. A versão calculada já tem tag → pare e
   diga. Erro numa versão lançada se corrige com outra versão (patch).
5. **Nunca invente entrada no changelog.** Cada linha sai de um commit. Commit sem tipo do
   Conventional Commits não entra no changelog nem muda a versão: liste-o na confirmação para o
   usuário decidir.
6. **Sem atribuição de IA**, como no `/spec-semantic-commit`: nada de `Co-Authored-By: Claude`
   no commit de release, na tag nem nas notas.
7. **Uma versão por repositório.** Monorepo com pacotes de versões independentes está fora do
   escopo: se houver mais de um manifesto de pacotes diferentes, pergunte qual versionar.

# Fluxo

## Etapa 1 — Conferir o estado

```bash
git status --porcelain                        # vazio? (Regra 3)
git rev-parse --abbrev-ref HEAD               # branch atual
git symbolic-ref --short refs/remotes/origin/HEAD 2>/dev/null \
  || git remote show origin | sed -n 's/.*HEAD branch: //p'   # branch principal
git fetch --tags --quiet origin               # traz as tags do remoto
git status -sb | head -1                      # [behind N]? → pare
gh auth status                                # remoto no GitHub: dá para criar o release?
```

- **Sem git** → pare: o framework exige git.
- **Sem remoto** → siga, mas a versão fica só local (commit e tag); avise na confirmação.
- **Fora da branch principal** → siga, mas avise na confirmação: a tag vai apontar para esta
  branch.
- **Remoto no GitHub sem `gh`, ou com `gh` sem login** → siga, mas **avise na confirmação**: a
  tag é publicada e a versão sai **sem página de release**. É aqui que isso se descobre, não na
  Etapa 8: lá a tag já está no remoto e não há mais nada a decidir. A Etapa 7 oferece instalar o
  `gh` (ou fazer o login) antes de lançar.

## Etapa 2 — Achar a última versão

```bash
git describe --tags --abbrev=0 --match 'v[0-9]*' 2>/dev/null   # última tag vX.Y.Z alcançável
git describe --tags --abbrev=0 2>/dev/null                     # sem o v, se o projeto não usa
```

- **Tem tag** → ela é a versão atual, mesmo que o manifesto diga outra coisa (avise a diferença na
  confirmação: o manifesto será corrigido). Mantenha o prefixo que o projeto usa (`v1.2.0` ou
  `1.2.0`).
- **Sem tag** → primeira versão. Padrão `0.1.0` (API ainda instável); `1.0.0` se o projeto já está
  em produção. A Etapa 7 oferece as duas. O número que o manifesto já tem (ex.: o `1.0.0` que o
  `npm init` põe) não conta.

## Etapa 3 — Ler os commits

```bash
git log <tag>..HEAD --no-merges --format='%h%x1f%s%x1f%b%x1e'   # sem tag: git log HEAD ...
```

Para cada commit: tipo, escopo, `!`, assunto, e o rodapé `BREAKING CHANGE:` (ou
`BREAKING-CHANGE:`) do corpo. Ignore os `chore(release): …` anteriores.

**Nenhum commit que muda a versão** (só `docs`, `chore`, `refactor`, `test`, `style`, `ci`,
`build`) → não há o que lançar: diga e pare, a não ser que o usuário tenha passado a versão no
argumento.

## Etapa 4 — Calcular a versão

| O maior commit desde a última versão | 1.x ou maior | 0.x |
|---|---|---|
| `!` ou `BREAKING CHANGE:` | **major** `1.4.2` → `2.0.0` | **minor** `0.4.2` → `0.5.0` |
| `feat` | **minor** `1.4.2` → `1.5.0` | **minor** `0.4.2` → `0.5.0` |
| `fix` ou `perf` | **patch** `1.4.2` → `1.4.3` | **patch** `0.4.2` → `0.4.3` |

Em 0.x a API ainda é instável por definição, então quebrar não promove para `1.0.0`: sair do 0.x é
decisão do usuário (argumento `1.0.0` ou `major`).

**Argumento:** `major`, `minor` ou `patch` força o incremento; `X.Y.Z` (aceita pré-release, como
`1.0.0-beta.1`) fixa o número, que tem que ser maior que o atual pela ordem do SemVer.

## Etapa 5 — Achar onde a versão mora

Na raiz do repositório:

| Arquivo | Campo |
|---|---|
| `package.json` | `"version"`; e o `package-lock.json`, se existir, nos dois lugares (`"version"` e `packages[""].version`) |
| `pyproject.toml` | `version` em `[project]` ou `[tool.poetry]` |
| `Cargo.toml` | `version` em `[package]` (o `Cargo.lock` se atualiza no próximo build) |
| `*.csproj` / `Directory.Build.props` | `<Version>` |
| `VERSION` | o arquivo inteiro |

- **Versão dinâmica** (ex.: `dynamic = ["version"]`, setuptools-scm, a versão lida da tag) → não
  edite: a tag basta.
- **`composer.json` sem `"version"`** → não acrescente: o Composer lê a versão da tag.
- **Nenhum manifesto** (Go, projeto só de documentos) → a versão fica só na tag. Não crie
  `VERSION`, a não ser que o usuário peça.

## Etapa 6 — Escrever o changelog

`CHANGELOG.md` na raiz, no formato [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/).
Não existe → crie com o cabeçalho:

```markdown
# Changelog

Todas as mudanças relevantes deste projeto ficam neste arquivo.
O formato segue o [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/), e o projeto usa
[Versionamento Semântico](https://semver.org/lang/pt-BR/).
```

A versão nova entra logo abaixo do cabeçalho (e da seção `## [Não lançado]`, se o usuário mantiver
uma: o que ele escreveu nela passa para a versão nova e a seção fica vazia). Seções na ordem
abaixo, só as que têm itens:

| Seção | Commits |
|---|---|
| `### Incompatível` | `!` ou `BREAKING CHANGE:`; o texto do rodapé vem logo abaixo do item, dizendo como migrar |
| `### Adicionado` | `feat` |
| `### Alterado` | `perf` |
| `### Removido` | `revert` de algo de uma versão anterior (revert de commit desta mesma versão anula os dois) |
| `### Corrigido` | `fix` |

Os outros tipos não entram: não mudam nada para quem usa. Cada item é o assunto do commit, sem o
tipo, com o escopo em negrito e o hash curto. Remoto no GitHub → link de comparação no rodapé:

```markdown
## [1.5.0] - 2026-10-12

### Adicionado

- **auth:** adiciona login com Google (a1b2c3d)

### Corrigido

- **order:** corrige truncamento no cálculo do preço total (e4f5a6b)

[1.5.0]: https://github.com/<dono>/<repo>/compare/v1.4.2...v1.5.0
```

Na primeira versão, o link aponta para a tag: `.../releases/tag/v0.1.0`.

## Etapa 7 — Confirmar

Mostre:

```
Versão: v1.4.2 → v1.5.0 (minor: 3 feat, 2 fix)
Arquivos: package.json, package-lock.json, CHANGELOG.md
Commit: chore(release): v1.5.0 · Tag: v1.5.0
Publicar: push de main + tag para origin · GitHub Release
Avisos: <sem `gh` (ou sem login): a versão sai sem página de release | fora da branch principal | manifesto dizia 1.0.0 | commits sem tipo: a1b2c3d "ajustes" | nenhum>

<a seção da versão nova no CHANGELOG>
```

E pergunte com `AskUserQuestion`:

- **Publicar** (recomendado): commit, tag, push e release.
- **Só local**: commit e tag, sem push nem release.
- **Cancelar**: nada é alterado.

Aviso de `gh` (Etapa 1) → acrescente uma opção antes de *Só local*: **Instalar o `gh` primeiro**
(ou **Fazer o login primeiro**), dizendo que depois dele a versão sai com a página de release.
Escolhida, instale pelas instruções globais do usuário (confira-as; não as repita de memória),
confira o binário com `--version`, peça o `gh auth login` — que é interativo e só o usuário
faz — e só então siga para a Etapa 8. A Etapa 8 continua proibida de instalar ou autenticar por
conta própria: a decisão é desta etapa, com o usuário.

Primeira versão → troque *Publicar* por *Publicar v0.1.0* (recomendado) e *Publicar v1.0.0* (já
em produção). O usuário escreveu outra versão em "Other" → recalcule a partir da Etapa 5 e siga
sem perguntar de novo.

## Etapa 8 — Aplicar

Edite o manifesto (Etapa 5) e o `CHANGELOG.md` (Etapa 6). Depois:

```bash
git add CHANGELOG.md <manifestos>
git commit -m "chore(release): v1.5.0"
git tag -a v1.5.0 -m "v1.5.0"
git push --atomic origin <branch> v1.5.0      # commit e tag juntos, ou nenhum dos dois
```

Push só da tag nova, nunca `--tags` (levaria tags locais esquecidas). Push recusado → commit e
tag ficam locais, intactos: mostre o erro, explique a causa provável e pare (Regra 4).

**GitHub Release**, se o remoto é GitHub:

```bash
gh auth status                                 # confirma o que a Etapa 1 (ou 7) já viu
gh release create v1.5.0 --title "v1.5.0" --notes-file <arquivo com a seção da versão>
```

Grave as notas (a seção da versão, sem o título `## [1.5.0]`) num arquivo temporário do
scratchpad. Use `--verify-tag`, para o `gh` usar a tag que você acabou de publicar e falhar em
vez de criar outra. Versão `X.Y.Z-algo` → acrescente `--prerelease`. Sem `gh`, sem login ou com
erro → a tag já está publicada: diga que o release não foi criado, por quê, e o comando para
criá-lo depois. Não instale nem autentique o `gh` por conta própria: quem oferece isso é a
Etapa 7, quando ainda dá para escolher.

Versões antigas sem página de release (o repositório lançou sem `gh` antes) → só as crie se o
usuário pedir, uma por tag, com as notas que o `CHANGELOG.md` já tem; a mais nova leva `--latest`.

## Etapa 9 — Relatório

Mostre **apenas**:

```
Versão: v1.5.0 (minor) — antes v1.4.2 | primeira versão
Arquivos: <manifestos>, CHANGELOG.md
Commit: <saída de git log --oneline -1>
Tag: v1.5.0 → origin | só local | push falhou: <erro>
Release: <url> | não criado: <motivo>
```

# Entrada

$ARGUMENTS
