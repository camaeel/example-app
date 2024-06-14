# example-app

Example go application. This application exposes simple CRUD for making notes and utilizes postgres database as storage. 

## Install

### Install prerequisites

### Using helm

1. Add repo: `helm repo add example-app https://camaeel.github.io/example-app/`
2. Update repos: `helm repo update`
3. Install: `helm upgrade --install --namespace example --create-namespace example-app/example-app`
   
## Local development

### Local run using docker-compose

1. Build docker image: `make docker`
2. Run `docker-compose up`

### Testing

Use [`tests.rest`](./tests.rest) file.
