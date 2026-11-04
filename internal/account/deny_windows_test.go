package account

import (
	"os/exec"
	"os/user"
	"testing"
)

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
