package models

import "fmt"

type QuestionType int

const (
	QuestionTypeText     QuestionType = 0
	QuestionTypeRadio    QuestionType = 1
	QuestionTypeCheckbox QuestionType = 2
)

type Question interface {
	Question() QuestionBase
	Validate() error
}

type QuestionBase struct {
	FormID      string
	FormVersion int
	Title       string
}

func (q QuestionBase) Question() QuestionBase {
	return q
}

func (q QuestionBase) Validate() error {
	if q.FormID == "" {
		return fmt.Errorf("form id is required")
	}

	if q.FormVersion == 0 {
		return fmt.Errorf("form version is required")
	}

	if q.Title == "" {
		return fmt.Errorf("text is required")
	}

	return nil

}

type TextQuestion struct {
	QuestionBase
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
