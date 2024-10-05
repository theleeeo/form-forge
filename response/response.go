package response

import (
	"time"

	"github.com/google/uuid"
)

type Response struct {
	// ID is the unique identifier of the response.
	Id uuid.UUID
	// FormId is the unique identifier of the form this response is for.
	FormVersionId uuid.UUID

	// Answers is the list of answers to the questions in the form.
	Answers []Answer

	SubmittedAt time.Time
}

type Answer interface {
	Question() uuid.UUID
}

type AnswerBase struct {
	// QuestionId is the unique identifier of the question this answer is for.
	QuestionId uuid.UUID
}

func (a AnswerBase) Question() uuid.UUID {
	return a.QuestionId
}

type CheckboxAnswer struct {
	AnswerBase
	// Values is the list of selected options.
	Values []int
}

type RadioAnswer struct {
	AnswerBase
	// Value is the selected option.
	Value int
}

type TextAnswer struct {
	AnswerBase
	// Value is the text answer.
	Value string
}
