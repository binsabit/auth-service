CREATE TABLE IF NOT EXISTS auth_table(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    auth_id TEXT NOT NULL,
    expires_at TIMESTAMP NOT NUll,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
            REFERENCES users(id)
                ON DELETE CASCADE
);