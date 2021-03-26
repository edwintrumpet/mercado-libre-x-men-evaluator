FROM golang:1.16-alpine as builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

COPY [".", "."]

EXPOSE 80
