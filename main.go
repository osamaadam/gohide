package main

import (
	"log"

	"github.com/osamaadam/gohide/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
