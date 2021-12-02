package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Direction int

const (
	FORWARD Direction = iota
	UP
	DOWN
)

func convertCommand(command string) (Direction, int, error) {
	split := strings.Split(command, " ")
	if len(split) != 2 {
		return 0, 0, fmt.Errorf("command string is too short: %s", command)
	}

	direction := FORWARD
	switch split[0] {
	case "forward":
		direction = FORWARD
	case "up":
		direction = UP
	case "down":
		direction = DOWN
	default:
		return 0, 0, fmt.Errorf("could not parse a valid direction: %s", split[0])
	}

	amount, err := strconv.Atoi(split[1])
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse a valid amaount: %s", split[1])
	}
	return direction, amount, nil
}

func moveAndMultiply(input []string) (int, error) {
	horizontal := 0
	depth := 0
	for _, command := range input {
		direction, amount, err := convertCommand(command)
		if err != nil {
			return 0, err
		}
		switch direction {
		case FORWARD:
			horizontal += amount
		case UP:
			depth -= amount
		case DOWN:
			depth += amount
		}
	}
	return horizontal * depth, nil
}

func moveWithAim(input []string) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0
	for _, command := range input {
		direction, amount, err := convertCommand(command)
		if err != nil {
			return 0, err
		}
		switch direction {
		case FORWARD:
			horizontal += amount
			depth += aim * amount
		case UP:
			aim -= amount
		case DOWN:
			aim += amount
		}
	}
	return horizontal * depth, nil
}
