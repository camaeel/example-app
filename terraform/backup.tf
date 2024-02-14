module "backup-bucket" {
#  source = "../../tf-modules/backup-bucket" #development
  source = "git@github.com:camaeel/tf-modules.git//backup-bucket?ref=backup-bucket-v1.2.0"
  aws_tags = var.aws_tags
  purpose = "postgres-backup"
  backup_expiration_days = 7
  destroy_backups = true
  vault_iam_identity_arn = data.aws_iam_user.vault.arn
  vault_aws_backend_path = "aws"
  vault_kubernetes_auth_path = "kubernetes"
  namespaces = ["example-app", "example-app-restore"]
  service_accounts = ["example-app-cnpgdb-db-backup-aws-creds"]
}
