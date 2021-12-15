package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Paper struct {
	matrix [][]int
}

func (p *Paper) countDots() (result int) {
	for _, row := range p.matrix {
		for _, val := range row {
			if val > 0 {
				result++
			}
		}
	}

	return result
}

func (p *Paper) printPaper() {
	fmt.Printf("\n")
	for _, row := range p.matrix {
		for _, value := range row {
			character := "."
			if value > 0 {
				character = "#"
			}

			fmt.Printf("%s ", character)
		}
		fmt.Printf("\n")
	}
}

func (p *Paper) foldOverY(position int) {
	tmp := make([][]int, len(p.matrix[:position]))
	for i := range tmp {
		tmp[i] = make([]int, len(p.matrix[0]))
	}

	// Upper part of the fold
	for i := 0; i < position; i++ {
		for q := 0; q < len(p.matrix[0]); q++ {
			tmp[i][q] += p.matrix[i][q]
		}
	}

	// Lower part of the fold
	for i := 0; (i + position + 1) < len(p.matrix); i++ {
		for q := 0; q < len(p.matrix[0]); q++ {
			foldedDelta := position - 1 - i
			toFoldDelta := i + position + 1

			tmp[foldedDelta][q] += p.matrix[toFoldDelta][q]
		}
	}

	p.matrix = tmp
}

// Broken
func (p *Paper) foldOverX(position int) {
	tmp := make([][]int, len(p.matrix))
	for i := range tmp {
		tmp[i] = make([]int, position)
	}

	// Left part of the fold.
	for i := range tmp {
		for q := 0; q < position; q++ {
			tmp[i][q] += p.matrix[i][q]
		}
	}

	// Right part of the fold.
	for i := range tmp {
		for q := 0; (q + position + 1) < len(p.matrix[0]); q++ {
			foldedDelta := position - 1 - q
			toFoldDelta := q + position + 1

			tmp[i][foldedDelta] += p.matrix[i][toFoldDelta]
		}
	}

	p.matrix = tmp
}

func (p *Paper) interpretOrder(order string) {
	parsedOrders := strings.Split(order, "=")
	axis := parsedOrders[0]
	value := extra.ConvertToInt(parsedOrders[1])

	if axis == "x" {
		p.foldOverX(value)
	} else {
		p.foldOverY(value)
	}
}

//loadInput returns a matrix of integers representing the sheet of transparent paper, displaying 0s as `.` and 1s as `#`
// this function will also return a slice of instructions about how to fold the paper.
func loadInput(scanner *bufio.Scanner) (paper *Paper, instructions []string) {
	instructions = make([]string, 0)

	coordinatesX := make([]int, 0)
	coordinatesY := make([]int, 0)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")

		if len(text) == 2 {
			// This is a matrix point.
			x, y := extra.ConvertToInt(text[0]), extra.ConvertToInt(text[1])

			coordinatesX = append(coordinatesX, x)
			coordinatesY = append(coordinatesY, y)

		} else if text[0] != "" {
			// This is a folding instruction.
			instructions = append(instructions, strings.Split(text[0], " ")[2])
		}

	}

	limitX := findLargest(coordinatesX) + 1
	limitY := findLargest(coordinatesY) + 1
	paperMatrix := make([][]int, limitY)
	for i := range paperMatrix {
		paperMatrix[i] = make([]int, limitX)
	}

	for i, x := range coordinatesX {
		y := coordinatesY[i]
		paperMatrix[y][x] = 1
	}

	return &Paper{matrix: paperMatrix}, instructions
}

func findLargest(numbers []int) (result int) {
	for _, elem := range numbers {
		if elem > result {
			result = elem
		}
	}

	return result
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day13_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	paper, orders := loadInput(bufio.NewScanner(file))

	paper.interpretOrder(orders[0])

	return paper.countDots(), nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day13_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)
	paper, orders := loadInput(bufio.NewScanner(file))

	for _, order := range orders {
		paper.interpretOrder(order)
	}

	paper.printPaper()

	return paper.countDots(), nil
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
