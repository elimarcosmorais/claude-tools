# Documentação do projeto

Onde cada documento mora, quem pode mexer nele e quando. Vale em qualquer tarefa.

## Onde cada coisa mora

| Lugar | Guarda | Quem mexe |
|---|---|---|
| `docs/specs/` | **como o sistema funciona hoje**, por módulo ou assunto | o agente, livremente |
| `docs/architecture.md` | a arquitetura base: camadas, módulos, fluxos e integrações principais | o agente, livremente |
| `docs/adr/NNN-<titulo>.md` + índice `docs/adr/README.md` | **por que** uma decisão transversal foi tomada | o `spec-executor`, no passo do plano que a registra, ou `/spec-adr`, sempre com a decisão já tomada |
| `docs/design-system.md` + `docs/design/` | a identidade visual: tokens, componentes e padrões | `/spec-design-system`; sincronizar com o código é livre |
| `docs/plans/AAAA-MM-DD-<slug>.md` | o plano de uma mudança, até ser executado | o `/spec-lead` escreve, o usuário aprova, o lead executa e apaga |
| `.temp/references/` | material de consulta temporário do usuário: marca, prints, prompts | o usuário |
| `.temp/design-system-formulario.md` | a especificação de entrada do design system, até ser preenchida e consumida | `/spec-design-system` cria, o usuário (ou um agente) preenche, a skill lê |
| `CLAUDE.md` | o contrato do agente com **este** projeto: stack, comandos reais, convenções, onde coisas novas vão | `/spec-claude` cria e audita, com o usuário; o **fato** que o repo responde, o agente sincroniza livremente (abaixo) |
| `CHANGELOG.md` | o que mudou para quem usa, em cada versão lançada | `/spec-release`, a cada versão |

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

## ADRs: pelo índice

`docs/adr/README.md` lista só as ADRs em vigor, com a coluna "Abra quando", agrupadas por tema
quando passam de 20. Substituída sai do índice: o histórico fica nas próprias ADRs (a nova
aponta a antiga em **Substitui**). Para achar a ADR de um assunto, leia o índice e abra só as
que a coluna aponta; o status que vale é o do arquivo da ADR. Sem índice, liste `docs/adr/`.
Quem escreve ADR e índice é o `spec-executor`, no passo do plano que a registra (a decisão já
veio tomada no plano aprovado), ou a skill `/spec-adr`, quando o usuário registra uma decisão
direto: os dois atualizam o índice a cada ADR salva, e `/spec-adr índice` o reorganiza. **Ninguém decide aqui** — a ADR
registra a escolha, não a faz. A única exceção ao dono único é a correção verificável, abaixo.

### Correção verificável: a exceção ao dono único

Quando o erro atrapalha a tarefa (`behavior.md`, regra 2) — o passo usa aquela ADR, ou o
índice que ele vai editar —, o `spec-executor` corrige numa ADR `proposto` ou `aceito`, e no
índice, **sem pedir confirmação** (`behavior.md`, regra 4), só o erro que se confere sem julgar
o conteúdo:

- link que não resolve, ou referência a uma ADR, arquivo ou âncora que não existe;
- erro de digitação, de acentuação ou de formatação Markdown que não muda o que a ADR afirma.

A prova é a conferência mecânica do `Pronto quando` — abrir o link e ver que resolve, conferir
que o arquivo citado existe (`behavior.md`, regra 3).

Fica **fora**, e vira proposta no relatório com o comando `/spec-adr` pronto:

- **o que a ADR afirma** — Contexto, Decisão, Consequências, Alternativas. Vale também para a
  consequência que se revelou errada: ela é juízo novo sobre a decisão, e se a decisão ainda se
  sustenta é o usuário que diz;
- **`status`, `Substitui` e `Substituída por`**, que são o contrato e o histórico: decisão nova
  continua sendo ADR nova, nunca edição da antiga;
- **toda ADR `substituído`**, que é histórico e não se reescreve.

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
| **Fato que um arquivo do repo responde:** comando de teste, build ou lint, versão da stack, pasta ou camada que passou a existir, ferramenta de teste visual instalada, `.temp/` fora do `.gitignore` | o agente, **livremente** — é sincronizar documento com código, como numa spec |
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
para `docs/design-system.md` e `docs/design/`, e nenhum documento permanente o cita.

## Nomes

kebab-case, minúsculas, sem acentos. Conteúdo em português com acentuação correta; nomes de
código ficam como no código.
