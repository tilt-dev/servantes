FROM golang:1.17-alpine

ADD . /go/src/github.com/tilt-dev/servantes/vigoda
RUN cd /go/src/github.com/tilt-dev/servantes/vigoda && go install .

ENTRYPOINT /go/bin/vigoda
