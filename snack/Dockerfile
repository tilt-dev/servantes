FROM golang:1.10

# TODO(dbentley): this is only relevant in devel; factor this out to a separate image
RUN git clone https://github.com/windmilleng/rerun-process-wrapper
RUN cp rerun-process-wrapper/*.sh /bin

ADD . /go/src/github.com/windmilleng/servantes/snack
RUN go install github.com/windmilleng/servantes/snack

ENTRYPOINT /bin/start.sh /go/bin/snack