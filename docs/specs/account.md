# Contas (`internal/account`)

Guarda logins do Claude Code com um nome (a "conta") e mostra qual está ativo. O pacote
`internal/account` faz as operações; o `internal/cli` as liga aos comandos
`claude-tools account …` e escreve as mensagens ([arquitetura](../architecture.md)).

Comandos: `save`, `use`, `rename`, `remove` e `status`.

## Arquivos

Todos os caminhos são relativos à home do usuário (`Store.Home`).

| Nome | Caminho | O que é |
|---|---|---|
| `CRED` | `~/.claude/.credentials.json` | o login ativo do Claude Code (tokens OAuth) |
| `STATE` | `~/.claude.json` | o estado do Claude Code; o campo `oauthAccount` diz de quem é o login |
| `STORE` | `~/.claude/accounts/` | as contas salvas |
| `CURRENT` | `~/.claude/accounts/current.txt` | o nome da conta atual, em texto puro |

Dentro de `STORE`:

| Arquivo | Conteúdo |
|---|---|
| `<conta>.credentials.json` | cópia byte a byte de `CRED` |
| `<conta>.account.json` | o objeto `oauthAccount` de `STATE`, sozinho (só existe se `STATE` tinha o campo) |
| `backup.credentials.json`, `backup.claude.json` | cópias de `CRED` e `STATE` de antes do último `use` |
| `current.txt` | o `CURRENT` |

O formato é o mesmo do antigo script `claude_switch.py`: as contas salvas por ele continuam
valendo sem salvar de novo. A troca depende só de `.credentials.json` e de `oauthAccount`, que
são formatos internos do Claude Code e podem mudar.

**E-mail de uma conta:** o campo `emailAddress` do bloco `oauthAccount`, ou `email` quando o
primeiro falta ou está vazio. Arquivo ausente, JSON inválido ou bloco sem os dois campos dão
"e-mail desconhecido", nunca erro.

**Escrita:** todo arquivo que o `claude-tools` grava em `CRED`, `STATE` ou `STORE` passa por
`writeFileAtomic`: grava um temporário `.<nome>.*.tmp` na mesma pasta e renomeia por cima do
destino. Se algo falha (pasta inexistente, destino que é um diretório, rename negado porque o
Claude Code está com o arquivo aberto), o destino fica como estava, o temporário é apagado e o
comando sai com `Erro: <detalhe>` e código 1. Basta rodar de novo.

## Nome de conta

- `[A-Za-z0-9_-]`, de 1 a 32 caracteres (`account.NormalizeName`).
- Gravado em minúsculas: `A_b-1` vira `a_b-1`. O sistema de arquivos do Windows não diferencia
  maiúsculas, então `Trabalho` e `trabalho` são a mesma conta.
- `backup` é reservado (em qualquer caixa), porque os arquivos `backup.*` são as cópias de
  segurança. `backup2` é um nome válido.
- Fora da regra, saída 2 e a mensagem no stderr:
  - `Nome de conta inválido: '<x>'. Use letras, números, '-' ou '_' (até 32 caracteres).`
    (com o nome como foi digitado);
  - `Nome de conta reservado: 'backup'.`

Os comandos que recebem `<conta>` normalizam o nome antes de tudo: as mensagens e os
arquivos usam o nome em minúsculas.

## `save <conta>`

Guarda o login ativo do Claude Code como `<conta>` (cria `STORE` se precisar):

1. Sem `CRED` → `Não achei <CRED>. Faça login no Claude Code antes de salvar.` no stderr,
   código 1, e nada é gravado.
2. Copia os bytes de `CRED` para `<conta>.credentials.json` (por cima de um save anterior).
3. Se `STATE` tem um `oauthAccount` que é um objeto com campos, grava esse bloco, com o texto
   exato que está em `STATE`, em `<conta>.account.json`. Sem o bloco (campo ausente, `null`,
   `{}`, `STATE` ausente ou inválido), **apaga** o `<conta>.account.json` de um save anterior:
   sobrando, ele daria à conta o e-mail e o `oauthAccount` de outro login.
4. Grava `<conta>` em `CURRENT` e imprime `Conta '<conta>' salva (<email>).`, ou
   `(email desconhecido)`.

## `use <conta>`

Ativa um login salvo. **Tudo é lido e validado antes da primeira escrita**; estes erros saem
com código 1 e não mudam nenhum arquivo:

- `<conta>.credentials.json` não existe → no stderr,
  `Conta '<conta>' não existe. Faça login nela e rode: claude-tools account save <conta>`;
- `<conta>.account.json` existe mas não é um objeto JSON → `Erro: <arquivo> não é um objeto
  JSON válido`;
- a conta tem `account.json` e `STATE` existe mas não é um objeto JSON →
  `Erro: <STATE> não é um objeto JSON válido; nada foi alterado` (regravá-lo perderia o resto
  do estado do Claude Code);
- `CRED`, `STATE` ou `CURRENT` existem mas não podem ser lidos.

Depois, nesta ordem:

1. **Devolve o token renovado.** O Claude Code renova o token dentro de `CRED`, então a cópia
   em `STORE` envelhece. Se `CURRENT` aponta uma conta salva `<cur>` (nome válido com
   `<cur>.credentials.json`; outro conteúdo é ignorado e nada é recriado) e `CRED` existe,
   copia `CRED` para `<cur>.credentials.json`. Vale também quando `<cur>` é a própria
   `<conta>`: aí o passo 4 copia esse token renovado de volta, e não a cópia velha.
2. **Login de fora.** Se o e-mail salvo de `<cur>` e o e-mail ativo (de `STATE`) existem e são
   diferentes, o login ativo é de outra pessoa: o passo 1 não acontece e sai no stdout
   `Aviso: o login ativo (<ativo>) não é o da conta '<cur>' (<salvo>). Não vou sobrescrever a
   conta '<cur>'. Use 'claude-tools account save' se quiser guardá-lo.` (uma linha). Com um dos
   dois desconhecido, o passo 1 acontece. A troca segue nos dois casos.
3. **Backup.** Copia `CRED` para `backup.credentials.json` e `STATE` para
   `backup.claude.json`, cada um só se existir. Num login de fora, é aqui que ele fica guardado.
4. **Troca.** Copia `<conta>.credentials.json` para `CRED` (cria `~/.claude` se precisar). Se
   existem `<conta>.account.json` e `STATE`, troca só o campo `oauthAccount` de `STATE`; sem um
   dos dois, `STATE` fica byte a byte.
5. Grava `<conta>` em `CURRENT` e imprime no stdout os avisos de ambiente (abaixo), o aviso do
   passo 2 e:

   ```
   Conta ativa: <conta> (<email | email desconhecido>)
   Reabra o 'claude' no terminal e use 'Developer: Reload Window' no VS Code.
   ```

**Regravação de `STATE`:** o objeto é lido como `map[string]json.RawMessage` e escrito por um
`json.Encoder` com `SetEscapeHTML(false)` e `SetIndent("", "  ")`. Os valores saem com o mesmo
texto (`12345678901234567890` e `1.50` não passam por `float64`; `<`, `>` e `&` não viram
`<`), com indentação de 2 espaços e uma quebra de linha no fim. Os campos do topo saem em
ordem alfabética, porque o Go não guarda a ordem das chaves de um map; o Claude Code não
depende dela.

**Proteção contra troca pela metade:** entre o backup e a troca, o `use` apaga `CURRENT`. Uma
escrita que falha depois disso (ex.: o Claude Code está com `CRED` aberto) para o `use` com
`Erro: <detalhe>` e deixa o arquivo daquela escrita como estava, mas `CRED` já pode ser o da
conta nova. Sem `CURRENT`, o próximo `use` não devolve nada, então não grava o login da conta
nova por cima da antiga (o que aconteceria quando `STATE` ainda tem o e-mail da antiga, como
numa conta sem `account.json`). A devolução e o backup que já aconteceram são seguros de
repetir: basta rodar de novo. Se a falha vem antes (leitura, devolução ou backup), `CURRENT`
fica como estava.

## `rename <conta> <nova>`

Os dois nomes passam pela regra de nome (saída 2 se um deles é recusado). Antes de mexer em
qualquer arquivo, sai com código 1 e a mensagem no stderr:

- `<conta>.credentials.json` não existe → `Conta '<conta>' não existe.`;
- `<nova>.credentials.json` já existe → `Já existe uma conta '<nova>'.` (inclusive quando os
  dois nomes são o mesmo, como `a` e `A`).

Depois, nesta ordem:

1. Renomeia `<conta>.credentials.json` para `<nova>.credentials.json`.
2. Renomeia `<conta>.account.json` para `<nova>.account.json`, se existir. Se não existe, apaga
   um `<nova>.account.json` que tenha sobrado sem credenciais: ele daria à conta o e-mail de
   outro login.
3. Se `CURRENT` aponta `<conta>` (sem diferenciar maiúsculas), grava `<nova>` nele.
4. Imprime `Conta '<conta>' renomeada para '<nova>'.`

**Em erro, nada muda:** se um passo falha, os renames já feitos são desfeitos na ordem inversa e
o comando sai com `Erro: <detalhe>` e código 1 (com `; ao desfazer: <detalhe>` no fim quando o
próprio desfazer falha). `CRED` e `STATE` não são tocados.

## `remove <conta>`

Apaga a conta salva, sem pedir confirmação (o login ativo não é tocado; perder a cópia se
resolve com `/login` e `save`):

1. `<conta>.credentials.json` não existe → `Conta '<conta>' não existe.` no stderr e código 1.
2. Apaga `<conta>.account.json` (se existir) e depois `<conta>.credentials.json`. Nessa ordem,
   uma falha no primeiro (`Erro: apagar …`, código 1) deixa a conta inteira; uma falha no
   segundo deixa a conta listada sem e-mail, e basta rodar de novo.
3. Imprime `Conta '<conta>' removida.` Se `CURRENT` apontava `<conta>` (sem diferenciar
   maiúsculas), apaga `CURRENT` e imprime também
   `Nenhuma conta ficou marcada como atual; o login ativo do Claude Code não mudou.`

`CRED`, `STATE` e os `backup.*` ficam como estão. Sem `CURRENT`, o próximo `use` não devolve o
login ativo para conta nenhuma (só o guarda em `backup.*`).

## `status`

`claude-tools account status` (sem argumentos; com argumento extra, mostra o uso e sai com 2):

```
Conta atual: <cur | (nenhuma)>
Email ativo no Claude Code: <email | (não encontrado)>
Contas salvas:
  <c>: <email | email desconhecido> (salva em DD/MM/AAAA HH:MM) <- ativa
```

- **Conta atual:** o conteúdo de `CURRENT` sem espaços e quebras de linha nas pontas;
  `(nenhuma)` quando o arquivo não existe ou está vazio.
- **Email ativo:** o e-mail do `oauthAccount` de `STATE`.
- **Contas salvas:** uma linha por `*.credentials.json` de `STORE`, em ordem alfabética do nome;
  `  (nenhuma)` quando não há nenhuma. `backup.*` não aparece. A data é o mtime do
  `.credentials.json`, no fuso local. O ` <- ativa` sai só na conta igual a `CURRENT`.
- Sem `STORE` (nunca salvou nada), sai a lista vazia e o código 0.
- `CURRENT` que existe mas não pode ser lido (ex.: é uma pasta) → `Erro: ler a conta atual: …` e
  código 1.

**Avisos de ambiente** (`status` e `use`, quando dão certo): antes das linhas acima, uma linha
para cada variável `ANTHROPIC_API_KEY` ou `ANTHROPIC_AUTH_TOKEN` definida e não vazia, nesta
ordem:

```
Aviso: <VAR> está definida e passa na frente do login; a troca de conta não terá efeito.
```

Com elas definidas, o Claude Code usa a chave da variável e ignora `CRED`, então trocar de conta
não muda nada.
