package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Astronaut struct {
	Name  string
	Craft string
}

type people struct {
	Number  int
	People  []Astronaut
	Message string
}

func main() {
	// the URL returns data about astronauts of the International Space Station
	url := "http://api.open-notify.org/astros.json"

	var netClient = http.Client{
		Timeout: time.Second * 10,
	}

	res, err := netClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(body)

	astros := people{}

	jsonErr := json.Unmarshal(body, &astros)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(astros)
}
