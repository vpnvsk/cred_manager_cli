package handler

import (
	"github.com/pterm/pterm"
	"p_s_cli/internal/models"
)

func (h *Handler) Auth() error {
	var err error
	result, _ := pterm.DefaultInteractiveSelect.
		WithOptions([]string{"log-in", "sign-up"}).
		Show()
	switch result {
	case "log-in":
		if err = h.LogIn(); err != nil {
			return err
		}
	case "sign-up":
		if err = h.SignUp(); err != nil {
			return err
		}
	}
	return err
}
func (h *Handler) LogIn() error {
	auth := models.User{}
	result, _ := pterm.DefaultInteractiveTextInput.Show("username")
	auth.UserName = result
	resultPas, _ := pterm.DefaultInteractiveTextInput.WithMask(" ").Show("password")
	auth.Password = resultPas
	_, err := h.repository.LogIn(auth)
	return err
}
func (h *Handler) SignUp() error {
	auth := models.User{}
	result, _ := pterm.DefaultInteractiveTextInput.Show("username")
	auth.UserName = result
	resultPas, _ := pterm.DefaultInteractiveTextInput.WithMask(" ").Show("password")
	auth.Password = resultPas
	err := h.repository.SignUp(auth)
	return err
}
