#! /bin/bash
curl -s "https://platform.zone01.gr/assets/superhero/all.json" | jq '.[] | select(.id == 170) | {name: .name, power: powerstats, gender: .appearance.gender}'