package app

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/theleeeo/form-forge/form"
)

func (t *TestSuiteRepo) Test_CreateForm() {
	t.Run("Create a new form", func() {
		lastDigit := 0
		form.UUIDNew = func() uuid.UUID {
			lastDigit++
			return uuid.MustParse(fmt.Sprintf("00000000-0000-0000-0000-%012d", lastDigit))
		}

		// Create a new form
		f, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
			Title:       "Test Form",
			Description: "This is a test form",
			Questions: []form.CreateQuestionParams{
				{Type: form.QuestionTypeText, Title: "Text question"},
				{Type: form.QuestionTypeRadio, Title: "Radio question", Options: []string{"Option 1", "Option 2"}},
				{Type: form.QuestionTypeCheckbox, Title: "Checkbox question", Options: []string{"Option 1", "Option 2"}},
			},
		})

		t.NoError(err)
		t.NotNil(f)

		t.Equal("Test Form", f.Title)
		t.Equal("This is a test form", f.Description)
		t.Equal(uint32(1), f.Version)
		t.Equal(f.BaseId, uuid.MustParse("00000000-0000-0000-0000-000000000001"))
		t.Equal(f.VersionId, uuid.MustParse("00000000-0000-0000-0000-000000000002"))
		t.Len(f.Questions, 3)
		t.Equal(form.Question(
			form.TextQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					Title: "Text question",
				},
			},
		), f.Questions[0])
		t.Equal(form.Question(
			form.RadioQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000004"),
					Title: "Radio question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[1])
		t.Equal(form.Question(
			form.CheckboxQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000005"),
					Title: "Checkbox question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[2])
	})
}

func (t *TestSuiteRepo) Test_GetForm() {
	form.TimeNow = func() time.Time {
		return time.Unix(6000, 0)
	}
	lastDigit := 0
	form.UUIDNew = func() uuid.UUID {
		lastDigit++
		return uuid.MustParse(fmt.Sprintf("00000000-0000-0000-0000-%012d", lastDigit))
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

	t.Run("Get a form", func() {
		// Get the form
		f2, err := t.app.GetForm(context.Background(), f.BaseId)
		t.NoError(err)

		t.Equal("Test Form", f.Title)
		t.Equal(uint32(1), f.Version)
		t.Equal(uuid.MustParse("00000000-0000-0000-0000-000000000001"), f.BaseId)
		t.Equal(uuid.MustParse("00000000-0000-0000-0000-000000000002"), f.VersionId)
		t.Equal(f, f2)
		t.Equal(time.Unix(6000, 0).UTC(), f2.CreatedAt)
		t.Len(f.Questions, 3)
		t.Equal(form.Question(
			form.TextQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					Title: "Text question",
				},
			},
		), f.Questions[0])
		t.Equal(form.Question(
			form.RadioQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000004"),
					Title: "Radio question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[1])
		t.Equal(form.Question(
			form.CheckboxQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000005"),
					Title: "Checkbox question",
				},
				Options: []string{"Option 1", "Option 2"},
			},
		), f.Questions[2])
	})
}

func (t *TestSuiteRepo) Test_ListForms() {
	lastTime := time.Unix(6000, 0)
	form.TimeNow = func() time.Time {
		lastTime = lastTime.Add(time.Second)
		return lastTime
	}
	lastDigit := 0
	form.UUIDNew = func() uuid.UUID {
		lastDigit++
		return uuid.MustParse(fmt.Sprintf("00000000-0000-0000-0000-%012d", lastDigit))
	}
	// Create a new form
	f1, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
		Title: "Form 1",
		Questions: []form.CreateQuestionParams{
			{Type: form.QuestionTypeText, Title: "TQ1"},
			{Type: form.QuestionTypeText, Title: "TQ2"},
		},
	})
	t.NoError(err)

	f2, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
		Title: "Form 2",
		Questions: []form.CreateQuestionParams{
			{Type: form.QuestionTypeRadio, Title: "RQ1", Options: []string{"O1"}},
			{Type: form.QuestionTypeRadio, Title: "RQ2", Options: []string{"O1"}},
		},
	})
	t.NoError(err)

	f3, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
		Title: "Form 3",
		Questions: []form.CreateQuestionParams{
			{Type: form.QuestionTypeCheckbox, Title: "CQ1", Options: []string{"O2"}},
			{Type: form.QuestionTypeCheckbox, Title: "CQ2", Options: []string{"O1"}},
		},
	})
	t.NoError(err)

	t.Run("Only base forms", func() {
		// Get the form
		resp, err := t.app.ListForms(context.Background(), form.ListFormsParams{})
		t.NoError(err)

		t.Len(resp, 3)
		t.Equal(f3, resp[0])
		t.Equal(f2, resp[1])
		t.Equal(f1, resp[2])
	})

	uf, err := t.app.UpdateForm(context.Background(), form.UpdateFormParams{
		Id: f2.BaseId,
		CreateFormParams: form.CreateFormParams{
			Title: "Updated form",
			Questions: []form.CreateQuestionParams{
				{Type: form.QuestionTypeCheckbox, Title: "CQ1", Options: []string{"O2"}},
				{Type: form.QuestionTypeCheckbox, Title: "CQ2", Options: []string{"O1"}},
			},
		},
	})
	t.NoError(err)

	t.Run("One updated form", func() {
		// Get the form
		resp, err := t.app.ListForms(context.Background(), form.ListFormsParams{})
		t.NoError(err)

		t.Len(resp, 3)
		t.Equal(uf, resp[0])
		t.Equal(f3, resp[1])
		t.Equal(f1, resp[2])
	})
}

func (t *TestSuiteRepo) Test_UpdateForms() {
	lastTime := time.Unix(6000, 0)
	form.TimeNow = func() time.Time {
		lastTime = lastTime.Add(time.Second)
		return lastTime
	}
	lastDigit := 0
	form.UUIDNew = func() uuid.UUID {
		lastDigit++
		return uuid.MustParse(fmt.Sprintf("00000000-0000-0000-0000-%012d", lastDigit))
	}

	f1, err := t.app.CreateNewForm(context.Background(), form.CreateFormParams{
		Title: "Test Form",
		Questions: []form.CreateQuestionParams{
			{Type: form.QuestionTypeText, Title: "TQ1"},
			{Type: form.QuestionTypeText, Title: "TQ2"},
		},
	})
	t.NoError(err)

	_, err = t.app.CreateNewForm(context.Background(), form.CreateFormParams{
		Title: "Test Form",
		Questions: []form.CreateQuestionParams{
			{Type: form.QuestionTypeRadio, Title: "RQ1", Options: []string{"O1"}},
			{Type: form.QuestionTypeRadio, Title: "RQ2", Options: []string{"O1"}},
		},
	})
	t.NoError(err)

	t.Run("Form not found", func() {
		_, err := t.app.UpdateForm(context.Background(), form.UpdateFormParams{
			Id: uuid.MustParse("00000000-0000-1234-0000-000000000000"),
			CreateFormParams: form.CreateFormParams{
				Title: "Test Form",
				Questions: []form.CreateQuestionParams{
					{Type: form.QuestionTypeText, Title: "TQ1"},
					{Type: form.QuestionTypeText, Title: "TQ2"},
				},
			},
		})
		t.Error(err)
		t.ErrorIs(err, ErrFormNotFound)
	})

	t.Run("Update successful", func() {
		uf, err := t.app.UpdateForm(context.Background(), form.UpdateFormParams{
			Id: f1.BaseId,
			CreateFormParams: form.CreateFormParams{
				Title: "Test Form 2",
				Questions: []form.CreateQuestionParams{
					{Type: form.QuestionTypeRadio, Title: "TQ1", Options: []string{"O1"}},
				},
			},
		})
		t.NoError(err)

		t.Equal("Test Form 2", uf.Title)
		t.Equal(uint32(2), uf.Version)
		t.Equal(f1.BaseId, uf.BaseId)
		t.Len(uf.Questions, 1)
		t.Equal(form.Question(
			form.RadioQuestion{
				QuestionBase: form.QuestionBase{
					Id:    uuid.MustParse("00000000-0000-0000-0000-000000000011"),
					Title: "TQ1",
				},
				Options: []string{"O1"},
			},
		), uf.Questions[0])
	})
}
