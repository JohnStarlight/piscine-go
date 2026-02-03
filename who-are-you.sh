#! /bin/bash
URL="https://platform.zone01.gr/assets/superhero/all.json"
curl -s "$URL" | jq '
.[]
| select(.id == 70)
| .name'