package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main() {
    f, err := os.Open("./csv/csv-files/users-delimiter.csv")
    if err != nil {
        log.Fatal(err)
    }
	defer f.Close()

	// set the separator and the comment character so that the package knows how to parse the file
    r := csv.NewReader(f)
    r.Comma = ';'
    r.Comment = '#'

    records, err := r.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(records)
}
