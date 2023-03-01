build:
	@go build -o bin/product

run: build
	./bin/product

test:
	go test -v ./...