FROM golang:1.10-alpine

ADD . /go/src/github.com/tilt-dev/servantes/snack
RUN go install github.com/tilt-dev/servantes/snack

ENTRYPOINT /go/bin/snack
