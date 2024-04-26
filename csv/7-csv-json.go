package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"strconv"
)

func main() {
	// open CSV file
	file, err := os.Open("./csv/csv-files/users.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// create CSV reader
	csvReader := csv.NewReader(file)

	// read CSV headers
	headers, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// read CSV data rows
	var data []map[string]interface{}
	for {
		row, err := csvReader.Read()
		if err != nil {
			break
		}

		// convert the row values to the appropriate types
		m := make(map[string]interface{})
		for i, val := range row {
			f, err := strconv.ParseFloat(val, 64)
			if err == nil {
				m[headers[i]] = f
				continue
			}

			b, err := strconv.ParseBool(val)
			if err == nil {
				m[headers[i]] = b
				continue
			}

			m[headers[i]] = val
		}

		data = append(data, m)
	}

	// encode JSON data
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(data); err != nil {
		log.Fatal(err)
	}
}
