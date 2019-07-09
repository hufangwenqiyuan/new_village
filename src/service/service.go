package service

import (
	"new_village/src/model"
)

type DbOrder interface {
	Create(order *model.Order) (*model.Order, error)
	UpdateById(order *model.Order) error
	Select(condition *model.QueryCondition) ([]*model.Order, error)
	SelectOrderById(order *model.Order) ([]*model.Order, error)
}

type ServiceOrder interface {
	DbOrder
	Delete() (string, error)
}

type ServiceMan struct {
	db      DbOrder
	message string
}

func ServiceManage(d DbOrder) *ServiceMan {
	return &ServiceMan{
		db:      d,
		message: "test",
	}
}

func (s *ServiceMan) Create(order *model.Order) (*model.Order, error) {
	return s.db.Create(order)
}

func (s *ServiceMan) UpdateById(order *model.Order) error {
	return s.db.UpdateById(order)
}

func (s *ServiceMan) Select(condition *model.QueryCondition) ([]*model.Order, error) {
	return s.db.Select(condition)
}

func (s *ServiceMan) SelectOrderById(order *model.Order) ([]*model.Order, error) {
	return s.db.SelectOrderById(order)
}

func (s *ServiceMan) Delete() (string, error) {

	return s.message, nil
}
