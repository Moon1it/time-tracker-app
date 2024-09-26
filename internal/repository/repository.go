package repository

import db "time-tracker-app/internal/db/sqlc"

type Repository struct {
	UserRepository *UserRepository
}

func NewRepository(querier db.Querier) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(querier),
	}
}
