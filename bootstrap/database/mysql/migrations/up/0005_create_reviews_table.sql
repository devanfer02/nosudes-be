CREATE TABLE IF NOT EXISTS reviews (
    user_id VARCHAR(255) REFERENCES users(id),
    attraction_id VARCHAR(255) REFERENCES attractions(id),
    comment TEXT,
    photo_url VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, attraction_id)
) Engine = InnoDB