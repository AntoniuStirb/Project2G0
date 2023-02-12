package controller

import (
	"Project2Go/models"
	"encoding/json"
	"io"
	"net/http"
)

type Client interface {
	Get(url string) (*http.Response, error)
}

// RealClient is the implementation of Client that uses the real http.Get
type RealClient struct{}

// Get implements the Client interface
func (c *RealClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

// FakeClient is a fake implementation of Client for testing
type FakeClient struct {
	Response *http.Response
	Error    error
}

// Get implements the Client interface
func (c *FakeClient) Get(url string) (*http.Response, error) {
	return c.Response, c.Error
}

func GetData(client RealClient, url string) ([]models.Person, error) {
	resp, err := client.Get(url)
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

func ReadAllRecords(client RealClient, numberOfRecords int, url string) ([]models.Person, error) {
	var persons []models.Person
	for len(persons) < numberOfRecords {
		result, err := GetData(client, url)
		if err != nil {
			return nil, err
		}
		persons = append(persons, result...)
	}
	persons = persons[:numberOfRecords]
	return persons, nil
}