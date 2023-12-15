package main

import (
	"fmt"
	"os"

	"github.com/trenton42/zettel/api"
	"github.com/trenton42/zettel/data"
)

func main() {
	db, err := data.New("zettel-api")
	if err != nil {
		fmt.Printf("Error setting up db connection: %s\n", err)
		os.Exit(1)
	}
	defer db.Close()
	s := api.New(db)
	s.Serve()
}
