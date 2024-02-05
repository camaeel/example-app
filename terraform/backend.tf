terraform {
  backend "s3" {
    #bucket         = "" // taken from aws ssm parameter: /terrraform/state-bucket
    key            = "example-app/dev/terraform.tfstate"
    region         = "eu-central-1"
    encrypt        = true
    kms_key_id     = "alias/main"
  }
}