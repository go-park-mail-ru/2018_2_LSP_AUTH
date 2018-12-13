FROM golang:alpine

ADD server.key /go/server.key
ADD server.crt /go/server.crt
ADD auth.linux.amd64 /go/auth.linux.amd64

ENTRYPOINT /go/auth.linux.amd64

EXPOSE 8080