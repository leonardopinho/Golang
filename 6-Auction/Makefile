.PHONY: run create show check test
run:
	docker compose up --build

create:
	curl -X POST http://localhost:8080/auction \
		-H "Content-Type: application/json" \
		-d '{"product_name": "Laptop", "category": "electronics", "description": "Laptop Top 32GB", "condition": 0}'

show:
	curl -X GET http://localhost:8080/auction?status=0 \
		-H "Content-Type: application/json"

check:
	curl -X GET http://localhost:8080/auction?status=1 \
		-H "Content-Type: application/json"

test:
	go test -v internal/infra/database/auction/create_auction_test.go
