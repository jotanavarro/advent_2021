package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
)

func firstExercise() int {
	file, err := os.Open("inputs/day01_exercise01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	var result int

	// We bootstrap the variables before looping through the input.
	var (
		previous int
		current  = extra.ConvertToInt(scanner.Text())
	)

	for scanner.Scan() {
		previous, current = current, extra.ConvertToInt(scanner.Text())
		if previous < current {
			result++
		}
	}

	return result
}

func secondExercise() int {
	file, err := os.Open("inputs/day01_exercise02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var result int

	// We populate the first three values to bootstrap the loop
	var first int
	scanner.Scan()
	var second = extra.ConvertToInt(scanner.Text())
	scanner.Scan()
	var third = extra.ConvertToInt(scanner.Text())
	scanner.Scan()
	var fourth = extra.ConvertToInt(scanner.Text())

	// Now we can easily traverse the file and generate all the subsegments.
	for scanner.Scan() {
		first, second, third, fourth = second, third, fourth, extra.ConvertToInt(scanner.Text())
		firstSegment := first + second + third
		secondSegment := second + third + fourth
		if firstSegment < secondSegment {
			result++
		}
	}

	return result
}

func main() {
	fmt.Printf("First exercise answer: %d.\n", firstExercise())
	fmt.Printf("Second exercise answer: %d.\n", secondExercise())
}
