package form

import "github.com/theleeeo/form-forge/models"

type CreateFormParams struct {
	Title     string
	Questions []CreateQuestionParams
}

type CreateQuestionParams struct {
	Type  models.QuestionType
	Title string
	// Options is only required for radio and checkbox questions
	Options []string
}
