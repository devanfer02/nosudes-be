CREATE TABLE IF NOT EXISTS reviews (
    review_id   VARCHAR(255) PRIMARY KEY,
    attraction_id VARCHAR(255),
    user_id VARCHAR(255),
    review_text TEXT,
    photo_url VARCHAR(255),
    date_created VARCHAR(255)
) Engine = InnoDB;