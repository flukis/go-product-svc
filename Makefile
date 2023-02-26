build:
	@go build -o ./build/product-svc

run: build
	./build/product-svc

test:
	go test -v ./...