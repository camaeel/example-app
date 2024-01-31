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

helm-install:
	helm upgrade --install -n example-app example-app charts/example-app --create-namespace \
		--set cnpgdb.enabled=true \
		--set cnpgdb.storage.storageClass=proxmox-data-ephemeral
