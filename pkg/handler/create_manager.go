package handler

import (
	"github.com/pterm/pterm"
	"p_s_cli/internal/models"
)

func (h *Handler) CreateManager() error {
	var input models.UpdateManager
	var err error
	result, _ := pterm.DefaultInteractiveTextInput.Show("title")
	input.Title = result
	result, _ = pterm.DefaultInteractiveTextInput.Show("username")
	input.Userlogin = result
	result, _ = pterm.DefaultInteractiveTextInput.WithMask(" ").Show("password")
	input.Password = result
	result, _ = pterm.DefaultInteractiveTextInput.Show("description")
	input.Description = result

	if err = h.repository.CreatePassword(input); err != nil {
		return err
	} else {
		pterm.Info.Printfln("Successfully updated")
	}
	return err
}
