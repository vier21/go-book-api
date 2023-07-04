build: docker-compose.yaml
	docker-compose -f docker-compose.yaml build
run: docker-compose.yaml
	docker-compose up --build backend
rund:
	docker-compose -f docker-compose.yaml up -d
down: docker-compose.yaml
	docker-compose -f docker-compose.yaml down
build: docker-compose.yaml
	docker-compose -f docker-compose.yaml build