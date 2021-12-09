package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// parsePoints will extract two Point from a line of text from our input.
func parsePoints(scanner *bufio.Scanner) (origin Point, destination Point) {
	rawInput := strings.Split(scanner.Text(), " -> ")
	rawOrigin := strings.Split(rawInput[0], ",")
	rawDestination := strings.Split(rawInput[1], ",")

	origin.x, origin.y = extra.ConvertToInt(rawOrigin[0]), extra.ConvertToInt(rawOrigin[1])
	destination.x, destination.y = extra.ConvertToInt(rawDestination[0]), extra.ConvertToInt(rawDestination[1])

	return origin, destination
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day05_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	scanner := bufio.NewScanner(file)
	diagram := Diagram{height: 0, width: 0, grid: make([][]int, 0)}

	for scanner.Scan() {
		origin, destination := parsePoints(scanner)
		diagram.drawLine(origin, destination, false)
	}

	//diagram.drawDiagram()
	return diagram.calculateDangerousPoints(), nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day05_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)
	scanner := bufio.NewScanner(file)
	diagram := Diagram{height: 0, width: 0, grid: make([][]int, 0)}

	for scanner.Scan() {
		origin, destination := parsePoints(scanner)
		diagram.drawLine(origin, destination, true)
	}

	//diagram.drawDiagram()
	return diagram.calculateDangerousPoints(), nil
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
