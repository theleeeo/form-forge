package main

import (
	"log"

	"github.com/theleeeo/form-forge/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}
