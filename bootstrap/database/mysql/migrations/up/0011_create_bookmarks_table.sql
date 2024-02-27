CREATE TABLE IF NOT EXISTS bookmarks (
    attraction_id VARCHAR(255),
    user_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(attraction_id) REFERENCES attractions(attraction_id) ON DELETE CASCADE,
    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
) Engine = InnoDB;