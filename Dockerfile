FROM golang:1.13-alpine

ENV GOPATH=/go

WORKDIR /go/src/github.com/mg/microgardener

RUN apk add git

COPY . /go/src/github.com/mg/microgardener
RUN go get "github.com/eclipse/paho.mqtt.golang"
RUN go get "github.com/lib/pq"

CMD [ "go", "run", "mg-main.go" ]