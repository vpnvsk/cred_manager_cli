package models

import (
	"encoding/json"
	"os"
)

type Config struct {
	JwtToken string `json:"token"`
}

func (c Config) WriteToken(filePath string) error {
	configJSON, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	configFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	_, err = configFile.Write(configJSON)
	if err != nil {
		panic(err)
	}
	return err
}
