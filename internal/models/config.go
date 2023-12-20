package models

import (
	"encoding/json"
	"os"
)

type Token struct {
	JwtToken string `json:"token"`
}

func (t Token) WriteToken(filePath string) error {
	configJSON, err := json.Marshal(t)
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
