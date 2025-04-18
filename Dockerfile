FROM golang:1.24-alpine

WORKDIR /usr/src/app

RUN apk add --no-cache bash

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

