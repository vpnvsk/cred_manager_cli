package handler

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/pterm/pterm"
	"p_s_cli/internal/models"
	"strconv"
	"unicode/utf8"
)

var emptyCredentialsList = errors.New("you have empty credentials list, please create some at first")

func (h *Handler) GetManager() error {
	list, err := h.repository.GetList()
	if err != nil {
		return err
	}
	var lst []string
	if len(list.Data) < 1 {
		return emptyCredentialsList
	}
	for i, e := range list.Data {
		formatString := fmt.Sprintf("%d %s(%s)", i+1, e.Title, e.Description)
		lst = append(lst, formatString)
	}
	result, _ := pterm.DefaultInteractiveSelect.
		WithOptions(lst).
		Show()
	var firstChar rune
	if len(result) > 0 {
		firstChar, _ = utf8.DecodeRuneInString(result)
	}
	num, _ := strconv.Atoi(string(firstChar))

	result, _ = pterm.DefaultInteractiveSelect.
		WithOptions([]string{"get", "update", "del"}).
		Show()
	switch result {
	case "get":
		res, err := h.GetCredentials(list.Data[num-1].Id)
		if err != nil {
			return err
		} else {
			pterm.Info.Printfln("Your login: %s\nYour password: %s", res.Userlogin, res.Password)
		}
	case "update":
		if err = h.UpdateManager(list.Data[num-1].Id); err != nil {
			return err
		} else {
			pterm.Info.Printfln("Successfully updated")
		}
	case "del":
		if err = h.DeleteManager(list.Data[num-1].Id); err != nil {
			return err
		} else {
			pterm.Info.Printfln("Successfully deleted")
		}

	}
	return err
}
func (h *Handler) GetCredentials(id uuid.UUID) (models.SingleManager, error) {
	res, err := h.repository.GetPassword(id)
	return res, err
}
func (h *Handler) UpdateManager(id uuid.UUID) error {
	var input models.UpdateManager
	result, _ := pterm.DefaultInteractiveTextInput.Show("title")
	input.Title = result
	result, _ = pterm.DefaultInteractiveTextInput.Show("username")
	input.Userlogin = result
	result, _ = pterm.DefaultInteractiveTextInput.WithMask(" ").Show("password")
	input.Password = result
	result, _ = pterm.DefaultInteractiveTextInput.Show("description")
	input.Description = result
	err := h.repository.UpdatePassword(id, input)
	return err
}
func (h *Handler) DeleteManager(id uuid.UUID) error {
	err := h.repository.DeletePassword(id)
	return err
}
