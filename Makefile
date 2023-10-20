PROJECTNAME := $(shell basename "$(PWD)")
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

init:
	podman-compose run --rm app go mod init github.com/jdssl/go-docker
	podman-compose run --rm app air init

dev:
	podman-compose up

down:
	podman-compose down

unit-test:
	go test -tags=unit ${PKG_LIST} -coverpkg=./... -coverprofile=./cover/coverage.out

coverage:
	go tool cover -html=./cover/coverage.out -o ./cover/test-coverage.html
