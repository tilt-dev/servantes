FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/doggos
RUN go install github.com/windmilleng/servantes/doggos

ENTRYPOINT /go/bin/doggos