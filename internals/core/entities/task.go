package entities

import "time"

type Task struct {
	ID          *int       `json:"id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
