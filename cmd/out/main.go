package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/mtharrison/github-pr-sha-comment-resource/internal/resource"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("missing arguments")
	}

	input, err := resource.GetInput(os.Stdin, true)
	if err != nil {
		log.Fatal("Could not read input: ", err)
	}

	sha, err := resource.GetShaFromDir(filepath.Join(os.Args[1], input.Params.Dir))
	if err != nil {
		log.Fatal("Could not get sha from dir: ", err)
	}

	pr, err := resource.GetPrNumberFromSha(input, sha)
	if err != nil {
		log.Println("Could not get PR from SHA: ", err)
		return // no exit - could be a non-pr commit
	}

	comment := resource.FormatComment(input.Params.Comment)

	url, err := resource.PostComment(input, comment, pr)
	if err != nil {
		log.Fatal("Could not post comment: ", err)
	}

	out, err := json.Marshal(resource.Output{Version: resource.Version{URL: url}})
	if err != nil {
		log.Fatal("Could not json marshal")
	}

	_, err = os.Stdout.Write(out)
	if err != nil {
		log.Fatal("Could not write to STDOUT")
	}
}
