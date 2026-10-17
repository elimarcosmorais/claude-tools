package account

import (
	"errors"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestNormalizeNameAccepts(t *testing.T) {
	cases := map[string]string{
		"A_b-1":                 "a_b-1",
		"a":                     "a",
		strings.Repeat("X", 32): strings.Repeat("x", 32),
		"backup2":               "backup2",
	}
	for in, want := range cases {
		got, err := NormalizeName(in)
		if err != nil {
			t.Errorf("NormalizeName(%q) devolveu erro %v", in, err)
			continue
		}
		if got != want {
			t.Errorf("NormalizeName(%q) = %q, quero %q", in, got, want)
		}
	}
}

func TestNormalizeNameRejects(t *testing.T) {
	long := strings.Repeat("a", 33)
	cases := []struct {
		in       string
		msg      string
		reserved bool
	}{
		{"a.b", "Nome de conta inválido: 'a.b'. Use letras, números, '-' ou '_' (até 32 caracteres).", false},
		{"", "Nome de conta inválido: ''. Use letras, números, '-' ou '_' (até 32 caracteres).", false},
		{long, "Nome de conta inválido: '" + long + "'. Use letras, números, '-' ou '_' (até 32 caracteres).", false},
		{"ação", "Nome de conta inválido: 'ação'. Use letras, números, '-' ou '_' (até 32 caracteres).", false},
		{"a b", "Nome de conta inválido: 'a b'. Use letras, números, '-' ou '_' (até 32 caracteres).", false},
		{"backup", "Nome de conta reservado: 'backup'.", true},
		{"BackUp", "Nome de conta reservado: 'backup'.", true},
	}
	for _, c := range cases {
		got, err := NormalizeName(c.in)
		if err == nil {
			t.Errorf("NormalizeName(%q) = %q, quero erro", c.in, got)
			continue
		}
		var nameErr *NameError
		if !errors.As(err, &nameErr) {
			t.Errorf("NormalizeName(%q): erro %T, quero *NameError", c.in, err)
			continue
		}
		if nameErr.Reserved != c.reserved {
			t.Errorf("NormalizeName(%q): Reserved = %v, quero %v", c.in, nameErr.Reserved, c.reserved)
		}
		if err.Error() != c.msg {
			t.Errorf("NormalizeName(%q): mensagem\n%q\nquero\n%q", c.in, err.Error(), c.msg)
		}
		if got != "" {
			t.Errorf("NormalizeName(%q) devolveu nome %q junto com o erro", c.in, got)
		}
	}
}

func TestWriteFileAtomicReplacesExisting(t *testing.T) {
	dir := tempDir(t)
	path := filepath.Join(dir, "alvo.json")
	writeFile(t, path, "antigo")

	if err := writeFileAtomic(path, []byte("novo")); err != nil {
		t.Fatalf("writeFileAtomic: %v", err)
	}

	if got := readFile(t, path); got != "novo" {
		t.Errorf("conteúdo = %q, quero %q", got, "novo")
	}
	assertDirEntries(t, dir, "alvo.json")
}

func TestWriteFileAtomicCreatesNew(t *testing.T) {
	dir := tempDir(t)
	path := filepath.Join(dir, "novo.json")

	if err := writeFileAtomic(path, []byte("{}")); err != nil {
		t.Fatalf("writeFileAtomic: %v", err)
	}

	if got := readFile(t, path); got != "{}" {
		t.Errorf("conteúdo = %q, quero %q", got, "{}")
	}
	assertDirEntries(t, dir, "novo.json")
}

func TestWriteFileAtomicMissingDir(t *testing.T) {
	dir := tempDir(t)
	path := filepath.Join(dir, "nao-existe", "alvo.json")

	err := writeFileAtomic(path, []byte("novo"))

	if !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("erro = %v, quero um os.ErrNotExist", err)
	}
	assertDirEntries(t, dir)
}

func TestWriteFileAtomicTargetIsDir(t *testing.T) {
	dir := tempDir(t)
	path := filepath.Join(dir, "alvo")
	if err := os.Mkdir(path, 0o755); err != nil {
		t.Fatal(err)
	}
	writeFile(t, filepath.Join(path, "dentro.txt"), "original")

	err := writeFileAtomic(path, []byte("novo"))

	if err == nil {
		t.Fatal("erro = nil, quero erro ao gravar por cima de um diretório")
	}
	if !strings.Contains(err.Error(), "alvo") {
		t.Errorf("erro %q não cita o arquivo de destino", err)
	}
	info, statErr := os.Stat(path)
	if statErr != nil || !info.IsDir() {
		t.Fatalf("o diretório de destino não ficou intacto: %v", statErr)
	}
	if got := readFile(t, filepath.Join(path, "dentro.txt")); got != "original" {
		t.Errorf("conteúdo do diretório mudou: %q", got)
	}
	assertDirEntries(t, dir, "alvo")
}

func TestStatusEmptyHome(t *testing.T) {
	st, err := Store{Home: tempDir(t)}.Status()
	if err != nil {
		t.Fatalf("Status: %v", err)
	}
	if st.Current != "" || st.ActiveEmail != "" || len(st.Accounts) != 0 {
		t.Errorf("Status = %+v, quero tudo vazio", st)
	}
}

func TestStatusWithAccounts(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(home, ".claude.json"),
		`{"numStartups": 3, "oauthAccount": {"emailAddress": "ativo@x.com", "email": "outro@x.com"}}`)
	writeFile(t, filepath.Join(store, "zeta.credentials.json"), `{"t":"z"}`)
	writeFile(t, filepath.Join(store, "zeta.account.json"), `{"email": "zeta@x.com"}`)
	writeFile(t, filepath.Join(store, "alfa.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "alfa.account.json"), `{"emailAddress": "alfa@x.com"}`)
	writeFile(t, filepath.Join(store, "meio.credentials.json"), `{"t":"m"}`)
	writeFile(t, filepath.Join(store, "backup.credentials.json"), `{"t":"b"}`)
	writeFile(t, filepath.Join(store, "backup.claude.json"), `{}`)
	writeFile(t, filepath.Join(store, "current.txt"), "zeta\r\n")
	when := time.Date(2026, 3, 5, 14, 7, 0, 0, time.Local)
	if err := os.Chtimes(filepath.Join(store, "alfa.credentials.json"), when, when); err != nil {
		t.Fatal(err)
	}

	st, err := Store{Home: home}.Status()
	if err != nil {
		t.Fatalf("Status: %v", err)
	}

	if st.Current != "zeta" {
		t.Errorf("Current = %q, quero %q", st.Current, "zeta")
	}
	if st.ActiveEmail != "ativo@x.com" {
		t.Errorf("ActiveEmail = %q, quero %q", st.ActiveEmail, "ativo@x.com")
	}
	var names, emails []string
	for _, a := range st.Accounts {
		names = append(names, a.Name)
		emails = append(emails, a.Email)
	}
	if got, want := strings.Join(names, ","), "alfa,meio,zeta"; got != want {
		t.Errorf("contas = %s, quero %s", got, want)
	}
	if got, want := strings.Join(emails, ","), "alfa@x.com,,zeta@x.com"; got != want {
		t.Errorf("e-mails = %s, quero %s", got, want)
	}
	if !st.Accounts[0].SavedAt.Equal(when) {
		t.Errorf("SavedAt = %v, quero %v", st.Accounts[0].SavedAt, when)
	}
}

func TestStatusIgnoresBrokenJSON(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(home, ".claude.json"), `{"oauthAccount": `)
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{}`)
	writeFile(t, filepath.Join(store, "a.account.json"), `não é json`)

	st, err := Store{Home: home}.Status()
	if err != nil {
		t.Fatalf("Status: %v", err)
	}
	if st.ActiveEmail != "" {
		t.Errorf("ActiveEmail = %q, quero vazio", st.ActiveEmail)
	}
	if len(st.Accounts) != 1 || st.Accounts[0].Name != "a" || st.Accounts[0].Email != "" {
		t.Errorf("Accounts = %+v, quero só 'a' sem e-mail", st.Accounts)
	}
}

func TestStatusCurrentUnreadable(t *testing.T) {
	home := tempDir(t)
	if err := os.MkdirAll(filepath.Join(home, ".claude", "accounts", "current.txt"), 0o755); err != nil {
		t.Fatal(err)
	}

	_, err := Store{Home: home}.Status()

	if err == nil {
		t.Fatal("erro = nil, quero erro quando current.txt não pode ser lido")
	}
	if !strings.HasPrefix(err.Error(), "ler a conta atual: ") || !strings.Contains(err.Error(), "current.txt") {
		t.Errorf("erro = %q, quero 'ler a conta atual: ...current.txt...'", err)
	}
}

func TestSaveNormalizesName(t *testing.T) {
	home := tempDir(t)
	writeFile(t, filepath.Join(home, ".claude", ".credentials.json"), `{"t":"a"}`)

	res, err := Store{Home: home}.Save("Trabalho")

	if err != nil {
		t.Fatalf("Save: %v", err)
	}
	if res != (SaveResult{Name: "trabalho"}) {
		t.Errorf("Save = %+v, quero {Name:trabalho}", res)
	}
	assertDirEntries(t, filepath.Join(home, ".claude", "accounts"), "current.txt", "trabalho.credentials.json")
	if got := readFile(t, filepath.Join(home, ".claude", "accounts", "current.txt")); got != "trabalho" {
		t.Errorf("current.txt = %q, quero %q", got, "trabalho")
	}
}

func TestSaveWithoutCredReturnsNoLoginError(t *testing.T) {
	home := tempDir(t)

	_, err := Store{Home: home}.Save("a")

	var noLogin *NoLoginError
	if !errors.As(err, &noLogin) {
		t.Fatalf("erro = %v (%T), quero *NoLoginError", err, err)
	}
	if want := filepath.Join(home, ".claude", ".credentials.json"); noLogin.Path != want {
		t.Errorf("Path = %q, quero %q", noLogin.Path, want)
	}
	assertDirEntries(t, home)
}

func TestUseMissingAccountReturnsError(t *testing.T) {
	_, err := Store{Home: tempDir(t)}.Use("B")

	var missing *MissingAccountError
	if !errors.As(err, &missing) || missing.Name != "b" {
		t.Fatalf("erro = %v (%T), quero *MissingAccountError{Name: b}", err, err)
	}
	if want := "Conta 'b' não existe."; err.Error() != want {
		t.Errorf("mensagem = %q, quero %q", err.Error(), want)
	}
}

// Sem login ativo não há o que devolver nem guardar em backup.credentials.json, e
// a pasta ~/.claude é criada para receber as credenciais.
func TestUseWithoutCred(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
	writeFile(t, filepath.Join(store, "current.txt"), "a")

	res, err := Store{Home: home}.Use("b")

	if err != nil {
		t.Fatalf("Use: %v", err)
	}
	if res != (UseResult{Name: "b"}) {
		t.Errorf("Use = %+v, quero {Name:b}", res)
	}
	if got := readFile(t, filepath.Join(home, ".claude", ".credentials.json")); got != `{"t":"b"}` {
		t.Errorf("credenciais = %q, quero as de b", got)
	}
	if got := readFile(t, filepath.Join(store, "a.credentials.json")); got != `{"t":"a"}` {
		t.Errorf("a.credentials.json = %q, quero sem mudar", got)
	}
	assertDirEntries(t, store, "a.credentials.json", "b.credentials.json", "current.txt")
	assertDirEntries(t, home, ".claude")
}

// current.txt apontando uma conta que não foi salva (ou um nome inválido) não
// recria a conta nem grava fora da pasta de contas.
func TestUseCurrentNotSaved(t *testing.T) {
	for _, current := range []string{"sumiu", `..\fora`} {
		t.Run(current, func(t *testing.T) {
			home := tempDir(t)
			store := filepath.Join(home, ".claude", "accounts")
			writeFile(t, filepath.Join(home, ".claude", ".credentials.json"), `{"t":"ativo"}`)
			writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
			writeFile(t, filepath.Join(store, "current.txt"), current)

			if _, err := (Store{Home: home}).Use("b"); err != nil {
				t.Fatalf("Use: %v", err)
			}

			assertDirEntries(t, store, "b.credentials.json", "backup.credentials.json", "current.txt")
			assertDirEntries(t, filepath.Join(home, ".claude"), ".credentials.json", "accounts")
			assertDirEntries(t, home, ".claude")
		})
	}
}

// Conta sem account.json troca só as credenciais: ~/.claude.json fica byte a byte.
func TestUseWithoutAccountJSONKeepsState(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	const state = "{\"oauthAccount\": {\"emailAddress\": \"a@x.com\"}, \"n\": 1.0}\n"
	writeFile(t, filepath.Join(home, ".claude.json"), state)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)

	if _, err := (Store{Home: home}).Use("b"); err != nil {
		t.Fatalf("Use: %v", err)
	}

	if got := readFile(t, filepath.Join(home, ".claude.json")); got != state {
		t.Errorf(".claude.json = %q, quero sem mudar", got)
	}
	if got := readFile(t, filepath.Join(store, "backup.claude.json")); got != state {
		t.Errorf("backup.claude.json = %q, quero %q", got, state)
	}
}

func TestUseInvalidAccountJSONChangesNothing(t *testing.T) {
	for name, content := range map[string]string{"inválido": `{"emailAddress": `, "null": `null`, "lista": `[]`} {
		t.Run(name, func(t *testing.T) {
			home := tempDir(t)
			store := filepath.Join(home, ".claude", "accounts")
			writeFile(t, filepath.Join(home, ".claude", ".credentials.json"), `{"t":"ativo"}`)
			writeFile(t, filepath.Join(home, ".claude.json"), `{"oauthAccount": {"email": "a@x.com"}}`)
			writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
			writeFile(t, filepath.Join(store, "b.account.json"), content)

			_, err := Store{Home: home}.Use("b")

			want := filepath.Join(store, "b.account.json") + " não é um objeto JSON válido"
			if err == nil || err.Error() != want {
				t.Fatalf("erro = %v, quero %q", err, want)
			}
			if got := readFile(t, filepath.Join(home, ".claude", ".credentials.json")); got != `{"t":"ativo"}` {
				t.Errorf("credenciais = %q, quero sem mudar", got)
			}
			if got := readFile(t, filepath.Join(home, ".claude.json")); got != `{"oauthAccount": {"email": "a@x.com"}}` {
				t.Errorf(".claude.json = %q, quero sem mudar", got)
			}
			assertDirEntries(t, store, "b.account.json", "b.credentials.json")
		})
	}
}

// Uma escrita que falha no meio para o Use com erro e deixa as credenciais ativas
// como estavam (CA-6).
func TestUseWriteFailureKeepsCred(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	cred := filepath.Join(home, ".claude", ".credentials.json")
	writeFile(t, cred, `{"t":"ativo"}`)
	writeFile(t, filepath.Join(home, ".claude.json"), `{}`)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
	if err := os.Mkdir(filepath.Join(store, "backup.claude.json"), 0o755); err != nil {
		t.Fatal(err)
	}

	_, err := Store{Home: home}.Use("b")

	if err == nil || !strings.HasPrefix(err.Error(), "gravar "+filepath.Join(store, "backup.claude.json")+": ") {
		t.Fatalf("erro = %v, quero 'gravar ...backup.claude.json: ...'", err)
	}
	if got := readFile(t, cred); got != `{"t":"ativo"}` {
		t.Errorf("credenciais = %q, quero sem mudar", got)
	}
	assertDirEntries(t, store, "b.credentials.json", "backup.claude.json", "backup.credentials.json")
}

// Se a troca falha depois de começar, current.txt já não aponta a conta antiga: um
// novo Use não pode devolver para ela um login que talvez já seja o da conta nova.
func TestUseFailedSwitchClearsCurrent(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	cred := filepath.Join(home, ".claude", ".credentials.json")
	writeFile(t, cred, `{"t":"a-renovado"}`)
	writeFile(t, filepath.Join(home, ".claude.json"), `{"oauthAccount": {"emailAddress": "a@x.com"}}`)
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a-velho"}`)
	writeFile(t, filepath.Join(store, "a.account.json"), `{"emailAddress": "a@x.com"}`)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
	writeFile(t, filepath.Join(store, "current.txt"), "a")
	// Nega criar arquivo em ~/.claude: a gravação de CRED (temporário + rename) falha,
	// e a leitura, a devolução e o backup (em accounts) continuam funcionando.
	denyCreate(t, filepath.Join(home, ".claude"))

	_, err := Store{Home: home}.Use("b")

	if err == nil || !strings.HasPrefix(err.Error(), "gravar "+cred+": ") {
		t.Fatalf("erro = %v, quero 'gravar %s: ...'", err, cred)
	}
	if got := readFile(t, cred); got != `{"t":"a-renovado"}` {
		t.Errorf("credenciais = %q, quero sem mudar", got)
	}
	if got := readFile(t, filepath.Join(store, "a.credentials.json")); got != `{"t":"a-renovado"}` {
		t.Errorf("a.credentials.json = %q, quero o token renovado", got)
	}
	if _, err := os.Stat(filepath.Join(store, "current.txt")); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("current.txt: %v, quero apagado", err)
	}
	assertDirEntries(t, store, "a.account.json", "a.credentials.json", "b.credentials.json", "backup.claude.json", "backup.credentials.json")
}

func TestRenameErrorTypes(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)

	_, err := Store{Home: home}.Rename("Z", "w")
	var missing *MissingAccountError
	if !errors.As(err, &missing) || missing.Name != "z" {
		t.Errorf("origem inexistente: erro = %v (%T), quero *MissingAccountError{Name: z}", err, err)
	}

	_, err = Store{Home: home}.Rename("a", "B")
	var existing *ExistingAccountError
	if !errors.As(err, &existing) || existing.Name != "b" {
		t.Errorf("destino existente: erro = %v (%T), quero *ExistingAccountError{Name: b}", err, err)
	}
	assertDirEntries(t, store, "a.credentials.json", "b.credentials.json")
}

// current.txt gravado com outra caixa (ex.: pelo claude_switch.py) ainda aponta a
// conta: o rename o atualiza.
func TestRenameCurrentIgnoresCase(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "current.txt"), "A\r\n")

	res, err := Store{Home: home}.Rename("a", "x")

	if err != nil {
		t.Fatalf("Rename: %v", err)
	}
	if res != (RenameResult{From: "a", To: "x"}) {
		t.Errorf("Rename = %+v, quero {From:a To:x}", res)
	}
	if got := readFile(t, filepath.Join(store, "current.txt")); got != "x" {
		t.Errorf("current.txt = %q, quero %q", got, "x")
	}
}

// Um account.json que sobrou sem credenciais no nome novo não pode dar à conta
// renomeada o e-mail de outro login.
func TestRenameRemovesOrphanAccountJSON(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "x.account.json"), `{"emailAddress": "outro@x.com"}`)

	if _, err := (Store{Home: home}).Rename("a", "x"); err != nil {
		t.Fatalf("Rename: %v", err)
	}

	assertDirEntries(t, store, "x.credentials.json")
	if got := readFile(t, filepath.Join(store, "x.credentials.json")); got != `{"t":"a"}` {
		t.Errorf("x.credentials.json = %q, quero os bytes de a", got)
	}
}

// Se o segundo rename falha, o primeiro é desfeito e current.txt não muda.
func TestRenameFailureUndoes(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "a.account.json"), `{"emailAddress": "a@x.com"}`)
	writeFile(t, filepath.Join(store, "current.txt"), "a")
	// Uma pasta não vazia no lugar de x.account.json faz o rename dele falhar.
	writeFile(t, filepath.Join(store, "x.account.json", "dentro.txt"), "ocupado")

	_, err := Store{Home: home}.Rename("a", "x")

	want := "renomear " + filepath.Join(store, "a.account.json") + " para x.account.json: "
	if err == nil || !strings.HasPrefix(err.Error(), want) {
		t.Fatalf("erro = %v, quero %q...", err, want)
	}
	if strings.Contains(err.Error(), "ao desfazer") {
		t.Errorf("erro = %q, o rollback não deveria falhar", err)
	}
	assertDirEntries(t, store, "a.account.json", "a.credentials.json", "current.txt", "x.account.json")
	if got := readFile(t, filepath.Join(store, "a.credentials.json")); got != `{"t":"a"}` {
		t.Errorf("a.credentials.json = %q, quero sem mudar", got)
	}
	if got := readFile(t, filepath.Join(store, "current.txt")); got != "a" {
		t.Errorf("current.txt = %q, quero %q", got, "a")
	}
}

func TestRemoveResult(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "b.credentials.json"), `{"t":"b"}`)
	writeFile(t, filepath.Join(store, "current.txt"), "A\n")

	res, err := Store{Home: home}.Remove("b")
	if err != nil || res != (RemoveResult{Name: "b"}) {
		t.Errorf("Remove(b) = %+v, %v; quero {Name:b WasCurrent:false}", res, err)
	}
	res, err = Store{Home: home}.Remove("A")
	if err != nil || res != (RemoveResult{Name: "a", WasCurrent: true}) {
		t.Errorf("Remove(A) = %+v, %v; quero {Name:a WasCurrent:true}", res, err)
	}
	assertDirEntries(t, store)

	_, err = Store{Home: home}.Remove("a")
	var missing *MissingAccountError
	if !errors.As(err, &missing) || missing.Name != "a" {
		t.Errorf("Remove(a) de novo: erro = %v (%T), quero *MissingAccountError{Name: a}", err, err)
	}
}

// Se o account.json não pode ser apagado, as credenciais e current.txt ficam.
func TestRemoveFailureKeepsAccount(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	writeFile(t, filepath.Join(store, "a.credentials.json"), `{"t":"a"}`)
	writeFile(t, filepath.Join(store, "a.account.json", "dentro.txt"), "ocupado")
	writeFile(t, filepath.Join(store, "current.txt"), "a")

	_, err := Store{Home: home}.Remove("a")

	want := "apagar " + filepath.Join(store, "a.account.json") + ": "
	if err == nil || !strings.HasPrefix(err.Error(), want) {
		t.Fatalf("erro = %v, quero %q...", err, want)
	}
	assertDirEntries(t, store, "a.account.json", "a.credentials.json", "current.txt")
}

// denyCreate nega ao usuário atual criar arquivos em dir até o fim do teste.
func denyCreate(t *testing.T, dir string) {
	t.Helper()
	u, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}
	if out, err := exec.Command("icacls", dir, "/deny", u.Username+":(WD)").CombinedOutput(); err != nil {
		t.Fatalf("icacls /deny: %v\n%s", err, out)
	}
	t.Cleanup(func() {
		if out, err := exec.Command("icacls", dir, "/remove:d", u.Username).CombinedOutput(); err != nil {
			t.Errorf("icacls /remove:d: %v\n%s", err, out)
		}
	})
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

func readFile(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func assertDirEntries(t *testing.T, dir string, want ...string) {
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
