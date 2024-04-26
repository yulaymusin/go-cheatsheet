package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	data := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}
	extraData := [][]string{
		{"Julian", "Stone", "developer"},
		{"Jack", "Black", "student"},
	}

	// create a file
	f, err := os.Create("./csv/csv-files/users-writed.csv")
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	defer f.Close()

	// initialize csv writer
	w := csv.NewWriter(f)
	defer w.Flush()

	// write row by row
	for _, record := range data {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	// write all rows at once
	err = w.WriteAll(extraData)
	if err != nil {
		log.Fatal(err)
	}

	// write single row
	extraDataRow := []string{"Benjamin", "Ethan", "driver"}
	w.Write(extraDataRow)
}
