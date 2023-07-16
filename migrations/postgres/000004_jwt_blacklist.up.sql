CREATE TABLE IF NOT EXISTS auth_table(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    token bytea NOT NULL,
    expires_at timestamp NOT NULL,
);