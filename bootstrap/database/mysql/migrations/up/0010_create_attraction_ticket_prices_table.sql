CREATE TABLE IF NOT EXISTS attraction_ticket_prices (
    attraction_id VARCHAR(255),
    price   INTEGER NOT NULL,
    day_type VARCHAR(100) NOT NULL,
    age_group VARCHAR(100) NOT NULL,
    FOREIGN KEY(attraction_id) REFERENCES attractions(attraction_id) ON DELETE CASCADE
) Engine = InnoDB;