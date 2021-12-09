package main

import (
	"advent_2021/extra"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func firstExercise() (int64, error) {
	file, err := os.Open("inputs/day03_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	var (
		counter       [12]int
		gammaString   string
		epsilonString string
		scanner       = bufio.NewScanner(file)
	)

	// We count the number or 1s and 0s per column
	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			if char == '1' {
				counter[pos]++
			} else {
				counter[pos]--
			}
		}
	}

	// We assemble the binary number with the values obtained before.  A positive number means there are more 1s while
	// a negative one means more 0s.
	for _, num := range counter {
		if num > 0 {
			gammaString += "1"
			epsilonString += "0"
		} else if num < 0 {
			gammaString += "0"
			epsilonString += "1"
		} else {
			return -1, errors.New("Same number of 1s and 0s.  Wrong input!")
		}
	}

	gammaValue, err := strconv.ParseInt(gammaString, 2, 0)
	if err != nil {
		return -1, err
	}

	epsilonValue, err := strconv.ParseInt(epsilonString, 2, 0)
	if err != nil {
		return -1, err
	}

	return gammaValue * epsilonValue, nil
}

func secondExercise() (int64, error) {
	file, err := os.Open("inputs/day03_exercise02.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	var (
		archive []string
		scanner = bufio.NewScanner(file)
	)

	// We store the contents of the file in an array, we will as well already count the digits in the first position
	for scanner.Scan() {
		archive = append(archive, scanner.Text())
	}

	binaryOxygenRating, err := findOxygenRating(archive, 0)
	if err != nil {
		return -1, err
	}

	binaryCO2Rating, err := findCO2Rating(archive, 0)
	if err != nil {
		return -1, err
	}

	oxygenRating, err := strconv.ParseInt(binaryOxygenRating[0], 2, 0)
	if err != nil {
		return -1, err
	}

	CO2Rating, err := strconv.ParseInt(binaryCO2Rating[0], 2, 0)
	if err != nil {
		return -1, err
	}

	return oxygenRating * CO2Rating, nil
}

// classify returns two arrays of strings, dividing elements depending on if they contain 0 or 1 in the given position.
// The first returned array is selected for 0s, while the other selects for 1s.
func classify(list []string, position int) ([]string, []string) {
	var beginWithZero []string
	var beginWithOne []string

	for _, elem := range list {
		if elem[position] == '0' {
			beginWithZero = append(beginWithZero, elem)
		} else {
			beginWithOne = append(beginWithOne, elem)
		}
	}

	return beginWithZero, beginWithOne
}

// findOxygenRating recursively filters a given string array, checking a given position on each of its strings.
// In this case, we select the most popular value in the given position.
func findOxygenRating(list []string, position int) ([]string, error) {
	if len(list) == 1 {
		return list, nil
	} else {
		beginWithZero, beginWithOne := classify(list, position)
		newPosition := position + 1

		if len(beginWithOne) >= len(beginWithZero) {
			return findOxygenRating(beginWithOne, newPosition)
		} else {
			return findOxygenRating(beginWithZero, newPosition)
		}
	}
}

// findCO2Rating recursively filters a given string array, checking a given position on each of its strings.
// In this case, we select the least popular value in the given position.
func findCO2Rating(list []string, position int) ([]string, error) {
	if len(list) == 1 {
		return list, nil
	} else {
		beginWithZero, beginWithOne := classify(list, position)
		newPosition := position + 1

		if len(beginWithOne) < len(beginWithZero) {
			return findCO2Rating(beginWithOne, newPosition)
		} else {
			return findCO2Rating(beginWithZero, newPosition)
		}
	}
}

func main() {
	firstResult, err := firstExercise()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The result of the first exercise is: %d.\n", firstResult)
	}

	secondResult, err := secondExercise()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The result of the second exercise is: %d.\n", secondResult)
	}
}
