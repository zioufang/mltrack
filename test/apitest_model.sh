#!/bin/sh

URL="http://localhost:8000"

createModel() {
    name=$1
    echo "testing POST on /models with ( $1 )"
    curl -X POST \
        "$URL/models" \
        --header 'Content-Type: application/json' \
        --header 'Accept: application/json' \
        -d '
        {
          "name": "'"$name"'"
        }'
}

createModel 'testmodel'
