CREATE TABLE IF NOT EXISTS questions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    form_version_id VARCHAR(36) NOT NULL,
    -- The order of the question in the form
    order_idx INT NOT NULL,
    title TEXT NOT NULL,
    -- The type of question, used to determine how to display and handle the question
    question_type INT NOT NULL,
    FOREIGN KEY (form_version_id) REFERENCES forms(version_id) ON DELETE CASCADE,
    UNIQUE (form_version_id, order_idx)
);
