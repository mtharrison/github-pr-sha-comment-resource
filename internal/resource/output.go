package resource

// Version is just a dummy version for concourse
type Version struct {
	URL string `json:"url"`
}

// Output is printed to stdout
type Output struct {
	Version Version `json:"version"`
}
