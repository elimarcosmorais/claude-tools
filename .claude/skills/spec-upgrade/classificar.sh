#!/bin/sh
# Compara o .claude/ do projeto com o do SpecFlow baixado e classifica cada arquivo.
#
# Uso: sh classificar.sh <clone-do-specflow> [raiz-do-projeto]
#
# Saída: uma linha por arquivo, "<situação> <caminho>", ordenada por caminho. Situações:
#   novo                    existe no framework, não no projeto
#   igual                   idêntico à versão atual do framework
#   desatualizado           idêntico a uma versão antiga do framework (nunca foi editado no projeto)
#   personalizado           difere de todas as versões do framework (editado no projeto)
#   removido                saiu do framework; o do projeto é uma versão antiga, sem edição
#   removido-personalizado  saiu do framework, mas foi editado no projeto
#   do-projeto              nunca existiu no framework (skill, agent ou rule do próprio projeto)
#
# "Do framework" = todo caminho que já existiu em .claude/ em qualquer commit do SpecFlow. É o
# histórico, e não o prefixo do nome, que separa o que é do framework do que é do projeto.
#
# Os settings ficam fora: o do projeto nunca se copia nem se sobrescreve; o SKILL.md só junta
# nele as regras de permissão do framework.
set -eu

up=${1:?uso: sh classificar.sh <clone-do-specflow> [raiz-do-projeto]}
proj=${2:-.}

tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT

# Todo par "blob caminho" que já existiu em .claude/ no framework.
git -C "$up" log --all --no-renames --format= --raw --no-abbrev -- .claude \
  | awk '{ if ($3 !~ /^0+$/) print $3, $6; if ($4 !~ /^0+$/) print $4, $6 }' \
  | grep -v -e ' .claude/settings.json$' -e ' .claude/settings.local.json$' \
  | sort -u > "$tmp/versoes"
cut -d' ' -f2 "$tmp/versoes" | sort -u > "$tmp/historicos"
git -C "$up" ls-files .claude \
  | grep -vxF -e .claude/settings.json -e .claude/settings.local.json \
  | sort -u > "$tmp/atuais"

# Pastas de .claude/ que o framework já usou (skills, agents, rules…): só elas são varridas no
# projeto, para não entrar em settings, worktrees ou o que mais o projeto guardar em .claude/.
: > "$tmp/locais"
for pasta in $(cut -d/ -f2 "$tmp/historicos" | sort -u); do
  [ -d "$proj/.claude/$pasta" ] || continue
  (cd "$proj" && find ".claude/$pasta" -type f) >> "$tmp/locais"
done
sort -u "$tmp/locais" -o "$tmp/locais"

# Hash do arquivo do projeto com fim de linha LF, como o framework guarda: CRLF do Windows não
# conta como edição.
hash_local() {
  tr -d '\r' < "$proj/$1" | git -C "$up" hash-object --stdin
}

ja_foi_do_framework() {
  grep -qxF "$1 $2" "$tmp/versoes"
}

sort -u "$tmp/atuais" "$tmp/locais" | while IFS= read -r p; do
  if grep -qxF "$p" "$tmp/atuais"; then
    if [ ! -f "$proj/$p" ]; then
      echo "novo $p"
    else
      h=$(hash_local "$p")
      if [ "$h" = "$(git -C "$up" rev-parse "HEAD:$p")" ]; then
        echo "igual $p"
      elif ja_foi_do_framework "$h" "$p"; then
        echo "desatualizado $p"
      else
        echo "personalizado $p"
      fi
    fi
  elif grep -qxF "$p" "$tmp/historicos"; then
    if ja_foi_do_framework "$(hash_local "$p")" "$p"; then
      echo "removido $p"
    else
      echo "removido-personalizado $p"
    fi
  else
    echo "do-projeto $p"
  fi
done
