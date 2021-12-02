package utils

import "testing"

func TestLoadInput(t *testing.T) {
	input, err := LoadInput("test_input.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	if len(input) != 2 {
		t.Errorf("expected error to be of length 2, but got %v", input)
	}
	if input[0] != "Hello" && input[1] != "World" {
		t.Errorf("expected input to contain [Hello, World], bu tgot %v", input)
	}
}
