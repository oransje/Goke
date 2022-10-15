# **SQLITE** example

This example shows how to run GOKE with [SQLITE](hhttps://www.sqlite.org/index.html).

## COMMANDS

goke --help

goke test

- --help is full command to see the help
- -h is shorthand to see the help

* all commands has a shorthand and full command.

  **OBS:** The shorthand is the first letter of the full command.

## Starting with **SQLITE**

**NOTE** : The table name should be the same as the `TABLE_NAME` name, otherwise can't manage the table.

```yaml
username: root #NOT USED
password: root #NOT USED
dialect: sqlite #DIALECT
sslmode: disable #NOT USED
dbname: Goke_test #NOT USED
host: localhost #NOT USED
sqlite_name: goke-test #NAME OF THE SQLITE FILE
port: 5432 #NOT USED
```

## Run GOKE

Run GOKE with the following command:

```bash
$ goke # check if path is valid
$ goke --init # create base structure
```

## Migrating from **dummy.sql** to SQLITE

```bash
$ goke --migrate # migrate dummy.sql to SQLITE
$ goke --migrate -n # add name to migration file
```

## Drop table for SQLITE

```bash
$ goke --drop --table ${TABLE_NAME} # drop table if exists
$ goke --drop --t ${TABLE_NAME} # drop table if exists

```

## Dump SQLITE schema to **JSON** file

```bash
$ goke dump --json # dump SQLITE schema to JSON file
$ goke dump -j # dump SQLITE schema to JSON file

```

## **FOLDER**./migrations

This folder contains all migrations files.

- JSON
- SQL
