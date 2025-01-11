# Description: Makefile for go project
# run project
run:
	air

# install dependencies
install:
	go mod tidy

# run docker
docker-run:
	docker compose -f docker/docker-compose.yaml up -d
