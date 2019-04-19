FROM golang

RUN go get gopkg.in/telegram-bot-api.v4

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

WORKDIR /go/src/github.com/heaptracetechnology/microservice-telegram

ADD . /go/src/github.com/heaptracetechnology/microservice-telegram

RUN go install github.com/heaptracetechnology/microservice-telegram

ENTRYPOINT microservice-telegram

EXPOSE 3000