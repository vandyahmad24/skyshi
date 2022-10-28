package migration

import "time"

type Todo struct {
	Id              int    `gorm:"column:id"`
	ActivityGroupId int    `gorm:"size:255"`
	Title           string `gorm:"size:255"`
	IsActive        string `gorm:"size:1"`
	Priority        string `gorm:"size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
