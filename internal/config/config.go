package config

import (
	"fmt"
	"os"
	"p_s_cli/internal/models"
	"path/filepath"
)

func Load() error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	fileName := "config.json"
	filePath := filepath.Join(currentDir, fileName)

	// Check if the configuration file exists.
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		// If the file does not exist, create a new configuration and write it.
		config := models.Config{
			JwtToken: "",
		}

		err = config.WriteToken(filePath)
		fmt.Printf("Configuration written to %s\n", filePath)
		return err
	} else if err == nil {
		// If the file exists, check if it's empty before writing.
		fileInfo, _ := os.Stat(filePath)
		if fileInfo.Size() == 0 {
			config := models.Config{
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
