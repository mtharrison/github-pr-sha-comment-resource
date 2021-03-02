package resource

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getInput() Input {
	return Input{
		Source: &Source{
			RepositoryString: "golang/go",
			AccessToken:      "abc",
			V3Endpoint:       "http://github.com/api/v3",
		},
		Params: &Params{
			Dir:     "123",
			Comment: "456",
		},
	}
}

func TestInputValidate(t *testing.T) {
	var i Input

	i = getInput()
	i.Source = nil
	err := i.Validate(false)
	if err.Error() != "source cannot be empty" {
		t.Error("should validate source cannot be empty")
	}

	i = getInput()
	i.Source.RepositoryString = ""
	err = i.Validate(false)
	if err.Error() != "source.repository cannot be empty" {
		t.Error("should validate source.repository cannot be empty")
	}

	i = getInput()
	i.Source.AccessToken = ""
	err = i.Validate(false)
	if err.Error() != "source.access_token cannot be empty" {
		t.Error("should validate source.access_token cannot be empty")
	}

	i = getInput()
	i.Source.V3Endpoint = ""
	err = i.Validate(false)
	if err != nil {
		t.Error("source.v3_endpoint should be optional")
	}

	i = getInput()
	i.Params.Dir = ""
	err = i.Validate(true)
	if err.Error() != "if set, version must include comment and pr" {
		t.Error("should not allow comment to be omitted")
	}

	i = getInput()
	i.Params.Comment = ""
	err = i.Validate(true)
	if err.Error() != "if set, version must include comment and pr" {
		t.Error("should not allow pr to be omitted")
	}
}

func TestSourceOwner(t *testing.T) {
	i := getInput()
	assert.Equal(t, i.Source.Owner(), "golang")
}

func TestSourceRepo(t *testing.T) {
	i := getInput()
	assert.Equal(t, i.Source.Repo(), "go")
}

func TestGetInput(t *testing.T) {
	inputString := "{\"source\":{\"repository\":\"golang/go\",\"access_token\":\"abc\",\"v3_endpoint\":\"https://github.com/api/v3\",\"regex\":\"do something\"},\"params\":{\"dir\":\"111\",\"comment\":\"222\"}}"
	reader := strings.NewReader(inputString)
	input, err := GetInput(reader, true)

	assert.Nil(t, err)

	expected := Input{
		Source: &Source{
			RepositoryString: "golang/go",
			AccessToken:      "abc",
			V3Endpoint:       "https://github.com/api/v3",
		},
		Params: &Params{
			Dir:     "111",
			Comment: "222",
		},
	}

	assert.Nil(t, err)
	assert.Equal(t, expected, input)
}
