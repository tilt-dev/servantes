FROM golang:1.10

ADD . /go/src/github.com/windmilleng/servantes/vigoda
RUN go install github.com/windmilleng/servantes/vigoda

ENV TEMPLATE_DIR /go/src/github.com/windmilleng/servantes/vigoda/web/templates

ENTRYPOINT /go/bin/vigoda
