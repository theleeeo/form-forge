-- A response is a submission of a form, it contains multiple answers
CREATE TABLE IF NOT EXISTS responses (
    id VARCHAR(36) PRIMARY KEY,
    -- The form that this response is for (not the base form)
    form_id VARCHAR(36) NOT NULL,
    form_version INT NOT NULL,
    -- The user that submitted this response
    -- This can be null if the response is anonymous
    user_id VARCHAR(36),
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (form_id, form_version) REFERENCES forms(id, version) ON DELETE CASCADE
);
