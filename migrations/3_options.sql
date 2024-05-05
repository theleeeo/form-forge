CREATE TABLE IF NOT EXISTS options (
    -- The question that this option belongs to
    question_id INT NOT NULL,
    -- The order of the option in the question
    order_idx INT NOT NULL,
    -- The text of the option
    option_text TEXT,
    FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE,
    PRIMARY KEY (question_id, order_idx)
);
