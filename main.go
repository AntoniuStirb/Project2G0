package main

import (
	"Project2Go/controller"
	"Project2Go/service"
	"fmt"
	"net/http"
	"os"
)

func main() {
	totalRecords := 100
	client := http.DefaultClient
	records, err := controller.ReadAllRecords(client, totalRecords,
		"https://randomapi.com/api/6de6abfedb24f889e0b5f675edc50deb?fmt=prettyjson&sole")
	if err != nil {
		fmt.Printf("Unable to read data, reason: %v", err)
		return
	}
	fmt.Println("Records were read successfully")

	uniqueResult := service.DeleteDuplicates(records)
	fmt.Println("Duplicates were deleted successfully")

	sortedUniqueResults := service.GroupPersonsByFirstLetter(uniqueResult)
	fmt.Println("Records were sorted successfully")

	dir := "files"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			fmt.Printf("Unable to create new directory for files, reason: %v", err)
			return
		}
	}

	err = service.WriteJsonFiles(sortedUniqueResults, dir)
	if err != nil {
		fmt.Printf("Unable to write data in json files, reason: %v", err)
		return
	}
	fmt.Printf("Data was written successfully")

}
