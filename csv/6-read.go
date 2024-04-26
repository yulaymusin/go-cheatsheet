package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	// open the file using the os.Open() function and assign it to the variable file
	file, err := os.Open("./csv/csv-files/users.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1   // allow variable number of fields
	data, err := reader.ReadAll() // read all the records from the file and assign them to the data variable
	if err != nil {
		log.Fatal(err)
	}

	// print CSV data
	for _, row := range data {
		for _, col := range row {
			fmt.Printf("%s,", col)
		}
		fmt.Println()
	}
}
