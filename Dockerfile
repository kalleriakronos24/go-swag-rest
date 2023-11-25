FROM golang:latest as base

RUN apt-get update && apt-get install -y \
    git \
    curl

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

RUN apt-get clean && rm -rf /var/lib/apt/lists/*

RUN go env -w GO111MODULE=auto

WORKDIR /opt/app/api

COPY . /opt/app/api

CMD ["go mod init github.com/kalleriakronos24/golang-experimental", "go mod tidy"]

CMD ["air"]