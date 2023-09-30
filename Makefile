init:
	podman-compose run --rm app go mod init github.com/jdssl/go-docker
	podman-compose run --rm app air init

dev:
	podman-compose up

down:
	podman-compose down

