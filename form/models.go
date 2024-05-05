package form

import "go.leeeo.se/form-forge/models"

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
