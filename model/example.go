package model

import "time"

// Example represents... insert a comment about the model here
type Example struct {
	ID            int64
	Name          string
	Useful        bool
	CreatedAt     time.Time
	DeactivatedAt time.Time
}
