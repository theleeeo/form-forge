CREATE TABLE IF NOT EXISTS forms (
    id VARCHAR(36),
    -- The version of the form, incremented when the form is updated
    version INT NOT NULL DEFAULT 1,
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(36),
    PRIMARY KEY (id, version)
);
