package templater

import (
	"bytes"
	"context"
	"html/template"

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
	questions := make([]expandedQuestion, 0, len(f.Questions))
	for questionOrder, q := range f.Questions {
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
			Title:   questionBase.Title,
			Type:    qType,
			Options: expOptions,
			Order:   questionOrder,
		})
	}

	return expandedForm{
		ID:        f.Id,
		Title:     f.Title,
		Questions: questions,
	}, nil
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
