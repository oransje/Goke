CREATE TABLE IF NOT EXISTS Post(
    post_id int NOT NULL,
    title varchar(255),
    sub_title varchar(255),
    created_at datetime,
    updated_at datetime,
    published boolean DEFAULT false,
    PRIMARY KEY(post_id),
    FOREIGN KEY (post_id) REFERENCES Persons(id)
)
