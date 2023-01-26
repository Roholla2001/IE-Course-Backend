
CREATE TABLE users (
    id BIGSERIAL,
    username VARCHAR(200),
    password VARCHAR(32)
);

CREATE TABLE urls (
    id BIGSERIAL,
    address VARCHAR(500) UNIQUE NOT NULL,
    user_id BIGINT REFERENCES users(id)
);