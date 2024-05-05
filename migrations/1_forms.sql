CREATE TABLE IF NOT EXISTS forms (
    id VARCHAR(36) PRIMARY KEY,
    -- The base_id is the id of the form that is version 1 of this form
    base_id VARCHAR(36),
    -- The version of the form, incremented when the form is updated
    version INT NOT NULL DEFAULT 1,
    title VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(36)
);
