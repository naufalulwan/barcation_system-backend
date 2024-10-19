FROM golang:alpine

LABEL MAINTAINER "Naufal Ulwan <naufalulwan63@gmail.com>"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . /app

RUN test -e .env

RUN go build -o main .

ENV HOST 0.0.0.0
ENV PORT 8080

EXPOSE 8080

CMD ./main
