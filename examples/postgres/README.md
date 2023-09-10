# **POSTGRES** example

This example shows how to run GOKE with [POSTGRES](https://www.postgresql.org/) database server in a Docker container.

## COMMANDS

goke --help

goke test

- --help is full command to see the help
- -h is shorthand to see the help

* all commands has a shorthand and full command.

  **OBS:** The shorthand is the first letter of the full command.

  **NOTE** : Docker is required to run the examples, otherwise you can run the examples directly on your machine or run the `SQLITE3` example.

## Starting **POSTGRES**

**NOTE** : The table name should be the same as the `TABLE_NAME` name, otherwise can't manage the table.

````bash
Start POSTGRES with the following command:

```bash
# docker-compose.yaml file contains POSTGRES service
$ docker-compose up -d
````

```yaml
# default docker-compose.yaml file
username: root
password: root
dialect: postgres
sslmode: disable
dbname: Goke_test
host: localhost
port: 5432
```

## Run GOKE

Run GOKE with the following command:

```bash
$ goke # check if path is valid
$ goke --init # create base structure
```

## Migrating from **dummy.sql** to POSTGRES

```bash
$ goke --migrate # migrate dummy.sql to MariaDB
$ goke --migrate -n # add name to migration file
```

## Drop table for POSTGRES

```bash
$ goke --drop --table ${TABLE_NAME} # drop table if exists
$ goke --drop --t ${TABLE_NAME} # drop table if exists

```

## Dump POSTGRES schema to **JSON** file

```bash
$ goke dump --json # dump POSTGRES schema to JSON file
$ goke dump -j # dump POSTGRES schema to JSON file

```

## **FOLDER** ./migrations

This folder contains all migrations files.

- JSON
- SQL
