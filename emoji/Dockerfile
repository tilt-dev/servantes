FROM golang:1.17-alpine

ADD . /go/src/github.com/tilt-dev/servantes/emoji
RUN cd /go/src/github.com/tilt-dev/servantes/emoji && go install .

ENTRYPOINT /go/bin/emoji
