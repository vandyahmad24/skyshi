package activity

import (
	"errors"
	"fmt"
)

type Service interface {
	GetAll() ([]Activity, error)
	// CountAll() (hasil int64)
	GetById(id int) (Activity, error)
	CreateActivity(input *InputActivity) (*ActivityResponse, error)
	UpdateCashier(id int, input *InputActivity) (*Activity, error)
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

func (s *service) GetAll() ([]Activity, error) {
	activity, err := s.repository.GetAll()
	if err != nil {
		return activity, err
	}
	return activity, err
}

func (s *service) GetById(id int) (Activity, error) {
	activity, err := s.repository.GetById(id)
	if err != nil {
		return activity, err
	}
	if activity.Id == 0 {
		return activity, errors.New(fmt.Sprintf("Activity with ID %d Not Found", id))
	}
	return activity, err
}

func (s *service) CreateActivity(input *InputActivity) (*ActivityResponse, error) {
	cashier := Activity{
		Title: input.Title,
		Email: input.Email,
	}
	newCashier, err := s.repository.Create(&cashier)
	if err != nil {
		return nil, err
	}
	cashierResponse := ActivityResponse{
		Id:        newCashier.Id,
		Title:     newCashier.Title,
		Email:     newCashier.Email,
		UpdatedAt: newCashier.UpdatedAt,
		CreatedAt: newCashier.CreatedAt,
	}
	return &cashierResponse, nil
}

func (s *service) UpdateCashier(id int, input *InputActivity) (*Activity, error) {
	cashier := Activity{
		Title: input.Title,
		Email: input.Email,
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
