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

func TestReadAllRecords(t *testing.T) {

	fakeResponse1 := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(strings.NewReader(`{"Results": [{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"}]}`)),
	}

	fakeResponse2 := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(strings.NewReader(`{"Results": [{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"}]}`)),
	}

	fakeResponse3 := &http.Response{
		StatusCode: http.StatusOK,
		Body: io.NopCloser(strings.NewReader(`{"Results": [{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"},{"First": "Stirb", "Last": "Antoniu", "Email": 
			"stirbantoniu@mail", "Address": "123 str", "Created": "feb2023", "Balance": "100"}]}`)),
	}

	testCases := []struct {
		name                  string
		client                models.Client
		numberOfRecordsNeeded int
		url                   string
	}{
		{
			name:                  "Test case 0 records",
			client:                &FakeClient{Response: fakeResponse1, Error: nil},
			numberOfRecordsNeeded: 0,
			url:                   "fakeUrl0",
		},
		{
			name:                  "Test case 4 records",
			client:                &FakeClient{Response: fakeResponse2, Error: nil},
			numberOfRecordsNeeded: 4,
			url:                   "fakeUrl1",
		},
		{
			name:                  "Test case 5 records",
			client:                &FakeClient{Response: fakeResponse3, Error: nil},
			numberOfRecordsNeeded: 5,
			url:                   "fakeUrl2",
		},
		//{
		//	name:           "Test case 750 records",
		//	numberInserted: 750,
		//},
		//{
		//	name:           "Test case 215 records",
		//	numberInserted: 215,
		//},
		//{
		//	name:           "Test case 5 records",
		//	numberInserted: 5,
		//},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual, _ := ReadAllRecords(test.client, test.numberOfRecordsNeeded,
				test.url)
			//if err != nil {
			//	fmt.Println(err)
			//}
			if diff := cmp.Diff(len(actual), test.numberOfRecordsNeeded); diff != "" {
				t.Errorf("TestedReadAllRecords() does not meet expectations, "+
					"\nactual=%#v, \nexpected=%#v, \nDIFF: %v", actual, test.numberOfRecordsNeeded, diff)
			}
		})
	}
}
