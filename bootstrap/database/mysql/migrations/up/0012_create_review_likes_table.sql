CREATE TABLE IF NOT EXISTS review_likes (
    review_id   VARCHAR(255),
    user_id     VARCHAR(255),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(review_id, user_id),
    FOREIGN KEY(review_id) REFERENCES reviews(review_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
) Engine = InnoDB;