-- An answer is a response to a single question
CREATE TABLE IF NOT EXISTS answers (
    -- The response that this answer belongs to
    response_id VARCHAR(36) NOT NULL,
    -- The question that this answer is for
    question_order INT NOT NULL,
    -- The answer to the question. The type of this field depends on the question type
    -- For example, a multiple choice question would have the option number here
    answer_text TEXT,
    FOREIGN KEY (response_id) REFERENCES responses(id) ON DELETE CASCADE
);
