package app

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/form"
	"github.com/theleeeo/form-forge/models"
)

func (t *TestSuiteRepo) TestCreateForm() {
	t.Run("Create a new form", func() {
		// Create a new form
		f, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
			Title: "Test Form",
			Questions: []form.CreateQuestionParams{
				{Type: models.QuestionTypeText, Title: "Text question"},
				{Type: models.QuestionTypeRadio, Title: "Radio question", Options: []string{"Option 1", "Option 2"}},
				{Type: models.QuestionTypeCheckbox, Title: "Checkbox question", Options: []string{"Option 1", "Option 2"}},
			},
		})
		t.NoError(err)
		t.NotNil(f)

		t.Equal("Test Form", f.Title)
		t.Equal(1, f.Version)
		t.NoError(uuid.Validate(f.ID))
		q, err := f.Questions(context.Background())
		t.NoError(err)
		t.Len(q, 3)
		t.Equal(models.Question(
			models.TextQuestion{
				QuestionBase: models.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Text question",
				},
			},
		), q[0])
		t.Equal(models.Question(
			models.RadioQuestion{
				QuestionBase: models.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Radio question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), q[1])
		t.Equal(models.Question(
			models.CheckboxQuestion{
				QuestionBase: models.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Checkbox question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), q[2])
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
			{Type: models.QuestionTypeText, Title: "Text question"},
			{Type: models.QuestionTypeRadio, Title: "Radio question", Options: []string{"Option 1", "Option 2"}},
			{Type: models.QuestionTypeCheckbox, Title: "Checkbox question", Options: []string{"Option 1", "Option 2"}},
		},
	})
	t.NoError(err)
	t.NoError(err)

	t.Run("Get a form", func() {
		// Get the form
		f2, err := t.app.GetForm(context.Background(), f.ID)
		t.NoError(err)

		_, err = f2.Questions(context.Background())
		t.NoError(err)

		t.Equal("Test Form", f.Title)
		t.Equal(1, f.Version)
		t.NoError(uuid.Validate(f.ID))
		t.Equal("00000000-0000-0000-0000-000000000001", f.ID)
		q, err := f.Questions(context.Background())
		t.NoError(err)
		t.Equal(f, f2)
		t.Equal(time.Unix(6000, 0).UTC(), f2.CreatedAt)
		t.Len(q, 3)
		t.Equal(models.Question(
			models.TextQuestion{
				QuestionBase: models.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Text question",
				},
			},
		), q[0])
		t.Equal(models.Question(
			models.RadioQuestion{
				QuestionBase: models.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Radio question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), q[1])
		t.Equal(models.Question(
			models.CheckboxQuestion{
				QuestionBase: models.QuestionBase{
					FormID:      f.ID,
					FormVersion: 1,
					Title:       "Checkbox question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), q[2])
	})
}
