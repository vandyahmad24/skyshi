package activity

import (
	"time"
)

type Activity struct {
	Id        int        `json:"id" gorm:"id"`
	Email     string     `json:"email"  gorm:"email"`
	Title     string     `json:"title"  gorm:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
