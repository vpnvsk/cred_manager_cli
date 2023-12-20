package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"p_s_cli/internal/models"
	"path/filepath"
	"time"
)

type Config struct {
	Url     string        `yaml:"url"`
	AuthUrl string        `yaml:"authUrl"`
	Retries int           `yaml:"retries"`
	TimeOut time.Duration `yaml:"timeOut"`
	AppId   int32         `yaml:"appId"`
}

func Load() *Config {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config/config.yaml"
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config path if incorrect or empty: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}
	return &cfg
}

func LoadToken() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := "token.json"
	filePath := filepath.Join(currentDir, fileName)

	// Check if the configuration file exists.
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		// If the file does not exist, create a new configuration and write it.
		config := models.Token{
			JwtToken: "",
		}

		err = config.WriteToken(filePath)
		fmt.Printf("Configuration written to %s\n", filePath)
		return err
	} else if err == nil {
		// If the file exists, check if it's empty before writing.
		fileInfo, _ := os.Stat(filePath)
		if fileInfo.Size() == 0 {
			config := models.Token{
				JwtToken: "",
			}
			err = config.WriteToken(filePath)
			fmt.Printf("Configuration written to %s\n", filePath)
		} else {
			// File exists and is not empty, no need to change it.
			fmt.Printf("Configuration file is not empty, skipping write.\n")
		}
	}

	return err
}
