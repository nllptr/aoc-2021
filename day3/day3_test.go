package day3

import (
	"reflect"
	"testing"

	"github.com/nllptr/aoc-2021/utils"
)

func TestParseGammaResult(t *testing.T) {
	input := []gammaCount{
		{5, 4},
		{3, 4},
		{1, 0},
	}
	result := parseGammaResult(input)
	if result != "101" {
		t.Errorf("expected result to be '101', but it was %s", result)
	}
}

func TestFindGamma(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	gammaBinary, err := findGammaBinary(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if gammaBinary != "10110" {
		t.Errorf("expected gamma to be 10110, but it was %s", gammaBinary)
	}
}

func TestBinaryToDecimal(t *testing.T) {
	decimal, err := binaryToInt("10110")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if decimal != 22 {
		t.Errorf("expected 10110 to be 22, but it was %d", decimal)
	}
}

func TestInvertBinary(t *testing.T) {
	input := "111000"
	result := invertBinary(input)
	if result != "000111" {
		t.Errorf("expected result to be 000111, but it was %s", result)
	}
}

func TestCalculatePowerConsumption(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	powerConsumption, err := CalculatePowerConsupmtion(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if powerConsumption != 198 {
		t.Errorf("expected power consumption to be 198, but it was %d", powerConsumption)
	}
}

func TestFilterMostCommon(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	expected := []string{
		"11110",
		"10110",
		"10111",
		"10101",
		"11100",
		"10000",
		"11001",
	}
	result := filterMostCommon(input, 0)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected result to be %v, but got %v", expected, result)
	}
}

func TestGetOxygenRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result := getOxygenRating(input)
	if result != "10111" {
		t.Errorf("expected oxygen rating to be 10111, but got %s", result)
	}
	decimal, err := binaryToInt(result)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if decimal != 23 {
		t.Errorf("expected decimal to be 23, but got %d", decimal)
	}
}

func TestGetCO2Rating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	result := getCO2Rating(input)
	if result != "01010" {
		t.Errorf("expected CO2 rating to be 01010, but got %s", result)
	}
	decimal, err := binaryToInt(result)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if decimal != 10 {
		t.Errorf("expected decimal to be 10, but got %d", decimal)
	}
}

func TestCalculateLifeSupportRating(t *testing.T) {
	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}
	rating, err := CalculateLifeSupportRating(input)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if rating != 230 {
		t.Errorf("expected rating to be 230, but got %d", rating)
	}
}

func TestDay2Part1(t *testing.T) {
	input, err := utils.LoadInput("input.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	powerConsumption, err := CalculatePowerConsupmtion(input)
	t.Logf("Day 3, part 1: %d", powerConsumption)
}

func TestDay2Part2(t *testing.T) {
	input, err := utils.LoadInput("input.txt")
	if err != nil {
		t.Error("failed to parse file")
	}
	lifeSupportRating, err := CalculateLifeSupportRating(input)
	t.Logf("Day 3, part 2: %d", lifeSupportRating)
}
