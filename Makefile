.PHONY: migrate run build

test-handler:
	go test ./internal/api/handler/...

docker-migrate:
	docker-compose run --rm migrate

docker-up:
	docker-compose up --build