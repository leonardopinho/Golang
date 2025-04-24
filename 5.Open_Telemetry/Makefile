.PHONY: run check

run:
	@docker-compose up --build

check:
	@echo -n "422: "; curl -s "http://127.0.0.1:8080/get_address/1234567"
	@echo -n "404: "; curl -s "http://127.0.0.1:8080/get_address/12345678"
	@echo -n "200: "; curl -s "http://127.0.0.1:8080/get_address/13330250"