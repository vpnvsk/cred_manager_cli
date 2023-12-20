package repository

import (
	"context"
	"github.com/google/uuid"
	"p_s_cli/internal/client"
	"p_s_cli/internal/models"
	"time"
)

type Repository struct {
	Authorization
	Manager
}

func NewRepository(
	ctx context.Context,
	url, authUrl string,
	timeout time.Duration,
	retriesCount int,
	appId int32,
) *Repository {
	return &Repository{
		Authorization: client.NewClient(ctx, authUrl, timeout, retriesCount, appId),
		Manager:       NewManager(url),
	}
}

type Authorization interface {
	LogIn(ctx context.Context, user models.User) error
	SignUp(ctx context.Context, user models.User) error
}
type Manager interface {
	GetList() (models.AllList, error)
	GetPassword(uuid uuid.UUID) (models.SingleManager, error)
	UpdatePassword(uuid uuid.UUID, object models.UpdateManager) error
	DeletePassword(uuid uuid.UUID) error
	CreatePassword(object models.UpdateManager) error
}
