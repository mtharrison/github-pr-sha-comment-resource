package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {

	_, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Could not read stdin")
	}

	_, err = os.Stdout.Write([]byte("[]"))
	if err != nil {
		log.Fatal("Could not write to STDOUT")
	}
}
