#! /bin/bash
URL="https://platform.zone01.gr/assets/superhero/all.json"
curl -s "$URL" | jq -r '
.[]
| select(.id == 170)
| .name, .powerstats.power, .appearance.gender
'