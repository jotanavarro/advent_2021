package main

import (
	"advent_2021/extra"
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
)

type Octopus struct {
	energy  int
	flashed bool
}

func loadInput(scanner *bufio.Scanner) [][]Octopus {
	result := make([][]Octopus, 0)

	for scanner.Scan() {
		line := scanner.Text()
		tmpSlice := make([]Octopus, len(line))
		for i, val := range line {
			tmpSlice[i] = Octopus{energy: extra.ConvertToInt(string(val)), flashed: false}
		}

		result = append(result, tmpSlice)
	}

	return result
}

//restartFlashMemory sets to logical false all the Octopus on the board.
func restartFlashMemory(matrix *[][]Octopus) {
	for x := 0; x < len(*matrix); x++ {
		for y := 0; y < len((*matrix)[x]); y++ {
			(*matrix)[x][y].flashed = false
		}
	}
}

// setLimits will return bounded values to traverse adjacent elements to a given position in the provided matrix, this
// allows us to loop through these adjacent positions without the risk of going out of bounds.
// NOTE: Copied form my solution to day 9.
func setLimits(matrix *[][]Octopus, x int, y int) (int, int, int, int) {
	startX, endX := x-1, x+1
	if startX < 0 {
		startX = x
	}
	if endX == len(*matrix) {
		endX = x
	}

	startY, endY := y-1, y+1
	if startY < 0 {
		startY = y
	}
	if endY == len((*matrix)[0]) {
		endY = y
	}
	return startX, endX, startY, endY
}

//splashFlash will increase the energy levels of all surrounding octopuses to a given location.
func splashFlash(matrix *[][]Octopus, x, y int) {
	startX, endX, startY, endY := setLimits(matrix, x, y)
	for i := startX; i <= endX; i++ {
		for q := startY; q <= endY; q++ {
			// We only raise the energy level of a non flashed octopus.
			if (*matrix)[i][q].flashed == false {
				(*matrix)[i][q].energy++
			}
		}
	}
}

//increaseEnergy will iterate over the board and increase energy accordingly.
func increaseEnergy(matrix *[][]Octopus) {
	for x := 0; x < len(*matrix); x++ {
		for y := 0; y < len((*matrix)[x]); y++ {
			(*matrix)[x][y].energy++
		}
	}
}

//triggerFlash goes over a given matrix of Octopus and trigger the flash on each one that has more than 9 energy and has
//not flashed already.
func triggerFlash(matrix *[][]Octopus) int {
	flashes := 0
	triggered := make([][]int, 0)

	for x := 0; x < len(*matrix); x++ {
		for y := 0; y < len((*matrix)[x]); y++ {
			currentOctopus := &(*matrix)[x][y]
			if currentOctopus.energy > 9 && currentOctopus.flashed == false {
				flashes++
				currentOctopus.flashed = true
				currentOctopus.energy = 0
				triggered = append(triggered, []int{x, y})
			}
		}
	}

	if flashes == 0 {
		return flashes
	} else {

		for i := range triggered {
			splashFlash(matrix, triggered[i][0], triggered[i][1])
		}

		flashes += triggerFlash(matrix)
		return flashes
	}
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day11_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	octopuses := loadInput(bufio.NewScanner(file))
	steps := 100
	flashes := 0

	for i := 0; i < steps; i++ {
		increaseEnergy(&octopuses)
		flashes += triggerFlash(&octopuses)
		restartFlashMemory(&octopuses)
	}

	return flashes, nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day11_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	octopuses := loadInput(bufio.NewScanner(file))
	numberOfOctopuses := len(octopuses) * len(octopuses[0])

	for i := 1; i < math.MaxInt; i++ {
		increaseEnergy(&octopuses)
		if triggerFlash(&octopuses) == numberOfOctopuses {
			return i, nil
		}
		restartFlashMemory(&octopuses)
	}

	return -1, errors.New("no solution found")
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
