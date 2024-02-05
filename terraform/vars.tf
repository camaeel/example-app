variable "aws_tags" {
  default = {
    project  = "example-app"
    source   = "https://github.com/camaeel/example-app"
    env_name = "dev"
  }
  type = map(string)
}

variable "vault_aws_iam_user_name" {
  default = "home-lab-vault-dev-vault"
  description = "IAM role used by vault"
  type = string
}

variable "aws_region" {
  default = "eu-north-1"
  description = "AWS region"
  type = string
}

variable "vault_aws_backend_path" {
  default = "aws"
  type = string
  description = "mount point for aws secret"
}

variable "vault_aws_backend_role_name" {
  default = "example-app-dev-database-backup"
}