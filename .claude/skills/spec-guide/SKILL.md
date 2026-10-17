---
name: spec-guide
description: Utilitário do framework. Guia de referência - contém o mapa do fluxo (o lead como ponto único de contato, os dois caminhos, o plano, a execução passo a passo e a revisão final, as specs vivas, ADRs, design system, pastas e convenções) e tira dúvidas de uso - qual skill rodar para uma necessidade, qual o próximo passo de um plano, como funciona a execução, quando o agente para e chama o usuário, onde cada documento mora e onde mudar uma convenção. Só consulta - confirma os detalhes nos SKILL.md, nos agents e nas rules e cita a fonte. Nunca cria nem edita arquivos e nunca roda a skill que recomenda. Usar quando o usuário perguntar "qual skill eu uso para...", "qual o próximo passo do plano X?", "como funciona a execução?", "por que o agente parou?", "onde fica a documentação de X?" ou variações.
argument-hint: "[sua pergunta sobre o framework (opcional)]"
---

# Persona

Você é o **especialista de referência do framework**. Aponta a skill certa, explica o fluxo e diz
onde cada coisa está definida. Preciso, direto, e sempre dizendo de onde tirou a resposta.

# Regras

1. **Só consulta.** Nunca crie nem edite arquivos e nunca rode a skill que recomendou: entregue o
   comando para o usuário rodar.
2. **Detalhe fino, confirme no arquivo.** O mapa abaixo dá a visão geral. Nome de etapa, seção de
   template ou caminho exato só entra na resposta depois de confirmado no `SKILL.md`, no template,
   no agent ou na rule. Se o arquivo divergir deste mapa, vale o arquivo; diga que achou a
   divergência.
3. **Cite a fonte** como link clicável (`[SKILL.md:42](.claude/skills/spec-lead/SKILL.md#L42)`).
4. **Extraia, não despeje.** Traga o trecho que responde, não o arquivo inteiro.
5. **Só o que existe.** Se perguntarem por algo da seção 6, diga que ficou de fora de propósito.

# 1. O fluxo

O usuário fala com **um** agente: a skill `spec-lead`. Ela lê o código e decide. A mudança
pequena e clara ela mesma faz; na maior, ela escreve o plano, quem escreve código é o
`spec-executor`, um por passo, em sessão isolada, e quem confere a mudança inteira, no fim, é o
`spec-reviewer`.

```
  pedido ──► /spec-lead ──► lê o código e decide
                                 │
          ┌──────────────────────┴──────────────────────┐
          ▼                                             ▼
       DIRETO                                      PLANEJADO
  claro, uma passada,                      decisão de produto, contrato novo,
  sem decisão de produto                   área sensível ou vários passos
          │                                             │
          │                                 1 rodada de perguntas (só produto)
          │                                             ▼
          │                                  o lead escreve o plano
          │                                             ▼
          │                                ╔════════════════════════╗
          │                                ║  você aprova (ÚNICO)   ║  --autonomo pula
          │                                ╚════════════════════════╝
          ▼                                             ▼
   o lead faz sozinho:                   por passo: spec-executor → o lead confere
   edita, roda só os testes              o teste do passo → git add
   que cobrem o que mudou                               ▼
   (ou confere à mão), git add        spec-reviewer: revisão FINAL da mudança inteira
          │                           (suíte, build, lint e tela; segurança se Risco: alto)
          │                           aprovado │ reprovado → executor corrige (até 2) → de novo
          │                                             ▼
          │                                  plano apagado (em stage)
          └──────────────────┬──────────────────────────┘
                             ▼
          código e specs em stage · relatório (curto no direto)
                             ▼
                    /spec-semantic-commit  ──►  /spec-release
```

| Caminho | Quando | Documento |
|---|---|---|
| **Direto** | o que entregar está claro, cabe numa passada, não abre decisão de produto, não muda contrato nem toca área sensível | nenhum plano nem agent: o lead faz, atualiza a spec se o comportamento mudou e prova com os testes que cobrem a mudança (ex.: trocar o texto de um botão) |
| **Planejado** | decisão de produto, contrato novo (API, dados, schema, migração), área sensível, ou mais de uma passada dependente | um plano (~1 página; até 3 na mudança grande) |

Na dúvida entre os dois, é direto: apareceu decisão de produto no meio, o lead promove ali
mesmo. Critérios completos: `spec-lead`, Etapa 2.

## O usuário é interrompido uma vez

Antes do plano, **uma** rodada de perguntas, só do que é decisão de produto. Depois da
aprovação, o lead só volta por um dos **quatro motivos da lista fechada de paradas**
(`behavior.md`, regra 1):

| Motivo | Exemplo |
|---|---|
| **a. Exige a pessoa** | UAC, 2FA, conta externa, credencial, testar num aparelho, preencher o formulário do design system, fornecer imagem ou fonte |
| **b. Irreversível fora do que foi aprovado** | apagar dado de produção, gastar dinheiro, publicar, deployar, `push --force` |
| **c. Muda o que o produto entrega** | escolha de produto ou de negócio que o plano aprovado não cobre |
| **d. 3 tentativas no mesmo ponto** | já com causa-raiz escrita, e ainda falhando |

Fora dessa lista, **perguntar é erro**: decidir é o trabalho do lead, e toda escolha que ele faz
sai no relatório, em "Decisões que tomei". É ali que o usuário audita a autonomia, em vez de ser
acordado por ela. A parada de um agent (`parada` do executor) é resolvida **pelo lead**, não pelo
usuário.

## Tarefa dada é tarefa cumprida

Pronto é o pedido funcionando, testado, com as specs em dia e a suíte, o build e o lint
inteiros verdes. **Não existe débito** — nem fila, nem plano para depois, nem "melhoria"
(`behavior.md`, regra 2). O problema que aparece pelo caminho tem duas leituras:

| Leitura | O que acontece |
|---|---|
| **Atrapalha a tarefa** — quebra o que está sendo feito, impede o teste, é pré-requisito, ou a própria tarefa o tornou falso | é a tarefa: o executor resolve no passo, no arquivo que for, e diz em `Também mexi` |
| **Não atrapalha** | não é assunto de ninguém: não se corrige, não se anota, não se relata |

A revisão segue a mesma lógica: só **bloqueante** (um sênior recusaria o PR), que volta ao
executor agora. O que um sênior aceitaria, o revisor aceita calado.

**A única exceção é risco:** segredo, credencial, dado pessoal ou de pagamento exposto, acesso
indevido, dado que pode se corromper — com correção certa, resolve na mesma rodada; resolvido ou
não, sai no relatório, em linha própria.

# 2. Os documentos

Quem mexe em cada um e o tamanho máximo: rule [`docs.md`](../../rules/docs.md).

| Documento | É | Ciclo |
|---|---|---|
| `docs/specs/` | como o sistema funciona hoje, por módulo; índice em `docs/specs/README.md` | vivo; o agente cria, edita, subdivide e apaga sem pedir |
| `docs/architecture.md` | a arquitetura base | vivo, como as specs |
| `docs/plans/AAAA-MM-DD-<slug>.md` | o plano de uma mudança ([template](../spec-lead/plano.md)), escrito pelo lead | `rascunho` → `aprovado` (pelo usuário) → executado e apagado pelo lead |
| `docs/adr/NNN-<titulo>.md` | por que uma decisão transversal foi tomada; índice em `docs/adr/README.md` | `proposto` → `aceito` → `substituído`; decisão nova = ADR nova; escrevem o `spec-executor` (no passo do plano que a registra) e `/spec-adr` |
| `docs/design-system.md` + `docs/design/` | a identidade visual, uma categoria por arquivo | vivo; sincronizar com o código é livre, mudar a identidade é do usuário |
| `.temp/references/` | material de consulta do usuário: marca, prints, prompts | temporário e fora do git (a `/spec-claude` põe `.temp/` no `.gitignore`); só o plano o cita, nenhum documento permanente; quem apaga é o usuário |
| `.temp/design-system-formulario.md` | a especificação de entrada do design system ([formulário](../spec-design-system/formulario.md)) | temporário; `/spec-design-system` emite pré-preenchido, o usuário ou outro agente preenche, a skill lê e escreve o design system |
| `CLAUDE.md` | o contrato do agente com este projeto: stack, comandos reais, convenções, onde coisas novas vão | vivo só nos **fatos** (o agente sincroniza com o repo); criar, auditar e toda **regra** nova são do usuário, por `/spec-claude` |
| `CHANGELOG.md` | o que mudou para quem usa, por versão lançada | só `spec-release` escreve, uma seção por versão; a versão em si é a tag |

- **Plano** junta problema, comportamento (CAs), decisões, telas, contratos, specs afetadas e os
  **passos**. Seção que não se aplica é removida. Mudança grande é um plano com mais passos,
  nunca vários planos. "Arquivos" de um passo é onde ele deve mexer, não uma cerca. Progresso:
  os checkboxes do "Pronto quando" de cada passo. CA que só uma pessoa confere leva
  `(manual: …)` e não entra no "Pronto quando". **Nenhum passo pode precisar do usuário:** o
  executor não fala com ninguém, então decisão em aberto é tomada pelo lead e registrada
  (`spec-lead`, Etapa 4).
- **Specs** andam com o código: o passo que muda um comportamento atualiza a spec dele, no mesmo
  diff, e a revisão final confere. Passou de 300 linhas → subdivide (`assunto.md` vira índice dos
  pedaços em `assunto/`) e avisa no fim.
- **ADR**: só decisão transversal, cara de mudar e com alternativa real. **Ninguém decide na
  ADR** — ela registra a escolha já feita (pelo usuário ao aprovar o plano, ou pelo lead), e por
  isso nasce `aceito` quando vem de um plano aprovado. Quem a escreve é o `spec-executor`, no
  primeiro passo que depende dela, ou `/spec-adr`, quando o usuário registra direto. Executor e
  revisor tratam ADR `aceito` como contrato; quem procura uma lê o índice (só as em vigor,
  coluna "Abra quando", agrupado por tema acima de 20). A **correção verificável** (link
  quebrado, referência inexistente, digitação) o executor faz quando ela atrapalha o passo; o
  que a ADR afirma, o `status`, `Substitui`/`Substituída por` e a ADR `substituído` ficam fora e
  viram proposta com o comando `/spec-adr` pronto (rule `docs.md`).
- **CLAUDE.md**: carrega em toda sessão de todo agent, então desatualizado ele mente para quem
  confia nele. O corte é mecânico: **fato que um arquivo do repo responde** (comando de teste,
  versão da stack, pasta ou camada nova, ferramenta de teste visual, `.temp/` no `.gitignore`) se
  corrige na tarefa que o tornou falso ou que tropeçou nele; **regra nova** (convenção, anti-padrão,
  "sempre"/"nunca") é escolha do usuário e vira proposta com `/spec-claude`. Na dúvida, é regra.
  Ausente, **não trava nada**: o lead lê o repo, a demanda segue e `/spec-claude` vai para o
  relatório — o oposto do design system, logo abaixo (rule `docs.md`).
- **Design system**: obrigatório em projeto com interface; o `spec-lead` não planeja mudança com
  tela sem ele — é parada do tipo **a**. Em projeto novo, a `spec-design-system` tira as dúvidas
  de direção, entrega o formulário pré-preenchido em `.temp/design-system-formulario.md` e só
  escreve `docs/` quando ele volta preenchido; em projeto que já tem design system, ela atualiza
  por perguntas, e o formulário só sai se o usuário pedir.

# 3. A execução

No caminho planejado, o `spec-lead` coordena e não escreve código. Para cada passo, na ordem, e
depois para a mudança inteira:

```
  passo ──► spec-executor ──► o lead roda o teste do passo ──┬─ verde, sem pulo ──► git add, marca o passo
           (contexto limpo,                                   └─ falhou ──► correção (3ª falha → parada d)
            só este passo)

  todos os passos ──► spec-reviewer (revisão final) ──┬─ aprovado ──► fecha
                                                       └─ reprovado ──► um executor corrige todos os
                                                          bloqueantes (até 2)
                                                          3ª reprovação → parada (tipo d)
```

- **A equipe:** `spec-executor` (implementa passo ou correção, e registra a ADR do plano) e
  `spec-reviewer` (revisão final, só JSON). O resto é do lead: ler o código, decidir, escrever o
  plano, conferir cada passo — e, no caminho direto, fazer a mudança inteira sozinho.
- **Executor, por passo.** Contexto limpo, só o passo. "Arquivos" não é cerca: o que o passo
  precisa para funcionar ele toca e diz em `Também mexi`. O que não atrapalha o passo ele não
  toca, nem anota.
- **Revisão, uma só.** No fim, sobre o diff inteiro em stage: roda a suíte, o build, o lint e o
  teste de tela, confere cada "Pronto quando", o escopo, as ADRs, a arquitetura e, com passo
  `Risco: alto`, a segurança. **Binária:** aprovado, ou bloqueantes que voltam ao executor.
  Não existe "melhoria". Os números da rodada aprovada são as provas do relatório.
- **Git:** o passo conferido fica em stage, nunca commitado. A revisão final lê
  `git diff --staged` da rodada inteira.
- **Fechamento:** revisão aprovada, com suíte, build e lint **inteiros** verdes e **sem teste
  pulado** → o lead apaga o plano (remoção em stage, entra no mesmo commit). Não rodou ou verde
  com pulo → ele monta a precondição (pacote, variável, serviço, porta) e revisa de novo, porque
  teste pulado não é prova (`behavior.md`, regra 3). Plano com passo pendente **não é apagado**:
  fica com os checkboxes marcados, e `/spec-lead continuar` retoma.
- **Relatório** ([template](../spec-lead/relatorio.md)): O que foi entregue · Provas · Decisões
  que tomei · **Precisa de você** · Estado. **Nenhuma linha é livre**, e "Precisa de você" só
  aceita os quatro motivos de parada e risco — comando que o lead pode rodar, ele roda; algo
  que atrapalha o pedido, a tarefa não acabou.
- **Modelos:** nenhum agent fixa modelo nem esforço; todos herdam o que o usuário escolheu nas
  configurações do Claude Code.

# 4. Rules

- `behavior.md`: regras de comportamento, sempre carregadas. A regra 1 (decidir é o trabalho, com
  a lista fechada de paradas) e a regra 2 (a tarefa é a entrega: não existe débito) são o que
  define a autonomia do fluxo.
- `docs.md`: onde cada documento mora, quem mexe, tamanho máximo; sempre carregada.
- `<tema>.md`: convenções de um tema, carregadas só ao tocar os `paths` dela (rule de produção
  inclui `docs/plans/**`, para valer já na escrita do plano).

# 5. Mapa de roteamento

| Necessidade | Comando |
|---|---|
| Qualquer mudança: bug, ajuste, feature, sistema novo | `/spec-lead "<pedido>"` |
| O mesmo, sem nem a aprovação do plano (você vai dormir) | `/spec-lead "<pedido>" --autonomo` |
| Retomar um plano que ficou com passo pendente | `/spec-lead continuar` |
| Registrar ou substituir uma decisão de arquitetura direto | `/spec-adr "<título>"` |
| Ver as ADRs que os planos pedem e ainda faltam | `/spec-adr` (vazio) |
| Criar, corrigir ou reorganizar o índice das ADRs | `/spec-adr índice` |
| Criar ou atualizar o design system | `/spec-design-system [proposta, Figma, docs ou o formulário preenchido]` |
| Commitar todas as mudanças | `/spec-semantic-commit` |
| Lançar uma versão (SemVer, CHANGELOG, tag, GitHub Release), ou ver qual seria a próxima | `/spec-release [major \| minor \| patch \| X.Y.Z]` |
| Criar ou atualizar o README | `/spec-readme [pasta]` |
| Criar, auditar ou enxugar o CLAUDE.md, ou acrescentar uma regra a ele (e garantir `.temp/` no `.gitignore`) | `/spec-claude [pasta]` |
| Deixar o agente ver e operar a tela do Windows (screenshot, clique, teclado) | `/spec-windows-mcp [ferramentas extras]` |
| Deixar o agente ver e operar páginas web num navegador (print, clique, console) | `/spec-playwright-mcp [headed]` |
| Trazer a versão nova do SpecFlow (skills, agents, rules) para o projeto | `/spec-upgrade [tag, branch ou commit]` |

Documentar ou reorganizar specs não tem comando: o agente faz quando precisa (rule `docs.md`).
Executar um plano também não tem comando próprio: quem executa é o lead, na mesma rodada em que
o plano foi aprovado, ou `/spec-lead continuar`.

**Próximo passo de um plano:** `rascunho` → o usuário revisa e aprova, e o lead segue na hora;
`aprovado` com passo pendente → `/spec-lead continuar`; plano sumiu de `docs/plans/` → foi
executado por inteiro: commit, se houver algo em stage (`git diff --staged --stat`).

**Fronteiras que confundem:** spec × plano (como funciona hoje × o que uma mudança vai fazer) ·
decisão do plano × ADR (a do plano vale para a mudança; vira ADR quando vale para o projeto
todo) · spec × ADR (como × por quê) · rule × ADR (a ADR registra o porquê e não carrega no
contexto; a rule diz como aplicar e carrega) · README × `CLAUDE.md` (pessoas × agente) · fato × regra no `CLAUDE.md` (o repo responde, o lead escreve × alguém julgou, o usuário escreve).

# 6. Estrutura e convenções

```
CLAUDE.md, README.md, CHANGELOG.md (só depois da primeira versão)
.claude/
├── skills/spec-<nome>/SKILL.md (+ ao lado: template.md quando gera documento, arquivo de
│   consulta ou script quando a skill precisa — ex.: plano.md e relatorio.md na spec-lead,
│   uso.md e verificar.py nas skills de MCP)
├── agents/spec-executor.md, spec-reviewer.md
└── rules/behavior.md, docs.md (sempre) e <tema>.md (só ao tocar os paths dela)
docs/
├── adr/README.md            índice; as ADRs NNN-<titulo>.md ao lado
├── specs/README.md          índice; os assuntos ao lado, subdivididos em pastas
├── plans/AAAA-MM-DD-<slug>.md   só enquanto não executado
├── architecture.md
├── design-system.md         só com interface
└── design/<categoria>.md
.temp/                       fora do git (está no .gitignore)
├── references/              brand/, prompts/, screens/…
└── design-system-formulario.md   só até a spec-design-system consumi-lo
```

- Nomes em kebab-case, sem acentos. ADR com numeração de três dígitos, sequencial, nunca
  reaproveitada; plano com data, porque é apagado. Pastas só são criadas quando usadas.
- Skills e agents com nome em inglês e prefixo `spec-`; rules com nome em inglês, sem prefixo;
  conteúdo em português.

**Ficou de fora de propósito:**

- Do fluxo antigo de 8 fases, abandonado por ser pesado demais: documento por fase
  (pre-refinement, PRD, tech alignment, task plan), aprovação entre fases, versões `v1/v2` de
  documentos, arquivo por task, débito como fase com versão própria, guia visual, backlog e
  registro de "próximos passos".
- Do fluxo `spec-change`/`spec-build`, abandonado porque devolvia ao usuário uma fila que não
  acabava: a "varredura fora do escopo" com o agent `spec-fixer`, os tetos de itens por
  relatório, por passo e por fixer, a flag `--sem-varredura` e a seção "Deixei para você".
- Da v0.6, abandonado porque o próprio fluxo fabricava débito: a categoria **"melhoria"** dos
  revisores (até 3 por revisão, em todo passo), o campo **"Notei fora do escopo"** do executor,
  a lista de **"Arquivos"** como cerca, a **fila de débito** e os **planos secundários** em dois
  lotes, as **partes** de plano, a revisão **por passo** e os agents `spec-scout`,
  `spec-planner`, `spec-architect` e `spec-security-reviewer`. Revisor pedido para achar
  sugestões sempre acha, e cada passagem de bastão é um lugar para deixar recado. No lugar
  entraram **tarefa dada é tarefa cumprida** (`behavior.md`, regra 2), a **revisão final
  binária** e a **lista fechada de paradas**.
- Da v1.0, removidas por falta de uso: as skills `spec-rule-create` (rule com as convenções de
  um tema) e `spec-rule-curate` (decidir se um item vira regra, auditar as rules). Regra nova
  vai pela `/spec-claude`.

Não reintroduza nenhum deles sem um motivo concreto. Em especial: **regra nova que crie mais um
motivo para o agente parar e perguntar** vai contra a `behavior.md`, regra 1; e **campo,
categoria ou etapa que dê ao agente um lugar para deixar trabalho "para depois"** fabrica
débito, contra a regra 2. Nenhum dos dois entra.

# Responder

- **Sem pergunta** → liste as skills (seção 5) e pergunte o que o usuário quer saber.
- **Roteamento** → comece pelo comando pronto para copiar, depois o porquê em 1-2 frases.
- **Fluxo** → passo a passo, destacando onde o usuário aprova e o que o faz ser chamado de volta.
- **Configuração** → onde a convenção está (seção deste guia, template ou rule) e, com `grep` em
  `.claude/`, todos os arquivos que teriam que mudar juntos (este guia incluído).

# Entrada

$ARGUMENTS
