#!/bin/bash
set -euo pipefail

# Map repo name -> remote URL (normalized to https)
declare -A URLMAP

to_https() {
  local u="$1"
  case "$u" in
    git@*:* )
      local host="${u#git@}"; host="${host%%:*}"
      local path="${u#*:}"; path="${path%.git}"
      echo "https://$host/$path"
      ;;
    ssh://* )
      local host="$(printf "%s" "$u" | sed -E 's@^ssh://[^@]+@ssh://@; s@^ssh://@@; s@/.*@@')"
      local path="$(printf "%s" "$u" | sed -E 's@^ssh://[^/]+/@@')"
      path="${path%.git}"
      echo "https://$host/$path"
      ;;
    http://*|https://* )
      printf "%s" "${u%.git}"
      ;;
    * )
      printf "%s" "${u%.git}"
      ;;
  esac
}

cd "$(dirname "$0")"

# build map from mirrors in repos/*.git
shopt -s nullglob
for bare in repos/*.git; do
  name="$(basename "${bare%.git}")"
  remote="$(git --git-dir="$bare" config --get remote.origin.url || true)"
  [[ -n "$remote" ]] || continue
  URLMAP["$name"]="$(to_https "$remote")"
done

# rewrite repository fields in results/*.trufflog -> in place
for f in results/*.trufflog; do
  [[ -e "$f" ]] || continue
  name="$(basename "$f" .trufflog)"
  url="${URLMAP[$name]:-}"
  [[ -n "$url" ]] || continue

  jq -c --arg repo "$url" '
    # nested path used by newer trufflehog
    ( .SourceMetadata as $s
      | if ($s|type=="object")
        then ( .SourceMetadata.Data.Git.repository ) = $repo
        else .
        end )
    |
    # top-level fallback used by some versions
    ( if has("repository") then .repository=$repo else . end )
  ' "$f" > "$f.tmp" && mv "$f.tmp" "$f"
done

echo "Repository URLs normalized in results/*.trufflog"

