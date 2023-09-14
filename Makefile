build:
	docker-compose build

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

generate-graphql:
	docker-compose exec web go generate -v ./...