package controller

import (
	"Project2Go/models"
	"encoding/json"
	"io"
	"net/http"
)

func GetData(url string) ([]models.Person, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var result models.Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Results, nil
}

func ReadAllRecords(numberOfRecords int, url string) ([]models.Person, error) {
	var persons []models.Person
	for len(persons) < numberOfRecords {
		result, err := GetData(url)
		if err != nil {
			return nil, err
		}
		persons = append(persons, result...)
	}
	persons = persons[:numberOfRecords]
	return persons, nil

}
