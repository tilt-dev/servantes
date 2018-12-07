FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/emoji
RUN go install github.com/windmilleng/servantes/emoji

ENTRYPOINT /go/bin/emoji