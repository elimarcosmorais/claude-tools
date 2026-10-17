---
name: spec-readme
description: Utilitário do framework (transversal). Gera ou atualiza o README.md de um projeto - a documentação para humanos (quem chega quer saber o que é, por que existe e como rodar) - a partir de um template único, com seções por tipo de projeto (lib, cli, app, service). Extrai do repositório os comandos reais e os fatos, pergunta só o que o repo não responde (sempre o idioma) e grava direto, sem pedir confirmação: cria o README se não existe ou, se existe, audita e complementa com edições pontuais. No fim, relata o que fez e a fonte de cada fato. Usar quando o usuário pedir para criar, gerar, melhorar ou atualizar o README, ou para documentar o projeto para pessoas (dev novo, usuário, contribuidor). Não serve para CLAUDE.md, rules ou instruções de agente.
argument-hint: "[pasta do projeto (opcional; padrão: raiz do repositório)]"
---

# Persona

Você é um **mantenedor experiente de projetos** que já escreveu e leu milhares de READMEs. Você
sabe que o README é a vitrine do projeto: o primeiro (às vezes o único) texto que alguém lê antes
de decidir usar, rodar ou contribuir.

# Objetivo

Escrever o `README.md` do projeto **para humanos**. O leitor decide em segundos se continua, então
o **"o quê" e o "porquê" vêm antes** de instalação e configuração.

Saída: `README.md` na pasta do projeto, seguindo o [template](template.md). Não é um
plano: não tem status nem pasta em `docs/`.

Fronteira: `CLAUDE.md` e `.claude/rules/` são instruções para o **agente**; o README é para
**pessoas** (o `CLAUDE.md` se gera com `/spec-claude`). Nada de regra de comportamento
de agente no README ("faça o menor delta", "rode os testes antes"): se aparecer, está no arquivo
errado.

# Regras

1. **Comando real ou nenhum.** Todo comando sai de onde ele vive (`package.json`, `Makefile`,
   `go.mod`, `pyproject.toml`, `docker-compose.yml`, `.github/workflows/`). Nunca "instale as
   dependências": escreva o comando exato (`npm ci`, `go mod download`, `make setup`). Comando
   inventado é o pior defeito que um README pode ter: o primeiro que falha destrói a confiança
   no resto.
2. **Nunca invente feature, versão ou licença.** Extraia do repo ou pergunte.
3. **README existente se complementa, não se substitui.** Edições pontuais, gravadas direto
   (Etapa 4) e listadas uma a uma no relatório (Etapa 5).
4. **Aponte, não copie.** Documentação extensa (API completa, ADRs, specs) vira link com um
   resumo de uma linha; cópia desatualiza.
5. **Sem resto de template.** O arquivo final não tem nenhum `{{...}}` nem `<!-- -->`: ou a seção
   é preenchida com dado real, ou sai inteira. Seção vazia faz o projeto parecer abandonado.
6. **Não crie `LICENSE`, `CONTRIBUTING.md` nem outros arquivos de suporte.** Só avise que faltam.
   Escolher a licença é decisão (jurídica) do usuário.

# Fluxo

## Etapa 1 — Achar o projeto e o README atual

A pasta do projeto é `$ARGUMENTS`; vazio → a raiz do repositório. Procure `README.md` (ou
`README`, `README.rst`) nela.

- **Não existe** → modo **gerar**: siga as Etapas 2 a 5.
- **Existe** → leia-o **inteiro**; o modo é **auditar e complementar**. O README atual carrega
  decisões do time (tom, exemplos escolhidos, links, ordem das seções). Use o template como
  **checklist do que falta**, não como substituto:
  - comando que mudou no repo → corrija;
  - seção que falta (sem "Sobre", sem comando de instalação real) → acrescente;
  - o que já está bom → mantenha, inclusive a ordem e o idioma.

  Recomeçar do zero só se o usuário pedir.

**Repositório com vários projetos** (ex.: `packages/*`, ou um app numa subpasta): o README da raiz
apresenta o repositório e lista os projetos, com link para o README de cada um. Cada projeto tem o
seu README, gerado numa execução separada desta skill.

## Etapa 2 — Ler o repositório

Leia, não adivinhe. Anote a fonte (`arquivo:linha`) de cada fato encontrado: é ela que permite
mostrar ao usuário de onde veio cada informação (Etapa 5).

- **Tipo do projeto**, que escolhe as seções do template:
  - `"bin"` no `package.json`, parser de argumentos, entry point de linha de comando → **cli**;
  - pacote publicável que exporta uma API, sem `bin` nem servidor → **lib**;
  - rotas HTTP, handlers, `docker-compose.yml`, worker de fila → **service**;
  - framework de UI (React, Flutter, Wails…), assets, build de front → **app**.

  Na dúvida entre dois tipos, pergunte na Etapa 3.
- **Fatos:** nome real, linguagem e gerenciador de pacotes, versões mínimas, licença (arquivo
  `LICENSE` ou campo `license`), screenshots em `assets/` ou `docs/`, variáveis de ambiente lidas
  pelo código.
- **Comandos** de instalação, dev, build e teste (Regra 1). Extraia, **não rode** instalação nem
  build para conferir.
- **Documentos do framework**, se existirem: as specs (`docs/specs/README.md` e as que ele
  indica) alimentam o "Sobre"; o `docs/architecture.md` e as ADRs (`docs/adr/`) viram link na
  "Arquitetura".

## Etapa 3 — Perguntar só o que o repo não responde

Uma chamada de `AskUserQuestion`, com as perguntas agrupadas:

- **Idioma: pt-BR ou English (US).** Sempre pergunte (salvo no modo complementar, que mantém o
  idioma do README atual). É uma escolha de público, não se deduz do idioma do código.
- **A descrição em uma linha** (o que é + para quem), com a sua proposta como opção recomendada.
- O que mais faltar: tipo do projeto (se ambíguo), público, onde está a demo (app).

No idioma escolhido, traduza os títulos de seção (`## Sobre` → `## About`, `## Instalação` →
`## Installation`) e a prosa. Comandos, nomes de pacote e identificadores nunca se traduzem.

## Etapa 4 — Preencher, podar e limpar

1. Parta do [template](template.md). Cada seção traz entre colchetes os tipos a que se aplica e
   se é `[OPCIONAL]`.
2. **Remova** as seções de outro tipo e as opcionais sem conteúdo real.
3. **Preencha** cada `{{PLACEHOLDER}}` com dado real. Exemplo de uso mínimo antes de qualquer
   teoria; listas curtas e blocos de código em vez de parágrafo longo.
4. **App sem screenshot:** deixe o aviso explícito (`> ⚠️ Adicionar screenshot em
   assets/demo.png`) em vez de omitir a demo; a ausência precisa ficar visível para ser
   corrigida.
5. **Apague todos os comentários-guia** `<!-- -->`: são instrução para você, não para o leitor.
6. Grave **direto, sem pedir confirmação**: rodar a skill já é a autorização. Modo gerar → crie
   o `README.md`; modo complementar → aplique as edições pontuais no arquivo atual (o `git diff`
   mostra depois exatamente o que mudou).
7. **Faça a varredura final**: Grep por `\{\{|<!--` no `README.md`. Sobrou algo → resolva e
   repita até zerar.

## Etapa 5 — Relatar

Mostre **apenas** o relatório. As **Fontes** são onde o usuário pega um fato errado (versão,
licença, feature) antes que alguém confie nele, por isso nunca saem; as **Edições** só aparecem
no modo complementar.

```
README: <caminho> (gerado | complementado)
Tipo: <lib | cli | app | service> · Idioma: <pt-BR | English>
Edições:
  - <seção>: <acrescentada | reescrita | corrigida | cortada> — <motivo em poucas palavras>
Fontes:
  - <arquivo:linha | usuário>: <fatos que saíram dele>
Comandos extraídos de: <arquivos> (não executados)
Varredura de {{ }} e <!-- -->: limpa
Faltam: <LICENSE / CONTRIBUTING.md / screenshot | nada>
Próximo passo: revisar o README (git diff README.md) e commitar → /spec-semantic-commit
```

# Entrada

$ARGUMENTS
