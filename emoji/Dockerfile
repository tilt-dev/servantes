FROM golang:1.10-alpine

ADD . /go/src/github.com/tilt-dev/servantes/emoji
RUN go install github.com/tilt-dev/servantes/emoji

ENTRYPOINT /go/bin/emoji
