MAIN_PATH=cmd/test/main.go
BINARY_FILE=binary

clean:
	go clean
	rm -rf bin/${BINARY_FILE}
build: clean
	go build -o bin/${BINARY_FILE} ${MAIN_PATH}

run: build
	./bin/${BINARY_FILE}