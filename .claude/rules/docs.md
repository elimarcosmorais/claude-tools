# Documentação do projeto

Onde cada documento mora, quem pode mexer nele e quando. Vale em qualquer tarefa.

## Onde cada coisa mora

| Lugar | Guarda | Quem mexe |
|---|---|---|
| `docs/specs/` | **como o sistema funciona hoje**, por módulo ou assunto | o agente, livremente |
| `docs/architecture.md` | a arquitetura base: camadas, módulos, fluxos e integrações principais | o agente, livremente |
| `docs/adr/NNN-<titulo>.md` + índice `docs/adr/README.md` | **por que** uma decisão transversal foi tomada | o agente: cria, atualiza e substitui conforme a tarefa pede (abaixo); `/spec-adr` quando o usuário registra direto |
| `docs/design-system.md` + `docs/design/` | a identidade visual: tokens, componentes e padrões | `/spec-design-system`; sincronizar com o código é livre |
| `docs/plans/AAAA-MM-DD-<slug>.md` | o plano de uma mudança, até ser executado | o `/spec-lead` escreve, o usuário aprova, o lead executa e apaga |
| `.temp/references/` | material de consulta temporário do usuário: marca, prints, prompts | o usuário |
| `.temp/design-system-formulario.md` | a especificação de entrada do design system, até ser preenchida e consumida | `/spec-design-system` cria, o usuário (ou um agente) preenche, a skill lê |
| `.temp/design-system-modelo-formulario.md` | o retrato de um design system que já existe, até virar modelo | `/spec-design-system` cria em branco, o usuário preenche, a skill o salva em `modelos/` |
| `.claude/skills/spec-design-system/modelos/<slug>.md` | os modelos de design system prontos para aplicar | `/spec-design-system` grava, ao cadastrar um modelo (abaixo) |
| `CLAUDE.md` | o contrato do agente com **este** projeto: stack, comandos reais, convenções, onde coisas novas vão | `/spec-claude` cria e audita, com o usuário; o **fato** que o repo responde, o agente sincroniza livremente (abaixo) |
| `CHANGELOG.md` | o que mudou para quem usa, em cada versão lançada | `/spec-release`, a cada versão |
| `.claude/skills/spec-*`, `.claude/agents/spec-*`, `.claude/rules/behavior.md` e `docs.md` | o framework SpecFlow | ninguém no projeto: a `/spec-upgrade` os traz do SpecFlow (abaixo) |

**Não duplique.** A spec diz **como funciona** e aponta para a ADR que diz **por quê**. O plano
diz o que uma mudança vai fazer e some depois dela: o que sobrevive vai para a spec (como
funciona), para uma ADR (por quê) ou para o `CLAUDE.md`/rules (regra para todo código novo).
Pastas e arquivos só são criados quando há conteúdo para eles.

## Specs e `architecture.md`: livres

Crie, edite, mova, subdivida e apague **sem pedir permissão**. Ao fim da resposta, uma linha:
`Specs: <arquivo> (criado | atualizado | subdividido | apagado), …`.

Atualize quando:

- uma mudança altera o comportamento que uma spec descreve: na **mesma** mudança, nunca depois;
- um assunto novo passa de um parágrafo;
- você descobre algo que custaria tempo para redescobrir: armadilha, regra de negócio implícita,
  detalhe de integração.

Não entra: o que o código diz sozinho, o porquê de uma decisão (ADR), plano ou pendência.

**Índice.** `docs/specs/README.md` lista cada arquivo com a coluna "Abra quando" (a situação em
que ele deve ser lido). Arquivo criado entra no índice; apagado sai. Antes de mexer num módulo,
leia o índice e abra só a spec dele.

## Tamanho máximo: 300 linhas

Vale para spec, `architecture.md`, `design-system.md` e os arquivos de `docs/design/`. Passou de
300 linhas (`wc -l`) → **subdivida na mesma tarefa, sem pedir permissão**, e diga no fim:

1. `<assunto>.md` vira índice: um resumo de poucas linhas e a tabela "Abra quando" dos pedaços.
2. O conteúdo vai para `<assunto>/<sub-assunto>.md`, um sub-assunto por arquivo.
3. Os links que apontavam para o trecho movido passam a apontar para o arquivo novo.

A profundidade é livre: um pedaço que passar de 300 linhas se subdivide do mesmo jeito.

## ADRs: do agente, pelo índice

`docs/adr/README.md` lista só as ADRs em vigor, com a coluna "Abra quando", agrupadas por tema
quando passam de 20. Substituída sai do índice: o histórico fica nas próprias ADRs (a nova
aponta a antiga em **Substitui**). Para achar a ADR de um assunto, leia o índice e abra só as
que a coluna aponta; o status que vale é o do arquivo da ADR. Sem índice, liste `docs/adr/`.

**ADR é do agente, não do usuário.** O usuário aprova o plano; a ADR registra o que a tarefa
decidiu, e o agente a cria e a modifica sempre que cumprir a tarefa pede — o `spec-executor` no
passo do plano, o `/spec-lead` no caminho direto. Nada disso pede confirmação (`behavior.md`,
regra 4) nem vira proposta para o usuário: ADR que a tarefa deixou errada é a tarefa por
terminar (`behavior.md`, regra 2). `/spec-adr` continua servindo ao usuário que quer registrar
uma decisão direto, e `/spec-adr índice` reorganiza o índice.

| A tarefa… | O agente |
|---|---|
| toma uma decisão transversal, cara de mudar e com alternativa real | **cria** a ADR (`aceito`) e a linha dela no índice |
| torna falso o que uma ADR afirma — caminho, mecanismo, nome, consequência —, e a decisão continua de pé | **atualiza** a ADR no lugar, na mesma tarefa, como faria com uma spec |
| muda a própria decisão | **substitui**: ADR nova com **Substitui**; na antiga, `status: substituído` e **Substituída por**; tira a antiga do índice |
| esbarra em link quebrado, referência inexistente ou erro de digitação numa ADR ou no índice | **corrige** |

Duas coisas seguram o histórico, e nenhuma é motivo de parar: **trocar a decisão é substituir,
nunca reescrever** — a ADR antiga é o registro de por que se fez assim, e a `substituído` não se
edita mais; e **ADR `aceito` é contrato** — o agente não a contraria por preferência, só quando a
tarefa pede outra coisa, e aí a substitui. Toda ADR criada, atualizada ou substituída sai no
relatório.

O índice `docs/adr/README.md` é recurso disputado: os passos rodam um de cada vez, e nenhum
agent o edita em paralelo.

## Design system: livre para sincronizar

Registrar o que o código **já implementa** é livre, como nas specs. Mudar a identidade (cor,
token, componente ou padrão que o código ainda não tem) é decisão do usuário: pergunte antes, ou
rode `/spec-design-system`.

## CLAUDE.md: fato é livre, regra é do usuário

O `CLAUDE.md` entra no contexto de **toda** sessão de **todo** agent, então desatualizado ele
não é inofensivo: ele mente para quem confia nele. Mas ele também é o único documento que
governa o próprio agente, e por isso o corte é mais estreito que o das specs:

| O que mudou | Quem resolve |
|---|---|
| **Fato que um arquivo do repo responde:** comando de teste, integração, build ou lint, versão da stack, pasta ou camada que passou a existir, ferramenta de teste visual instalada, `.temp/` fora do `.gitignore` | o agente, **livremente** — é sincronizar documento com código, como numa spec |
| **Regra nova:** convenção que ninguém escreveu, anti-padrão, "sempre faça X", "nunca edite Y" | o **usuário**: `/spec-claude` |

O teste é mecânico: **deu para extrair de um arquivo do repo?** É fato. **Veio do juízo de alguém
sobre como o time deve trabalhar?** É regra, e regra é escolha do usuário (`behavior.md`,
regra 1, motivo **c**). Na dúvida, é regra.

O fato que a tarefa tornou falso, ou que fez um agente tropeçar nela, é corrigido na própria
tarefa (`behavior.md`, regra 2).

Três coisas seguram o risco de o agente editar as próprias amarras, e valem juntas: ele escreve
**só fato com fonte**, a edição sai no relatório em "Também resolvi", e nada é commitado —
o usuário vê o `git diff --staged` antes.

**Criar o `CLAUDE.md` que não existe é do usuário**, porque a `spec-claude` precisa do que o repo
não responde (o idioma, os anti-padrões que o time mais corrige). Mas a ausência dele **não
trava** demanda nenhuma: o agente lê o repo e trabalha sem ele, e `/spec-claude` vai para o
relatório como proposta. Nisso ele é o oposto do design system, que trava de propósito — tela
sem identidade visual sai para ser jogada fora; código sem `CLAUDE.md` sai um pouco pior.

## Arquivos do framework: não se editam no projeto

As skills e os agents `spec-*` e as rules `behavior.md` e `docs.md` vêm do SpecFlow, e a
`/spec-upgrade` os traz de novo a cada versão: editado no projeto, um deles perde a edição na
atualização seguinte ou fica preso a uma versão velha. Por isso o "resolva no arquivo que
precisar" (`behavior.md`, regra 2) não os alcança.

Faltou algo numa skill do framework para cumprir o pedido → cumpra-o sem editá-la: rode você os
comandos que ela rodaria, com o que o projeto pede. O que o projeto tem de próprio vai para um
arquivo dele: o fato para o `CLAUDE.md` (como a linha `Release:` que a `/spec-release` lê), o
procedimento para uma skill ou rule do projeto, com nome sem o prefixo `spec-`. O
`.claude/settings.json` é do projeto e fica fora desta regra. No repositório do próprio SpecFlow
esses arquivos são a fonte, e editá-los é o trabalho.

A exceção é `.claude/skills/spec-design-system/modelos/`: a `/spec-design-system` grava ali o
modelo que o usuário cadastra. Modelo criado no projeto nunca existiu no framework, e por isso a
`/spec-upgrade` o trata como do projeto e não o toca; cadastrado no repositório do SpecFlow, ele
chega a todos os projetos na atualização seguinte.

## Planos: temporários

Plano não vira documentação: ele é executado e **apagado pelo `/spec-lead`** no fechamento, e a
remoção fica em stage junto com o código, no mesmo commit. Existe um tipo só: o plano da
mudança, `AAAA-MM-DD-<slug>.md`, que nasce `rascunho` e vira `aprovado` com a aprovação do
usuário — a única do fluxo. Mudança grande é um plano com mais passos, não vários planos.

**Só é apagado o plano cumprido:** todo checkbox `[x]`, a revisão final aprovada e a suíte, o
build e o lint inteiros verdes, sem teste pulado. Com passo pendente, o plano fica — com os
checkboxes marcados — e `/spec-lead continuar` o retoma.

## Referências: temporárias

`.temp/references/` é material de consulta do usuário e pode sumir a qualquer momento: `.temp/`
não é versionado, porque esse material é criado, alterado e apagado toda hora (a `/spec-claude`
garante a linha `.temp/` no `.gitignore`). Leia à vontade, mas **só o plano** o cita, porque
também é temporário. Nenhum documento permanente (specs, `architecture.md`, ADR, design system,
`CLAUDE.md`, README, rules) cita, linka ou documenta esses arquivos. O que ele precisar de uma
referência (uma cor, uma regra, um fluxo) é escrito nele como fato do projeto, sem o caminho.
Quem apaga é o usuário; o agente não apaga nem move nada ali.

O formulário do design system (`.temp/design-system-formulario.md`) mora em `.temp/` pela mesma
razão: a `/spec-design-system` o cria pré-preenchido e o lê de volta, o usuário (ou outro agente)
o preenche, e quem o apaga é o usuário. Ele é entrada, não documentação: o que ele responde vai
para `docs/design-system.md` e `docs/design/`, e nenhum documento permanente o cita. O formulário
de modelo (`.temp/design-system-modelo-formulario.md`) segue a mesma regra: o que ele responde
vai para `modelos/<slug>.md`.

## Nomes

kebab-case, minúsculas, sem acentos. Conteúdo em português com acentuação correta; nomes de
código ficam como no código.
