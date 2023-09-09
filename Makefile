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

test:
	docker-compose exec web go test -v ./...