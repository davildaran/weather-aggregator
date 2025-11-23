build-bin:
	rm weatheragg
	go build -o weatheragg .

build:
	docker build -t weatheragg:latest -f Dockerfile .

up:
	docker run -p 8080:8080 weatheragg:latest
