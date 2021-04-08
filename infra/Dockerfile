FROM golang:latest

COPY ../app /go/src/app

WORKDIR /go/src/app

RUN apt-get update \
  && go mod download

EXPOSE 8080

CMD go run main.go