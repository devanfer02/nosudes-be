CREATE TABLE IF NOT EXISTS reviews (
    user_id VARCHAR(255),
    attraction_id VARCHAR(255),
    comment TEXT,
    photo_url VARCHAR(200),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, attraction_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY(attraction_id) REFERENCES attractions(attraction_id) ON DELETE CASCADE
) Engine = InnoDB