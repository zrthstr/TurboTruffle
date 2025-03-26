#!/bin/bash
source target

REPOS=$(curl -s -H "Authorization: token $GH_TOKEN" "https://api.github.com/orgs/$ORG/repos?per_page=100" | jq -r '.[].clone_url')

echo $REPOS | sed 's/ /\n/g'
