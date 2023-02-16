package service

import (
	"Project2Go/models"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []models.Person
		expected []models.Person
	}{
		{
			name: "Test case 1 == normal scenario, 1 duplicate",
			input: []models.Person{
				{
					First:   "Unique",
					Last:    "Emmerich",
					Email:   "Unique.Emmerich@constantin.biz",
					Address: "5104 Schoen Inlet",
					Created: "June 2, 2018",
					Balance: "$138.92",
				},
				{
					First:   "Dayna",
					Last:    "Wunsch",
					Email:   "plumturtle97@gmail.com",
					Address: "890 Ally Forks",
					Created: "December 31, 2014",
					Balance: "$1,002.81",
				},
				{
					First:   "Destiny",
					Last:    "Reinger",
					Email:   "skybluegiraffe32@gmail.com",
					Address: "33161 Weimann Ford",
					Created: "October 11, 2013",
					Balance: "$1,249.18",
				},
				{
					First:   "John",
					Last:    "Dan",
					Email:   "johndan@yahoo.com",
					Address: "33161 Weimann Ford",
					Created: "December 31, 2014",
					Balance: "$1,249.18",
				},
				{
					First:   "Dayna",
					Last:    "Wunsch",
					Email:   "plumturtle97@gmail.com",
					Address: "890 Ally Forks",
					Created: "December 31, 2014",
					Balance: "$1,002.81",
				},
			},
			expected: []models.Person{
				{
					First:   "Unique",
					Last:    "Emmerich",
					Email:   "Unique.Emmerich@constantin.biz",
					Address: "5104 Schoen Inlet",
					Created: "June 2, 2018",
					Balance: "$138.92",
				},
				{
					First:   "Dayna",
					Last:    "Wunsch",
					Email:   "plumturtle97@gmail.com",
					Address: "890 Ally Forks",
					Created: "December 31, 2014",
					Balance: "$1,002.81",
				},
				{
					First:   "Destiny",
					Last:    "Reinger",
					Email:   "skybluegiraffe32@gmail.com",
					Address: "33161 Weimann Ford",
					Created: "October 11, 2013",
					Balance: "$1,249.18",
				},
				{
					First:   "John",
					Last:    "Dan",
					Email:   "johndan@yahoo.com",
					Address: "33161 Weimann Ford",
					Created: "December 31, 2014",
					Balance: "$1,249.18",
				},
			},
		},
		{
			name:     "Test case 2 == Empty input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "Test case 3 == Empty slice",
			input:    []models.Person{},
			expected: nil,
		},
		{
			name: "Test case 4 == All records identical",
			input: []models.Person{
				{
					First:   "Ocie",
					Last:    "Nader",
					Email:   "azurewolf08@gmail.com",
					Address: "87486 Gibson Turnpike",
					Created: "April 20, 2015",
					Balance: "$8,188.76",
				},
				{
					First:   "Ocie",
					Last:    "Nader",
					Email:   "azurewolf08@gmail.com",
					Address: "87486 Gibson Turnpike",
					Created: "April 20, 2015",
					Balance: "$8,188.76",
				},
				{
					First:   "Ocie",
					Last:    "Nader",
					Email:   "azurewolf08@gmail.com",
					Address: "87486 Gibson Turnpike",
					Created: "April 20, 2015",
					Balance: "$8,188.76",
				},
			},
			expected: []models.Person{
				{
					First:   "Ocie",
					Last:    "Nader",
					Email:   "azurewolf08@gmail.com",
					Address: "87486 Gibson Turnpike",
					Created: "April 20, 2015",
					Balance: "$8,188.76",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := DeleteDuplicates(test.input)
			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("TestedDeleteDuplicates() does not meet expectations, "+
					"\nactual=%#v, \nexpected=%#v, \nDIFF: %v", actual, test.expected, diff)
			}
		})
	}
}

func TestGroupPersonsByFirstLetter(t *testing.T) {
	testCases := []struct {
		name     string
		input    []models.Person
		expected []models.PersonInfo
	}{
		{
			name:     "Test case 1: No persons in input slice",
			input:    []models.Person{},
			expected: nil,
		},
		{
			name: "Test case 2: Multiple persons with same first letter",
			input: []models.Person{
				{
					First:   "Dayna",
					Last:    "Wunsch",
					Email:   "plumturtle97@gmail.com",
					Address: "890 Ally Forks",
					Created: "December 31, 2014",
					Balance: "$1,002.81",
				},
				{
					First:   "Destiny",
					Last:    "Reinger",
					Email:   "skybluegiraffe32@gmail.com",
					Address: "33161 Weimann Ford",
					Created: "October 11, 2013",
					Balance: "$1,249.18",
				},
				{
					First:   "Dock",
					Last:    "Price",
					Email:   "Dock.Price@lucas.info",
					Address: "2800 Marks Hills",
					Created: "October 23, 2015",
					Balance: "$4,240.26",
				},
			},
			expected: []models.PersonInfo{
				{
					FirstLetter: "D",
					Persons: []models.Person{
						{
							First:   "Dayna",
							Last:    "Wunsch",
							Email:   "plumturtle97@gmail.com",
							Address: "890 Ally Forks",
							Created: "December 31, 2014",
							Balance: "$1,002.81",
						},
						{
							First:   "Destiny",
							Last:    "Reinger",
							Email:   "skybluegiraffe32@gmail.com",
							Address: "33161 Weimann Ford",
							Created: "October 11, 2013",
							Balance: "$1,249.18",
						},
						{
							First:   "Dock",
							Last:    "Price",
							Email:   "Dock.Price@lucas.info",
							Address: "2800 Marks Hills",
							Created: "October 23, 2015",
							Balance: "$4,240.26",
						},
					},
					TotalRecords: 3,
				},
			},
		},
		{
			name: "Test case 3: Input with multiple persons with different initial",
			input: []models.Person{
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
				{
					First:   "Unique",
					Last:    "Emmerich",
					Email:   "Unique.Emmerich@constantin.biz",
					Address: "5104 Schoen Inlet",
					Created: "June 2, 2018",
					Balance: "$138.92",
				},
			},
			expected: []models.PersonInfo{
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
					TotalRecords: 1,
				},
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := GroupPersonsByFirstLetter(test.input)
			if diff := cmp.Diff(actual, test.expected); diff != "" {
				t.Errorf("TestedGroupPersonByFirstLetter() does not meet expectations, "+
					"\nactual=%#v, \nexpected=%#v, \nDIFF: %v", actual, test.expected, diff)
			}
		})
	}
}
