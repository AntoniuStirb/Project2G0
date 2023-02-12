package service

import (
	"Project2Go/models"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteJsonFiles(t *testing.T) {
	dirName := "testdir"

	testData := []models.PersonInfo{
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
				{
					First:   "Anna",
					Last:    "Keeling",
					Email:   "Anna.Keeling@sandra.biz",
					Address: "2197 Kreiger Forest",
					Created: "August 10, 2021",
					Balance: "$2,967.44",
				},
			},
			TotalRecords: 2,
		},
		{
			FirstLetter: "R",
			Persons: []models.Person{
				{
					First:   "Raul",
					Last:    "Lueilwitz",
					Email:   "redrabbit76@gmail.com",
					Address: "1797 Gorczany Drive",
					Created: "March 6, 2021",
					Balance: "$1,025.28",
				},
				{
					First:   "Rachel",
					Last:    "Robel",
					Email:   "greyturtle86@gmail.com",
					Address: "71706 Kaya Views",
					Created: "December 23, 2017",
					Balance: "$5,429.99",
				},
			},
			TotalRecords: 2,
		},
		{
			FirstLetter: "U",
			Persons: []models.Person{
				{
					First:   "Unique",
					Last:    "Emmerich",
					Email:   "Unique.Emmerich@constantin.biz",
					Address: "5104 Schoen Inlet",
					Created: "June 2, 2018",
					Balance: "$138.92",
				},
			},
			TotalRecords: 2,
		},
	}

	err := WriteJsonFiles(testData, dirName)
	if err != nil {
		t.Errorf("Error writing JSON files: %v", err)
	}

	// Check the generated files
	for _, record := range testData {
		filename := record.FirstLetter + ".json"
		filePath := filepath.Join(dirName, filename)

		// Read the contents of file
		fileContents, err := os.ReadFile(filePath)
		if err != nil {
			t.Errorf("Error reading file %s: %v", filePath, err)
		}

		// Unmarshal the contents of the file
		var personInfo models.PersonInfo
		err = json.Unmarshal(fileContents, &personInfo)
		if err != nil {
			t.Errorf("Error unmarshalling file %s: %v", filePath, err)
		}

		if diff := cmp.Diff(personInfo, record); diff != "" {
			t.Errorf("TestedDeleteDuplicates() does not meet expectations, "+
				"\nactual=%#v, \nexpected=%#v, \nDIFF: %v", personInfo, record, diff)

			// Remove files from directory
			err = os.RemoveAll(filePath)
			if err != nil {
				t.Errorf("Error cleaning test directory: %v", err)
			}
		}
	}
}
