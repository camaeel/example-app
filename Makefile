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
		--set bitnamipostgres.enabled=false \
		--set cnpgdb.storage.storageClass=proxmox-data-ephemeral \
		--set cnpgdb.eso.databaseBackup.vaultUrl=`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'`

helm-template:
	helm template -n example-app example-app charts/example-app \
		--set cnpgdb.enabled=true \
		--set bitnamipostgres.enabled=false \
		--set cnpgdb.storage.storageClass=proxmox-data-ephemeral \
		--set cnpgdb.eso.databaseBackup.vaultUrl=`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'`

goreleaser:
	goreleaser release  --clean --snapshot
