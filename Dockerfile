FROM golang

RUN go get gopkg.in/telegram-bot-api.v4

RUN go get github.com/gorilla/mux

RUN go get github.com/cloudevents/sdk-go

WORKDIR /go/src/github.com/oms-services/telegram

ADD . /go/src/github.com/oms-services/telegram

RUN go install github.com/oms-services/telegram

ENTRYPOINT telegram

EXPOSE 3000