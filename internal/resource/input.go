package resource

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

// Input is provided to this resource via STDIN
type Input struct {
	Source *Source `json:"source"`
	Params *Params `json:"params"`
}

// Validate checks if the input is sound
func (i *Input) Validate(requireVersion bool) error {

	if i.Source == nil {
		return errors.New("source cannot be empty")
	}

	if i.Source.RepositoryString == "" {
		return errors.New("source.repository cannot be empty")
	}

	if i.Source.AccessToken == "" {
		return errors.New("source.access_token cannot be empty")
	}

	if i.Params != nil {
		if i.Params.Comment == "" {
			return errors.New("must include comment")
		}

		if i.Params.Dir == "" {
			return errors.New("must include dir")
		}
	}

	return nil
}

// Source is the configurable settings a user provides to this resource
type Source struct {
	RepositoryString string `json:"repository"`
	AccessToken      string `json:"access_token"`
	V3Endpoint       string `json:"v3_endpoint"`
}

// Owner returns the repo owner
func (s *Source) Owner() string {
	return strings.Split(s.RepositoryString, "/")[0]
}

// Repo returns the repo name
func (s *Source) Repo() string {
	return strings.Split(s.RepositoryString, "/")[1]
}

// Params represents a the params in a put operation
type Params struct {
	Dir     string `json:"dir"`
	Comment string `json:"comment"`
}

// GetInput takes a reader and constructs an Input
func GetInput(reader io.Reader, requireVersion bool) (Input, error) {
	input, err := readInput(reader)
	if err != nil {
		return input, err
	}

	err = input.Validate(requireVersion)
	if err != nil {
		return input, err
	}

	return input, nil
}

func readInput(reader io.Reader) (Input, error) {
	var input Input

	rawInput, err := ioutil.ReadAll(reader)
	if err != nil {
		return input, err
	}

	err = json.Unmarshal(rawInput, &input)
	if err != nil {
		return input, err
	}

	return input, nil
}
