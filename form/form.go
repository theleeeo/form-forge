package form

import (
	"time"
)

type FormBase struct {
	ID      string
	Version int

	Title     string
	CreatedAt time.Time
	CreatedBy string
}

type Form struct {
	FormBase
	Questions []Question
}
