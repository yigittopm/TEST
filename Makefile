MAIN_PATH=cmd/test/main.go # Contains main.go file
BINARY_FILE=binary 		   # Binary file name

# Removes binary
clean:
	go clean
	rm -rf bin/${BINARY_FILE}

# Builds main.go
build: clean
	go build -o bin/${BINARY_FILE} ${MAIN_PATH}

# First removes the binary file than rebuilds main.go and runs the binary
run: build
	./bin/${BINARY_FILE}

# The project starts and ends with this command
# Removes running composes
down:
	docker compose down
# Starts compose file
up: down
	docker compose up --build

# Swagger docs update
swag:
	swag init -g ${MAIN_PATH} -o ./docs  