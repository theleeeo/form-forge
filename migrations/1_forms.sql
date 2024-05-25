CREATE TABLE IF NOT EXISTS forms (
    -- The id of the form, it is the same for all versions of the form
    id VARCHAR(36),
    --  The id of the version of the form, it is unique for each version of the form
    version_id VARCHAR(36),
    -- The version of the form, incremented when the form is updated
    version INT NOT NULL DEFAULT 1,
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (version_id),
    UNIQUE (id, version)
);
