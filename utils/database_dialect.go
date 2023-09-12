package utils

import (
	"regexp"
)

func ReturnDialectSource(dialect string) string {
	var SQLITE_REGEX = "sqlite[3]|sqlite|SQLITE|SQLITE[3]"
	var POSTGRES_REGEX = "postgresql|postgres|psql|POSTGRESQL|POSTGRES|PSQL"
	var MYSQL_REGEX = "mysql|mariadb|maria|MYSQL|MARIADB|MARIA"
	const msg = "[ERROR]: invalid dialect check goke-config.yaml -> dialect"

	IS_SQLITE, sqlite_err := regexp.Match(SQLITE_REGEX, []byte(dialect))
	HandleError(sqlite_err, msg)
	if IS_SQLITE {
		return "sqlite3"
	}

	IS_POSTGRES, postgres_err := regexp.Match(POSTGRES_REGEX, []byte(dialect))
	HandleError(postgres_err, msg)
	if IS_POSTGRES {
		return "postgres"
	}

	IS_MYSQL, mysql_err := regexp.Match(MYSQL_REGEX, []byte(dialect))
	HandleError(mysql_err, msg)
	if IS_MYSQL {
		return "mysql"
	}

	return "Invalid dialect"
}
