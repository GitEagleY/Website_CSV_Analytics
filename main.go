package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCSV(filePath string) map[string]map[string]bool {
	//open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	//create reader
	reader := csv.NewReader(file)
	//read records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		os.Exit(1)
	}

	//map(UserID, ProductID:bool)
	visits := make(map[string]map[string]bool)
	for _, record := range records {
		userID := record[0]
		productID := record[1]
		// if there is no entry for the userID in the visits map (ok is false),
		//the code initializes a new map for that userID
		if _, ok := visits[userID]; !ok {
			visits[userID] = make(map[string]bool)
		}
		visits[userID][productID] = true
	}

	return visits
}

func findUsersWithNewVisits(daylyVisits map[string]map[string]bool) []string {
	var usersWithNewVisits []string

	for userID, productsDay := range daylyVisits {
		// Assume the user visited all products until proven otherwise
		newVisit := true

		// Check if the user visited all products on the given day
		for _, visited := range productsDay {
			// If the user didn't visit a product, set newVisit to false
			if !visited {
				newVisit = false
				// Exit the loop early if any product wasn't visited
				break
			}
		}

		// If newVisit is still true, add the user ID to the result slice
		if newVisit {
			usersWithNewVisits = append(usersWithNewVisits, userID)
		}
	}

	// Return the slice containing user IDs who visited all products
	return usersWithNewVisits
}

func main() {
	file1 := "file1.csv"
	file2 := "file2.csv"

	visits1 := readCSV(file1)
	visits2 := readCSV(file2)

	usersWithNewVisits1 := findUsersWithNewVisits(visits1)
	usersWithNewVisits2 := findUsersWithNewVisits(visits2)

	// Find common users who visited all products on both days
	commonUsers := merge(usersWithNewVisits1, usersWithNewVisits2)

	for _, userID := range commonUsers {
		fmt.Println(userID)
	}
}

func merge(slice1, slice2 []string) []string {
	merge := make(map[string]bool)
	result := make([]string, 0)

	for _, item := range slice1 {
		merge[item] = true
	}

	for _, item := range slice2 {
		if merge[item] {
			result = append(result, item)
		}
	}

	return result
}
