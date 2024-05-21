package templater

import (
	"bytes"
	"context"
	"html/template"

	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/models"
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
	ID        string
	Title     string
	Questions []expandedQuestion
}

type expandedQuestion struct {
	Order   int
	Type    string
	Title   string
	Options []expandedOption
}

type expandedOption struct {
	Label string
	Order int
}

func ResolveForm(ctx context.Context, f form.Form) (expandedForm, error) {
	// Resolve the questions
	qs, err := f.Questions(ctx)
	if err != nil {
		return expandedForm{}, err
	}

	questions := make([]expandedQuestion, 0, len(qs))
	for i, q := range qs {
		qb := q.Question()

		var options []string
		var qType string
		switch q := q.(type) {
		case models.RadioQuestion:
			options = q.Options
			qType = "radio"
		case models.CheckboxQuestion:
			options = q.Options
			qType = "checkbox"
		case models.TextQuestion:
			qType = "text"
		}

		expOptions := make([]expandedOption, 0, len(options))
		for j, o := range options {
			expOptions = append(expOptions, expandedOption{
				Label: o,
				Order: j,
			})
		}

		cq := expandedQuestion{
			Title:   qb.Title,
			Type:    qType,
			Options: expOptions,
			Order:   i,
		}

		questions = append(questions, cq)
	}

	fr := expandedForm{
		ID:        f.ID,
		Title:     f.Title,
		Questions: questions,
	}

	return fr, nil
}

func (t *Templater) Generate(ctx context.Context, f form.Form) ([]byte, error) {
	template := template.Must(template.New("test").ParseFiles("templates/test.html"))
	template = template.Lookup("test.html")

	fr, err := ResolveForm(ctx, f)
	if err != nil {
		return nil, err
	}

	var tpl bytes.Buffer
	if err := template.Execute(&tpl, fr); err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}
