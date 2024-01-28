all: clean test build

clean: 
	rm -rf bin/*

test: generate
	go test ./...

generate:
	go generate ./...

build:
	go build -o bin/app github.com/camaeel/example-app/cmd/app

docker:
	docker buildx build -t example-app:local .
