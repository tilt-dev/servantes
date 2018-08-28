.PHONY: all proto test tilt up watch docker-ci

all: proto test

test: proto
	go build ./...

proto:
	cd fortune && make proto

watch:
	tilt up --watch servantes

up:
	tilt up servantes

docker-ci:
	cd .circleci && docker build -t gcr.io/windmill-public-containers/servantes-ci .
	docker push gcr.io/windmill-public-containers/servantes-ci
