CREATE TABLE IF NOT EXISTS articles (
    article_id VARCHAR(255) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    summary VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    photo TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) Engine = InnoDB