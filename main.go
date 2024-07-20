package main

import (
	"fmt"
	"os"

	"github.com/trenton42/zettel/pkg/data"
	"github.com/trenton42/zettel/pkg/server"
)

func main() {
	db, err := data.New("~/Documents/zettel")
	if err != nil {
		fmt.Printf("Error setting up db connection: %s\n", err)
		os.Exit(1)
	}
	s := server.New(db)
	s.Start(":1123")
}
