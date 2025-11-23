CREATE TABLE IF NOT EXISTS versions (
    space_id TEXT NOT NULL ,
    id TEXT NOT NULL,
    name TEXT NOT NULL,
    channel TEXT NOT NULL,
    published_at DATETIME NOT NULL,
    changelog TEXT,
    supported_platform_versions TEXT,
    downloads INTEGER NOT NULL DEFAULT 0,
    PRIMARY KEY (space_id, id)
);