ARG APP_NAME=api.fernandopaiva.dev

FROM golang:latest as compile

ARG APP_NAME
ENV APP_NAME=$APP_NAME
ENV CGO_ENABLED=1

WORKDIR /usr/app
COPY . .

RUN	go mod download && \
	go build -o $APP_NAME

FROM busybox:latest

ARG APP_NAME
ENV APP_NAME=$APP_NAME

COPY --from=compile /usr/app /usr/app
WORKDIR /usr/app

EXPOSE 3000

CMD ./$APP_NAME
