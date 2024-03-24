CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    email VARCHAR (50) UNIQUE NOT NULL,
    password_hash VARCHAR (100) NOT NULL
);
