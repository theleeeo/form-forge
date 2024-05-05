package models

import "fmt"

type QuestionType int

const (
	maxQuestionTypeNumber int          = 2
	QuestionTypeText      QuestionType = 0
	QuestionTypeRadio     QuestionType = 1
	QuestionTypeCheckbox  QuestionType = 2
)

type Question interface {
	Question() QuestionBase
	Validate() error
}

type QuestionBase struct {
	FormID string
	Title  string
	Type   QuestionType
}

func (q QuestionBase) Validate() error {
	if q.FormID == "" {
		return fmt.Errorf("form id is required")
	}

	if q.Title == "" {
		return fmt.Errorf("text is required")
	}

	if q.Type < 0 || int(q.Type) > maxQuestionTypeNumber {
		return fmt.Errorf("invalid question type")
	}

	return nil

}

type TextQuestion struct {
	QuestionBase
}

func (q TextQuestion) Question() QuestionBase {
	return q.QuestionBase
}

func (q TextQuestion) Validate() error {
	if err := q.QuestionBase.Validate(); err != nil {
		return err
	}

	return nil
}

type RadioQuestion struct {
	QuestionBase
	Options []string
}

func (q RadioQuestion) Question() QuestionBase {
	return q.QuestionBase
}

func (q RadioQuestion) Validate() error {
	if err := q.QuestionBase.Validate(); err != nil {
		return err
	}

	if len(q.Options) == 0 {
		return fmt.Errorf("options are required for radio questions")
	}

	for _, o := range q.Options {
		if o == "" {
			return fmt.Errorf("empty option is not allowed")
		}
	}

	return nil
}

type CheckboxQuestion struct {
	QuestionBase
	Options []string
}

func (q CheckboxQuestion) Question() QuestionBase {
	return q.QuestionBase
}

func (q CheckboxQuestion) Validate() error {
	if err := q.QuestionBase.Validate(); err != nil {
		return err
	}

	if len(q.Options) == 0 {
		return fmt.Errorf("options are required for checkbox questions")
	}

	for _, o := range q.Options {
		if o == "" {
			return fmt.Errorf("empty option is not allowed")
		}
	}

	return nil
}
