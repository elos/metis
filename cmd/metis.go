package main

import (
	"log"
	"os"
	"strings"

	"github.com/elos/metis"
)

// Note: the first arg is _always_ "metis"
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide a file name, i.e., metis model.json or metis \"*.json\"")
	}

	arg := os.Args[1]
	models := make([]*metis.Model, 1)

	if strings.Contains(arg, "*") { // glob
		var err error
		models, err = metis.ParseGlob(arg)

		if err != nil {
			log.Fatalf("Error parsing %s: %s", arg, err)
		}

	} else {
		model, err := metis.ParseModelFile(arg)

		// Print the Error
		if err != nil {
			log.Fatalf("Error parsing %s: %s", arg, err)
		}

		models[0] = model
	}

	s := metis.BuildSchema(models...)
	if err := s.Valid(); err != nil {
		log.Fatal("Error validating schema: %s", err)
	}

	log.Print("Schema Valid")
}
