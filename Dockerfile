# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /btc-app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
RUN go get rate-api

COPY . .





RUN go build -o /docker-btc-rate-api

EXPOSE 8080

CMD [ "/docker-btc-rate-api" ]