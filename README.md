# example-app

Example go application. This application exposes simple CRUD for making notes and utilizes postgres database as storage. 

# Install

## Using helm

1. Add repo: `helm repo add example-app https://camaeel.github.io/example-app/`
2. Update repos: `helm repo update`
3. Install: `helm upgrade --install --namespace example --create-namespace example-app/example-app`
   
