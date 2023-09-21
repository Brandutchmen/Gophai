build:
	docker-compose build && make gqlgen && make sqlc

start:
	docker-compose up -d

stop:
	docker-compose down

restart:
	docker-compose restart

logs:
	docker-compose logs -f --tail=100

run-tests:
	docker-compose exec -e APP_ENV=testing web go test -v ./test/...

gqlgen:
	go generate -v ./internal/graph/...

sqlc:
	sqlc generate

migration:
	docker-compose exec web migrate create -ext sql -dir internal/database/migrations -format 00_2006_01_02_15_0405 $(name)

migrate-up:
	docker-compose exec web go run cli/main.go migrate

# migrate-down:
# 	docker-compose exec web migrate -path internal/database/migrations -database $(DB_URL) -verbose down

