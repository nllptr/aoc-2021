package aoc

import "strconv"

func findIncreases(input []string) (int, error) {
	numerics := []int{}
	for _, item := range input {
		numeric, err := strconv.Atoi(item)
		if err != nil {
			return 0, err
		}
		numerics = append(numerics, numeric)
	}

	increases := 0
	for i := 1; i < len(numerics); i++ {
		if numerics[i] > numerics[i-1] {
			increases++
		}
	}
	return increases, nil
}

func findIncreasesSlidingWindow(input []string) (int, error) {
	numerics := []int{}
	for _, item := range input {
		numeric, err := strconv.Atoi(item)
		if err != nil {
			return 0, err
		}
		numerics = append(numerics, numeric)
	}

	increases := 0
	for i := 3; i < len(numerics); i++ {
		if windowSum(&numerics, i) > windowSum(&numerics, i-1) {
			increases++
		}
	}
	return increases, nil
}

func windowSum(input *[]int, index int) int {
	return (*input)[index] + (*input)[index-1] + (*input)[index-2]
}
