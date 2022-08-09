package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	IsDone    bool               `json:"is_done" validate:"required"`
	Assignee  string             `json:"assignee,omitempty" validate:"required"`
	DueDate   time.Time          `json:"due_date,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" validate:"required"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" validate:"required"`
}
