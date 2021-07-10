package models

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"isDone" gorm:"type:tinyint(1) NOT NULL DEFAULT 0"`
	CreatedAt   time.Time `json:"createdAt,omitempty" gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty" gorm:"type:timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP"`
}
