package service

import (
	"Project2Go/models"
	"sort"
)

// DeleteDuplicates function take as an input a slice of structs of type Person, creates a map of the
// same type which acts like a filter to store unique elements of the input. The output of the function
// is also a slice of structs of type Person, but duplicates are removed.
func DeleteDuplicates(s []models.Person) []models.Person {
	inResult := make(map[models.Person]bool)
	var result []models.Person
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

// GroupPersonsByFirstLetter function takes as input a slie of Persons type. This slice is sorted
// based on the first letter of the First field, this being the initial of the name. The sorting is done using
// the "sort.Slice" function, which sorts the slice in place. The function returns a slice of type PersonInfo
// containing the FirstLetter, sorted Records after initial of First field and TotalRecords
func GroupPersonsByFirstLetter(s []models.Person) []models.PersonInfo {
	sort.Slice(s, func(i, j int) bool {
		return s[i].First[0:1] < s[j].First[0:1]
	})
	var collections []models.PersonInfo
	for i := 0; i < len(s); i++ {
		cnt := i
		for cnt < len(s) && s[cnt].First[0:1] == s[i].First[0:1] {
			cnt++
		}
		collections = append(collections, models.PersonInfo{
			FirstLetter:  s[i].First[0:1],
			Persons:      s[i:cnt],
			TotalRecords: cnt - i,
		})
		i = cnt - 1
	}
	return collections
}
