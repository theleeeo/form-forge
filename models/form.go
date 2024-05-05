package models

import "time"

type Form struct {
	ID      string
	BaseID  string
	Version int

	Title     string
	CreatedAt time.Time
	CreatedBy string
}
