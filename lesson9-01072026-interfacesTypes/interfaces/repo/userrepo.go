package repo

import "database/sql"

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (ur *UserRepo) CreateUser(id int, name string) error {
	//query INSERT INTO user VALUES()
}
