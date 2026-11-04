package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/elimarcosmorais/claude-tools/internal/account"
)

const wantUsage = `Uso: claude-tools account <comando>

Comandos:
  save <conta>            guarda o login atual do Claude Code como <conta>
  use <conta>             ativa o login guardado em <conta>
  rename <conta> <nova>   renomeia <conta> para <nova>
  remove <conta>          apaga o login guardado em <conta>
  status                  mostra a conta ativa e as contas salvas
`

// run chama Run com uma home vazia e sem as variáveis que passam na frente do login.
func run(t *testing.T, home string, args ...string) (code int, stdout, stderr string) {
	t.Helper()
	t.Setenv("ANTHROPIC_API_KEY", "")
	t.Setenv("ANTHROPIC_AUTH_TOKEN", "")
	return runWithEnv(home, args...)
}

func runWithEnv(home string, args ...string) (int, string, string) {
	var out, errOut bytes.Buffer
	code := Run(args, home, &out, &errOut)
	return code, out.String(), errOut.String()
}

func TestRunUsage(t *testing.T) {
	cases := []struct {
		args       []string
		code       int
		stdout     string
		stderr     string
		unknownCmd string
	}{
		{args: nil, code: 0, stdout: wantUsage},
		{args: []string{"help"}, code: 0, stdout: wantUsage},
		{args: []string{"account"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "help"}, code: 0, stdout: wantUsage},
		{args: []string{"foo"}, code: 2, stderr: "Comando desconhecido: foo\n\n" + wantUsage},
		{args: []string{"account", "foo"}, code: 2, stderr: "Comando desconhecido: foo\n\n" + wantUsage},
		{args: []string{"account", "status", "x"}, code: 2, stderr: wantUsage},
		{args: []string{"help", "x"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "help", "x"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "save"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "save", "a", "b"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "use"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "use", "a", "b"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "rename", "a"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "rename", "a", "b", "c"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "remove"}, code: 2, stderr: wantUsage},
		{args: []string{"account", "remove", "a", "b"}, code: 2, stderr: wantUsage},
	}
	for _, c := range cases {
		t.Run(strings.Join(c.args, "_"), func(t *testing.T) {
			code, stdout, stderr := run(t, tempDir(t), c.args...)
			if code != c.code {
				t.Errorf("código = %d, quero %d", code, c.code)
			}
			if stdout != c.stdout {
				t.Errorf("stdout =\n%q\nquero\n%q", stdout, c.stdout)
			}
			if stderr != c.stderr {
				t.Errorf("stderr =\n%q\nquero\n%q", stderr, c.stderr)
			}
		})
	}
}

func TestReportNameErrors(t *testing.T) {
	long := strings.Repeat("a", 33)
	cases := map[string]string{
		"a.b":    "Nome de conta inválido: 'a.b'. Use letras, números, '-' ou '_' (até 32 caracteres).\n",
		"":       "Nome de conta inválido: ''. Use letras, números, '-' ou '_' (até 32 caracteres).\n",
		long:     "Nome de conta inválido: '" + long + "'. Use letras, números, '-' ou '_' (até 32 caracteres).\n",
		"backup": "Nome de conta reservado: 'backup'.\n",
	}
	for name, want := range cases {
		_, err := account.NormalizeName(name)
		var stderr bytes.Buffer
		if code := report(err, &stderr); code != 2 {
			t.Errorf("%q: código = %d, quero 2", name, code)
		}
		if stderr.String() != want {
			t.Errorf("%q: stderr =\n%q\nquero\n%q", name, stderr.String(), want)
		}
	}
}

func TestReportOperationError(t *testing.T) {
	var stderr bytes.Buffer
	_, err := os.ReadFile(filepath.Join(tempDir(t), "nada.json"))

	code := report(err, &stderr)

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	if want := "Erro: " + err.Error() + "\n"; stderr.String() != want {
		t.Errorf("stderr = %q, quero %q", stderr.String(), want)
	}
}

func TestStatusEmptyHome(t *testing.T) {
	code, stdout, stderr := run(t, tempDir(t), "account", "status")

	want := "Conta atual: (nenhuma)\n" +
		"Email ativo no Claude Code: (não encontrado)\n" +
		"Contas salvas:\n" +
		"  (nenhuma)\n"
	if code != 0 {
		t.Errorf("código = %d, quero 0", code)
	}
	if stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
}

func TestStatusWithAccounts(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(home, ".claude.json"), `{"oauthAccount": {"emailAddress": "b@x.com"}}`)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "a.account.json"), `{"emailAddress": "a@x.com"}`)
	writeFile(t, filepath.Join(store, "backup.credentials.json"), `{"t":"x"}`)
	writeFile(t, filepath.Join(store, "backup.claude.json"), `{}`)
	writeFile(t, filepath.Join(store, "current.txt"), "b")
	setMtime(t, filepath.Join(store, "a.credentials.json"), time.Date(2026, 3, 5, 14, 7, 0, 0, time.Local))
	setMtime(t, filepath.Join(store, "b.credentials.json"), time.Date(2026, 12, 31, 9, 30, 59, 0, time.Local))

	code, stdout, stderr := run(t, home, "account", "status")

	want := "Conta atual: b\n" +
		"Email ativo no Claude Code: b@x.com\n" +
		"Contas salvas:\n" +
		"  a: a@x.com (salva em 05/03/2026 14:07)\n" +
		"  b: email desconhecido (salva em 31/12/2026 09:30) <- ativa\n"
	if code != 0 {
		t.Errorf("código = %d, quero 0", code)
	}
	if stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
}

func TestStatusEnvWarnings(t *testing.T) {
	const lines = "Conta atual: (nenhuma)\n" +
		"Email ativo no Claude Code: (não encontrado)\n" +
		"Contas salvas:\n" +
		"  (nenhuma)\n"
	const apiKey = "Aviso: ANTHROPIC_API_KEY está definida e passa na frente do login; a troca de conta não terá efeito.\n"
	const authToken = "Aviso: ANTHROPIC_AUTH_TOKEN está definida e passa na frente do login; a troca de conta não terá efeito.\n"
	cases := []struct {
		name, apiKey, authToken, want string
	}{
		{"api key definida", "sk-ant-x", "", apiKey + lines},
		{"as duas definidas", "sk-ant-x", "tok", apiKey + authToken + lines},
		{"auth token definida", "", "tok", authToken + lines},
		{"as duas vazias", "", "", lines},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Setenv("ANTHROPIC_API_KEY", c.apiKey)
			t.Setenv("ANTHROPIC_AUTH_TOKEN", c.authToken)

			code, stdout, stderr := runWithEnv(tempDir(t), "account", "status")

			if code != 0 {
				t.Errorf("código = %d, quero 0", code)
			}
			if stdout != c.want {
				t.Errorf("stdout =\n%s\nquero\n%s", stdout, c.want)
			}
			if stderr != "" {
				t.Errorf("stderr = %q, quero vazio", stderr)
			}
		})
	}
}

func TestStatusError(t *testing.T) {
	home := tempDir(t)
	if err := os.MkdirAll(filepath.Join(home, ".claude", "accounts", "current.txt"), 0o755); err != nil {
		t.Fatal(err)
	}

	code, stdout, stderr := run(t, home, "account", "status")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, quero vazio", stdout)
	}
	if !strings.HasPrefix(stderr, "Erro: ler a conta atual: ") || !strings.HasSuffix(stderr, "\n") {
		t.Errorf("stderr = %q, quero 'Erro: ler a conta atual: ...'", stderr)
	}
}

const reloadHint = "Reabra o 'claude' no terminal e use 'Developer: Reload Window' no VS Code.\n"

// paths são os arquivos do Claude Code e das contas dentro de uma home de teste.
type paths struct {
	home, cred, state, store, current string
}

func newPaths(home string) paths {
	store := filepath.Join(home, ".claude", "accounts")
	return paths{
		home:    home,
		cred:    filepath.Join(home, ".claude", ".credentials.json"),
		state:   filepath.Join(home, ".claude.json"),
		store:   store,
		current: filepath.Join(store, "current.txt"),
	}
}

func (p paths) inStore(name string) string { return filepath.Join(p.store, name) }

func TestSaveWithOAuthAccount(t *testing.T) {
	p := newPaths(tempDir(t))
	const cred = "{\"claudeAiOauth\":{\"accessToken\":\"tok-a\",\"expiresAt\":12345678901234567890}}\r\n"
	const acc = `{"emailAddress": "a@x.com", "accountUuid": "u-a", "n": 12345678901234567890}`
	writeFile(t, p.cred, cred)
	writeFile(t, p.state, `{"numStartups": 3, "oauthAccount": `+acc+`, "tip": "a<b>"}`)

	code, stdout, stderr := run(t, p.home, "account", "save", "A")

	if code != 0 {
		t.Errorf("código = %d, quero 0", code)
	}
	if want := "Conta 'a' salva (a@x.com).\n"; stdout != want {
		t.Errorf("stdout = %q, quero %q", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
	assertFile(t, p.inStore("a.credentials.json"), cred)
	assertFile(t, p.inStore("a.account.json"), acc)
	assertFile(t, p.current, "a")
	assertFiles(t, p.store, "a.account.json", "a.credentials.json", "current.txt")
}

func TestSaveWithoutOAuthAccount(t *testing.T) {
	cases := map[string]string{
		"STATE sem o campo":  `{"numStartups": 3}`,
		"oauthAccount vazio": `{"oauthAccount": {}}`,
		"oauthAccount null":  `{"oauthAccount": null}`,
		"STATE inválido":     `{"oauthAccount": `,
		"sem STATE":          "",
	}
	for name, state := range cases {
		t.Run(name, func(t *testing.T) {
			p := newPaths(tempDir(t))
			writeFile(t, p.cred, `{"t":"a"}`)
			if state != "" {
				writeFile(t, p.state, state)
			}

			code, stdout, stderr := run(t, p.home, "account", "save", "a")

			if code != 0 {
				t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
			}
			if want := "Conta 'a' salva (email desconhecido).\n"; stdout != want {
				t.Errorf("stdout = %q, quero %q", stdout, want)
			}
			assertFile(t, p.inStore("a.credentials.json"), `{"t":"a"}`)
			assertFiles(t, p.store, "a.credentials.json", "current.txt")
		})
	}
}

// Um save sem oauthAccount por cima de uma conta que tinha o bloco não pode deixar
// o account.json antigo: ele daria o e-mail e o oauthAccount de outra pessoa.
func TestSaveWithoutOAuthAccountRemovesStaleAccount(t *testing.T) {
	p := newPaths(tempDir(t))
	writeFile(t, p.cred, `{"t":"novo"}`)
	writeFile(t, p.inStore("a.credentials.json"), `{"t":"velho"}`)
	writeFile(t, p.inStore("a.account.json"), `{"emailAddress": "velho@x.com"}`)

	code, stdout, _ := run(t, p.home, "account", "save", "a")

	if code != 0 || stdout != "Conta 'a' salva (email desconhecido).\n" {
		t.Errorf("código = %d, stdout = %q", code, stdout)
	}
	assertFile(t, p.inStore("a.credentials.json"), `{"t":"novo"}`)
	assertFiles(t, p.store, "a.credentials.json", "current.txt")
}

func TestSaveWithoutCred(t *testing.T) {
	p := newPaths(tempDir(t))
	writeFile(t, p.state, `{"oauthAccount": {"emailAddress": "a@x.com"}}`)
	writeFile(t, p.inStore("b.credentials.json"), `{"t":"b"}`)
	before := snapshot(t, p.home)

	code, stdout, stderr := run(t, p.home, "account", "save", "a")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, quero vazio", stdout)
	}
	want := "Não achei " + p.cred + ". Faça login no Claude Code antes de salvar.\n"
	if stderr != want {
		t.Errorf("stderr = %q, quero %q", stderr, want)
	}
	assertSnapshot(t, p.home, before)
}

func TestSaveInvalidName(t *testing.T) {
	p := newPaths(tempDir(t))
	writeFile(t, p.cred, `{"t":"a"}`)
	before := snapshot(t, p.home)

	code, stdout, stderr := run(t, p.home, "account", "save", "backup")

	if code != 2 || stdout != "" || stderr != "Nome de conta reservado: 'backup'.\n" {
		t.Errorf("código = %d, stdout = %q, stderr = %q", code, stdout, stderr)
	}
	assertSnapshot(t, p.home, before)
}

func TestUseMissingAccount(t *testing.T) {
	p := newPaths(tempDir(t))
	writeFile(t, p.cred, `{"t":"a"}`)
	writeFile(t, p.state, `{"oauthAccount": {"emailAddress": "a@x.com"}}`)
	writeFile(t, p.inStore("a.credentials.json"), `{"t":"a"}`)
	writeFile(t, p.inStore("a.account.json"), `{"emailAddress": "a@x.com"}`)
	writeFile(t, p.current, "a")
	before := snapshot(t, p.home)

	code, stdout, stderr := run(t, p.home, "account", "use", "X")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, quero vazio", stdout)
	}
	if want := "Conta 'x' não existe. Faça login nela e rode: claude-tools account save x\n"; stderr != want {
		t.Errorf("stderr = %q, quero %q", stderr, want)
	}
	assertSnapshot(t, p.home, before)
}

func TestUseInvalidName(t *testing.T) {
	p := newPaths(tempDir(t))
	before := snapshot(t, p.home)

	code, stdout, stderr := run(t, p.home, "account", "use", `..\a`)

	want := "Nome de conta inválido: '..\\a'. Use letras, números, '-' ou '_' (até 32 caracteres).\n"
	if code != 2 || stdout != "" || stderr != want {
		t.Errorf("código = %d, stdout = %q, stderr = %q", code, stdout, stderr)
	}
	assertSnapshot(t, p.home, before)
}

// useFixture monta a conta a como atual (com o login ativo dela já renovado) e a
// conta b salva. activeEmail é o e-mail do oauthAccount de STATE.
func useFixture(t *testing.T, activeEmail string) paths {
	t.Helper()
	p := newPaths(tempDir(t))
	writeFile(t, p.cred, `{"t":"a-renovado"}`)
	writeFile(t, p.state, `{
  "numStartups": 12345678901234567890,
  "oauthAccount": {"emailAddress": "`+activeEmail+`", "accountUuid": "u-a"},
  "tip": "a<b>&c",
  "projects": {"C:\\dev\\x": {"cost": 1.50, "tags": [], "on": true, "none": null}}
}`)
	writeFile(t, p.inStore("a.credentials.json"), `{"t":"a-velho"}`)
	writeFile(t, p.inStore("a.account.json"), `{"emailAddress": "a@x.com", "accountUuid": "u-a"}`)
	writeFile(t, p.inStore("b.credentials.json"), `{"t":"b"}`)
	writeFile(t, p.inStore("b.account.json"), `{"emailAddress": "b@x.com", "accountUuid": "u-b"}`)
	writeFile(t, p.current, "a")
	return p
}

// wantStateB é o STATE da useFixture depois da troca para b: só o oauthAccount
// muda, os valores ficam com o mesmo texto e a indentação é de 2 espaços.
const wantStateB = `{
  "numStartups": 12345678901234567890,
  "oauthAccount": {
    "emailAddress": "b@x.com",
    "accountUuid": "u-b"
  },
  "projects": {
    "C:\\dev\\x": {
      "cost": 1.50,
      "tags": [],
      "on": true,
      "none": null
    }
  },
  "tip": "a<b>&c"
}
`

func TestUseSwitchesAccount(t *testing.T) {
	p := useFixture(t, "a@x.com")
	stateBefore := readFile(t, p.state)

	code, stdout, stderr := run(t, p.home, "account", "use", "b")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	if want := "Conta ativa: b (b@x.com)\n" + reloadHint; stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
	assertFile(t, p.inStore("a.credentials.json"), `{"t":"a-renovado"}`)
	assertFile(t, p.cred, `{"t":"b"}`)
	assertFile(t, p.inStore("backup.credentials.json"), `{"t":"a-renovado"}`)
	assertFile(t, p.inStore("backup.claude.json"), stateBefore)
	assertFile(t, p.state, wantStateB)
	assertFile(t, p.current, "b")
	assertFile(t, p.inStore("b.credentials.json"), `{"t":"b"}`)
	assertFile(t, p.inStore("b.account.json"), `{"emailAddress": "b@x.com", "accountUuid": "u-b"}`)
	assertFiles(t, p.store, "a.account.json", "a.credentials.json", "b.account.json", "b.credentials.json",
		"backup.claude.json", "backup.credentials.json", "current.txt")
	assertFiles(t, p.home, ".claude", ".claude.json")
}

func TestUseForeignLogin(t *testing.T) {
	p := useFixture(t, "intruso@x.com")
	stateBefore := readFile(t, p.state)

	code, stdout, stderr := run(t, p.home, "account", "use", "b")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	want := "Aviso: o login ativo (intruso@x.com) não é o da conta 'a' (a@x.com). Não vou sobrescrever a conta 'a'. Use 'claude-tools account save' se quiser guardá-lo.\n" +
		"Conta ativa: b (b@x.com)\n" + reloadHint
	if stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
	assertFile(t, p.inStore("a.credentials.json"), `{"t":"a-velho"}`)
	assertFile(t, p.inStore("backup.credentials.json"), `{"t":"a-renovado"}`)
	assertFile(t, p.inStore("backup.claude.json"), stateBefore)
	assertFile(t, p.cred, `{"t":"b"}`)
	assertFile(t, p.current, "b")
}

// Com um dos dois e-mails desconhecido, o login ativo é devolvido para a conta atual.
func TestUseUnknownEmailGivesTokenBack(t *testing.T) {
	cases := map[string]func(p paths){
		"e-mail salvo desconhecido": func(p paths) { removeFile(t, p.inStore("a.account.json")) },
		"e-mail ativo desconhecido": func(p paths) { writeFile(t, p.state, `{"numStartups": 1}`) },
		"sem STATE":                 func(p paths) { removeFile(t, p.state) },
	}
	for name, change := range cases {
		t.Run(name, func(t *testing.T) {
			p := useFixture(t, "intruso@x.com")
			change(p)

			code, stdout, stderr := run(t, p.home, "account", "use", "b")

			if code != 0 {
				t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
			}
			if want := "Conta ativa: b (b@x.com)\n" + reloadHint; stdout != want {
				t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
			}
			assertFile(t, p.inStore("a.credentials.json"), `{"t":"a-renovado"}`)
			assertFile(t, p.cred, `{"t":"b"}`)
		})
	}
}

func TestUseCurrentAccountKeepsRenewedToken(t *testing.T) {
	p := useFixture(t, "a@x.com")

	code, stdout, stderr := run(t, p.home, "account", "use", "a")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	if want := "Conta ativa: a (a@x.com)\n" + reloadHint; stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
	assertFile(t, p.inStore("a.credentials.json"), `{"t":"a-renovado"}`)
	assertFile(t, p.cred, `{"t":"a-renovado"}`)
	assertFile(t, p.current, "a")
}

func TestUseEnvWarning(t *testing.T) {
	p := useFixture(t, "a@x.com")
	t.Setenv("ANTHROPIC_API_KEY", "")
	t.Setenv("ANTHROPIC_AUTH_TOKEN", "tok")

	code, stdout, stderr := runWithEnv(p.home, "account", "use", "b")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	want := "Aviso: ANTHROPIC_AUTH_TOKEN está definida e passa na frente do login; a troca de conta não terá efeito.\n" +
		"Conta ativa: b (b@x.com)\n" + reloadHint
	if stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
}

func TestUseInvalidStateChangesNothing(t *testing.T) {
	p := useFixture(t, "a@x.com")
	writeFile(t, p.state, `{"oauthAccount": `)
	before := snapshot(t, p.home)

	code, stdout, stderr := run(t, p.home, "account", "use", "b")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	if stdout != "" {
		t.Errorf("stdout = %q, quero vazio", stdout)
	}
	if want := "Erro: " + p.state + " não é um objeto JSON válido; nada foi alterado\n"; stderr != want {
		t.Errorf("stderr = %q, quero %q", stderr, want)
	}
	assertSnapshot(t, p.home, before)
}

// storeFixture monta as contas a (com account.json, atual) e b (sem account.json),
// e um login ativo em CRED e STATE.
func storeFixture(t *testing.T) paths {
	t.Helper()
	p := newPaths(tempDir(t))
	writeFile(t, p.cred, `{"t":"ativo"}`)
	writeFile(t, p.state, `{"oauthAccount": {"emailAddress": "a@x.com"}, "n": 12345678901234567890}`)
	writeFile(t, p.inStore("a.credentials.json"), `{"t":"a"}`)
	writeFile(t, p.inStore("a.account.json"), `{"emailAddress": "a@x.com"}`)
	writeFile(t, p.inStore("b.credentials.json"), `{"t":"b"}`)
	writeFile(t, p.inStore("backup.credentials.json"), `{"t":"backup"}`)
	writeFile(t, p.current, "a")
	return p
}

func TestRenameCurrentAccount(t *testing.T) {
	p := storeFixture(t)

	code, stdout, stderr := run(t, p.home, "account", "rename", "A", "X")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	if want := "Conta 'a' renomeada para 'x'.\n"; stdout != want {
		t.Errorf("stdout = %q, quero %q", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
	assertFile(t, p.inStore("x.credentials.json"), `{"t":"a"}`)
	assertFile(t, p.inStore("x.account.json"), `{"emailAddress": "a@x.com"}`)
	assertFile(t, p.current, "x")
	assertFiles(t, p.store, "b.credentials.json", "backup.credentials.json", "current.txt",
		"x.account.json", "x.credentials.json")
	assertFile(t, p.cred, `{"t":"ativo"}`)
}

func TestRenameAccountWithoutAccountJSON(t *testing.T) {
	p := storeFixture(t)

	code, stdout, stderr := run(t, p.home, "account", "rename", "b", "y")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	if want := "Conta 'b' renomeada para 'y'.\n"; stdout != want {
		t.Errorf("stdout = %q, quero %q", stdout, want)
	}
	assertFile(t, p.inStore("y.credentials.json"), `{"t":"b"}`)
	assertFile(t, p.current, "a")
	assertFiles(t, p.store, "a.account.json", "a.credentials.json", "backup.credentials.json", "current.txt",
		"y.credentials.json")
}

func TestRenameErrorsChangeNothing(t *testing.T) {
	cases := []struct {
		name     string
		from, to string
		code     int
		stderr   string
	}{
		{"origem inexistente", "z", "w", 1, "Conta 'z' não existe.\n"},
		{"destino existente", "a", "B", 1, "Já existe uma conta 'b'.\n"},
		{"mesmo nome", "a", "A", 1, "Já existe uma conta 'a'.\n"},
		{"destino inválido", "a", "x.y", 2, "Nome de conta inválido: 'x.y'. Use letras, números, '-' ou '_' (até 32 caracteres).\n"},
		{"destino reservado", "a", "backup", 2, "Nome de conta reservado: 'backup'.\n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := storeFixture(t)
			before := snapshot(t, p.home)

			code, stdout, stderr := run(t, p.home, "account", "rename", c.from, c.to)

			if code != c.code {
				t.Errorf("código = %d, quero %d", code, c.code)
			}
			if stdout != "" {
				t.Errorf("stdout = %q, quero vazio", stdout)
			}
			if stderr != c.stderr {
				t.Errorf("stderr = %q, quero %q", stderr, c.stderr)
			}
			assertSnapshot(t, p.home, before)
		})
	}
}

func TestRemoveOtherAccount(t *testing.T) {
	p := storeFixture(t)
	writeFile(t, p.inStore("b.account.json"), `{"emailAddress": "b@x.com"}`)

	code, stdout, stderr := run(t, p.home, "account", "remove", "B")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	if want := "Conta 'b' removida.\n"; stdout != want {
		t.Errorf("stdout = %q, quero %q", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
	assertFile(t, p.current, "a")
	assertFiles(t, p.store, "a.account.json", "a.credentials.json", "backup.credentials.json", "current.txt")
}

func TestRemoveCurrentAccount(t *testing.T) {
	p := storeFixture(t)
	cred, state := readFile(t, p.cred), readFile(t, p.state)

	code, stdout, stderr := run(t, p.home, "account", "remove", "a")

	if code != 0 {
		t.Errorf("código = %d, quero 0 (stderr %q)", code, stderr)
	}
	want := "Conta 'a' removida.\n" +
		"Nenhuma conta ficou marcada como atual; o login ativo do Claude Code não mudou.\n"
	if stdout != want {
		t.Errorf("stdout =\n%s\nquero\n%s", stdout, want)
	}
	if stderr != "" {
		t.Errorf("stderr = %q, quero vazio", stderr)
	}
	assertFiles(t, p.store, "b.credentials.json", "backup.credentials.json")
	assertFile(t, p.cred, cred)
	assertFile(t, p.state, state)
}

func TestRemoveErrorsChangeNothing(t *testing.T) {
	cases := []struct {
		name, account string
		code          int
		stderr        string
	}{
		{"conta inexistente", "z", 1, "Conta 'z' não existe.\n"},
		{"backup", "backup", 2, "Nome de conta reservado: 'backup'.\n"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := storeFixture(t)
			before := snapshot(t, p.home)

			code, stdout, stderr := run(t, p.home, "account", "remove", c.account)

			if code != c.code {
				t.Errorf("código = %d, quero %d", code, c.code)
			}
			if stdout != "" {
				t.Errorf("stdout = %q, quero vazio", stdout)
			}
			if stderr != c.stderr {
				t.Errorf("stderr = %q, quero %q", stderr, c.stderr)
			}
			assertSnapshot(t, p.home, before)
		})
	}
}

func assertFile(t *testing.T, path, want string) {
	t.Helper()
	if got := readFile(t, path); got != want {
		t.Errorf("%s =\n%s\nquero\n%s", filepath.Base(path), got, want)
	}
}

// assertFiles confere os nomes (em ordem alfabética) dentro de dir.
func assertFiles(t *testing.T, dir string, want ...string) {
	t.Helper()
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	var got []string
	for _, e := range entries {
		got = append(got, e.Name())
	}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Errorf("arquivos em %s = %v, quero %v", dir, got, want)
	}
}

// snapshot devolve o conteúdo de cada arquivo dentro de root, pelo caminho relativo.
func snapshot(t *testing.T, root string) map[string]string {
	t.Helper()
	files := map[string]string{}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, _ := filepath.Rel(root, path)
		files[rel] = readFile(t, path)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return files
}

func assertSnapshot(t *testing.T, root string, want map[string]string) {
	t.Helper()
	got := snapshot(t, root)
	for name, content := range want {
		if g, ok := got[name]; !ok {
			t.Errorf("%s sumiu", name)
		} else if g != content {
			t.Errorf("%s mudou:\n%s\nera\n%s", name, g, content)
		}
	}
	for name := range got {
		if _, ok := want[name]; !ok {
			t.Errorf("%s apareceu", name)
		}
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func removeFile(t *testing.T, path string) {
	t.Helper()
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
}

// tempDir é um t.TempDir() que tolera o antivírus do Windows na limpeza: ele abre os
// arquivos recém-gravados para escanear, o arquivo apagado fica "delete pending" e a
// pasta só esvazia quando ele solta. O t.TempDir() não repete o RemoveAll nesse erro
// (ERROR_DIR_NOT_EMPTY) e reprova o teste.
func tempDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("", "claude-tools-test-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		deadline := time.Now().Add(5 * time.Second)
		for {
			err := os.RemoveAll(dir)
			if err == nil {
				return
			}
			if time.Now().After(deadline) {
				t.Errorf("apagar a pasta temporária: %v", err)
				return
			}
			time.Sleep(20 * time.Millisecond)
		}
	})
	return dir
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func setMtime(t *testing.T, path string, when time.Time) {
	t.Helper()
	if err := os.Chtimes(path, when, when); err != nil {
		t.Fatal(err)
	}
}
