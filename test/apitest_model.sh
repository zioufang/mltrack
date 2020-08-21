#!/bin/sh

URL="http://localhost:8000"

createModel() {
    echo "testing POST on /models with ( $1 )"
    curl -X POST \
        "$URL/models" \
        --header 'Content-Type: application/json' \
        --header 'Accept: application/json' \
        -d $1
}

getModelByName() {
    echo "testing GET /models with model name"
    curl -X "$URL/models/$1"
}

deleteModel() {
    echo "testing DELETE /models"
    curl -X DELETE "$URL/models/$1"
}

createModel '{"name":"testmodel"}'
createModel '{"name":"testmodel"}'
createModel '{"somthing":"somevalue"}'

