package aoc

import "testing"

func TestFindIncreases(t *testing.T) {

	type test struct {
		intput []string
		want   int
	}

	tests := []test{
		{[]string{"199", "200", "201"}, 2},
		{[]string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}, 7},
	}

	for _, testCase := range tests {
		got, err := findIncreases(testCase.intput)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != testCase.want {
			t.Errorf("wanted %d, but got %d", testCase.want, got)
		}
	}
}

func TestSlidingWindow(t *testing.T) {
	type test struct {
		intput []string
		want   int
	}

	tests := []test{
		{[]string{"199", "200", "208", "210", "200", "207", "240", "269", "260", "263"}, 5},
	}

	for _, testCase := range tests {
		got, err := findIncreasesSlidingWindow(testCase.intput)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if got != testCase.want {
			t.Errorf("wanted %d, but got %d", testCase.want, got)
		}
	}
}

func TestDay1Part1(t *testing.T) {
	input, err := loadInput("./input/day1.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	increases, err := findIncreases(input)
	t.Logf("Day 1, part 1: %d", increases)
}

func TestDay1Part2(t *testing.T) {
	input, err := loadInput("./input/day1.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	increases, err := findIncreasesSlidingWindow(input)
	t.Logf("Day 1, part 2: %d", increases)
}
