CREATE TABLE IF NOT EXISTS users (
    user_id     VARCHAR(255) PRIMARY KEY,
    fullname    VARCHAR(255) NOT NULL,
    username    VARCHAR(150) NOT NULL UNIQUE,
    email       VARCHAR(150) NOT NULL UNIQUE,
    password    VARCHAR(255) NOT NULL,
    photo_url   VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) Engine = InnoDB
