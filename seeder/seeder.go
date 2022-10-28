package seeder

import (
	"fmt"
	"vandyahmad/skyshi/activity"
	"vandyahmad/skyshi/migration"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Seeder interface {
	CashierSeeder() error
}

type seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *seeder {
	return &seeder{
		db: db,
	}
}

func (s *seeder) ActivitySeeder() error {
	fmt.Println("activity cashier")
	var count int64
	var oneActivity activity.Activity
	var allActivity []activity.Activity
	err := s.db.Model(&oneActivity).Select("count(id)").Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		for i := 1; i <= 11; i++ {
			oneActivity.Title = fmt.Sprintf("coba %d", i)
			oneActivity.Email = uuid.New().String() + "@skyshi.com"

			allActivity = append(allActivity, oneActivity)
		}
		err = s.db.Create(&allActivity).Error
		if err != nil {
			return err
		}

	}

	fmt.Println("selesai seeder activity")

	return nil
}

func (s *seeder) TodoSeeder() error {
	fmt.Println("todo cashier")
	var count int64
	var oneActivity migration.Todo
	var allActivity []migration.Todo
	err := s.db.Model(&oneActivity).Select("count(id)").Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		for i := 1; i <= 5; i++ {
			oneActivity = migration.Todo{
				ActivityGroupId: i,
				Title:           uuid.NewString() + " title",
				IsActive:        "1",
				Priority:        "very-hight",
			}

			allActivity = append(allActivity, oneActivity)
		}
		err = s.db.Create(&allActivity).Error
		if err != nil {
			return err
		}

	}

	fmt.Println("selesai seeder todo")

	return nil
}
