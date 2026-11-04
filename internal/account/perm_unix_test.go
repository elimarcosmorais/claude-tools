//go:build !windows

package account

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// No Linux, os arquivos de login gravados pelo claude-tools são só do dono (0600) e a
// pasta de contas que ele cria também (0700), mesmo quando o CRED de antes era legível
// por todos.
func TestSaveAndUseWriteOwnerOnlyPermissions(t *testing.T) {
	home := tempDir(t)
	store := filepath.Join(home, ".claude", "accounts")
	cred := filepath.Join(home, ".claude", ".credentials.json")
	writeFile(t, cred, `{"t":"a"}`)
	writeFile(t, filepath.Join(home, ".claude.json"), `{"oauthAccount": {"emailAddress": "a@x.com"}}`)
	chmod(t, cred, 0o644)

	if _, err := (Store{Home: home}).Save("a"); err != nil {
		t.Fatalf("Save: %v", err)
	}

	assertMode(t, store, fs.ModeDir|0o700)
	assertMode(t, filepath.Join(store, "a.credentials.json"), 0o600)
	assertMode(t, filepath.Join(store, "a.account.json"), 0o600)

	// O CRED que o Claude Code deixou com outra permissão sai do use com 0600.
	chmod(t, cred, 0o644)
	assertMode(t, cred, 0o644)

	if _, err := (Store{Home: home}).Use("a"); err != nil {
		t.Fatalf("Use: %v", err)
	}

	assertMode(t, cred, 0o600)
	assertMode(t, filepath.Join(store, "backup.credentials.json"), 0o600)
	assertMode(t, filepath.Join(store, "a.credentials.json"), 0o600)
}

func chmod(t *testing.T, path string, mode os.FileMode) {
	t.Helper()
	if err := os.Chmod(path, mode); err != nil {
		t.Fatal(err)
	}
}

func assertMode(t *testing.T, path string, want fs.FileMode) {
	t.Helper()
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := info.Mode(); got != want {
		t.Errorf("modo de %s = %v, quero %v", path, got, want)
	}
}
