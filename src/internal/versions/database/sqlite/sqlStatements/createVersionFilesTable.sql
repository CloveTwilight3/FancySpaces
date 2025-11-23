CREATE TABLE IF NOT EXISTS version_files (
    space_id TEXT NOT NULL ,
    version_id TEXT NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    size INTEGER NOT NULL,
    PRIMARY KEY (space_id, version_id, name)
);