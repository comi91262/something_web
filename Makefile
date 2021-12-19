SERVICE := app-rest app-orm

all: build run

build-dependencies:
	docker build -t dependencies -f dependencies/Dockerfile dependencies	

build: build-dependencies
	docker-compose build $(SERVICE)

create-db: build-dependencies
	docker-compose up --build db-world-creation

migrate: build-dependencies
	docker-compose up --build db-world-migration

run:
	docker-compose up $(SERVICE)

down:
	docker-compose down

rm:
	docker-compose down --rmi all --volumes
