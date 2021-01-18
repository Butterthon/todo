FROM golang:1.14.4-alpine

ENV GO111MODULE on

ADD . /todo

WORKDIR /todo

RUN apk update \
    && apk add --no-cache git \
    && go build -o todo

RUN go get github.com/cosmtrek/air
