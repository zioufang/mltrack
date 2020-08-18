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

createModel '{"name":"testmodel"}'
createModel '{"name":"testmodel","somthing":"somevalue"}'
createModel '{"somthing":"somevalue"}'

