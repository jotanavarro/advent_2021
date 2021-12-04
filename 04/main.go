package main

import (
	"advent_2021/extra"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Matrix is the type that defined a bingo board.  The board has as property a boolean that marks if it has already
// won, so we can quickly dismiss it.
type Matrix struct {
	grid  [][]string
	valid bool
}

// createMatrix will return a matrix expressed as a slice of slices when fed a pointer to a scanner.
func createMatrix(scanner *bufio.Scanner) (matrix Matrix) {
	temporalMatrix := make([][]string, 0)

	for i := 0; i < 5; i++ {
		row := strings.Fields(scanner.Text())
		temporalMatrix = append(temporalMatrix, [][]string{row}...)
		scanner.Scan()
	}

	matrix.grid = temporalMatrix
	matrix.valid = true

	return matrix
}

// stitchMatrices will read the given file and generate a slice of matrices from it.
func stitchMatrices(scanner *bufio.Scanner) (matrices []Matrix) {

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		} else {
			newMatrix := createMatrix(scanner)
			matrices = append(matrices, newMatrix)
		}
	}

	return matrices
}

// evaluateWinCondition will check if the matrix is a winner one and return the result.  It will also invalidate this
// matrix so later on we can tell beforehand that it is already marked.
func (matrix *Matrix) evaluateWinCondition() bool {

	for i := 0; i < 5; i++ {
		// Checking the rows
		if matrix.grid[i][0] == "X" &&
			matrix.grid[i][1] == "X" &&
			matrix.grid[i][2] == "X" &&
			matrix.grid[i][3] == "X" &&
			matrix.grid[i][4] == "X" {

			matrix.valid = false
			return true
		}

		// Checking the columns
		if matrix.grid[0][i] == "X" &&
			matrix.grid[1][i] == "X" &&
			matrix.grid[2][i] == "X" &&
			matrix.grid[3][i] == "X" &&
			matrix.grid[4][i] == "X" {

			matrix.valid = false
			return true
		}
	}

	return false
}

// calculatePoints returns the points of a given matrix with the rules of the first exercise.
func (matrix *Matrix) calculatePoints(number int) int {
	var points int

	for _, row := range matrix.grid {
		for _, value := range row {
			if value != "X" {
				points += extra.ConvertToInt(value)
			}
		}
	}

	return points * number
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day04_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	matrices := stitchMatrices(scanner)

	for i, number := range numbers {
		for x, matrix := range matrices {
			for y, row := range matrix.grid {
				for z, value := range row {
					if value == number {
						matrices[x].grid[y][z] = "X"
					}
				}
			}

			if matrices[x].evaluateWinCondition() {
				fmt.Printf("BINGO!\nFound after %d numbers (%s) in matrix %d:\n%s\n", i, number, x, matrices[x])
				return matrices[x].calculatePoints(extra.ConvertToInt(number)), nil
			}
		}
	}

	return -1, errors.New("there is no good board")
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day04_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbers := strings.Split(scanner.Text(), ",")
	matrices := stitchMatrices(scanner)
	solution, lastNumber := Matrix{valid: true}, 0

	for _, number := range numbers {
		for x, matrix := range matrices {
			if matrix.valid {
				for y, row := range matrix.grid {
					for z, value := range row {
						if value == number {
							matrices[x].grid[y][z] = "X"
						}
					}
				}
				if matrices[x].evaluateWinCondition() {
					solution, lastNumber = matrices[x], extra.ConvertToInt(number)
				}
			}
		}
	}

	fmt.Printf("BINGO!\nFound after number %d in matrix :\n%s\n", lastNumber, solution)
	return solution.calculatePoints(lastNumber), nil
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
