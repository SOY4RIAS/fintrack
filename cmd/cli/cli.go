package cli

import (
	"log"
)

func Run() (func(), error) {
	log.Println("Executing run function")

	return func() {
		log.Println("Cleaning up run function")
	}, nil
}
