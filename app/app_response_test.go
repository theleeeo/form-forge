package app

import (
	"context"

	"github.com/theleeeo/form-forge/form"
)

func (t *TestSuiteRepo) Test_SubmitResponse() {
	// Create a new form
	f, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
		Title: "Test Form",
		Questions: []form.CreateQuestionParams{
			{Type: form.QuestionTypeText, Title: "Text question"},
			{Type: form.QuestionTypeRadio, Title: "Radio question", Options: []string{"Option 1", "Option 2"}},
			{Type: form.QuestionTypeCheckbox, Title: "Checkbox question", Options: []string{"Option 1", "Option 2"}},
		},
	})
	t.NoError(err)

	t.Run("Form not found", func() {
		err := t.app.SubmitResponse(context.Background(), "00000000-0000-123-0000-000000000000", map[string][]string{})
		t.ErrorIs(err, ErrFormNotFound)

	})

	t.Run("Successful submit", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"0": {"An answer"},
			"1": {"1"},
		})
		t.NoError(err)
	})

	t.Run("Out of bound key", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"-1": {"An answer", "Another answer"},
		})
		t.Error(err)

		err = t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"100": {"An answer", "Another answer"},
		})
		t.Error(err)
	})

	t.Run("Non-int key", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"hello": {"An answer"},
		})
		t.Error(err)
	})

	t.Run("Text, multiple values", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"0": {"An answer", "Another answer"},
		})
		t.Error(err)
	})

	t.Run("Radio, Multiple values", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"1": {"0", "1"},
		})
		t.Error(err)
	})

	t.Run("Radio, non-int value", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"1": {"hello"},
		})
		t.Error(err)
	})

	t.Run("Out of bound radio, negative", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"1": {"-1"},
		})
		t.Error(err)
	})

	t.Run("Out of bound radio, too big", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"1": {"0", "7"},
		})
		t.Error(err)
	})

	t.Run("Out of bound checkbox, negative", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"2": {"-1"},
		})
		t.Error(err)
	})

	t.Run("Out of bound checkbox, too big", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"2": {"0", "7"},
		})
		t.Error(err)
	})

	t.Run("Checkbox, non-int value", func() {
		err := t.app.SubmitResponse(context.Background(), f.Id, map[string][]string{
			"2": {"hello"},
		})
		t.Error(err)
	})
}
