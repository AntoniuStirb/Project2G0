package controller

import (
	"Project2Go/models"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

// FakeClient is a fake implementation of Client for testing
type FakeClient struct {
	Response *http.Response
	Error    error
}

// Get implements the Client interface
func (c *FakeClient) Get(url string) (*http.Response, error) {
	return c.Response, c.Error
}

func TestGetData(t *testing.T) {
	// fake response for testing
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
			name:   "success scenario",
			client: &FakeClient{Response: fakeResponse, Error: nil},
			expected: []models.Person{{First: "Stirb", Last: "Antoniu", Email: "stirbantoniu@mail",
				Address: "123 str", Created: "feb2023", Balance: "100"}},
			url: "good url",
		},
		{
			name:        "failure scenario",
			client:      &FakeClient{Response: nil, Error: errors.New("failed to make request")},
			expectedErr: errors.New("failed to make request"),
			url:         "bad url",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			result, err := GetData(test.client, test.url)
			if err != nil && test.expectedErr == nil {
				t.Errorf("unexpected error: %s", err)
			}
			if err == nil && test.expectedErr != nil {
				t.Errorf("expected error: %s but got nil", test.expectedErr)
			}
			if err != nil && test.expectedErr != nil && err.Error() != test.expectedErr.Error() {
				t.Errorf("expected error: %s but got: %s", test.expectedErr, err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("expected: %+v but got: %+v", test.expected, result)
			}
		})
	}
}
