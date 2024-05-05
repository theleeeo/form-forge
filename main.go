package main

import (
	"log"

	"go.leeeo.se/form-forge/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}
