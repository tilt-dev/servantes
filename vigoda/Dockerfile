FROM golang:1.10-alpine

ADD . /go/src/github.com/tilt-dev/servantes/vigoda
RUN go install github.com/tilt-dev/servantes/vigoda

ENTRYPOINT /go/bin/vigoda
