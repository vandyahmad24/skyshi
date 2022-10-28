package todo

import (
	"time"
)

type Todo struct {
	Id              int        `json:"id"`
	ActivityGroupId string     `json:"activity_group_id"`
	IsActive        string     `json:"is_active"`
	Title           string     `json:"title"`
	Priority        string     `json:"priority"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
}
