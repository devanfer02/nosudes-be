CREATE TABLE IF NOT EXISTS reviews (
    review_id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255),
    attraction_id VARCHAR(255),
    review_text TEXT,
    photo_url VARCHAR(255),
    date_created VARCHAR(255),
    UNIQUE(user_id, attraction_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    FOREIGN KEY(attraction_id) REFERENCES attractions(attraction_id) ON DELETE CASCADE
) Engine = InnoDB