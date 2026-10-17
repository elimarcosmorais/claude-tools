// Comando claude-tools: troca a conta do Claude Code. Ver `claude-tools help`.
package main

import (
	"fmt"
	"os"

	"github.com/elimarcosmorais/claude-tools-win/internal/cli"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro: não achei a pasta do usuário: %v\n", err)
		os.Exit(1)
	}
	os.Exit(cli.Run(os.Args[1:], home, os.Stdout, os.Stderr))
}
