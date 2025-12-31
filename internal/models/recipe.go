package models

import "github.com/google/uuid"

type Recipe struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	TimeMin int       `json:"time_min"`
	Tags    []string  `json:"tags"`
}
