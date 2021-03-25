FROM golang:1.16-alpine as builder

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY ["go.mod", "go.sum", "./"]

RUN go mod download

COPY . .

RUN go build -o ./bin/server ./src

FROM alpine

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY --from=builder ["/usr/src/app/bin/server", "/usr/src/app/"]

EXPOSE 80

CMD ["./server"]
