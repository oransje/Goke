# **MARIADB** example

This example shows how to run GOKE with [MariaDB](https://mariadb.org/) database server in a Docker container.

## COMMANDS

goke --help

- --help is full command to see the help
- -h is shorthand to see the help

* all commands has a shorthand and full command.

  **OBS:** The shorthand is the first letter of the full command.

  **NOTE** : Docker is required to run the examples, otherwise you can run the examples directly on your machine or run the `SQLITE3` example.

## Starting **MariaDB**

**NOTE** : The table name should be the same as the `TABLE_NAME` name, otherwise can't manage the table.

````bash
Start MariaDB with the following command:

```bash
# docker-compose.yaml file contains MARIADB service
$ docker-compose up -d
````

```yaml
# default docker-compose.yaml file
username: root
password: root
dialect: mariadb
sslmode: disable
dbname: Goke_test
host: localhost
port: 3306
```

## Run GOKE

Run GOKE with the following command:

```bash
$ goke # check if path is valid
$ goke --init # create base structure
```

## Migrating from **dummy.sql** to MariaDB

```bash
$ goke --migrate # migrate dummy.sql to MariaDB
$ goke --migrate -n # add name to migration file
```

## Drop table for MariaDB

```bash
$ goke --drop --table ${TABLE_NAME} # drop table if exists
$ goke --drop --t ${TABLE_NAME} # drop table if exists

```

## Dump MariaDB schema to **JSON** file

```bash
$ goke dump --json # dump MariaDB schema to JSON file
$ goke dump -j # dump MariaDB schema to JSON file

```

## ./migrations

This folder contanis all migrations files.

- JSON
- SQL
