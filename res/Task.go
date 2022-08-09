package res

import (
	"time"
)

type CreateTaskBody struct {
	Assignee string    `json:"assignee" validate:"required"`
	DueDate  time.Time `json:"due_date" validate:"required"`
	Title    string    `json:"title" validate:"required"`
}

type EditTaskBody struct {
	ID string `json:"id" validate:"required"`
	CreateTaskBody
}
