FROM golang:alpine

ADD . /go/auth.linux.amd64

ENTRYPOINT /go/auth.linux.amd64

EXPOSE 8080