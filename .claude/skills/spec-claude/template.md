# CLAUDE.md

## Visão do projeto
<!-- [sempre] 2–4 linhas: o que é, para quem, o que otimiza, restrições principais. Sem marketing
nem história da empresa. Se há specs em docs/specs/, o problema pode já estar lá. -->
{{VISAO}}

## Stack
<!-- [sempre] Linguagem, framework, versões, gerenciador de pacotes, libs principais (state,
banco, testes), cada item tirado do manifest/lockfile. A linha "NÃO usar" é metade do valor:
libs e padrões proibidos sem aprovação. Sem negativo real conhecido, remover a linha. -->
- {{ITEM_DA_STACK}}
- NÃO usar: {{LIBS_OU_PADROES_PROIBIDOS}}

## Comandos
<!-- [sempre] Comandos REAIS, copiados de onde vivem (package.json, Makefile, Taskfile, CI).
Só os que existem: sem script de lint no repo, sem linha de lint. -->
- Instalar: `{{COMANDO}}`
- Dev: `{{COMANDO}}`
- Testes: `{{COMANDO}}`
<!-- [OPCIONAL · integração fora da suíte padrão (build tag, config ou script próprio)] A linha que o
spec-reviewer e o spec-executor leem para rodar só o que a mudança alcança, e a /spec-release para
rodá-la inteira. Sem como achar os alvos, fica só o comando inteiro. -->
- Integração: `{{COMANDO_INTEIRO}}` · por impacto: `{{COMANDO_COM_<alvos>}}` · alvos: {{CRITERIO}}
- Lint / typecheck / build: `{{COMANDO}}`
<!-- [OPCIONAL · versão fora dos manifestos ou instalável para o release (.exe, .apk)] A linha que a
/spec-release lê; o formato está nela. -->
- Release: versão em `{{ARQUIVO}}` (`{{CAMPO}}`) · anexos: `{{COMANDO}}` → `{{ANEXO}}`

## Convenções
<!-- [quase sempre] Regras objetivas, que dá para conferir olhando um diff. Nunca "escreva código
limpo". Prefira negativos ("NÃO edite os arquivos gerados em db/gen/") e cite um arquivo
exemplar em vez de copiar código ("handler novo segue src/users/handler.go"). Só as
transversais: as de uma área só (banco, HTTP, um módulo) vão para .claude/rules/<tema>.md. -->
- {{CONVENCAO}}

## Onde coisas novas vão
<!-- [quase sempre] Onde nasce cada tipo de arquivo (rota, componente, migração, teste). O agente
já vê a estrutura; ele precisa saber onde colocar o PRÓXIMO arquivo. Inclua a regra contra
quase-duplicata: "prefira editar o componente existente a criar um parecido". -->
- {{TIPO_DE_ARQUIVO}} → `{{PASTA}}`

## Arquitetura
<!-- [OPCIONAL · camadas ou decisões não óbvias] Regras de decisão, não nomes de pastas ("regra
de negócio só em internal/<dominio>/service.go; handler não acessa o banco"). Decisões
registradas: linkar o índice docs/adr/README.md, não copiar. Segredos ficam no .env (nunca no código). -->
- {{REGRA_DE_ARQUITETURA}}

## UI
<!-- [OPCIONAL · só projeto com interface] Estilo traduzido em implementação: "espaçamento só
pelos tokens de src/theme/tokens.ts", "botão novo = variant do Button, não componente novo".
Nunca "deixe moderno". Se existe docs/design-system.md, aponte para ele em vez de repetir. -->
- {{REGRA_DE_UI}}

## Teste visual
<!-- [OPCIONAL · só projeto com interface] Como o agente confere uma mudança de tela: a
ferramenta (a do repo ou a de teste-visual.md), o comando, onde o teste nasce e a regra "toda
mudança de tela passa por aqui antes de dizer que terminou". Ex.: "tela nova ou alterada → teste
Dusk em tests/Browser/ e `php artisan dusk` verde; na falha, abrir o print de
tests/Browser/screenshots/". -->
- {{REGRA_DE_TESTE_VISUAL}}

## Pronto significa
<!-- [OPCIONAL · projeto com testes ou CI] O que precisa passar antes de dizer que terminou, com
os comandos reais: "testes verdes (`make test`) e lint limpo (`make lint`)". -->
- {{CRITERIO}}

## Mudanças sensíveis
<!-- [OPCIONAL · há superfície pública ou sensível] O que exige confirmação antes: renomear rota
pública, mudar schema ou auth, quebrar compatibilidade. Só o específico deste projeto: as
confirmações gerais já estão em .claude/rules/behavior.md. -->
- {{REGRA}}

Como evoluir: erro que este arquivo já cobria → fortaleça a regra; erro novo → acrescente uma regra.
