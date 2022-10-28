package todo

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(param string) ([]Todo, error)
	GetById(id int) (Todo, error)
	Create(todo *Todo) (*Todo, error)
	Update(id int, todo *Todo) (*Todo, error)
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

func (r *repository) GetAll(param string) ([]Todo, error) {
	var allTodo []Todo

	query := r.db.Debug().Model(&Todo{})
	if param != "" {
		query.Where("todos.activity_group_id = ? ", param)
	}
	err := query.Scan(&allTodo).Error
	if err != nil {
		return allTodo, err
	}
	return allTodo, nil
}

func (r *repository) GetById(id int) (Todo, error) {
	var activity Todo
	err := r.db.Model(&activity).Where("todos.id = ?", id).Find(&activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *repository) Create(activity *Todo) (*Todo, error) {
	err := r.db.Create(activity).Error
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (r *repository) Update(id int, activity *Todo) (*Todo, error) {
	var oldActivity Todo
	err := r.db.Where("id = ?", id).First(&oldActivity).Error
	if err != nil {
		return activity, err
	}

	oldActivity.Title = activity.Title
	oldActivity.ActivityGroupId = activity.ActivityGroupId
	err = r.db.Save(&oldActivity).Error
	if err != nil {
		return &oldActivity, err
	}

	return &oldActivity, nil
}

func (r *repository) Delete(id int) error {
	var activity Todo
	err := r.db.Where("id = ?", id).Delete(&activity).Error
	if err != nil {
		return err
	}
	return nil
}
