// Package account guarda e troca os logins do Claude Code.
//
// O armazenamento é o mesmo do antigo claude_switch.py: as contas ficam em
// ~/.claude/accounts como <conta>.credentials.json (cópia de
// ~/.claude/.credentials.json) e <conta>.account.json (o bloco oauthAccount de
// ~/.claude.json), e current.txt guarda o nome da conta atual.
package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	credSuffix    = ".credentials.json"
	accountSuffix = ".account.json"
	reservedName  = "backup"
)

var validName = regexp.MustCompile(`^[A-Za-z0-9_-]{1,32}$`)

// Store opera sobre os arquivos do Claude Code e as contas salvas dentro de Home.
type Store struct {
	Home string
}

// Saved é uma conta salva em ~/.claude/accounts.
type Saved struct {
	Name    string
	Email   string // vazio quando não há e-mail conhecido
	SavedAt time.Time
}

// Status descreve a conta atual, o login ativo e as contas salvas.
type Status struct {
	Current     string // vazio quando nenhuma conta está marcada como atual
	ActiveEmail string // vazio quando o login ativo não tem e-mail conhecido
	Accounts    []Saved
}

// NameError é um nome de conta recusado pela regra de nomes.
type NameError struct {
	Name     string
	Reserved bool
}

func (e *NameError) Error() string {
	if e.Reserved {
		return fmt.Sprintf("Nome de conta reservado: '%s'.", e.Name)
	}
	return fmt.Sprintf("Nome de conta inválido: '%s'. Use letras, números, '-' ou '_' (até 32 caracteres).", e.Name)
}

// NormalizeName valida o nome de uma conta e o devolve em minúsculas.
func NormalizeName(name string) (string, error) {
	if !validName.MatchString(name) {
		return "", &NameError{Name: name}
	}
	name = strings.ToLower(name)
	if name == reservedName {
		return "", &NameError{Name: name, Reserved: true}
	}
	return name, nil
}

// NoLoginError é o save sem o arquivo de login do Claude Code.
type NoLoginError struct {
	Path string
}

func (e *NoLoginError) Error() string {
	return fmt.Sprintf("Não achei %s. Faça login no Claude Code antes de salvar.", e.Path)
}

// MissingAccountError é uma operação sobre uma conta que não foi salva.
type MissingAccountError struct {
	Name string
}

func (e *MissingAccountError) Error() string {
	return fmt.Sprintf("Conta '%s' não existe.", e.Name)
}

// ExistingAccountError é um rename para um nome que já tem conta salva.
type ExistingAccountError struct {
	Name string
}

func (e *ExistingAccountError) Error() string {
	return fmt.Sprintf("Já existe uma conta '%s'.", e.Name)
}

// RenameResult são os nomes, já normalizados, de um Rename feito.
type RenameResult struct {
	From, To string
}

// RemoveResult é a conta que o Remove apagou.
type RemoveResult struct {
	Name string
	// WasCurrent diz se a conta era a de current.txt, que então foi apagado.
	WasCurrent bool
}

// SaveResult é a conta que o Save gravou.
type SaveResult struct {
	Name  string
	Email string // vazio quando o login ativo não tem e-mail conhecido
}

// UseResult é a conta que o Use ativou.
type UseResult struct {
	Name  string
	Email string // vazio quando a conta não tem e-mail conhecido
	// Foreign não é nil quando o login ativo era de outra pessoa e por isso não
	// foi devolvido para a conta atual.
	Foreign *ForeignLogin
}

// ForeignLogin é um login ativo cujo e-mail não é o salvo da conta atual.
type ForeignLogin struct {
	Current     string // a conta de current.txt, que ficou como estava
	ActiveEmail string
	SavedEmail  string
}

func (s Store) credPath() string    { return filepath.Join(s.Home, ".claude", ".credentials.json") }
func (s Store) statePath() string   { return filepath.Join(s.Home, ".claude.json") }
func (s Store) storeDir() string    { return filepath.Join(s.Home, ".claude", "accounts") }
func (s Store) currentPath() string { return filepath.Join(s.storeDir(), "current.txt") }

func (s Store) credentialsPath(name string) string {
	return filepath.Join(s.storeDir(), name+credSuffix)
}

func (s Store) accountPath(name string) string {
	return filepath.Join(s.storeDir(), name+accountSuffix)
}

// Save guarda o login ativo do Claude Code como a conta name: copia o arquivo de
// credenciais, grava o bloco oauthAccount de ~/.claude.json (ou apaga o de um
// save anterior, quando o bloco não existe) e marca a conta como atual.
func (s Store) Save(name string) (SaveResult, error) {
	name, err := NormalizeName(name)
	if err != nil {
		return SaveResult{}, err
	}
	cred, ok, err := readOptional(s.credPath())
	if err != nil {
		return SaveResult{}, err
	}
	if !ok {
		return SaveResult{}, &NoLoginError{Path: s.credPath()}
	}
	state, err := s.readState()
	if err != nil {
		return SaveResult{}, err
	}

	if err := os.MkdirAll(s.storeDir(), 0o700); err != nil {
		return SaveResult{}, fmt.Errorf("criar %s: %w", s.storeDir(), err)
	}
	if err := writeFileAtomic(s.credentialsPath(name), cred); err != nil {
		return SaveResult{}, err
	}
	acc := state["oauthAccount"]
	if isNonEmptyObject(acc) {
		if err := writeFileAtomic(s.accountPath(name), acc); err != nil {
			return SaveResult{}, err
		}
	} else if err := os.Remove(s.accountPath(name)); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return SaveResult{}, fmt.Errorf("apagar %s: %w", s.accountPath(name), err)
	}
	if err := writeFileAtomic(s.currentPath(), []byte(name)); err != nil {
		return SaveResult{}, err
	}
	return SaveResult{Name: name, Email: emailOf(acc)}, nil
}

// Use ativa a conta name. Antes devolve o login ativo (com o token renovado) para
// a conta atual, se ele for dela, e guarda cópias de ~/.claude/.credentials.json
// e ~/.claude.json em backup.*. Depois copia as credenciais da conta para o
// Claude Code, troca só o oauthAccount de ~/.claude.json e marca a conta como
// atual. Tudo é lido e validado antes da primeira escrita: conta inexistente ou
// JSON inválido não mudam nenhum arquivo.
func (s Store) Use(name string) (UseResult, error) {
	name, err := NormalizeName(name)
	if err != nil {
		return UseResult{}, err
	}
	target, ok, err := readOptional(s.credentialsPath(name))
	if err != nil {
		return UseResult{}, err
	}
	if !ok {
		return UseResult{}, &MissingAccountError{Name: name}
	}
	targetAcc, hasTargetAcc, err := readOptional(s.accountPath(name))
	if err != nil {
		return UseResult{}, err
	}
	if hasTargetAcc && !isObject(targetAcc) {
		return UseResult{}, fmt.Errorf("%s não é um objeto JSON válido", s.accountPath(name))
	}
	current, err := s.current()
	if err != nil {
		return UseResult{}, err
	}
	cred, hasCred, err := readOptional(s.credPath())
	if err != nil {
		return UseResult{}, err
	}
	stateBytes, hasState, err := readOptional(s.statePath())
	if err != nil {
		return UseResult{}, err
	}
	var newState []byte
	if hasTargetAcc && hasState {
		if newState, err = replaceOAuthAccount(s.statePath(), stateBytes, targetAcc); err != nil {
			return UseResult{}, err
		}
	}

	res := UseResult{Name: name, Email: emailOf(targetAcc)}
	if cur := s.savedAccount(current); cur != "" && hasCred {
		saved := emailOf(readFileOrNil(s.accountPath(cur)))
		active := emailOf(objectOrNil(stateBytes)["oauthAccount"])
		if saved != "" && active != "" && saved != active {
			res.Foreign = &ForeignLogin{Current: cur, ActiveEmail: active, SavedEmail: saved}
		} else {
			if err := writeFileAtomic(s.credentialsPath(cur), cred); err != nil {
				return UseResult{}, err
			}
			if cur == name {
				target = cred
			}
		}
	}

	if hasCred {
		if err := writeFileAtomic(filepath.Join(s.storeDir(), reservedName+credSuffix), cred); err != nil {
			return UseResult{}, err
		}
	}
	if hasState {
		if err := writeFileAtomic(filepath.Join(s.storeDir(), reservedName+".claude.json"), stateBytes); err != nil {
			return UseResult{}, err
		}
	}

	// Daqui em diante CRED e STATE passam a ser da conta nova. Se uma escrita falhar
	// no meio, current.txt não pode continuar apontando a conta antiga: o próximo Use
	// devolveria o login da conta nova para ela. Sem current.txt, ele não devolve nada.
	if err := os.Remove(s.currentPath()); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return UseResult{}, fmt.Errorf("apagar %s: %w", s.currentPath(), err)
	}
	if err := os.MkdirAll(filepath.Dir(s.credPath()), 0o700); err != nil {
		return UseResult{}, fmt.Errorf("criar %s: %w", filepath.Dir(s.credPath()), err)
	}
	if err := writeFileAtomic(s.credPath(), target); err != nil {
		return UseResult{}, err
	}
	if newState != nil {
		if err := writeFileAtomic(s.statePath(), newState); err != nil {
			return UseResult{}, err
		}
	}
	if err := writeFileAtomic(s.currentPath(), []byte(name)); err != nil {
		return UseResult{}, err
	}
	return res, nil
}

// Rename renomeia a conta from para to: o arquivo de credenciais, o account.json
// (se existir) e, quando from é a conta atual, current.txt. Um account.json de to
// que sobrou sem credenciais é apagado quando from não tem o seu, para to não
// herdar o e-mail de outro login. Se um passo falha, os anteriores são desfeitos.
func (s Store) Rename(from, to string) (RenameResult, error) {
	from, err := NormalizeName(from)
	if err != nil {
		return RenameResult{}, err
	}
	to, err = NormalizeName(to)
	if err != nil {
		return RenameResult{}, err
	}
	if err := s.requireSaved(from); err != nil {
		return RenameResult{}, err
	}
	taken, err := exists(s.credentialsPath(to))
	if err != nil {
		return RenameResult{}, err
	}
	if taken {
		return RenameResult{}, &ExistingAccountError{Name: to}
	}
	hasAcc, err := exists(s.accountPath(from))
	if err != nil {
		return RenameResult{}, err
	}
	current, err := s.current()
	if err != nil {
		return RenameResult{}, err
	}

	var undo []func() error
	fail := func(err error) (RenameResult, error) {
		for i := len(undo) - 1; i >= 0; i-- {
			if uerr := undo[i](); uerr != nil {
				err = fmt.Errorf("%w; ao desfazer: %v", err, uerr)
			}
		}
		return RenameResult{}, err
	}

	if err := rename(s.credentialsPath(from), s.credentialsPath(to)); err != nil {
		return fail(err)
	}
	undo = append(undo, func() error { return rename(s.credentialsPath(to), s.credentialsPath(from)) })
	if hasAcc {
		if err := rename(s.accountPath(from), s.accountPath(to)); err != nil {
			return fail(err)
		}
		undo = append(undo, func() error { return rename(s.accountPath(to), s.accountPath(from)) })
	} else if err := removeOptional(s.accountPath(to)); err != nil {
		return fail(err)
	}
	if sameName(current, from) {
		if err := writeFileAtomic(s.currentPath(), []byte(to)); err != nil {
			return fail(err)
		}
	}
	return RenameResult{From: from, To: to}, nil
}

// Remove apaga a conta name (o account.json e o arquivo de credenciais) e, se ela
// era a atual, current.txt. Não toca no login ativo do Claude Code. O account.json
// sai primeiro: se o arquivo de credenciais não puder ser apagado, a conta continua
// listada e basta rodar de novo.
func (s Store) Remove(name string) (RemoveResult, error) {
	name, err := NormalizeName(name)
	if err != nil {
		return RemoveResult{}, err
	}
	if err := s.requireSaved(name); err != nil {
		return RemoveResult{}, err
	}
	current, err := s.current()
	if err != nil {
		return RemoveResult{}, err
	}
	if err := removeOptional(s.accountPath(name)); err != nil {
		return RemoveResult{}, err
	}
	if err := removeOptional(s.credentialsPath(name)); err != nil {
		return RemoveResult{}, err
	}
	res := RemoveResult{Name: name, WasCurrent: sameName(current, name)}
	if res.WasCurrent {
		if err := removeOptional(s.currentPath()); err != nil {
			return RemoveResult{}, err
		}
	}
	return res, nil
}

// requireSaved devolve *MissingAccountError quando name não tem o arquivo de
// credenciais salvo.
func (s Store) requireSaved(name string) error {
	info, err := os.Stat(s.credentialsPath(name))
	if errors.Is(err, fs.ErrNotExist) || (err == nil && !info.Mode().IsRegular()) {
		return &MissingAccountError{Name: name}
	}
	if err != nil {
		return fmt.Errorf("ler a conta '%s': %w", name, err)
	}
	return nil
}

// sameName diz se o conteúdo de current.txt aponta a conta name, sem diferenciar
// maiúsculas, como os nomes de conta.
func sameName(current, name string) bool {
	cur, err := NormalizeName(current)
	return err == nil && cur == name
}

// savedAccount devolve o nome normalizado da conta gravada em current.txt quando
// ela é um nome válido com credenciais salvas, ou "" quando não é.
func (s Store) savedAccount(current string) string {
	name, err := NormalizeName(current)
	if err != nil {
		return ""
	}
	if info, err := os.Stat(s.credentialsPath(name)); err != nil || !info.Mode().IsRegular() {
		return ""
	}
	return name
}

// readState lê ~/.claude.json de forma tolerante, como o Status: ausente ou
// inválido vira um objeto vazio. Só um erro de leitura do arquivo que existe é
// devolvido.
func (s Store) readState() (map[string]json.RawMessage, error) {
	b, _, err := readOptional(s.statePath())
	if err != nil {
		return nil, err
	}
	return objectOrNil(b), nil
}

// replaceOAuthAccount devolve o JSON de state com o campo oauthAccount trocado por
// acc. Os demais valores saem com o mesmo texto (json.RawMessage): números não
// passam por float64 e <, > e & não viram <; a indentação é de 2 espaços.
func replaceOAuthAccount(path string, state, acc []byte) ([]byte, error) {
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(state, &obj); err != nil || obj == nil {
		return nil, fmt.Errorf("%s não é um objeto JSON válido; nada foi alterado", path)
	}
	obj["oauthAccount"] = json.RawMessage(acc)
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	if err := enc.Encode(obj); err != nil {
		return nil, fmt.Errorf("montar %s: %w", path, err)
	}
	return buf.Bytes(), nil
}

// Status lê a conta atual, o e-mail do login ativo e as contas salvas, em ordem
// alfabética e sem o backup. Pasta de contas inexistente é uma lista vazia.
func (s Store) Status() (Status, error) {
	current, err := s.current()
	if err != nil {
		return Status{}, err
	}
	st := Status{
		Current:     current,
		ActiveEmail: emailOf(readObject(s.statePath())["oauthAccount"]),
	}

	entries, err := os.ReadDir(s.storeDir())
	if errors.Is(err, fs.ErrNotExist) {
		return st, nil
	}
	if err != nil {
		return Status{}, fmt.Errorf("ler as contas salvas: %w", err)
	}
	for _, e := range entries {
		name, ok := strings.CutSuffix(e.Name(), credSuffix)
		if !ok || name == "" || e.IsDir() || strings.EqualFold(name, reservedName) {
			continue
		}
		info, err := e.Info()
		if err != nil {
			return Status{}, fmt.Errorf("ler a conta '%s': %w", name, err)
		}
		st.Accounts = append(st.Accounts, Saved{
			Name:    name,
			Email:   emailOf(readFileOrNil(s.accountPath(name))),
			SavedAt: info.ModTime(),
		})
	}
	sort.Slice(st.Accounts, func(i, j int) bool { return st.Accounts[i].Name < st.Accounts[j].Name })
	return st, nil
}

// current devolve o nome gravado em current.txt, ou "" quando ele não existe.
func (s Store) current() (string, error) {
	b, err := os.ReadFile(s.currentPath())
	if errors.Is(err, fs.ErrNotExist) {
		return "", nil
	}
	if err != nil {
		return "", fmt.Errorf("ler a conta atual: %w", err)
	}
	return strings.TrimSpace(string(b)), nil
}

// readObject lê um objeto JSON; arquivo ausente ou inválido vira um objeto vazio,
// como no claude_switch.py.
func readObject(path string) map[string]json.RawMessage {
	return objectOrNil(readFileOrNil(path))
}

func objectOrNil(b []byte) map[string]json.RawMessage {
	var obj map[string]json.RawMessage
	if json.Unmarshal(b, &obj) != nil {
		return nil
	}
	return obj
}

// isObject diz se b é um objeto JSON (não null, não outro tipo).
func isObject(b []byte) bool {
	return objectOrNil(b) != nil
}

// isNonEmptyObject diz se b é um objeto JSON com ao menos um campo, a mesma regra
// do claude_switch.py para gravar o oauthAccount.
func isNonEmptyObject(b []byte) bool {
	return len(objectOrNil(b)) > 0
}

// readOptional lê path e diz se ele existe. Arquivo ausente não é erro; qualquer
// outra falha de leitura é.
func readOptional(path string) ([]byte, bool, error) {
	b, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("ler %s: %w", path, err)
	}
	return b, true, nil
}

// exists diz se há algo (arquivo ou pasta) em path.
func exists(path string) (bool, error) {
	_, err := os.Lstat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("ler %s: %w", path, err)
	}
	return true, nil
}

func rename(from, to string) error {
	if err := os.Rename(from, to); err != nil {
		return fmt.Errorf("renomear %s para %s: %w", from, filepath.Base(to), err)
	}
	return nil
}

// removeOptional apaga path; arquivo ausente não é erro.
func removeOptional(path string) error {
	if err := os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("apagar %s: %w", path, err)
	}
	return nil
}

func readFileOrNil(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	return b
}

// emailOf devolve o e-mail de um bloco oauthAccount: emailAddress, ou email quando
// o primeiro falta ou está vazio. Bloco ausente ou inválido devolve "".
func emailOf(raw json.RawMessage) string {
	var acc map[string]any
	if json.Unmarshal(raw, &acc) != nil {
		return ""
	}
	for _, key := range []string{"emailAddress", "email"} {
		if v, ok := acc[key].(string); ok && v != "" {
			return v
		}
	}
	return ""
}

// writeFileAtomic grava data num temporário na pasta de path e o renomeia por cima
// de path. Se algo falha, path fica como estava e o temporário é apagado.
func writeFileAtomic(path string, data []byte) (err error) {
	tmp, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".*.tmp")
	if err != nil {
		return fmt.Errorf("gravar %s: %w", path, err)
	}
	tmpName := tmp.Name()
	defer func() {
		if err != nil {
			tmp.Close()
			os.Remove(tmpName)
		}
	}()
	if _, err = tmp.Write(data); err != nil {
		return fmt.Errorf("gravar %s: %w", path, err)
	}
	if err = tmp.Sync(); err != nil {
		return fmt.Errorf("gravar %s: %w", path, err)
	}
	if err = tmp.Close(); err != nil {
		return fmt.Errorf("gravar %s: %w", path, err)
	}
	if err = os.Rename(tmpName, path); err != nil {
		return fmt.Errorf("gravar %s: %w", path, err)
	}
	return nil
}
