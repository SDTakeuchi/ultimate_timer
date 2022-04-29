CREATE TABLE IF NOT EXISTS users (
    id varchar(36) PRIMARY KEY,
    password varchar(255) NOT NULL,
    email varchar(128) NOT NULL,
    sound_on boolean NOT NULL
);
