//go:build !windows

package account

import (
	"os"
	"testing"
)

// denyCreate nega ao usuário atual criar arquivos em dir até o fim do teste: tira a
// permissão de escrita da pasta (0o500) e a devolve (0o700) na limpeza. O root ignora
// a permissão da pasta, e aí o teste não provaria nada: ele reprova em vez de pular.
func denyCreate(t *testing.T, dir string) {
	t.Helper()
	if err := os.Chmod(dir, 0o500); err != nil {
		t.Fatalf("chmod 0500 %s: %v", dir, err)
	}
	t.Cleanup(func() {
		if err := os.Chmod(dir, 0o700); err != nil {
			t.Errorf("chmod 0700 %s: %v", dir, err)
		}
	})
	f, err := os.CreateTemp(dir, "deny-*.tmp")
	if err == nil {
		f.Close()
		os.Remove(f.Name())
		t.Fatalf("criar arquivo em %s ainda funciona depois do chmod 0500: os testes não podem rodar como root", dir)
	}
}
