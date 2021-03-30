.PHONY: all proto test tilt up watch docker-ci integration

all: proto test

test: proto
	go build ./...

integration: test
	tilt ci

proto:
	cd fortune && make proto

watch:
	tilt up --watch servantes

up:
	tilt up servantes

