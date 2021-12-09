package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// sortIntoFishPool Given an array with the program's input, an array of strings.  It will produce another integer
// array where each position represents the number of fish in that state of their reproductive cycle.  For example when
// a fish is in the position 0, it means it will reproduce in the next cycle.
func sortIntoFishPool(input []string) (result [9]int) {
	for i := range input {
		result[extra.ConvertToInt(input[i])]++
	}

	return result
}

// countFish will just iterate over the array of fish and aggregate how many are on each life cycle.
func countFish(input [9]int) (result int) {
	for i := 0; i < len(input); i++ {
		result += input[i]
	}

	return result
}

// fishLife will simulate any number of days in their life, given a buffer from where to read the initial population.
func fishLife(scanner *bufio.Scanner, days int) [9]int {
	input := strings.Split(scanner.Text(), ",")
	fishPool := sortIntoFishPool(input)

	for i := 0; i < days; i++ {
		//fmt.Printf("Day %d:\t%+v\n", i, fishPool)
		var tmpPool [9]int

		for i := 0; i < 8; i++ {
			tmpPool[i] = fishPool[i+1]
		}

		tmpPool[6] += fishPool[0]
		tmpPool[8] = fishPool[0]

		fishPool = tmpPool
	}

	return fishPool
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day06_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	fishPool := fishLife(scanner, 80)

	return countFish(fishPool), nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day06_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	fishPool := fishLife(scanner, 256)

	return countFish(fishPool), nil
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
