FROM golang:1.17-alpine

ADD . /go/src/github.com/tilt-dev/servantes/snack
RUN cd /go/src/github.com/tilt-dev/servantes/snack && go install .

ENTRYPOINT /go/bin/snack
