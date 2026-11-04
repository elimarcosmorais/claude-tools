#!/bin/sh
# Instala o claude-tools (Linux) na pasta de programas e o põe no PATH pelo ~/.profile.
#
# Copia bin/claude-tools-<VERSION> (gerado por 'go tool task build') para
# <root>/claude-tools-v<VERSION>/claude-tools, recria o link simbólico <root>/current
# apontando para essa pasta e acrescenta ao ~/.profile, se ainda não estiver lá, a linha
# export PATH="<root>/current:$PATH". Rodar de novo dá o mesmo resultado.
#
# Uso: sh scripts/install.sh --root <pasta> [--skip-path]
#   --root <pasta>  pasta geral do programa (ex.: $HOME/programs/claude-tools).
#                   Obrigatório: não há pasta padrão.
#   --skip-path     não mexe no ~/.profile.

set -eu

fail() {
	printf 'Erro: %s\n' "$1" >&2
	exit 1
}

root=''
skip_path=0
while [ $# -gt 0 ]; do
	case $1 in
	--root)
		[ $# -ge 2 ] || break
		root=$2
		shift 2
		;;
	--skip-path)
		skip_path=1
		shift
		;;
	*)
		fail "argumento desconhecido: $1 (uso: --root <pasta> [--skip-path])."
		;;
	esac
done
[ -n "$root" ] || fail 'informe a pasta de instalação: --root <pasta> (pelo Task: go tool task install DIR=<pasta>).'

# O repositório é a pasta acima da deste script, de onde quer que ele seja chamado.
script_dir=$(CDPATH='' cd -- "$(dirname -- "$0")" && pwd)
repo=$(dirname -- "$script_dir")

# tr tira o \r do checkout do Windows (CRLF); a substituição de comando, o \n final.
version=$(tr -d '\r' <"$repo/VERSION") || fail "não consegui ler $repo/VERSION."
case $version in
'' | *[!0-9.]*) fail "VERSION inválida: '$version'." ;;
esac
expr "$version" : '[0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*$' >/dev/null || fail "VERSION inválida: '$version'."

bin="$repo/bin/claude-tools-$version"
[ -f "$bin" ] || fail "não achei $bin. Rode 'go tool task build' antes."
if [ "$skip_path" -eq 0 ] && [ -z "${HOME:-}" ]; then
	fail 'a variável HOME não está definida; não sei onde fica o ~/.profile.'
fi

mkdir -p -- "$root" || fail "não consegui criar $root."
root=$(CDPATH='' cd -- "$root" && pwd) || fail "não consegui abrir $root."
version_dir="$root/claude-tools-v$version"
current="$root/current"
profile_line="export PATH=\"$current:\$PATH\""

# Confere tudo antes de gravar qualquer coisa.
if [ -e "$current" ] && [ ! -L "$current" ]; then
	fail "$current existe e não é um link simbólico; não vou apagá-lo."
fi
if [ "$skip_path" -eq 0 ]; then
	# Dentro das aspas duplas da linha do ~/.profile, estes caracteres mudariam o caminho.
	case $current in
	*[\"\$\`\\]*) fail "a pasta $root tem \", \$, \` ou \\, que não cabem na linha do ~/.profile; use outra pasta ou --skip-path." ;;
	esac
fi

mkdir -p -- "$version_dir" || fail "não consegui criar $version_dir."
# Cópia num temporário e rename: gravar por cima de um binário em execução falharia
# (Text file busy), e o rename troca o arquivo de uma vez.
tmp="$version_dir/.claude-tools.$$"
if ! cp -- "$bin" "$tmp" || ! chmod 755 -- "$tmp" || ! mv -f -- "$tmp" "$version_dir/claude-tools"; then
	rm -f -- "$tmp"
	fail "não consegui copiar $bin para $version_dir."
fi

# rm num link simbólico remove só o link: a pasta para onde ele aponta fica.
if [ -L "$current" ]; then
	rm -f -- "$current" || fail "não consegui remover o link $current."
fi
ln -s -- "$version_dir" "$current" || fail "não consegui criar o link $current."

printf 'claude-tools %s instalado em %s\n' "$version" "$version_dir"
printf 'Link: %s -> %s\n' "$current" "$version_dir"

if [ "$skip_path" -eq 1 ]; then
	echo 'PATH: não mexi (--skip-path).'
	exit 0
fi
profile="$HOME/.profile"
if [ -f "$profile" ] && grep -Fqx -- "$profile_line" "$profile"; then
	printf 'PATH: %s já estava no ~/.profile.\n' "$current"
	exit 0
fi
{
	# Sem o \n final, o comentário grudaria na última linha do arquivo.
	if [ -s "$profile" ] && [ -n "$(tail -c 1 -- "$profile")" ]; then
		printf '\n'
	fi
	printf '# claude-tools\n%s\n' "$profile_line"
} >>"$profile" || fail "não consegui gravar em $profile."
printf "PATH: acrescentei %s ao ~/.profile. Abra um terminal novo para usar o 'claude-tools'.\n" "$current"
