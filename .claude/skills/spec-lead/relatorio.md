<!-- LLM-ONLY: template do relatório final do /spec-lead. Preencha e mostre no chat, não salve em
arquivo. Seção sem conteúdo some, menos "Precisa de você", que vira "Precisa de você: nada." —
essa linha é a informação mais importante do relatório. Nenhuma linha é livre: toda linha vem de
um campo do executor, do JSON da revisão final ou da Etapa 7. O caminho direto usa a forma
curta, no fim deste arquivo. -->

## <pedido em uma frase>

**Status:** concluído · <N> passos · revisão aprovada na tentativa <T>
<!-- ou: parado no passo <N> de `<plano>` — <qual dos quatro motivos da lista de paradas> -->
<!-- ou: parado na revisão final de `<plano>` — d: reprovada nas 2 tentativas -->
<!-- LLM-ONLY: <T> é 1 ou 2: a revisão final não tem terceira tentativa (spec-lead, Etapa 6). -->


### O que foi entregue

- <1 a 4 linhas do que o usuário ganhou, do ponto de vista dele; não a lista de arquivos>
- **Arquivos:** <N> criados, <M> modificados (`git diff --staged --stat`)
- **Specs:** `<arquivo>` (criado | atualizado | subdividido) | nenhuma
- **ADRs:** <NNN — título> (registrada | atualizada: <o que mudou> | substituída por NNN) | nenhuma
- **Também resolvi:** <o que atrapalhava o pedido, em 1 linha cada, `arquivo`> | omita a linha
- **Riscos fechados:** <o que estava exposto, a quem, `arquivo:linha`> · <como fechou, 1 linha> | omita a linha

<!-- LLM-ONLY: "Também resolvi" é só o que ATRAPALHAVA a tarefa (behavior.md, regra 2, leitura 1)
e por isso virou parte dela. O que não atrapalhava não foi feito e não aparece aqui.
Risco fechado mora em "Riscos fechados", NUNCA em "Precisa de você": é entrega, não tarefa do
usuário. Limite que só alcança quem já tem um acesso que vale mais (ex.: exige gravar na pasta
servida, que já permite rodar código) não é risco: fica na ADR e não ganha linha. "Continua
aberto" sem ação para o usuário não existe: risco aberto é a linha "Decidir (risco)". -->

### Provas

| O que | Comando | Resultado |
|---|---|---|
| suíte | `<comando>` | <N> passando, 0 falhando, 0 pulados |
| build | `<comando>` | ok |
| lint | `<comando>` | ok |
| integração | `<comando>` | <alcance: por impacto em <alvos> \| inteira \| nenhum alvo> · ok |
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

<!-- LLM-ONLY: SÓ o que o usuário tem que FAZER e você não pode — um dos quatro motivos da
lista de paradas (behavior.md, regra 1). Cada item diz a ação e traz o comando pronto; item sem
ação para o usuário não entra aqui (informação vai para "O que foi entregue"). Nada aqui →
escreva literalmente "Precisa de você: nada." -->

- **Conferir:** <CA-N (manual: …)> → <roteiro em 1 linha>
- **Decidir:** <o que muda no que o produto entrega> · Recomendo: <…> → `/spec-lead "<pedido pronto>"`
- **Decidir (risco):** <quem explora, como, e o que ganha, `arquivo:linha`> · (a) <correção> — custa <o que o produto perde> (b) <…> · Recomendo: <x>, porque <…> → `/spec-lead "<pedido pronto>"`
- **Fazer:** <exige UAC, 2FA, conta externa ou aparelho> → `<comando pronto>`
- **Destravar:** <passo N | revisão final, B1> `arquivo:linha` — <o que falha ou o que o revisor recusa> · tentei: <a CAUSA-RAIZ de cada tentativa> · não bastou porque <…> · (a) <saída> — custa <…> (b) <…> · Recomendo: <x>, porque <…> → `<comando pronto>`

<!-- LLM-ONLY: risco aberto só chega aqui quando TODA correção muda o que o produto entrega
(behavior.md, regra 1, motivo c) — e chega como escolha, com o ataque concreto, o custo de cada
opção e a recomendação. Existe correção que não muda o produto → a tarefa não terminou: volte à
Etapa 5 ou 6 e corrija.
"Destravar" é a parada do tipo d (behavior.md, regra 1): um item por passo que falhou 3 vezes
ou por bloqueante que restou da revisão final reprovada nas 2 tentativas. Ele é o relatório do
que fazer — nunca "continua reprovado" sem saída, recomendação e comando. -->


### Estado

Código e specs em **stage, sem commit**. Revise com `git diff --staged` e commite com
`/spec-semantic-commit`.

**Plano:** apagado (remoção em stage) | mantido `<arquivo>` — faltam os passos <N, M>; retome
com `/spec-lead continuar`.

<!-- LLM-ONLY: forma curta, só no caminho direto (spec-lead, Etapa 3). Quatro linhas; risco fechado,
quando houver, entra como linha "Riscos fechados:" antes de "Precisa de você", e risco aberto
entra em "Precisa de você" como "Decidir (risco)", iguais às da forma completa. -->

## <pedido em uma frase>

**Feito:** <o que mudou, em 1 linha> · `<arquivo>` (+ `<teste ou spec que mudou junto>`)
**Prova:** `<comando>` → <N> passando, 0 pulados | conferência: <o que conferiu e o resultado>
**Precisa de você:** nada.
**Estado:** em stage, sem commit — `/spec-semantic-commit`.
