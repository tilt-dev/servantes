FROM golang:1.10-alpine

ADD . /go/src/github.com/tilt-dev/servantes/doggos
RUN go install github.com/tilt-dev/servantes/doggos

ENTRYPOINT /go/bin/doggos
