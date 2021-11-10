FROM golang:1.17-alpine

ADD . /go/src/github.com/tilt-dev/servantes/doggos
RUN cd /go/src/github.com/tilt-dev/servantes/doggos && go install .

ENTRYPOINT /go/bin/doggos
