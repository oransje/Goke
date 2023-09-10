package cmd

const TODO_DUMMY = `CREATE TABLE IF NOT EXISTS todo(
    id int NOT NULL,
    title varchar(255),
    content varchar(255),
    created_at datetime,
    updated_at datetime,
    done boolean DEFAULT false,
    PRIMARY KEY(ID)
)
`

const VERSION = "v0.1.3"
