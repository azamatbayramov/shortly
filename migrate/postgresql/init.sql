CREATE TABLE IF NOT EXISTS links (
    id   bigserial PRIMARY KEY,
    link VARCHAR(2048) UNIQUE NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_link_id  ON links USING hash (id);
CREATE INDEX IF NOT EXISTS idx_link_url ON links USING hash (link);
