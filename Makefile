# Golang
build-bin:
	rm weatheragg
	go build -o weatheragg .

# Docker
build:
	docker build -t weatheragg:latest -f Dockerfile .

run:
	docker run -p 8080:8080 weatheragg:latest

# Docker Compose
up:
	docker compose -f docker-compose.yaml --profile redis --profile api up 
