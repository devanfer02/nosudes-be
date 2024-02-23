CREATE TABLE IF NOT EXISTS attraction_photos (
    attraction_id VARCHAR(255) REFERENCES attractions(attraction_id) ON DELETE CASCADE,
    photo_url TEXT NOT NULL
) Engine = InnoDB;