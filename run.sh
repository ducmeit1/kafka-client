#!/usr/bin/env bash

function run() {
    ./main -config=/app/config.yaml -log-level=debug -http-port=8888 -config-type=yaml
}

run