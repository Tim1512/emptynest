package main

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// Name is a descriptive name of the
func Name() string {
	return "proxy"
}

// ID returns a unique integer across plugins
func ID() int {
	return 3
}

// String returns a friendly version of the payload.
func String(data []byte) string {
	return string(data)
}

// Help returns documentation for the plugin.
func Help() string {
	return "proxy <url>"
}

// Generate generates a payload based on the arguments provided.
func Generate(args []string) ([]byte, error) {
	url := string(args[0])
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

// Process prepares a payload based on the arguments provided.
func Process(args []string) ([]byte, error) {
	if len(args) < 1 {
		return []byte{}, errors.New("missing required argument")
	}
	return []byte(args[0]), nil
}
