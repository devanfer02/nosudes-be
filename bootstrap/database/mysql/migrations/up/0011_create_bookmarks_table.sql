CREATE TABLE IF NOT EXISTS bookmarks (
    attraction_id VARCHAR(255) REFERENCES attractions(attraction_id),
    user_id VARCHAR(255) REFERENCES users(user_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) Engine = InnoDB;