
CREATE TABLE userModel (
    id BIGSERIAL,
    username VARCHAR(200) UNIQUE,
    password VARCHAR(32)
);

CREATE TABLE url (
    id BIGSERIAL,
    url VARCHAR(500) UNIQUE NOT NULL,
    user_id BIGINT REFERENCES users(id),
    success_count BIGINT,
    fail_count BIGINT
);