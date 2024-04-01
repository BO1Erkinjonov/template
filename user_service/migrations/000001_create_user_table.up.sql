CREATE TABLE users
(
    first_name   VARCHAR(30),
    last_name    VARCHAR(30),
    username     VARCHAR(30),
    role VARCHAR(30),
    password     TEXT,
    email        TEXT,
    id           UUID PRIMARY KEY,
    refreshtoken TEXT,
    created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP
)