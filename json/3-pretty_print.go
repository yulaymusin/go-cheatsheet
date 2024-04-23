package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	birds := map[string]interface{}{
		"sounds": map[string]string{
			"pigeon":  "coo",
			"eagle":   "squak",
			"owl":     "hoot",
			"duck":    "quack",
			"cuckoo":  "ku-ku",
			"raven":   "cruck-cruck",
			"chicken": "cluck",
			"rooster": "cock-a-doodle-do",
		},
	}

	data, err := json.MarshalIndent(birds, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// output is indented and hence more readable
	fmt.Println(string(data))
}
