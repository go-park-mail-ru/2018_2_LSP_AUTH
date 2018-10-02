FROM golang:alpine

RUN apk add --no-cache git

ADD . /go/src/github.com/go-park-mail-ru/2018_2_LSP_AUTH

RUN cd /go/src/github.com/go-park-mail-ru/2018_2_LSP_AUTH && go get ./...

RUN go install github.com/go-park-mail-ru/2018_2_LSP_AUTH

ENTRYPOINT /go/bin/2018_2_LSP_AUTH

EXPOSE 8080