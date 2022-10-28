package migration

import "time"

type Activity struct {
	Id        int    `gorm:"column:id"`
	Email     string `gorm:"size:255"`
	Title     string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
