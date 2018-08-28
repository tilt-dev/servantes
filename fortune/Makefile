.PHONY: all proto test

all: proto test

test: proto
	go build ./...

proto:
	protoc --go_out=. -I. api/fortune.proto
