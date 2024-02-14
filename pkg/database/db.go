package database

import (
	"database/sql"
	"github.com/camaeel/example-app/pkg/config"
	_ "github.com/lib/pq"
)

func SetupDriver() (*sql.DB, error) {
	dbCfg := config.GetConfig()
	// connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"

	// normalize driverName
	driverName := dbCfg.DatasourceUrl.Scheme
	if driverName == "postgresql" {
		driverName = "postgres"
	}

	db, err := sql.Open(driverName, dbCfg.DatasourceUrl.String())
	return db, err
}
