package service

import (
	"Project2Go/models"
	"testing"
)

func TestWriteJsonFiles(t *testing.T) {
	testCases := []struct {
		collections []models.PersonInfo
		dirName     string
		expectedErr bool
	}{
		{
			collections: []models.PersonInfo{},
			dirName:     "testdata",
			expectedErr: false,
		},
		{
			collections: []models.PersonInfo{
				{
					FirstLetter: "A",
					Persons: []models.Person{
						{
							First:   "Abel",
							Last:    "Kris",
							Email:   "pinkturtle55@gmail.com",
							Address: "179 Theron Fork",
							Created: "July 18, 2018",
							Balance: "$9,891.57",
						},
					},
					TotalRecords: 1,
				},
			},
			dirName:     "testdata",
			expectedErr: false,
		},
		{
			collections: []models.PersonInfo{
				{
					FirstLetter: "A",
					Persons: []models.Person{
						{
							First:   "Abel",
							Last:    "Kris",
							Email:   "pinkturtle55@gmail.com",
							Address: "179 Theron Fork",
							Created: "July 18, 2018",
							Balance: "$9,891.57",
						},
					},
					TotalRecords: 1,
				},
			},
			dirName:     "/path/invalidDir",
			expectedErr: true,
		},
		//{
		//	collections: []models.PersonInfo{
		//		{
		//			FirstLetter: "A",
		//			Persons: []models.Person{
		//				{
		//					First:   "Abel",
		//					Last:    "Kris",
		//					Email:   "pinkturtle55@gmail.com",
		//					Address: "179 Theron Fork",
		//					Created: "July 18, 2018",
		//					Balance: "$9,891.57",
		//				},
		//			},
		//			TotalRecords: 1,
		//		},
		//	},
		//	dirName:     "testdata",
		//	expectedErr: true,
		//},
	}
	for _, testCases := range testCases {
		actualErr := false
		err := WriteJsonFiles(testCases.collections, testCases.dirName)
		if err != nil {
			actualErr = true
		}
		if actualErr != testCases.expectedErr {
			t.Errorf("expected error %v but got %v", testCases.expectedErr, actualErr)
		}
	}
}
