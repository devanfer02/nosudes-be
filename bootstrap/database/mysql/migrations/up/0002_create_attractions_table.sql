CREATE TABLE IF NOT EXISTS attractions (
    attraction_id   VARCHAR(255) PRIMARY KEY, 
    name            VARCHAR(150) NOT NULL, 
    category_id     INTEGER REFERENCES attraction_categories(id),
    description     TEXT NOT NULL,
    opening_hours   VARCHAR(200) NOT NULL,
    maps_embed_url  VARCHAR(200) NOT NULL, 
    location        VARCHAR(250) NOT NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) Engine = InnoDB