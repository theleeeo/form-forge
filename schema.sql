CREATE TABLE IF NOT EXISTS forms (
    -- The base id of the form, it is the same for all versions of the form
    base_id UUID NOT NULL,
    --  The id of this version of the form, it is unique for each version of the form
    version_id UUID NOT NULL PRIMARY KEY,
    -- The version of the form, incremented when the form is updated
    version INT NOT NULL,
    title VARCHAR(255),
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    UNIQUE (base_id, version)
);

CREATE TABLE IF NOT EXISTS questions (
    id UUID PRIMARY KEY,
    form_version_id UUID NOT NULL REFERENCES forms(version_id) ON DELETE CASCADE,
    -- The order of the question in the form
    order_idx INT NOT NULL,
    title TEXT NOT NULL,
    -- The type of question, used to determine how to display and handle the question
    question_type INT NOT NULL,
    UNIQUE (form_version_id, order_idx)
);

CREATE TABLE IF NOT EXISTS options (
    -- The question that this option belongs to
    question_id UUID NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    -- The order of the option in the question
    order_idx INT NOT NULL,
    -- The text of the option
    option_text TEXT,
    PRIMARY KEY (question_id, order_idx)
);

-- A response is a submission of a form, it contains multiple answers
CREATE TABLE IF NOT EXISTS responses (
    id UUID PRIMARY KEY,
    -- The form that this response is for (not the base form)
    form_version_id UUID NOT NULL REFERENCES forms(version_id) ON DELETE CASCADE,
    submitted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- An answer is a response to a single question
CREATE TABLE IF NOT EXISTS answers (
    -- The response that this answer belongs to
    response_id UUID NOT NULL REFERENCES responses(id) ON DELETE CASCADE,
    -- The question that this answer is for
    question_id UUID NOT NULL REFERENCES questions(id) ON DELETE CASCADE,
    -- The answer to the question. The type of this field depends on the question type
    -- For example, a multiple choice question would have the option number here
    answer_text TEXT
);
-- Indexes?
