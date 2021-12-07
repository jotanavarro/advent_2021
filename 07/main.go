package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func loadValues(scanner *bufio.Scanner) []int {
	scanner.Scan()
	stringInput := strings.Split(scanner.Text(), ",")
	result := make([]int, len(stringInput))

	for i, value := range stringInput {
		result[i] = extra.ConvertToInt(value)
	}

	return result
}

// calculateMedian returns a rounded value for the median of the elements in an integer slice.
func calculateMedian(positions []int) (result int) {
	tmp := make([]int, len(positions))
	copy(tmp, positions)
	sort.Ints(tmp)

	mid := len(tmp) / 2

	if mid%2 == 0 {
		return (tmp[mid-1] + tmp[mid]) / 2
	} else {
		return tmp[mid]
	}
}

// calculateMean returns a rounded mean for a given integer slice.
func calculateMean(positions []int) (result int) {
	for _, value := range positions {
		result += value
	}

	return result / len(positions)
}

func calculateFuel(positions []int, position int) (result int) {
	for i := range positions {
		result += int(math.Abs(float64(positions[i] - position)))
	}
	return
}

func calculateFuelIncrementally(positions []int, target int) (result int) {
	for _, position := range positions {
		fuelNeeded := 0
		distance := int(math.Abs(float64(target - position)))

		// We need to stack extra fuel consumption as it does not grow linearly anymore.
		for d := 1; d <= distance; d++ {
			fuelNeeded += d
		}
		result += fuelNeeded
	}
	return result
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day07_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	// We use the median as its value separates in two equal half the values in the provided array.
	input := loadValues(bufio.NewScanner(file))
	median := calculateMedian(input)
	return calculateFuel(input, median), nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day07_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	// Since we are vulnerable to the extremes (They cost a lot of fuel) we can use the mean here, which is the average
	// position for the crabs.  There is some margin of error here because we are rounding it, on a bright note we were
	// lucky enough that the rounded value was good to obtain the solution of the exercise, otherwise we should have
	// tried with the floor/ceiling of the float value of the mean.
	input := loadValues(bufio.NewScanner(file))
	mean := calculateMean(input)
	return calculateFuelIncrementally(input, mean), nil
}

func main() {
	firstResult, err := firstExercise()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The firstResult of the first exercise is: %d.\n", firstResult)
	}

	secondResult, err := secondExercise()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The firstResult of the second exercise is: %d.\n", secondResult)
	}
}
