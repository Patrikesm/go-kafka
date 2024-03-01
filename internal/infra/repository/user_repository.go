package repository

import (
	"database/sql"

	"github.com/Patrikesm/kafka-with-go/internal/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) Create(user *entity.User) error {
	_, err := u.DB.Exec("insert into user (name) values (?)", user.Name)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) FindByName(name string) (*entity.User, error) {
	row := u.DB.QueryRow("select * from user where name = ?", name)

	var user *entity.User

	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		return nil, err
	}

	return user, nil
}
