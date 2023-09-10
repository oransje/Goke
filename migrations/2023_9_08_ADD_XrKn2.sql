CREATE TABLE IF NOT EXISTS Persons(
    id int NOT NULL,
    last_name varchar(255),
    first_name varchar(255),
    city varchar(255),
    PRIMARY KEY(ID)
);


INSERT INTO Persons (id, last_name, first_name, city) 
VALUES 
(1, 'Roberts', 'Joseph', 'Caruaru -CANADA'),
(2,"Bento","Golias","Rio de nojeira"),
(3,"Lanches","Benilson","Acido paulo");
