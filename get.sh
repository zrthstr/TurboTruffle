#!/bin/bash

source target

#REPOS=$(curl -s -H "Authorization: token $GH_TOKEN" "https://api.github.com/orgs/$ORG/repos?per_page=100" | jq -r '.[].clone_url')

PAGE=1
ALL_REPOS=()

while :; do
    REPOS=$(curl -s -H "Authorization: token $GH_TOKEN" \
        "https://api.github.com/orgs/$ORG/repos?per_page=100&page=$PAGE" | jq -r '.[].clone_url')

    [[ -z "$REPOS" ]] && break
    ALL_REPOS+=($REPOS)
    ((PAGE++))
done


REPOS=$(printf "%s\n" "${ALL_REPOS[@]}")
echo $REPOS

mkdir -p repos
cd repos

for REPO in $REPOS; do
  REPO_SSH_URL=$(echo $REPO | sed "s|https://github.com|git@$REPLACEMENT_HOST:|")
  REPO_NAME=$(basename $REPO_SSH_URL .git)
  echo "Cloning $REPO_NAME from $REPO_SSH_URL"

  git clone --mirror $REPO_SSH_URL
done

echo "Repositories cloned successfully."
