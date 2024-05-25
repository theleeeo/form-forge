package form

import (
	"time"
)

type FormBase struct {
	Id        string
	VersionId string
	Version   uint32

	Title     string
	CreatedAt time.Time
	CreatedBy string
}

type Form struct {
	FormBase
	Questions []Question
}
