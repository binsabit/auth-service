CREATE TABLE IF NOT EXISTS otps(
    id SERIAL,
    phone TEXT NOT NULL,
    code TEXT NOT NULL, 
    expires_at TIMESTAMP 
);