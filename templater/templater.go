package templater

import (
	"bytes"
	"context"
	"html/template"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/form"
)

// type Config struct {
// TemplateDir string
// }

func New() *Templater {
	return &Templater{}
}

type Templater struct {
}

type expandedForm struct {
	ID        uuid.UUID
	Title     string
	Questions []expandedQuestion
}

type expandedQuestion struct {
	Id      uuid.UUID
	Order   int
	Type    string
	Title   string
	Options []expandedOption
}

type expandedOption struct {
	Label string
	Order int
}

func constructExpandedForm(f form.Form, qs []form.Question) expandedForm {
	questions := make([]expandedQuestion, 0, len(qs))
	for questionOrder, q := range qs {
		questionBase := q.Question()

		var options []string
		var qType string
		switch q := q.(type) {
		case form.RadioQuestion:
			options = q.Options
			qType = "radio"
		case form.CheckboxQuestion:
			options = q.Options
			qType = "checkbox"
		case form.TextQuestion:
			qType = "text"
		}

		expOptions := make([]expandedOption, 0, len(options))
		for j, o := range options {
			expOptions = append(expOptions, expandedOption{
				Label: o,
				Order: j,
			})
		}

		questions = append(questions, expandedQuestion{
			Id:      questionBase.Id,
			Title:   questionBase.Title,
			Type:    qType,
			Options: expOptions,
			Order:   questionOrder,
		})
	}

	return expandedForm{
		ID:        f.BaseId,
		Title:     f.Title,
		Questions: questions,
	}
}

func (t *Templater) Generate(ctx context.Context, f form.Form, qs []form.Question) ([]byte, error) {
	template := template.Must(template.New("test").ParseFiles("templates/test.html"))
	template = template.Lookup("test.html")

	var tpl bytes.Buffer
	if err := template.Execute(&tpl, constructExpandedForm(f, qs)); err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}
