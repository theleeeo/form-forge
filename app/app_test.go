package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/form"
)

func (t *TestSuiteRepo) TestCreateForm() {
	t.Run("Create a new form", func() {
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
		t.NotNil(f)

		t.Equal("Test Form", f.Title)
		t.Equal(1, f.Version)
		t.NoError(uuid.Validate(f.ID))
		t.Len(f.Questions, 3)
		t.Equal(form.Question(
			form.TextQuestion{
				QuestionBase: form.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Text question",
				},
			},
		), f.Questions[0])
		t.Equal(form.Question(
			form.RadioQuestion{
				QuestionBase: form.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Radio question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[1])
		t.Equal(form.Question(
			form.CheckboxQuestion{
				QuestionBase: form.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Checkbox question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[2])
	})
}

func (t *TestSuiteRepo) TestGetForm() {
	form.TimeNow = func() time.Time {
		return time.Unix(6000, 0)
	}
	form.UUIDNew = func() uuid.UUID {
		return uuid.MustParse("00000000-0000-0000-0000-000000000001")
	}
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
	t.NoError(err)

	t.Run("Get a form", func() {
		// Get the form
		f2, err := t.app.GetForm(context.Background(), f.ID)
		t.NoError(err)

		t.Equal("Test Form", f.Title)
		t.Equal(1, f.Version)
		t.NoError(uuid.Validate(f.ID))
		t.Equal("00000000-0000-0000-0000-000000000001", f.ID)
		t.Equal(f, f2)
		t.Equal(time.Unix(6000, 0).UTC(), f2.CreatedAt)
		t.Len(f.Questions, 3)
		t.Equal(form.Question(
			form.TextQuestion{
				QuestionBase: form.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Text question",
				},
			},
		), f.Questions[0])
		t.Equal(form.Question(
			form.RadioQuestion{
				QuestionBase: form.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Radio question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[1])
		t.Equal(form.Question(
			form.CheckboxQuestion{
				QuestionBase: form.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Checkbox question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[2])
	})
}
