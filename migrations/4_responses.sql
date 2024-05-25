-- A response is a submission of a form, it contains multiple answers
CREATE TABLE IF NOT EXISTS responses (
    id VARCHAR(36) PRIMARY KEY,
    -- The form that this response is for (not the base form)
    form_version_id VARCHAR(36) NOT NULL,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (form_version_id) REFERENCES forms(version_id) ON DELETE CASCADE
);
