init:
	docker compose run --rm app go mod init github.com/JonatanLima/go-docker
	docker compose run --rm app air init

dev:
	docker compose up

down:
	docker compose down

