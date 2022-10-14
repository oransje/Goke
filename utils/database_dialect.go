package utils

import (
	"log"
	"regexp"
)

func checkErr(e error) {
	if e != nil {
		log.Fatalf("error %v", e)
	}
}

func ReturnDialectSource(dialect string) string {
	var SQLITE_REGEX = "sqlite[3]|sqlite|SQLITE|SQLITE[3]"
	var POSTGRES_REGEX = "postgresql|postgres|psql|POSTGRESQL|POSTGRES|PSQL"
	var MYSQL_REGEX = "mysql|mariadb|maria|MYSQL|MARIADB|MARIA"

	IS_SQLITE, sqlite_err := regexp.Match(SQLITE_REGEX, []byte(dialect))
	checkErr(sqlite_err)
	if IS_SQLITE {
		return "sqlite3"
	}

	IS_POSTGRES, postgres_err := regexp.Match(POSTGRES_REGEX, []byte(dialect))
	checkErr(postgres_err)
	if IS_POSTGRES {
		return "postgres"
	}

	IS_MYSQL, mysql_err := regexp.Match(MYSQL_REGEX, []byte(dialect))
	checkErr(mysql_err)
	if IS_MYSQL {
		return "mysql"
	}

	return "Invalid dialect"
}
