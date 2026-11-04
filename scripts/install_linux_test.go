//go:build linux

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

// TestInstallTwice roda o install.sh duas vezes numa raiz temporária (sem mexer no
// ~/.profile) e confere a pasta da versão, o binário executável e o link current.
func TestInstallTwice(t *testing.T) {
	repo := buildBin(t)
	version := readVersion(t, repo)
	root := t.TempDir()
	versionDir := filepath.Join(root, "claude-tools-v"+version)
	current := filepath.Join(root, "current")
	want := "claude-tools " + version + " instalado em " + versionDir + "\n" +
		"Link: " + current + " -> " + versionDir + "\n" +
		"PATH: não mexi (--skip-path).\n"

	for i := 1; i <= 2; i++ {
		code, stdout, stderr := runInstall(t, repo, t.TempDir(), "--root", root, "--skip-path")
		if code != 0 {
			t.Fatalf("execução %d: código = %d, quero 0\nstdout: %s\nstderr: %s", i, code, stdout, stderr)
		}
		if stdout != want {
			t.Errorf("execução %d: stdout = %q, quero %q", i, stdout, want)
		}
	}

	info, err := os.Lstat(filepath.Join(versionDir, "claude-tools"))
	if err != nil {
		t.Fatalf("binário na pasta da versão: %v", err)
	}
	if !info.Mode().IsRegular() || info.Mode().Perm() != 0o755 {
		t.Errorf("binário: modo = %v, quero um arquivo comum com 0755", info.Mode())
	}
	entries, err := os.ReadDir(versionDir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Errorf("pasta da versão tem %d itens, quero só o claude-tools (sobrou temporário?)", len(entries))
	}
	target, err := os.Readlink(current)
	if err != nil {
		t.Fatalf("current não é um link simbólico: %v", err)
	}
	if target != versionDir {
		t.Errorf("current aponta para %s, quero %s", target, versionDir)
	}

	cmd := exec.Command(filepath.Join(current, "claude-tools"), "help")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		t.Fatalf("current/claude-tools help: %v", err)
	}
	if !strings.HasPrefix(out.String(), "Uso: claude-tools account <comando>\n") {
		t.Errorf("saída do help = %q, quero o uso", out.String())
	}
}

// TestInstallAddsPathOnce instala duas vezes sem --skip-path, com HOME numa pasta
// temporária, e confere que o ~/.profile ganha a linha do PATH uma vez só, sem grudar
// na última linha que já existia.
func TestInstallAddsPathOnce(t *testing.T) {
	repo := buildBin(t)
	home := t.TempDir()
	root := t.TempDir()
	current := filepath.Join(root, "current")
	profile := filepath.Join(home, ".profile")
	// Conteúdo de antes, sem o \n final.
	if err := os.WriteFile(profile, []byte("umask 022"), 0o644); err != nil {
		t.Fatal(err)
	}

	wantLast := []string{
		"PATH: acrescentei " + current + " ao ~/.profile. Abra um terminal novo para usar o 'claude-tools'.",
		"PATH: " + current + " já estava no ~/.profile.",
	}
	for i, want := range wantLast {
		code, stdout, stderr := runInstall(t, repo, home, "--root", root)
		if code != 0 {
			t.Fatalf("execução %d: código = %d, quero 0\nstdout: %s\nstderr: %s", i+1, code, stdout, stderr)
		}
		lines := strings.Split(strings.TrimSuffix(stdout, "\n"), "\n")
		if got := lines[len(lines)-1]; got != want {
			t.Errorf("execução %d: última linha = %q, quero %q", i+1, got, want)
		}
	}

	b, err := os.ReadFile(profile)
	if err != nil {
		t.Fatal(err)
	}
	wantProfile := "umask 022\n# claude-tools\nexport PATH=\"" + current + ":$PATH\"\n"
	if string(b) != wantProfile {
		t.Errorf(".profile = %q, quero %q", b, wantProfile)
	}
}

// TestInstallRefusesRealCurrentDir confere que uma pasta comum no lugar do link
// current não é apagada.
func TestInstallRefusesRealCurrentDir(t *testing.T) {
	repo := buildBin(t)
	root := t.TempDir()
	current := filepath.Join(root, "current")
	keep := filepath.Join(current, "guardar.txt")
	if err := os.MkdirAll(current, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(keep, []byte("dado"), 0o600); err != nil {
		t.Fatal(err)
	}

	code, stdout, stderr := runInstall(t, repo, t.TempDir(), "--root", root, "--skip-path")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	wantErr := "Erro: " + current + " existe e não é um link simbólico; não vou apagá-lo.\n"
	if stderr != wantErr || stdout != "" {
		t.Errorf("stdout = %q, stderr = %q; quero stdout vazio e stderr %q", stdout, stderr, wantErr)
	}
	if b, err := os.ReadFile(keep); err != nil || string(b) != "dado" {
		t.Errorf("a pasta current não ficou intacta: %q, %v", b, err)
	}
	if _, err := os.Stat(filepath.Join(root, "claude-tools-v"+readVersion(t, repo))); !errors.Is(err, os.ErrNotExist) {
		t.Errorf("pasta da versão criada apesar do erro: %v", err)
	}
}

// TestInstallRequiresRoot confere que, sem --root, o install.sh falha em vez de
// instalar numa pasta padrão.
func TestInstallRequiresRoot(t *testing.T) {
	repo, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}

	code, stdout, stderr := runInstall(t, repo, t.TempDir(), "--skip-path")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	wantErr := "Erro: informe a pasta de instalação: --root <pasta> (pelo Task: go tool task install DIR=<pasta>).\n"
	if stderr != wantErr || stdout != "" {
		t.Errorf("stdout = %q, stderr = %q; quero stdout vazio e stderr %q", stdout, stderr, wantErr)
	}
}

// TestInstallRequiresBuild roda uma cópia do install.sh num repositório temporário sem
// bin/: ele acha o repositório pela pasta do próprio script e pede o build.
func TestInstallRequiresBuild(t *testing.T) {
	repo, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	fake := t.TempDir()
	script, err := os.ReadFile(filepath.Join(repo, "scripts", "install.sh"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(fake, "scripts"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(fake, "scripts", "install.sh"), script, 0o644); err != nil {
		t.Fatal(err)
	}
	// Em CRLF, como no checkout do Windows: o \r não pode entrar na versão.
	if err := os.WriteFile(filepath.Join(fake, "VERSION"), []byte("9.8.7\r\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()

	code, stdout, stderr := runInstall(t, fake, t.TempDir(), "--root", root, "--skip-path")

	if code != 1 {
		t.Errorf("código = %d, quero 1", code)
	}
	wantErr := "Erro: não achei " + filepath.Join(fake, "bin", "claude-tools-9.8.7") + ". Rode 'go tool task build' antes.\n"
	if stderr != wantErr || stdout != "" {
		t.Errorf("stdout = %q, stderr = %q; quero stdout vazio e stderr %q", stdout, stderr, wantErr)
	}
}

// buildBin compila bin/claude-tools-<VERSION>, que é de onde o install.sh copia, e
// devolve a raiz do repositório.
func buildBin(t *testing.T) string {
	t.Helper()
	repo, err := filepath.Abs("..")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "build", "-o", filepath.Join("bin", "claude-tools-"+readVersion(t, repo)), "./cmd/claude-tools")
	cmd.Dir = repo
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("go build: %v\n%s", err, out)
	}
	return repo
}

// runInstall roda o scripts/install.sh de repo com HOME=home, para nunca tocar no
// ~/.profile de quem roda os testes, e devolve o código de saída, o stdout e o stderr.
func runInstall(t *testing.T, repo, home string, args ...string) (int, string, string) {
	t.Helper()
	cmd := exec.Command("sh", append([]string{filepath.Join(repo, "scripts", "install.sh")}, args...)...)
	cmd.Env = append(os.Environ(), "HOME="+home)
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
		t.Fatalf("sh: %v", err)
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
