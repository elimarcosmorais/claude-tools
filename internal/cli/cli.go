// Package cli interpreta a linha de comando do claude-tools: o parse, as
// mensagens e os códigos de saída. Sucesso e avisos vão para o stdout; erro e uso
// incorreto, para o stderr.
package cli

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/elimarcosmorais/claude-tools/internal/account"
)

const usage = `Uso: claude-tools account <comando>

Comandos:
  save <conta>            guarda o login atual do Claude Code como <conta>
  use <conta>             ativa o login guardado em <conta>
  rename <conta> <nova>   renomeia <conta> para <nova>
  remove <conta>          apaga o login guardado em <conta>
  status                  mostra a conta ativa e as contas salvas
`

// Códigos de saída.
const (
	exitOK    = 0
	exitError = 1
	exitUsage = 2
)

// Variáveis que, definidas, passam na frente do login do Claude Code.
var overridingEnv = []string{"ANTHROPIC_API_KEY", "ANTHROPIC_AUTH_TOKEN"}

// Run executa o comando em args (sem o nome do programa) sobre a home dada e
// devolve o código de saída.
func Run(args []string, home string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stdout, usage)
		return exitOK
	}
	switch args[0] {
	case "help":
		return help(args[1:], stdout, stderr)
	case "account":
		return runAccount(args[1:], account.Store{Home: home}, stdout, stderr)
	default:
		return unknownCommand(args[0], stderr)
	}
}

func runAccount(args []string, store account.Store, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		return usageError(stderr)
	}
	switch args[0] {
	case "help":
		return help(args[1:], stdout, stderr)
	case "save":
		if len(args) != 2 {
			return usageError(stderr)
		}
		return save(store, args[1], stdout, stderr)
	case "use":
		if len(args) != 2 {
			return usageError(stderr)
		}
		return use(store, args[1], stdout, stderr)
	case "rename":
		if len(args) != 3 {
			return usageError(stderr)
		}
		return rename(store, args[1], args[2], stdout, stderr)
	case "remove":
		if len(args) != 2 {
			return usageError(stderr)
		}
		return remove(store, args[1], stdout, stderr)
	case "status":
		if len(args) != 1 {
			return usageError(stderr)
		}
		return status(store, stdout, stderr)
	default:
		return unknownCommand(args[0], stderr)
	}
}

func save(store account.Store, name string, stdout, stderr io.Writer) int {
	res, err := store.Save(name)
	if err != nil {
		return report(err, stderr)
	}
	fmt.Fprintf(stdout, "Conta '%s' salva (%s).\n", res.Name, orDefault(res.Email, "email desconhecido"))
	return exitOK
}

func use(store account.Store, name string, stdout, stderr io.Writer) int {
	res, err := store.Use(name)
	var missing *account.MissingAccountError
	if errors.As(err, &missing) {
		fmt.Fprintf(stderr, "%s Faça login nela e rode: claude-tools account save %s\n", missing, missing.Name)
		return exitError
	}
	if err != nil {
		return report(err, stderr)
	}
	warnOverridingEnv(stdout)
	if f := res.Foreign; f != nil {
		fmt.Fprintf(stdout, "Aviso: o login ativo (%s) não é o da conta '%s' (%s). Não vou sobrescrever a conta '%s'. Use 'claude-tools account save' se quiser guardá-lo.\n",
			f.ActiveEmail, f.Current, f.SavedEmail, f.Current)
	}
	fmt.Fprintf(stdout, "Conta ativa: %s (%s)\n", res.Name, orDefault(res.Email, "email desconhecido"))
	fmt.Fprintln(stdout, "Reabra o 'claude' no terminal e use 'Developer: Reload Window' no VS Code.")
	return exitOK
}

func rename(store account.Store, from, to string, stdout, stderr io.Writer) int {
	res, err := store.Rename(from, to)
	if err != nil {
		return report(err, stderr)
	}
	fmt.Fprintf(stdout, "Conta '%s' renomeada para '%s'.\n", res.From, res.To)
	return exitOK
}

func remove(store account.Store, name string, stdout, stderr io.Writer) int {
	res, err := store.Remove(name)
	if err != nil {
		return report(err, stderr)
	}
	fmt.Fprintf(stdout, "Conta '%s' removida.\n", res.Name)
	if res.WasCurrent {
		fmt.Fprintln(stdout, "Nenhuma conta ficou marcada como atual; o login ativo do Claude Code não mudou.")
	}
	return exitOK
}

func status(store account.Store, stdout, stderr io.Writer) int {
	st, err := store.Status()
	if err != nil {
		return report(err, stderr)
	}
	warnOverridingEnv(stdout)
	fmt.Fprintf(stdout, "Conta atual: %s\n", orDefault(st.Current, "(nenhuma)"))
	fmt.Fprintf(stdout, "Email ativo no Claude Code: %s\n", orDefault(st.ActiveEmail, "(não encontrado)"))
	fmt.Fprintln(stdout, "Contas salvas:")
	if len(st.Accounts) == 0 {
		fmt.Fprintln(stdout, "  (nenhuma)")
	}
	for _, a := range st.Accounts {
		mark := ""
		if a.Name == st.Current {
			mark = " <- ativa"
		}
		fmt.Fprintf(stdout, "  %s: %s (salva em %s)%s\n",
			a.Name, orDefault(a.Email, "email desconhecido"), a.SavedAt.Format("02/01/2006 15:04"), mark)
	}
	return exitOK
}

func warnOverridingEnv(stdout io.Writer) {
	for _, name := range overridingEnv {
		if os.Getenv(name) != "" {
			fmt.Fprintf(stdout, "Aviso: %s está definida e passa na frente do login; a troca de conta não terá efeito.\n", name)
		}
	}
}

// report escreve o erro no stderr e devolve o código de saída: nome de conta
// recusado é uso incorreto; o resto é erro de operação, com a mensagem do próprio
// erro quando ela já é a frase para o usuário.
func report(err error, stderr io.Writer) int {
	var nameErr *account.NameError
	if errors.As(err, &nameErr) {
		fmt.Fprintln(stderr, nameErr)
		return exitUsage
	}
	var noLogin *account.NoLoginError
	var missing *account.MissingAccountError
	var existing *account.ExistingAccountError
	if errors.As(err, &noLogin) || errors.As(err, &missing) || errors.As(err, &existing) {
		fmt.Fprintln(stderr, err)
		return exitError
	}
	fmt.Fprintf(stderr, "Erro: %v\n", err)
	return exitError
}

func help(extra []string, stdout, stderr io.Writer) int {
	if len(extra) != 0 {
		return usageError(stderr)
	}
	fmt.Fprint(stdout, usage)
	return exitOK
}

func unknownCommand(name string, stderr io.Writer) int {
	fmt.Fprintf(stderr, "Comando desconhecido: %s\n\n", name)
	return usageError(stderr)
}

func usageError(stderr io.Writer) int {
	fmt.Fprint(stderr, usage)
	return exitUsage
}

func orDefault(s, def string) string {
	if s == "" {
		return def
	}
	return s
}
