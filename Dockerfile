FROM golang:1.23.0-alpine3.20 AS build

WORKDIR /go/src

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o server

FROM alpine:3.20.2

COPY --from=build /go/src/server /usr/bin/server
