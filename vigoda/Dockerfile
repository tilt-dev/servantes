FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/vigoda
RUN go install github.com/windmilleng/servantes/vigoda

ENTRYPOINT /go/bin/vigoda