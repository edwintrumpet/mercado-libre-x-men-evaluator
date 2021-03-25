#!/bin/bash

sufix=""

if [[ $OSTYPE == "msys" ]]; then
  sufix=".exe"
fi

function main() {
  go run ./src
}

function runTest() {
  go \test -v ./...
}

function notFound() {
  echo This \command does not exist \in run.sh
}

case $1 in
  '')
    main
    ;;
  run)
    main
    ;;
  test)
    runTest
    ;;
  *)
    notFound
    ;;
esac