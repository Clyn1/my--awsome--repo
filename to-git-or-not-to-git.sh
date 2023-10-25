curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' . [] | select(.id==170).name'
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' . [] | select(.id==170).powerstats | .power'
curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jq ' . [] | select(.id==170).appearance | .gender'


