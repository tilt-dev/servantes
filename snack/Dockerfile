FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/snack
RUN go install github.com/windmilleng/servantes/snack

ENTRYPOINT /go/bin/snack