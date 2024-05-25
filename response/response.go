package response

import "time"

type Response struct {
	// ID is the unique identifier of the response.
	Id string
	// FormId is the unique identifier of the form this response is for.
	FormVersionId string

	// Answers is the list of answers to the questions in the form.
	Answers []Answer

	SubmittedAt time.Time
}

type Answer interface {
	Question() int
}

type AnswerBase struct {
	// QuestionOrder is the order of the question this answer is for.
	QuestionOrder int
}

func (a AnswerBase) Question() int {
	return a.QuestionOrder
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
