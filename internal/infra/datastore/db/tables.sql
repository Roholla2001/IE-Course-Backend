CREATE TABLE user_model (
    id BIGSERIAL PRIMARY KEY,
    user_name VARCHAR(200) UNIQUE,
    password VARCHAR(500)
);

CREATE TABLE url_model (
    id BIGSERIAL,
    url VARCHAR(500) UNIQUE NOT NULL,
    user_id BIGINT REFERENCES user_model (id) ON DELETE CASCADE,
    success_count BIGINT,
    fail_count BIGINT
);