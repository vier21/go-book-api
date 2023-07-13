build: docker-compose.yaml
	docker-compose -f docker-compose.yaml build
run:
	sudo go run cmd/app/main.go
rund:
	docker-compose -f docker-compose.yaml up -d
down: docker-compose.yaml
	docker-compose -f docker-compose.yaml down
build: docker-compose.yaml
	docker-compose -f docker-compose.yaml build