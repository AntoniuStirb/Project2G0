package models

import "net/http"

type Person struct {
	First   string
	Last    string
	Email   string
	Address string
	Created string
	Balance string
}

type Response struct {
	Results []Person
}

type PersonInfo struct {
	FirstLetter  string
	Persons      []Person
	TotalRecords int
}

type Client interface {
	Get(url string) (*http.Response, error)
}
