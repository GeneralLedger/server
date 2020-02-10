CREATE TABLE IF NOT EXISTS users(
    id            serial        PRIMARY KEY,
    created_at    timestamptz   NOT NULL DEFAULT CURRENT_TIMESTAMP
);
