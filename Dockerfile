FROM golang:1.22.0-bullseye as compile

ENV APP_NAME=api.fernandopaiva.dev
ENV CGO_ENABLED=1

WORKDIR /usr/app
COPY . .

RUN apt install gcc && \
	go mod tidy && \
	go build -o $APP_NAME

FROM ubuntu:latest

ENV APP_NAME=api.fernandopaiva.dev

COPY --from=compile /usr/app /usr/app
WORKDIR /usr/app

EXPOSE 3000

CMD ./$APP_NAME
