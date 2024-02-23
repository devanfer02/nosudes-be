CREATE TABLE IF NOT EXISTS operation_hours (
    op_hour_id  VARCHAR(255) PRIMARY KEY,
    attraction_id VARCHAR(255) REFERENCES attraction(attraction_id),
    day VARCHAR(255) NOT NULL,
    timespan VARCHAR(255) NOT NULL  
) Engine = InnoDB;