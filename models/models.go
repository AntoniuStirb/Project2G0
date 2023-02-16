package models

import "net/http"

type Person struct {
	First   string `json:"first"`
	Last    string `json:"last"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Created string `json:"created"`
	Balance string `json:"balance"`
}

type Response struct {
	Results []Person `json:"results"`
}

type PersonInfo struct {
	FirstLetter  string `json:"index"`
	Persons      []Person
	TotalRecords int `json:"total_records"`
}

type Client interface {
	Get(url string) (*http.Response, error)
}
