// Package scripts guarda os testes dos scripts do repositório.
package scripts

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestInstallTwice roda o install.ps1 duas vezes numa raiz temporária (sem mexer
// no PATH) e confere a pasta da versão, a junction current e o executável.
func TestInstallTwice(t *testing.T) {
	repo := buildExe(t)
	root := t.TempDir()

	for i := 1; i <= 2; i++ {
		code, stdout, stderr := runInstall(t, repo, root)
		if code != 0 {
			t.Fatalf("execução %d: código = %d, quero 0\nstdout: %s\nstderr: %s", i, code, stdout, stderr)
		}
	}

	versionDir := filepath.Join(root, "claude-tools-v"+readVersion(t, repo))
	if info, err := os.Stat(filepath.Join(versionDir, "claude-tools.exe")); err != nil || !info.Mode().IsRegular() {
		t.Fatalf("executável na pasta da versão: %v", err)
	}
	current := filepath.Join(root, "current")
	target, err := os.Readlink(current)
	if err != nil {
		t.Fatalf("current não é um link: %v", err)
	}
	// O temporário pode vir com nome curto (ELIMAR~1) e o alvo, com o longo:
	// compara o arquivo, não o texto.
	targetInfo, err := os.Stat(target)
	if err != nil {
		t.Fatalf("alvo de current: %v", err)
	}
	versionInfo, err := os.Stat(versionDir)
	if err != nil {
		t.Fatalf("pasta da versão: %v", err)
	}
	if !os.SameFile(targetInfo, versionInfo) {
		t.Errorf("current aponta para %s, quero %s", target, versionDir)
	}

	cmd := exec.Command(filepath.Join(current, "claude-tools.exe"), "help")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("current\\claude-tools.exe help: %v", err)
	}
	if !strings.HasPrefix(out.String(), "Uso: claude-tools account <comando>\n") {
		t.Errorf("saída do help = %q, quero o uso", out.String())
	}
}

// TestInstallRefusesRealCurrentDir confere que uma pasta comum no lugar da
// junction current não é apagada.
func TestInstallRefusesRealCurrentDir(t *testing.T) {
	repo := buildExe(t)
	root := t.TempDir()
	current := filepath.Join(root, "current")
	keep := filepath.Join(current, "guardar.txt")
	if err := os.MkdirAll(current, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keep, []byte("dado"), 0o600); err != nil {
		t.Fatal(err)
	}

	code, _, stderr := runInstall(t, repo, root)

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	// O pwsh redirecionado escreve na página de código do console (não em UTF-8), e
	// o caminho pode sair com o nome longo da pasta temporária: a asserção usa só os
	// trechos sem acento da mensagem.
	if !strings.HasPrefix(stderr, "Erro: ") || !strings.Contains(stderr, `\current existe e n`) ||
		!strings.Contains(stderr, " uma junction;") {
		t.Errorf("stderr = %q, quero o erro de junction", stderr)
	}
	if b, err := os.ReadFile(keep); err != nil || string(b) != "dado" {
		t.Errorf("a pasta current não ficou intacta: %q, %v", b, err)
	}
}

// TestInstallRequiresRoot confere que, sem -Root, o install.ps1 falha em vez de
// instalar numa pasta padrão.
func TestInstallRequiresRoot(t *testing.T) {
	repo, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("pwsh", "-NoProfile", "-File", filepath.Join(repo, "scripts", "install.ps1"), "-SkipPath")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	var exitErr *exec.ExitError
	if !errors.As(err, &exitErr) || exitErr.ExitCode() != 1 {
		t.Fatalf("erro = %v, quero saída 1", err)
	}
	// Só os trechos sem acento: o pwsh redirecionado não escreve em UTF-8.
	if !strings.HasPrefix(stderr.String(), "Erro: informe a pasta de instala") ||
		!strings.Contains(stderr.String(), "task install DIR=<pasta>") {
		t.Errorf("stderr = %q, quero o pedido da pasta", stderr.String())
	}
}

// buildExe compila bin/claude-tools-<VERSION>.exe, que é de onde o install.ps1
// copia, e devolve a raiz do repositório.
func buildExe(t *testing.T) string {
	t.Helper()
	repo, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "build", "-o", filepath.Join("bin", "claude-tools-"+readVersion(t, repo)+".exe"), "./cmd/claude-tools")
	cmd.Dir = repo
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return repo
}

func runInstall(t *testing.T, repo, root string) (int, string, string) {
	t.Helper()
	cmd := exec.Command("pwsh", "-NoProfile", "-File", filepath.Join(repo, "scripts", "install.ps1"),
		"-Root", root, "-SkipPath")
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err := cmd.Run()
	var exitErr *exec.ExitError
	switch {
	case err == nil:
		return 0, stdout.String(), stderr.String()
	case errors.As(err, &exitErr):
		return exitErr.ExitCode(), stdout.String(), stderr.String()
	default:
		t.Fatalf("pwsh: %v", err)
		return 0, "", ""
	}
}

func readVersion(t *testing.T, repo string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(repo, "VERSION"))
	if err != nil {
		t.Fatal(err)
	}
	return strings.TrimSpace(string(b))
}
