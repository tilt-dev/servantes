FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/random
RUN go install github.com/windmilleng/servantes/random

ENTRYPOINT /go/bin/random