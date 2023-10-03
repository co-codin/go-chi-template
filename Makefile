DB_DOCKER_CONTAINER=coffee_db
BINARY_NAME=coffeapi


# creating the container with postgres software
postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine


# creating the coffee db inside the postgres container
createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=root --owner=root coffedb


# stop other docker containers
stop_containers:
	@echo "Stopping other docker containers"
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers..."; \
		docker stop $$(docker ps -q); \
	else \
		echo "no active containers found..."; \
	fi


# start docker container
start-docker:
	docker start ${DB_DOCKER_CONTAINER}


build:
	@echo "Building backend api binary"
	go build -o ${BINARY_NAME} cmd/server/*.go
	@echo "Binary built!"


run:
	go run cmd/server/main.go
