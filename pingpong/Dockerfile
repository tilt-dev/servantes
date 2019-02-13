FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/pingpong
RUN go install github.com/windmilleng/servantes/pingpong

ENTRYPOINT /go/bin/pingpong
