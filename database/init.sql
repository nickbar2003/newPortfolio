CREATE TABLE IF NOT EXISTS figure (
    id int(10) NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    birthdate Date,
    deathdate Date,
    PRIMARY KEY(id)
);

INSERT INTO figure 
    (name, birthdate, deathdate)
VALUES 
    ('Abraham Lincoln', '1809-02-12', '1865-04-15'),
    ('George Washington', '1732-02-22', '1799-12-14'),
    ('Julius Caesar', '0100-07-12', '0044-03-15'),
    ('Napoleon Bonaparte', '1769-08-15', '1821-05-05');