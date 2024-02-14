# CloudNative PG operator

Cloudnative PG operator does restore during creation of new cluster.

Taking backup:
* configured cron
* with kubectl-cnpg plugin `kubectl cnpg backup ${CLUSTER_NAME} [--backup-name ${OPTIONAL_BACKUP_NAME}]`

Removing cluster:
`kubectl delete cluster example-app-cnpgdb` (warning scheduled backups are deleted as well)

Restore - didn't succeeded. 
