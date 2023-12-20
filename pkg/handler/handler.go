package handler

import (
	"context"
	"github.com/pterm/pterm"
	"p_s_cli/pkg/repository"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repository: repo}
}

func (h *Handler) Init(ctx context.Context) error {
	for {
		result, _ := pterm.DefaultInteractiveSelect.
			WithOptions([]string{"auth", "get list of services", "create new", "quit"}).
			Show()
		switch result {
		case "auth":
			if err := h.Auth(ctx); err != nil {
				return err
			}
		case "get list of services":
			if err := h.GetManager(); err != nil {
				return err
			}
		case "create new":
			if err := h.CreateManager(); err != nil {
				return err
			}
		case "quit":
			return nil

		}
	}
}
