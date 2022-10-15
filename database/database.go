package database

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vsantos1/Goke/config"
	"github.com/vsantos1/Goke/utils"
)

func WichDatabase(cfg *config.ConfigYaml) (string, string) {
	var DB_POSTGRES = fmt.Sprintf("host=%s sslmode=%s port=%d user=%s password=%s dbname=%s", cfg.Host, cfg.SslMode, cfg.Port, cfg.Username, cfg.Password, cfg.DbName)

	var DB_MYSQL = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

	var DB_MARIADB = fmt.Sprintf("host=%s sslmode=%s port=%d user=%s password=%s dbname=%s", cfg.Host, cfg.SslMode, cfg.Port, cfg.Username, cfg.Password, cfg.DbName)

	var DB_SQLITE3 = fmt.Sprintf("./%s.sqlite3", cfg.SqliteName)

	var dial = strings.ToLower(cfg.Dialect)

	dialect := utils.ReturnDialectSource(dial)

	if dialect == "postgres" {
		return DB_POSTGRES, dialect
	} else if dialect == "mysql" {
		return DB_MYSQL, dialect
	} else if dialect == "sqlite3" {
		return DB_SQLITE3, dialect
	} else if dialect == "mariadb" {
		return DB_MARIADB, dialect
	} else {
		return dialect, dialect
	}

}

func TestDatabaseconnection(cfg *config.ConfigYaml) string {
	ConnectionDatabase(cfg)
	message := "Everything Working fine, ready to use."
	return message
}

func ConnectionDatabase(cfg *config.ConfigYaml) *sqlx.DB {

	connection_string, dialect := WichDatabase(cfg)
	if dialect == "Invalid dialect" {
		log.Fatalf("Error, invalid dialect check goke-config.yaml -> dialect")
	}
	db, err := sqlx.Connect(dialect, connection_string)

	if err != nil {
		log.Fatalf("%v", err.Error())
	}

	db.SetConnMaxLifetime(5)
	return db
}
