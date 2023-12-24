package main

import (
	"log"

	"github.com/gomodul/try"
)

func main() {
	errTry := try.Do(func(attempt int) error {
		return try.ErrSomethingWrong
	})
	if errTry != nil {
		log.Fatal(errTry)
	}
}
