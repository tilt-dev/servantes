FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/secrets
RUN go install github.com/windmilleng/servantes/secrets

ENTRYPOINT /go/bin/secrets