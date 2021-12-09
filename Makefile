all: build run

build-dependencies:
	docker build -t dependencies .

build: build-dependencies
	docker-compose build

run:
	docker-compose up

down:
	docker-compose down

rm:
	docker-compose down --rmi all --volumes
