package services

import (
	"database/sql"
	"github.com/justyork/api-template/internal/models"
)

type UserService struct {
	DB *sql.DB
}

func (s *UserService) CreateUser(user models.User) (int64, error) {
	stmt, err := s.DB.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (s *UserService) GetUserByID(id int64) (*models.User, error) {
	row := s.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}

	return &user, nil
}
