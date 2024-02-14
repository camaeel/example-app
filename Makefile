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
	docker buildx build -t example-app:local --load .

helm-install:
	helm upgrade --install -n example-app example-app charts/example-app --create-namespace \
		-f charts/example-app/init-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'`

helm-restore:
	helm upgrade --install -n example-app-restore example-app charts/example-app --create-namespace \
		-f charts/example-app/restore-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'`


helm-template:
	helm template -n example-app example-app charts/example-app \
		-f charts/example-app/init-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'`

helm-template-restore:
	helm template -n example-app-restore example-app charts/example-app \
		-f charts/example-app/restore-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'`


terraform-init:
	terraform -chdir=./terraform init -backend-config=bucket=`aws ssm get-parameter --name /terrraform/state-bucket | jq -r '.Parameter.Value'`

goreleaser:
	goreleaser release  --clean --snapshot
