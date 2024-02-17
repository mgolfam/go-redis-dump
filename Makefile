# Variables
BINARY_NAME=redis2csv
APP_DIR=app
BUILD_DIR=app/bin
BIN_DIR=bin

# Targets
all: build run

clean:
	ls
	rm -rf $(BUILD_DIR)

build:
	mkdir -p bin
	cp ./config.json ./bin
	mkdir -p $(BUILD_DIR)
	cd $(APP_DIR) && go mod tidy && go build -o bin/$(BINARY_NAME) .
	mv $(BUILD_DIR)/* bin/
	rm -rf $(BUILD_DIR)

run:
	cd ./$(BIN_DIR)/ && ./$(BINARY_NAME)

docker-build:
	docker build -t myapp:latest -f $(APP_DIR)/Dockerfile .

docker-run:
	docker-compose up

.PHONY: all clean build run docker-build docker-run