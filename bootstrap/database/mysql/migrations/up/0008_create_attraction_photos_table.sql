CREATE TABLE IF NOT EXISTS attraction_photos (
    attraction_id VARCHAR(255),
    photo_url TEXT NOT NULL,
    FOREIGN KEY(attraction_id) REFERENCES attractions(attraction_id) ON DELETE CASCADE
) Engine = InnoDB;