FROM golang:1.10

ENV PROTOC_VERSION 3.5.1

ADD . /go/src/github.com/windmilleng/servantes/fe
RUN go install github.com/windmilleng/servantes/fe
ENTRYPOINT /go/bin/fe
