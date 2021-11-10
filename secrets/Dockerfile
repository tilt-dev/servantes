FROM golang:1.17-alpine

ADD . /go/src/github.com/tilt-dev/servantes/secrets
RUN cd /go/src/github.com/tilt-dev/servantes/secrets && go install .

ENTRYPOINT /go/bin/secrets
