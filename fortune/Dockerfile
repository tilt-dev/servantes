FROM golang:1.17

RUN apt update && apt install -y unzip time make protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

ADD . /go/src/github.com/tilt-dev/servantes/fortune
RUN cd /go/src/github.com/tilt-dev/servantes/fortune && make proto
RUN cd /go/src/github.com/tilt-dev/servantes/fortune && go install .

ENTRYPOINT /go/bin/fortune
