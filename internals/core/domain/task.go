package domain

import (
	"errors"
	"strings"
)

type Priority int

const (
	Attended Priority = 0
	Low      Priority = 1
	Medium   Priority = 2
	High     Priority = 3
	Urgent   Priority = 4
	Critical Priority = 5
)

type Task struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Priority    Priority `json:"priority"`
}

func NewTask(title, description string, priority Priority) (*Task, error) {
	title = strings.TrimSpace(title)
	description = strings.TrimSpace(description)

	if title == "" {
		return nil, errors.New("el título no puede estar vacío")
	}

	if priority < Attended || priority > Critical {
		return nil, errors.New("prioridad inválida")
	}

	return &Task{
		Title:       title,
		Description: description,
		Priority:    priority,
	}, nil
}
