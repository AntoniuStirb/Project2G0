package service

import (
	"Project2Go/models"
	"encoding/json"
	"fmt"
	"os"
)

func WriteJsonFiles(collections []models.PersonInfo, dirName string) error {
	for _, record := range collections {
		filename := record.FirstLetter + ".json"
		file, err := os.Create(fmt.Sprintf("%s/%s", dirName, filename))
		if err != nil {
			return err
		}

		jsonData, err := json.MarshalIndent(record, "", "  ")
		if err != nil {
			return err
		}

		_, err = file.Write(jsonData)
		if err != nil {
			return err
		}

		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}
