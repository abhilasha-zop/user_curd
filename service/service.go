package service

import (
	"errors"
	"strings"
	"user_crud/model"
	"user_crud/store"
)

type UserService struct {
	Store *store.UserStore
}

// simple custom logic function
func (s *UserService) FormatName(name string) string {
	return strings.ToUpper(name)
}

func (s *UserService) CreateUser(u model.User) error {
	if u.Name == "" {
		return errors.New("name required")
	}

	u.Name = s.FormatName(u.Name) // using custom func
	return s.Store.Create(u)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.Store.GetAll()
}

func (s *UserService) GetUser(id int) (model.User, error) {
	return s.Store.GetByID(id)
}

func (s *UserService) UpdateUser(u model.User) error {
	if u.ID == 0 {
		return errors.New("invalid id")
	}
	return s.Store.Update(u)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Store.Delete(id)
}
