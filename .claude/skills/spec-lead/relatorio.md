<!-- LLM-ONLY: template do relatório final do /spec-lead. Preencha e mostre no chat, não salve em
arquivo. Seção sem conteúdo some, menos "Precisa de você", que vira "Precisa de você: nada." —
essa linha é a informação mais importante do relatório. Nenhuma linha é livre: toda linha vem de
um campo do executor, do JSON da revisão final ou da Etapa 7. O caminho direto usa a forma
curta, no fim deste arquivo. -->

## <pedido em uma frase>

**Status:** concluído · <N> passos · revisão aprovada na tentativa <T>
<!-- ou: parado no passo <N> de `<plano>` — <qual dos quatro motivos da lista de paradas> -->

### O que foi entregue

- <1 a 4 linhas do que o usuário ganhou, do ponto de vista dele; não a lista de arquivos>
- **Arquivos:** <N> criados, <M> modificados (`git diff --staged --stat`)
- **Specs:** `<arquivo>` (criado | atualizado | subdividido) | nenhuma
- **ADRs:** <NNN — título> (registrada) | nenhuma
- **Também resolvi:** <o que atrapalhava o pedido, em 1 linha cada, `arquivo`> | omita a linha

<!-- LLM-ONLY: "Também resolvi" é só o que ATRAPALHAVA a tarefa (behavior.md, regra 2, leitura 1)
e por isso virou parte dela. O que não atrapalhava não foi feito e não aparece aqui. -->

### Provas

| O que | Comando | Resultado |
|---|---|---|
| suíte | `<comando>` | <N> passando, 0 falhando, 0 pulados |
| build | `<comando>` | ok |
| lint | `<comando>` | ok |
| tela | <teste de tela ou MCP> | <o que viu> |

**Sem prova:** <o que ficou> — <por quê> | nada

<!-- LLM-ONLY: os números saem da rodada da revisão final aprovada. "0 pulados" é literal
(behavior.md, regra 3). Pulo que você não conseguiu resolver entra em "Sem prova", com o nome do
teste e o motivo, e nunca na linha de verde. -->

### Decisões que tomei

| Decisão | Escolhi | Por quê |
|---|---|---|
| <o que precisou ser decidido> | <o que foi escolhido> | <motivo em 1 linha; alternativa descartada> |

<!-- LLM-ONLY: é aqui que o usuário audita a autonomia em vez de ser interrompido por ela. Entra
toda escolha que mudaria o resultado se fosse outra — sua ou do executor. Não entra o que o
código, uma spec ou uma ADR já respondiam. -->

### Precisa de você

<!-- LLM-ONLY: SÓ o que você não pode fazer — um dos quatro motivos da lista de paradas
(behavior.md, regra 1). Cada item traz o que fazer e o comando pronto. Nada aqui → escreva
literalmente "Precisa de você: nada." -->

- **Conferir:** <CA-N (manual: …)> → <roteiro em 1 linha>
- **Decidir:** <o que muda no que o produto entrega> · Recomendo: <…> → `/spec-lead "<pedido pronto>"`
- **Fazer:** <exige UAC, 2FA, conta externa ou aparelho> → `<comando pronto>`
- **Risco:** <segredo, credencial, dado pessoal ou de pagamento exposto, `arquivo:linha`> · <resolvido nesta rodada | Recomendo: … → `<comando>`>

<!-- LLM-ONLY: a linha "Risco" aparece mesmo resolvida; risco nunca é descartado (behavior.md,
regra 2, a exceção única). -->

### Estado

Código e specs em **stage, sem commit**. Revise com `git diff --staged` e commite com
`/spec-semantic-commit`.

**Plano:** apagado (remoção em stage) | mantido `<arquivo>` — faltam os passos <N, M>; retome
com `/spec-lead continuar`.

<!-- LLM-ONLY: forma curta, só no caminho direto (spec-lead, Etapa 3). Quatro linhas; risco, quando
houver, entra como quinta linha, igual à da forma completa. -->

## <pedido em uma frase>

**Feito:** <o que mudou, em 1 linha> · `<arquivo>` (+ `<teste ou spec que mudou junto>`)
**Prova:** `<comando>` → <N> passando, 0 pulados | conferência: <o que conferiu e o resultado>
**Precisa de você:** nada.
**Estado:** em stage, sem commit — `/spec-semantic-commit`.
