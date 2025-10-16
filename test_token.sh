#!/bin/bash
echo "======================================="
echo [+] Running test_token
echo "======================================="

source target

REPOS=$(curl -s -H "Authorization: token $GH_TOKEN" "https://api.github.com/orgs/$ORG/repos?per_page=100" | jq -r '.[].clone_url')

echo "Found repo(s):"
echo $REPOS | sed 's/ /\n/g'
echo "---"
echo "Found Org(s):"
curl -s -H "Authorization: Bearer $GH_TOKEN" https://api.github.com/user/orgs | jq -r '.[].login'
echo "---"
