package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Primer42/marv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file to interpret")
	}

	fileBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines, err := marv.Preprocess(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", lines)

	p := marv.NewParser()

	smts, err := p.Parse(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", smts)

}
