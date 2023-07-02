run: docker-compose.yaml
	docker-compose -f docker-compose.yaml up
rund:
	docker-compose -f docker-compose.yaml up -d
down: docker-compose.yaml
	docker-compose -f docker-compose.yaml down
build: docker-compose.yaml
	docker-compose -f docker-compose.yaml build