# Variables
DOCKER_IMAGE_NAME := curconv
DOCKER_CONTAINER_NAME := curconv-app
DOCKER_COMPOSE_FILE := docker-compose.yml
POSTGRES_IMAGE := postgres:latest

# Default target
all: build

# Build the Docker image
build:
	docker build -t $(DOCKER_IMAGE_NAME) .

# Run Docker Compose
up: build
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop Docker Compose
down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Restart the application container
restart:
	docker-compose -f $(DOCKER_COMPOSE_FILE) restart $(DOCKER_CONTAINER_NAME)

# Tail the logs of the application container
logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f $(DOCKER_CONTAINER_NAME)

# Clean up the environment
clean: down
	docker-compose -f $(DOCKER_COMPOSE_FILE) rm -f
	docker rmi $(DOCKER_IMAGE_NAME)

.PHONY: all build up down restart logs clean
