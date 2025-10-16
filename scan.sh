#!/bin/bash
set -euo pipefail

echo "======================================="
echo "[+] Running scan (mirrors)"
echo "======================================="

IMAGE="ghcr.io/trufflesecurity/trufflehog:latest"
mkdir -p results
cd repos
shopt -s nullglob

for BARE in *.git/; do
  BARE="${BARE%/}"
  NAME="${BARE%.git}"
  echo "Preparing to scan $NAME"
  docker run --rm --entrypoint /bin/sh -v "$PWD:/pwd:ro" "$IMAGE" -c '
    set -e
    git clone /pwd/'"$BARE"' /tmp/wt >/dev/stderr 2>&1
    cd /tmp/wt
    git fetch --all --prune >/dev/stderr 2>&1
    git checkout -f >/dev/stderr 2>&1
    trufflehog -j git file:///tmp/wt
  ' > "../results/$NAME.trufflog"


done


