curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq '.[] | select (.id == '$HERO_ID')' | jq '.connections.relatives' | sed 's/"//g'
