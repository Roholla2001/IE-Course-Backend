package main

import (
	"log"

	"github.com/Roholla2001/ie-course-backend/cmd/boot"
)

func main() {
	// starting our golang web server
	err := boot.BootServer()
	if err != nil {
		log.Fatal(err)
	}
}
