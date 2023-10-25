curl -s https://learn.zone01kisumu.ke/assets/superhero/all.json | jp ' . [] | select(.id==170).name'
