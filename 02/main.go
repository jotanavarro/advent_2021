package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day02_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	reg, err := regexp.Compile("^(?P<command>[a-z]+) (?P<distance>[0-9]+)$")
	if err != nil {
		return -1, err
	}

	var (
		scanner                   = bufio.NewScanner(file)
		horizontalDistance, depth int
	)

	for scanner.Scan() {
		match := reg.FindStringSubmatch(scanner.Text())
		var command, distance = match[1], extra.ConvertToInt(match[2])

		switch command {
		case "forward":
			horizontalDistance += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	return horizontalDistance * depth, nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day02_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	reg, err := regexp.Compile("^(?P<command>[a-z]+) (?P<units>[0-9]+)$")
	if err != nil {
		return -1, err
	}

	var (
		scanner                                = bufio.NewScanner(file)
		horizontalDistance, inclination, depth int
	)

	for scanner.Scan() {
		match := reg.FindStringSubmatch(scanner.Text())
		var command, units = match[1], extra.ConvertToInt(match[2])

		// We choose the `switch` over an `if`-`else` as it is more performant in this case.
		switch command {
		case "forward":
			horizontalDistance += units
			depth += units * inclination
		case "up":
			inclination -= units
		case "down":
			inclination += units
		}
	}

	return horizontalDistance * depth, nil
}

func main() {
	firstResult, err := firstExercise()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The distance is: %d.\n", firstResult)
	}

	secondResult, err := secondExercise()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("The distance is: %d.\n", secondResult)
	}
}
