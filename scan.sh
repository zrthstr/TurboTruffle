#!/bin/bash
source target

test -d results || mkdir results
test -d repos || { echo "repos dir not found"; exit 1; }

cd repos

for REPO_DIR in */; do
  REPO_NAME=$(basename $REPO_DIR)
  echo "Preparing to scan $REPO_NAME"


  docker run --rm -it \
    -v "$PWD:/pwd" \
    trufflesecurity/trufflehog:latest -j git file:///pwd/$REPO_NAME 2>&1 | tee ../results/$REPO_NAME.trufflog

done
