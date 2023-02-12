package controller

import (
	"Project2Go/models"
	"errors"
	"github.com/google/go-cmp/cmp"
	"io"
	"net/http"
	"strings"
	"testing"
)

type FakeClient struct {
	Response *http.Response
	Error    error
}

func (c *FakeClient) Get(string) (*http.Response, error) {
	return c.Response, c.Error
}

func TestGetData(t *testing.T) {
	fakeResponse := &http.Response{
		Body: io.NopCloser(strings.NewReader(`{"Results": [{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"}]}`)),
	}

	tests := []struct {
		name        string
		client      models.Client
		expected    []models.Person
		expectedErr error
		url         string
	}{
		{
			name:   "Successfully request",
			client: &FakeClient{Response: fakeResponse, Error: nil},
			expected: []models.Person{{First: "Stirb", Last: "Antoniu", Email: "stirbantoniu@mail",
				Address: "123 str", Created: "feb2023", Balance: "100"}},
			url: "fakeUrl1",
		},
		{
			name:        "Failed request",
			client:      &FakeClient{Response: nil, Error: errors.New("failed to make request")},
			expectedErr: errors.New("failed to make request"),
			url:         "fakeUrl2",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := GetData(test.client, test.url)
			if err != nil && test.expectedErr == nil {
				t.Errorf("unexpected error: %s", err)
			}
			if err == nil && test.expectedErr != nil {
				t.Errorf("error expected: %s but got nil", test.expectedErr)
			}
			if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("expected error: %s but got: %s", test.expectedErr, err)
			}
			if diff := cmp.Diff(result, test.expected); diff != "" {
				t.Errorf("expected: %+v but got:\n %+v", test.expected, result)
			}
		})
	}
}
