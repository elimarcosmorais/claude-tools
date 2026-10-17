# {{NOME_DO_PROJETO}}

> {{DESCRICAO_EM_UMA_LINHA}}
<!-- [todos] Uma frase: o que é + para quem. Sem jargão de marketing. -->

{{BADGES}}
<!-- [OPCIONAL · todos] Até 3 badges shields.io (build, versão no registry, licença), só do que o
repo realmente tem. Ex.: [![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE).
Badge apontando para CI ou registry que não existe é pior que nenhum: sem badge real, remover. -->

![{{NOME_DO_PROJETO}}]({{CAMINHO_DA_DEMO}})
<!-- [app] Screenshot ou GIF no topo: para um app, o visual convence antes do texto. Sem mídia,
trocar por: > ⚠️ Adicionar screenshot em <caminho sugerido>. Nos outros tipos, remover. -->

## Sobre
<!-- [todos] 1–2 parágrafos: que problema resolve, para quem, e por que existe. O "porquê" vem
antes do "como". Se há specs em docs/specs/, o problema pode já estar lá. -->
{{POR_QUE_EXISTE}}

## Features
<!-- [OPCIONAL · app, service] 3–6 bullets dos recursos principais. Verbos no presente. -->
- {{FEATURE}}

## Arquitetura
<!-- [OPCIONAL · app, service] Bullets curtos dos componentes (front, motor, banco, filas) ou um
diagrama. Decisões de arquitetura: linkar as ADRs (docs/adr/), não copiá-las. -->
{{ARQUITETURA}}

## Pré-requisitos
<!-- [OPCIONAL · todos] Ferramentas e versões mínimas, tiradas do manifest (go.mod, engines,
pubspec). Ex.: Go 1.26+, Node 20+, Docker. -->
- {{PRE_REQUISITO}}

## Instalação
<!-- [todos] Título por tipo: lib e cli → "Instalação"; app → "Como rodar"; service → "Como subir".
Comandos REAIS, copiados de onde vivem (package.json, Makefile, go.mod, docker-compose.yml):
lib → comando do gerenciador de pacotes; cli → instalação global; app → instalar deps + subir em
dev; service → docker compose up / make run. -->
```bash
{{COMANDOS_DE_INSTALACAO}}
```

## Uso
<!-- [lib, cli] O menor exemplo que roda de verdade: lib → import + a chamada principal + o
resultado; cli → comando + output esperado. Mostrar, não descrever. -->
```{{LINGUAGEM}}
{{EXEMPLO_MINIMO}}
```

## Comandos
<!-- [cli] Comandos e flags principais, tirados do --help ou do parser de argumentos. -->
| Comando / flag | Descrição |
|----------------|-----------|
| {{COMANDO}} | {{DESCRICAO}} |

## Configuração
<!-- [OPCIONAL · todos] Variáveis de ambiente ou opções. Tabela OU link para a doc. Nunca vazar
segredo: valor de exemplo, nunca o real. -->
| Variável / opção | Default | Descrição |
|------------------|---------|-----------|
| {{OPCAO}} | {{DEFAULT}} | {{DESCRICAO}} |

## API
<!-- [lib, service] lib → principais funções/classes exportadas, com assinatura; service →
principais endpoints. API extensa: resumo + link (OpenAPI, doc gerada). Título no service:
"Endpoints". -->
{{API_OU_LINK}}

## Build / Deploy
<!-- [OPCIONAL · app, service] Comando real de build de produção e para onde vai o artefato. -->
```bash
{{COMANDO_DE_BUILD}}
```

## Testes
<!-- [OPCIONAL · todos] Comando real que roda a suíte. -->
```bash
{{COMANDO_DE_TESTE}}
```

## Contribuindo
<!-- [OPCIONAL · todos] Só se o projeto aceita contribuição: link para CONTRIBUTING.md ou 2–3
linhas (rodar local, testar, abrir PR). -->
{{CONTRIBUINDO_OU_LINK}}

## Licença
<!-- [todos, se existir LICENSE] Sem arquivo LICENSE nem campo license no manifest: remover a
seção e avisar o usuário (não inventar licença). -->
{{NOME_DA_LICENCA}} — ver [LICENSE](LICENSE).
