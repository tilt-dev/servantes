.PHONY: all proto test tilt up watch docker-ci

all: proto test

test: proto
	go build ./...

proto:
	protoc --go_out=. -I. api/fortune/fortune.proto

watch:
	tilt up --watch servantes

up:
	tilt up servantes

docker-ci:
	cd .circleci && docker build -t gcr.io/windmill-public-containers/servantes-ci .
	docker push gcr.io/windmill-public-containers/servantes-ci
