package models

import "time"

type Form struct {
	ID      string
	Version int

	Title     string
	CreatedAt time.Time
	CreatedBy string
}
