up:
	docker-compose -f docker-compose.dev.yaml up -d
down:
	docker-compose -f docker-compose.dev.yaml down --volumes
