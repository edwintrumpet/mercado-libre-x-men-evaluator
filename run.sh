#!/bin/bash

source ./.env

sufix=""

if [[ $OSTYPE == "msys" ]]; then
  sufix=".exe"
fi

function main() {
  CompileDaemon -exclude-dir=.git -exclude=".#*" -directory=./src -build="go build -o ../bin/server${sufix}" -command="./bin/server${sufix}" --color=true
}

function dockerDev() {
  docker-compose -f ./docker/docker-compose.dev.yml up
}

function dockerDown() {
  docker-compose -f ./docker/docker-compose.dev.yml down
}

function dockerBuildDev() {
  docker-compose -f ./docker/docker-compose.dev.yml build
}

function build() {
  docker build -f ./docker/dockerfiles/prod.Dockerfile -t ${ECR_URL}/${REPOSITORY_NAME}:${TAG} \.
}

function ecrAuth() {
  aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 251265670580.dkr.ecr.us-east-1.amazonaws.com 
}

function push() {
  docker build -f ./docker/dockerfiles/prod.Dockerfile -t ${ECR_URL}/${REPOSITORY_NAME}:${TAG} \. && docker push ${ECR_URL}/${REPOSITORY_NAME}:${TAG}
}

function runDB() {
  docker \exec -it docker_db_1 psql -U postgres
}

function runTest() {
  go \test -v ./...
}

function runCoverage() {
  go test -cover ./...
}

function notFound() {
  echo This \command does not exist \in run.sh
}

case $1 in
  '')
    dockerDev
    ;;
  buildDev)
    dockerBuildDev
    ;;
  down)
    dockerDown
    ;;
  build)
    build
    ;;
  ecrAuth)
    ecrAuth
    ;;
  push)
    push
    ;;
  runDB)
    runDB
    ;;
  test)
    runTest
    ;;
  coverage)
    runCoverage
    ;;
  *)
    notFound
    ;;
esac
