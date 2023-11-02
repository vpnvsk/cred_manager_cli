package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/pterm/pterm"
	"io"
	"net/http"
	"p_s_cli/internal/models"
)

type AuthorizationRepo struct {
	baseURL string
}

func NewAuthorizationRepo(url string) *AuthorizationRepo {
	return &AuthorizationRepo{baseURL: url}
}

func (r *AuthorizationRepo) LogIn(user models.User) error {
	url := r.baseURL + "/auth/log-in"
	// Marshal the data into a JSON byte slice.
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if resp.StatusCode != 200 {
		return errors.New("bad credentials")
	}

	if err = getAndSaveToken(resp); err != nil {
		return err
	}
	pterm.Success.Printfln("LogIn successfully")

	return err
}
func getAndSaveToken(resp *http.Response) error {
	var accessToken models.JWTToken

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err = json.Unmarshal(responseBody, &accessToken); err != nil {
		return err
	}
	config := models.Config{
		JwtToken: accessToken.Token,
	}
	filePath := "config.json"
	go config.WriteToken(filePath)
	return err
}
func (r *AuthorizationRepo) SignUp(user models.User) error {
	url := r.baseURL + "/auth/sign-up"
	// Marshal the data into a JSON byte slice.
	jsonData, err := json.Marshal(user)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	loginResult := make(chan error)

	// Start a goroutine to perform the login.
	go func() {
		loginErr := r.LogIn(user)
		if loginErr != nil {
			loginResult <- loginErr
		} else {
			loginResult <- nil // Login succeeded
		}
	}()

	// Wait for the login result from the goroutine.
	loginErr := <-loginResult

	return loginErr
}
