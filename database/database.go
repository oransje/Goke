package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/vsantos1/Goke/config"
)

func TestDatabaseconnection(cfg *config.ConfigYaml) string {
	ConnectionDatabase(cfg)
	message := "Everything Working fine, ready to use."

	return message
}

func ConnectionDatabase(cfg *config.ConfigYaml) *sqlx.DB {

	db, err := sqlx.Connect(cfg.Dialect, cfg.Datasource)
	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	return db
}

func CloseDatabase(cfg *config.ConfigYaml) {
	db := ConnectionDatabase(cfg)
	db.Close()
}
