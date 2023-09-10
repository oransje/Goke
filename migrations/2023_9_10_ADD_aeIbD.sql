CREATE TABLE IF NOT EXISTS Addresses(
    id int NOT NULL,
    cep varchar(255) NOT NULL,
    logradouro varchar(255),
    complemento varchar(255),
    bairro varchar(255),
    localidade varchar(255),
    uf varchar(255),    
    PRIMARY KEY(ID)
);

INSERT INTO Addresses (id,cep, logradouro, complemento, bairro, localidade, uf) VALUES 
    (1,'01001-000', 'Praça da Sé', 'lado ímpar', 'Sé', 'São Paulo', 'SP'),
    (2,'12345-678', '123 Main St', 'Apt 4B', 'Downtown', 'Cityville', 'CA'),
    (3,'98765-432', '456 Elm St', '', 'Suburbia', 'Villagetown', 'NY'),
    (4,'54321-876', '789 Oak St', 'Unit 5', 'Outskirts', 'Countryside', 'TX'),
    (5,'24680-135', 'Élysée Boulevard', 'Appt Ȃ3', 'Champs-Élysées', 'Paris', 'Île-de-France'),
    (6,'99999-999', '700 Null St', NULL, 'Emptyville', 'Nowhere', 'NA'),
    (7,'11111-111', '555 State St', 'Suite 2A', 'Downtown', 'Citysville', 'TX'),
    (8,'88888-888', '777 Park Ave', 'Unit 5', 'Central Park', 'City Center', 'NY'),
    (9,'22222-222', '123 Missing State St', 'Apt 1', 'Undefined', 'Unknownville', 'NA'),
    (10,'33333-333', '456 Short St', '', 'Suburb', 'Townton', 'AZ'),
    (11,'55555-555', '987 Mountain Rd', 'Suite 10', 'Hillside', 'Scenicville', 'CO');
