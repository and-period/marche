package main

import (
	"log"
	"os"

	cmd "github.com/and-period/marche/api/internal/gateway/cmd/admin"
)

func main() {
	if err := cmd.Exec(); err != nil {
		log.Printf("An error has occurred: %v", err)
		os.Exit(1)
	}
}
