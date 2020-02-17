package main

import "fmt"

import "time"

func main() {
	_, cleanup, err := initUserStore("posgtres://toto:toto@host.domain/db")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	defer cleanup()

	time.Sleep(4 * time.Second)
}
