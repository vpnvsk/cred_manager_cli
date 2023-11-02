package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"p_s_cli/internal/models"
)

var (
	unauthorizedError = errors.New("unauthorized")
)

type ManagerRepo struct {
	baseURL string
}

func NewManager(url string) *ManagerRepo {
	return &ManagerRepo{baseURL: url}
}

func (r *ManagerRepo) makeRequest(method, endpoint string, body []byte) (*http.Response, error) {
	url := r.baseURL + endpoint
	token, err := readConfig("config.json")
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token.Token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		return nil, unauthorizedError
	}

	return resp, nil
}

func (r *ManagerRepo) GetList() (models.AllList, error) {
	var list models.AllList
	resp, err := r.makeRequest("GET", "/api/ps/", nil)
	if err != nil {
		return list, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(responseBody, &list)
	return list, err
}

func (r *ManagerRepo) GetPassword(uuid uuid.UUID) (models.SingleManager, error) {
	var singleManager models.SingleManager
	resp, err := r.makeRequest("GET", fmt.Sprintf("/api/ps/%s", uuid), nil)
	if err != nil {
		return singleManager, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return singleManager, err
	}

	err = json.Unmarshal(responseBody, &singleManager)
	return singleManager, err
}

func (r *ManagerRepo) UpdatePassword(uuid uuid.UUID, object models.UpdateManager) error {
	jsonData, err := json.Marshal(object)
	if err != nil {
		return err
	}

	resp, err := r.makeRequest("PUT", fmt.Sprintf("/api/ps/%s", uuid), jsonData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	return err
}

func (r *ManagerRepo) DeletePassword(uuid uuid.UUID) error {
	resp, err := r.makeRequest("DELETE", fmt.Sprintf("/api/ps/%s", uuid), nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	return err
}

func (r *ManagerRepo) CreatePassword(object models.UpdateManager) error {
	jsonData, _ := json.Marshal(object)

	resp, err := r.makeRequest("POST", "/api/ps/", jsonData)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	return err
}

func readConfig(configFile string) (models.JWTToken, error) {
	var token models.JWTToken

	data, err := os.ReadFile(configFile)
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(data, &token)
	return token, err
}
