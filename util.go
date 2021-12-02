package aoc

import (
	"io/ioutil"
	"strings"
)

func loadInput(path string) ([]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(bytes)), "\n"), nil
}
