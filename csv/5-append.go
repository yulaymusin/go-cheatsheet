package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// open CSV file for appending
	file, err := os.OpenFile("./csv/csv-files/users-writed.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write a new row to CSV file
	row := []string{"Martin", "Yellow", "engineer"}
	err = writer.Write(row)
	if err != nil {
		log.Fatal(err)
	}
}
