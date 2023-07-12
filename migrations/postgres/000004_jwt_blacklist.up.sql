CREATE TABLE IF NOT EXISTS jwt_blacklist(
    id SERIAL PRIMARY KEY,
    token TEXT NOT NULL,
    expires_at TIMESTAMP NOT NUll
);