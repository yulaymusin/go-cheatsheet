package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendMessage(ch)
	message := <-ch // receive message from channel
	fmt.Println(message)
}

func sendMessage(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello World!"
}
