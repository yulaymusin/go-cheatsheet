package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Person struct {
	Name   string `csv:"name"`
	Age    int    `csv:"age"`
	Gender string `csv:"gender"`
}

func main() {
	f, err := os.Create("./csv/csv-files/persons-writed.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	people := []*Person{
		{"Alice", 25, "Female"},
		{"Bob", 30, "Male"},
		{"Charlie", 35, "Male"},
	}

	if err := gocsv.MarshalFile(&people, f); err != nil {
		log.Fatal(err)
	}
}
