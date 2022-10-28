package todo

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll(c *gin.Context) ([]Todo, error)
	GetById(id int) (Todo, error)
	CreateActivity(input *InputTodo) (*Todo, error)
	UpdateTodo(id int, input *InputTodo) (*Todo, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll(c *gin.Context) ([]Todo, error) {
	activity_group_id := c.Query("activity_group_id")
	activity, err := s.repository.GetAll(activity_group_id)
	if err != nil {
		return activity, err
	}

	return activity, err
}

func (s *service) GetById(id int) (Todo, error) {
	activity, err := s.repository.GetById(id)
	if err != nil {
		return activity, err
	}
	if activity.Id == 0 {
		return activity, errors.New(fmt.Sprintf("Todo with ID %d Not Found", id))
	}
	return activity, err
}

func (s *service) CreateActivity(input *InputTodo) (*Todo, error) {

	cashier := Todo{
		Title:           input.Title,
		ActivityGroupId: input.ActivityGroupId,
		IsActive:        "1",
		Priority:        "very-high",
	}
	newCashier, err := s.repository.Create(&cashier)
	if err != nil {
		return nil, err
	}

	return newCashier, nil
}

func (s *service) UpdateTodo(id int, input *InputTodo) (*Todo, error) {
	cashier := Todo{
		Title:           input.Title,
		ActivityGroupId: input.ActivityGroupId,
	}
	newCashier, err := s.repository.Update(id, &cashier)
	if err != nil {
		return nil, err
	}

	return newCashier, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
