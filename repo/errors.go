package repo

import "errors"

var (
	// ErrNotFound is returned when the requested resource is not found.
	ErrNotFound = errors.New("not found")
)
