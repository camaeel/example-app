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

run: build
	GIN_MODE=debug bin/app -configFile config/local.yml

# age password injected into macos keychain:
# security add-generic-password -a AGE-PUBLIC-KEY -s AGE-PUBLIC-KEY -w CONTENTS
helm-install:
	helm upgrade --install -n example-app example-app charts/example-app --create-namespace \
		-f charts/example-app/init-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'` \
		--set ingress.enabled=true \
		--set ingress.hosts[0].host=example-app.`kubectl get certificate -n kong kong-gateway-proxy -ojson | jq -r '.spec.commonName'` \
		--set ingress.hosts[0].paths[0].path="/" \
		--set cnpgdb.logicalBackup.encrypt.agePublicKey=`security find-generic-password -a AGE-PUBLIC-KEY -s AGE-PUBLIC-KEY -w`


helm-restore:
	helm upgrade --install -n example-app-restore example-app charts/example-app --create-namespace \
		-f charts/example-app/restore-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'` \
		--set ingress.enabled=true \
		--set ingress.hosts[0].host=example-app-restore.`kubectl get certificate -n kong kong-gateway-proxy -ojson | jq -r '.spec.commonName'` \
		--set ingress.hosts[0].paths[0].path="/"

helm-template:
	helm template -n example-app example-app charts/example-app \
		-f charts/example-app/init-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'` \
		--set ingress.enabled=true \
		--set ingress.hosts[0].host=example-app.`kubectl get certificate -n kong kong-gateway-proxy -ojson | jq -r '.spec.commonName'` \
		--set ingress.hosts[0].paths[0].path="/" \
		--debug -s templates/ingress.yaml \
		--set cnpgdb.logicalBackup.encrypt.agePublicKey=`security find-generic-password -a AGE-PUBLIC-KEY -s AGE-PUBLIC-KEY -w`

helm-template-restore:
	helm template -n example-app-restore example-app charts/example-app \
		-f charts/example-app/restore-values.yml \
		--set cnpgdb.eso.databaseBackup.vaultUrl=https://`kubectl get ingress -n vault vault -ojson | jq -r '.spec.rules[0].host'` \
		--set ingress.enabled=true \
        --set ingress.hosts[0].host=example-app-restore.`kubectl get certificate -n kong kong-gateway-proxy -ojson | jq -r '.spec.commonName'` \
        --set ingress.hosts[0].paths[0].path="/" \
        --debug


terraform-init:
	terraform -chdir=./terraform init -backend-config=bucket=`aws ssm get-parameter --name /terrraform/state-bucket | jq -r '.Parameter.Value'`

goreleaser:
	goreleaser release  --clean --snapshot
