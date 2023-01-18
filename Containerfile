FROM ubuntu:latest

ENV APP_NAME=qr-code-generator-api
ENV CGO_ENABLED=1

RUN apt update -y && \
	apt install -y tar && \
	apt install -y wget && \
	apt install -y gcc && \
	wget https://go.dev/dl/go1.20.2.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go1.20.2.linux-amd64.tar.gz && \
	export PATH=$PATH:/usr/local/go/bin && \
	go version

WORKDIR /usr/server/$APP_NAME
COPY . .

RUN export PATH=$PATH:/usr/local/go/bin && \
	go mod tidy && \
	go build -o $APP_NAME

EXPOSE 3000

CMD ./$APP_NAME
