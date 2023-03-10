package controller

import (
	"Project2Go/models"
	"encoding/json"
	"io"
	"log"
)

// GetData function makes an HTTP GET request to the specified URL using the provided client.
// If the request is successful, the response body is read and unmarshalled into a
// models.Response struct. The result is a slice of models.Person which is then returned.
func GetData(client models.Client, url string) ([]models.Person, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and \nbody: %s\n", resp.StatusCode, body)
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	//not necessarry to treat the error when closing
	//if err != nil {
	//	return nil, err
	//}

	var result models.Response
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result.Results, nil
}

// ReadAllRecords function retrieves a set of records from a data source using an HTTP GET request.
// The function takes in three arguments: a Client object that implements the Get method to make HTTP requests,
// a numberOfRecords that specifies the number of records to retrieve,
// and a URL that specifies the endpoint to retrieve the data from.
// The function repeatedly calls the GetData function with the provided Client
// and url arguments, until the desired number of records has been retrieved.
// The returned result from each call to GetData is then appended to the persons slice.
func ReadAllRecords(client models.Client, numberOfRecords int, url string) ([]models.Person, error) {
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
