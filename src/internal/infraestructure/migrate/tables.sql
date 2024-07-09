-- CREATE DATABASE IF NOT EXISTS unisync;
-- USE unisync;

-- DROP TABLE IF EXISTS users;

-- CREATE TABLE users(
--     id int AUTO_INCREMENT PRIMARY KEY,
--     name varchar(50) not null,
--     username varchar(50) not null unique,
--     email varchar(50) not null unique,
--     password varchar(50) not null unique,
--      timestamp default CURRENT_TIMESTAMP()
-- )ENGINE=INNODB;