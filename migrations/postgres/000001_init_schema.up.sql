CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    phone TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    password bytea NOT NULL
);

