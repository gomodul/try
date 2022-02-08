package main

import (
	"log"

	"github.com/gomodul/try"
)

func main() {
	errTry := try.Do(func(attempt int) error {
		log.Printf("counter: %v", attempt)
		panic("some panic")
	})
	if errTry != nil {
		log.Fatal(errTry)
	}
}
