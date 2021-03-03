package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Could not read stdin")
	}

	_, err = os.Stdout.Write(input)
	if err != nil {
		log.Fatal("Could not write to STDOUT")
	}
}
