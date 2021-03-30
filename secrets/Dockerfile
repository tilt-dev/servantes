FROM golang:1.10-alpine

ADD . /go/src/github.com/tilt-dev/servantes/secrets
RUN go install github.com/tilt-dev/servantes/secrets

ENTRYPOINT /go/bin/secrets
