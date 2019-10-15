curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq '.[] | select (.id == 1) | .connections | .relatives' | sed 's/"/ /g' | sed 's/ //'
