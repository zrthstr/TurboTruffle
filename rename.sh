#!/bin/bash

cd repos

for REPO_DIR in *.git; do
  REPO_NAME="${REPO_DIR%.git}"
  mkdir -p "$REPO_NAME"
  mv "$REPO_DIR" "$REPO_NAME/.git"
done

echo "All .git folders moved successfully."
