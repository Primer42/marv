package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a file to interpret")
	}

	fileBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	lines, err := preprocess(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", lines)

}

func preprocess(dirty []byte) (clean []string, err error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(dirty))

	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			clean = append(clean, scanner.Text())
		}
	}

	return

}
