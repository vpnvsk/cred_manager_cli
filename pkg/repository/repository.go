package repository

import (
	"github.com/google/uuid"
	"p_s_cli/internal/models"
)

type Repository struct {
	Authorization
	Manager
}

func NewRepository(url string) *Repository {
	return &Repository{
		Authorization: NewAuthorizationRepo(url),
		Manager:       NewManager(url),
	}
}

type Authorization interface {
	LogIn(user models.User) (string, error)
	SignUp(user models.User) error
}
type Manager interface {
	GetList() (models.AllList, error)
	GetPassword(uuid uuid.UUID) (models.SingleManager, error)
	UpdatePassword(uuid uuid.UUID, object models.UpdateManager) error
	DeletePassword(uuid uuid.UUID) error
	CreatePassword(object models.UpdateManager) error
}
