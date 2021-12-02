package aoc

import "testing"

func TestConvertToCommand(t *testing.T) {
	direction, amount, err := convertCommand("forward 3")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if direction != FORWARD {
		t.Errorf("expected direction to be FORWARD, but got %v", direction)
	}
	if amount != 3 {
		t.Errorf("expected amount to be 3, but got %d", amount)
	}
}

func TestMoveAndMultiply(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	product, err := moveAndMultiply(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if product != 150 {
		t.Errorf("expected product to be 150, but got %d", product)
	}
}

func TestMoveWithAim(t *testing.T) {
	input := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}
	product, err := moveWithAim(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if product != 900 {
		t.Errorf("expected product to be 900, but got %d", product)
	}
}

func TestDay2Part1(t *testing.T) {
	input, err := loadInput("./input/day2.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	product, err := moveAndMultiply(input)
	t.Logf("Day 2, part 1: %d", product)
}

func TestDay2Part2(t *testing.T) {
	input, err := loadInput("./input/day2.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	product, err := moveWithAim(input)
	t.Logf("Day 2, part 2: %d", product)
}
