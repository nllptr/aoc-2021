package day3

import (
	"strconv"
	"strings"
)

type gammaCount struct {
	ones   int
	zeroes int
}

func parseGammaResult(input []gammaCount) string {
	result := ""
	for _, bit := range input {
		if bit.ones > bit.zeroes {
			result = result + "1"
		} else {
			result = result + "0"
		}
	}
	return result
}

func binaryToInt(input string) (int, error) {
	decimal, err := strconv.ParseInt(input, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(decimal), nil
}

func findGammaBinary(input []string) (string, error) {
	rowLength := len(input[0])
	result := make([]gammaCount, rowLength)
	for _, row := range input {
		for bitIndex, bit := range strings.TrimSpace(row) {
			if bit == '1' {
				result[bitIndex].ones++
			} else {
				result[bitIndex].zeroes++
			}
		}
	}

	return parseGammaResult(result), nil
}

func invertBinary(input string) string {
	result := ""
	for _, bit := range input {
		if bit == '1' {
			result = result + "0"
		} else {
			result = result + "1"
		}
	}
	return result
}

func filterMostCommon(input []string, index int) []string {
	count := gammaCount{}
	for _, row := range input {
		if row[index] == '1' {
			count.ones++
		} else {
			count.zeroes++
		}
	}
	var mostCommon rune
	if count.ones >= count.zeroes {
		mostCommon = '1'
	} else {
		mostCommon = '0'
	}

	result := []string{}
	for _, row := range input {
		if row[index] == byte(mostCommon) {
			result = append(result, row)
		}
	}
	return result
}

func getOxygenRating(input []string) string {
	result := input
	for i := 0; i < len(input[0]); i++ {
		result = filterMostCommon(result, i)
		if len(result) == 1 {
			break
		}
	}
	return result[0]
}

func filterLeastCommon(input []string, index int) []string {
	count := gammaCount{}
	for _, row := range input {
		if row[index] == '1' {
			count.ones++
		} else {
			count.zeroes++
		}
	}
	var leastCommon rune
	if count.zeroes <= count.ones {
		leastCommon = '0'
	} else {
		leastCommon = '1'
	}

	result := []string{}
	for _, row := range input {
		if row[index] == byte(leastCommon) {
			result = append(result, row)
		}
	}
	return result
}

func getCO2Rating(input []string) string {
	result := input
	for i := 0; i < len(input[0]); i++ {
		result = filterLeastCommon(result, i)
		if len(result) == 1 {
			break
		}
	}
	return result[0]
}

func CalculatePowerConsupmtion(input []string) (int, error) {
	gammaBinary, err := findGammaBinary(input)
	if err != nil {
		return 0, err
	}
	epsilonBinary := invertBinary(gammaBinary)
	gammaDecimal, err := binaryToInt(gammaBinary)
	if err != nil {
		return 0, err
	}
	epsilonDecimal, err := binaryToInt(epsilonBinary)
	if err != nil {
		return 0, err
	}
	return gammaDecimal * epsilonDecimal, nil
}

func CalculateLifeSupportRating(input []string) (int, error) {
	oxygenBinary := getOxygenRating(input)
	oxygen, err := binaryToInt(oxygenBinary)
	if err != nil {
		return 0, err
	}
	co2Binary := getCO2Rating(input)
	co2, err := binaryToInt(co2Binary)
	if err != nil {
		return 0, err
	}
	return oxygen * co2, nil
}
