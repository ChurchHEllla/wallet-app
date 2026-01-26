.PHONY: migrate run build

migrate:
	go run cmd/migrate/main.go

run:
	go run cmd/service/main.go

build:
	go build -o wallet-app cmd/service/main.go
	go build -o migrate-tool cmd/migrate/main.go

docker-migrate:
	docker-compose run --rm app go run cmd/migrate/main.go