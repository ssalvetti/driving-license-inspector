-- this file creates the database and the table(s) to store driving licesnes data from http://dati.mit.gov.it/catalog/dataset/patenti

-- CREATE db
CREATE DATABASE driving_licenses;
GRANT ALL PRIVILEGES ON DATABASE driving_licenses TO postgres;

-- CREATE table
CREATE TABLE IF NOT EXISTS patenti (
    id bigint PRIMARY KEY,
    anno_nascita int,
    regione_residenza varchar(40),
    provincia_residenza varchar(40),
    comune_residenza varchar(60),
    sesso char(1),
    categoria_patente varchar(3),
    data_rilascio varchar(30),
    abilitato_a char(1),
    data_abilitazione_a varchar(30),
    data_scadenza varchar(30),
    punti_patente int
);

-- test an insert
INSERT INTO patenti VALUES (6133015, 1960, 'LOMBARDIA', 'LODI', 'LODI', 'F', 'B', '1979-08-22 00:00:00', 'S', '1979-08-22 00:00:00', '2019-07-21 00:00:00', 30);
TRUNCATE TABLE patenti;