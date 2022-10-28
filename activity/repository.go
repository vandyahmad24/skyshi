package activity

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Activity, error)
	// CountAll() (hasil int64)
	GetById(id int) (Activity, error)
	Create(activity *Activity) (*Activity, error)
	Update(id int, activity *Activity) (*Activity, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Activity, error) {
	var allCashier []Activity
	err := r.db.Find(&allCashier).Error
	if err != nil {
		return allCashier, err
	}
	return allCashier, nil
}

// func (r *repository) CountAll() (hasil int64) {
// 	var cashier Cashier
// 	var total int64
// 	r.db.Model(&cashier).Select("count(distinct(name))").Count(&total)

// 	return total
// }

func (r *repository) GetById(id int) (Activity, error) {
	var activity Activity
	err := r.db.Model(&activity).Where("id = ?", id).Find(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *repository) Create(activity *Activity) (*Activity, error) {
	err := r.db.Create(activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *repository) Update(id int, activity *Activity) (*Activity, error) {
	var oldActivity Activity
	err := r.db.Where("id = ?", id).First(&oldActivity).Error
	if err != nil {
		return activity, err
	}

	oldActivity.Title = activity.Title
	oldActivity.Email = activity.Email
	err = r.db.Save(&oldActivity).Error
	if err != nil {
		return &oldActivity, err
	}

	return &oldActivity, nil
}

func (r *repository) Delete(id int) error {
	var activity Activity
	err := r.db.Where("id = ?", id).Delete(&activity).Error
	if err != nil {
		return err
	}
	return nil
}
