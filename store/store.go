package store

import (
	"database/sql"
	"user_crud/model"
)

type UserStore struct {
	DB *sql.DB
}

func (s *UserStore) Create(u model.User) error {
	_, err := s.DB.Exec("INSERT INTO users(name,age) VALUES (?,?)", u.Name, u.Age)
	return err
}

func (s *UserStore) GetAll() ([]model.User, error) {
	rows, err := s.DB.Query("SELECT id,name,age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.Age)
		users = append(users, u)
	}
	return users, nil
}

func (s *UserStore) GetByID(id int) (model.User, error) {
	var u model.User
	err := s.DB.QueryRow("SELECT id,name,age FROM users WHERE id=?", id).
		Scan(&u.ID, &u.Name, &u.Age)
	return u, err
}

func (s *UserStore) Update(u model.User) error {
	_, err := s.DB.Exec("UPDATE users SET name=?, age=? WHERE id=?", u.Name, u.Age, u.ID)
	return err
}

func (s *UserStore) Delete(id int) error {
	_, err := s.DB.Exec("DELETE FROM users WHERE id=?", id)
	return err
}
