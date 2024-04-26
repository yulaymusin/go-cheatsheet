package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Record struct {
	Name   string `csv:"name"`
	Gender string `csv:"gender"`
}

func main() {
	// open the CSV file
	file, err := os.Open("./csv/csv-files/persons-writed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// read CSV file into a slice of Record structs
	var records []Record
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		log.Fatal(err)
	}

	// print records
	for _, record := range records {
		fmt.Printf("Name: %s, Gender: %s\n", record.Name, record.Gender)
	}
}
