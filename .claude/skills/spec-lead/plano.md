---
status: rascunho      # rascunho | aprovado
---

# <Título da mudança>

- **Data:** AAAA-MM-DD · **Pedido:** <texto curto do pedido | `.temp/references/...`>
- **ADRs que se aplicam:** [ADR NNN](../adr/NNN-<titulo>.md) | nenhuma
- **Specs afetadas:** `docs/specs/<assunto>.md` (cria | altera) · `docs/architecture.md` (altera) | nenhuma

<!-- LLM-ONLY: plano é temporário — o /spec-lead o escreve, o executa e o apaga. O que vale
depois da mudança (como funciona) vai para as specs afetadas, nos passos. Seção que não se
aplica é REMOVIDA, não fica "N/A". Cabe em ~1 página (até 3 na mudança grande). Mudança grande
é um plano com mais passos, nunca vários planos. -->

## 1. Problema e objetivo

<2 a 5 linhas: o que está errado ou faltando hoje, para quem, e o que muda quando estiver pronto.>

## 2. Comportamento esperado

- **CA-1** — <comportamento verificável, do ponto de vista de quem usa: dado X, quando Y, então Z>
- **CA-2** — <caso de erro ou de borda: o que o usuário vê, com a mensagem literal quando importa>

**Fora do escopo:** <o que parece fazer parte, mas fica de fora, e por quê> | nada

<!-- LLM-ONLY: cada CA é verificável por um teste. CA que SÓ uma pessoa consegue conferir
(responder ao UAC, usar outro aparelho, olhar a tela num projeto sem a seção Teste visual no
CLAUDE.md) termina com "(manual: <roteiro em 1 linha>)": ele não entra no "Pronto quando" de
nenhum passo, e o relatório o recolhe para o usuário. Só o que muda entra. -->

## 3. Decisões

| Decisão | Escolha | Por quê |
|---|---|---|
| <o que precisou ser decidido> | <o que foi escolhido> | <motivo em 1 linha; alternativa descartada> |

**ADR a registrar:** <título> | nenhuma

<!-- LLM-ONLY: só decisão não óbvia — a que o usuário respondeu, ou a que o lead tomou e foge do
padrão do código. O que o código, uma spec ou uma ADR já respondem não entra. ADR só para
decisão que vale para o projeto todo; o executor a registra no primeiro passo que depende dela
(Arquivos: `docs/adr/NNN-<titulo>.md` e `docs/adr/README.md`). -->

## 4. Telas

<Por tela: como se chega, o que mostra, ações, estados (carregando, vazio, erro) com as mensagens
literais. Componentes e tokens do design system; o que for novo nele, dito aqui.>

## 5. Contratos e dados

<API (método, rota, entrada, saída, erros), formato de arquivo, tabelas e migrações (com o que
desfaz cada uma, ou "irreversível: <plano>").>

## 6. Passos

### Passo 1 — <nome>

- **Faz:** <o que muda, em nível de plano: nomes, assinaturas, códigos de erro; não o corpo das
  funções>
- **Arquivos:** `<caminho>` (novo) · `<caminho>` (modifica) · `docs/specs/<assunto>.md` (modifica)
- **Pronto quando:**
  - [ ] <teste: entrada → resultado concreto> (CA-1)
  - [ ] `docs/specs/<assunto>.md` descreve <o comportamento novo>
  - [ ] `<comando de testes e build do projeto>` sem erros

### Passo 2 — <nome>

- **Faz:** …
- **Arquivos:** …
- **Risco:** alto — <área: autenticação, pagamentos…>
- **Pronto quando:**
  - [ ] …

<!-- LLM-ONLY: passos na ordem de execução; cada um entregável e testável sozinho, numa passada.
"Arquivos" é onde o passo deve mexer, não uma cerca: o que ele precisar a mais para funcionar
(chamador, teste, config) o executor toca e diz. "Pronto quando" só tem o que um agente prova —
teste com resultado concreto, comando, trecho da spec; nunca "funciona". Todo CA que não é
"(manual: …)" aparece em algum "Pronto quando", e toda spec afetada é atualizada pelo passo que
muda aquele comportamento. "Risco: alto" só no passo que toca autenticação, autorização,
pagamentos, segredos ou dados pessoais: com ele, a revisão final aplica o checklist de
segurança. Projeto sem suíte naquela camada → o "Pronto quando" descreve a conferência mecânica
(behavior.md, regra 3), nunca fica sem prova. -->

## 7. Riscos e dúvidas

- <risco ou `[HIPÓTESE]`> → <como validar ou mitigar>

<!-- LLM-ONLY: risco que o próprio plano resolve não entra aqui — entra como passo. -->
