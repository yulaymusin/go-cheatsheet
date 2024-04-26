package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type User struct {
	firstName  string
	lastName   string
	occupation string
}

func main() {
	users, err := readData("./csv/csv-files/users.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range users {
		user := User{
			firstName:  record[0],
			lastName:   record[1],
			occupation: record[2],
		}
		fmt.Printf("%s %s is a %s\n", user.firstName, user.lastName, user.occupation)
	}
}

func readData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// skip the first line, which contains the column names
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	// get all the records in one shot with ReadAll
	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
