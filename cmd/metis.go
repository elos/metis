package main

import (
	"log"
	"os"
	"strings"

	"github.com/elos/metis"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please provide a file name, i.e., metis model.json or metis \"*.json\"")
	}

	arg := os.Args[1]
	models := make([]*metis.Model, 1)
	if strings.Contains(arg, "*") {
		var err error
		models, err = metis.ParseGlob(arg)
		if err != nil {
			log.Fatalf("error parsing %s: %s", arg, err)
		}
	} else {
		model := metis.ParseModelFile(arg)
		models[0] = model
	}

	s := metis.BuildSchema(models...)
	if err := s.Valid(); err != nil {
		log.Fatal(err)
	}

	log.Print("schema valid")
}
