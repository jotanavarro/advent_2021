package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

// loadMatrix will split the input into an array of integer arrays, so we can easily process the data.
func loadMatrix(scanner *bufio.Scanner) [][]int {
	result := make([][]int, 0)

	for scanner.Scan() {
		currentLine := scanner.Text()
		temporalArray := make([]int, len(currentLine))

		for i := range currentLine {
			temporalArray[i] = extra.ConvertToInt(string(currentLine[i]))
		}

		result = append(result, temporalArray)
	}

	return result
}

// validatePosition returns true if the element chosen in a matrix is lesser than its adjacent ones.
func validatePosition(matrix [][]int, x int, y int) (result bool) {
	startX, endX, startY, endY := setLimits(matrix, x, y)
	result = true

	for i := startX; i <= endX; i++ {
		for q := startY; q <= endY; q++ {
			validation := !(x == i && y == q)
			if validation {
				source := matrix[x][y]
				target := matrix[i][q]
				result = result && (source < target)
			}
		}
	}

	return result
}

// setLimits will return bounded values to traverse adjacent elements to a given position in the provided matrix, this
// allows us to loop through these adjacent positions without the risk of going out of bounds.
func setLimits(matrix [][]int, x int, y int) (int, int, int, int) {
	startX, endX := x-1, x+1
	if startX < 0 {
		startX = x
	}
	if endX == len(matrix) {
		endX = x
	}

	startY, endY := y-1, y+1
	if startY < 0 {
		startY = y
	}
	if endY == len(matrix[0]) {
		endY = y
	}
	return startX, endX, startY, endY
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day09_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	matrix := loadMatrix(bufio.NewScanner(file))
	riskLevel := 0

	for i := range matrix {
		for q := range matrix[i] {
			if validatePosition(matrix, i, q) {
				riskLevel += matrix[i][q] + 1
			}
		}
	}

	return riskLevel, nil
}

// findNeighbours will check the top, left, right and bottom locations to a given point and evaluate if they belong to
// the basin of the given point.
func findNeighbours(matrix [][]int, point Point) (result []Point) {
	result = []Point{point}
	startX, endX, startY, endY := setLimits(matrix, point.x, point.y)
	pointValue := matrix[point.x][point.y]

	if startX != point.x {
		targetValue := matrix[startX][point.y]
		if targetValue > pointValue && targetValue != 9 {
			result = append(result, Point{x: startX, y: point.y})
		}
	}
	if endX != point.x {
		targetValue := matrix[endX][point.y]
		if targetValue > pointValue && targetValue != 9 {
			result = append(result, Point{x: endX, y: point.y})
		}
	}
	if startY != point.y {
		targetValue := matrix[point.x][startY]
		if targetValue > pointValue && targetValue != 9 {
			result = append(result, Point{x: point.x, y: startY})
		}
	}
	if endY != point.y {
		targetValue := matrix[point.x][endY]
		if targetValue > pointValue && targetValue != 9 {
			result = append(result, Point{x: point.x, y: endY})
		}
	}

	return result
}

// pointInBasin will simply return if a given Point exists in a provided slice of Point.
func pointInBasin(basin []Point, point Point) bool {
	for _, current := range basin {
		if current.x == point.x && current.y == point.y {
			return true
		}
	}
	return false
}

// removeDuplicates removes all duplicated Point form a provided Point slice.
func removeDuplicates(basin []Point) []Point {
	collection := make(map[Point]int)
	result := make([]Point, 0)

	for i := 0; i < len(basin); i++ {
		collection[basin[i]]++
	}

	for key := range collection {
		result = append(result, key)
	}

	return result
}

// recursiveBasinDetection will find out, for a given point, which new adjacent points belong to the basin
// and recursively extend to the adjacent ones to those.
func recursiveBasinDetection(matrix [][]int, basin []Point, point Point) []Point {
	candidates := findNeighbours(matrix, point)

	// We clean the candidates slice by dropping duplicates we already have in the basin.
	for i := 0; i < len(candidates); i++ {
		candidate := candidates[i]
		if pointInBasin(basin, candidate) {
			candidates = append(candidates[:i], candidates[i+1:]...)
		}
	}

	// For each potential candidate, we explore the basin in that direction.
	tmpCandidates := make([]Point, len(candidates))
	copy(tmpCandidates, candidates)
	for i := 0; i < len(tmpCandidates); i++ {
		candidate := tmpCandidates[i]
		if candidate.x == point.x && candidate.y == point.y {
			continue
		}
		candidates = removeDuplicates(append(candidates, recursiveBasinDetection(matrix, append(basin, candidates...), candidate)...))
	}

	return candidates
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day09_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	matrix := loadMatrix(bufio.NewScanner(file))
	lowPoints := make([]Point, 0)
	basinSizes := make([]int, 0)
	result := 1

	for i := range matrix {
		for q := range matrix[i] {
			if validatePosition(matrix, i, q) {
				lowPoints = append(lowPoints, Point{x: i, y: q})
			}
		}
	}

	for i := range lowPoints {
		basinTopography := recursiveBasinDetection(matrix, make([]Point, 0), lowPoints[i])
		basinSizes = append(basinSizes, len(basinTopography))
	}

	// We sort our basin sizes and get the largest ones.
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	for i := 0; i < len(basinSizes) && i < 3; i++ {
		result *= basinSizes[i]
	}

	return result, nil
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
