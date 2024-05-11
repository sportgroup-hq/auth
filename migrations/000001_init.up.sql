BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_credentials
(
    user_id       uuid PRIMARY KEY,
    password_hash VARCHAR(255) NOT NULL
);

COMMIT;
