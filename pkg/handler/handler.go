package handler

import (
	"github.com/pterm/pterm"
	"p_s_cli/pkg/repository"
)

type Handler struct {
	repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{repository: repository}
}

func (h *Handler) Init() error {

	for {
		result, _ := pterm.DefaultInteractiveSelect.
			WithOptions([]string{"auth", "get list of services", "create new", "quit"}).
			Show()
		switch result {
		case "auth":
			if err := h.Auth(); err != nil {
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
