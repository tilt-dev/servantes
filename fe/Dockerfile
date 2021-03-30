FROM golang:1.10

RUN apt update && apt install -y unzip time make

ENV PROTOC_VERSION 3.5.1

RUN wget https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
  unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d protoc && \
  mv protoc/bin/protoc /usr/bin/protoc

RUN go get github.com/golang/protobuf/protoc-gen-go

ADD . /go/src/github.com/tilt-dev/servantes/fe
RUN go install github.com/tilt-dev/servantes/fe
ENTRYPOINT /go/bin/fe
