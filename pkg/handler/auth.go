package handler

import (
	"context"
	"github.com/pterm/pterm"
	"p_s_cli/internal/models"
)

func (h *Handler) Auth(ctx context.Context) error {
	var err error
	result, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"log-in", "sign-up"}).
		Show()
	switch result {
	case "log-in":
		if err = h.LogIn(ctx); err != nil {
			return err
		}
	case "sign-up":
		if err = h.SignUp(ctx); err != nil {
			return err
		}
	}
	return err
}
func (h *Handler) LogIn(ctx context.Context) error {
	auth := models.User{}
	result, _ := pterm.DefaultInteractiveTextInput.Show("username")
	auth.UserName = result
	resultPas, _ := pterm.DefaultInteractiveTextInput.WithMask(" ").Show("password")
	auth.Password = resultPas
	err := h.repository.LogIn(ctx, auth)
	return err
}
func (h *Handler) SignUp(ctx context.Context) error {
	auth := models.User{}
	result, _ := pterm.DefaultInteractiveTextInput.Show("username")
	auth.UserName = result
	resultPas, _ := pterm.DefaultInteractiveTextInput.WithMask(" ").Show("password")
	auth.Password = resultPas
	err := h.repository.SignUp(ctx, auth)
	return err
}
