package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

// loadInput loads the puzzle input.  We will for simplicity avoid registering paths to the "start" node, as for no
// possible path we would like to visit that node.
func loadInput(scanner *bufio.Scanner) map[string][]string {
	result := make(map[string][]string)

	for scanner.Scan() {
		tmpLine := strings.Split(scanner.Text(), "-")
		leftNode := tmpLine[0]
		rightNode := tmpLine[1]

		result[leftNode] = append(result[leftNode], rightNode)
		result[rightNode] = append(result[rightNode], leftNode)
	}

	for key, connections := range result {
		tmpCon := make([]string, 0)
		tmpMap := make(map[string]bool)

		for _, conn := range connections {
			if _, ok := tmpMap[conn]; !ok && conn != "start" {
				tmpMap[conn] = true
				tmpCon = append(tmpCon, conn)
			}
		}

		result[key] = tmpCon
	}
	return result
}

var largeCave, _ = regexp.Compile("[A-Z]+")
var smallCave, _ = regexp.Compile("[a-z]+")

//existsInSlice simply lets us know if a value exists in a slice.
func existsInSlice(value string, slice []string) (result bool) {
	for _, elem := range slice {
		if value == elem {
			return true
		}
	}

	return result
}

// firstSolution will return all the possible paths from a given path, in which no small cave (lowercase letter node) is
// visited more than once.
func firstSolution(data map[string][]string, path []string) (solution int) {
	for _, value := range data[path[len(path)-1]] {
		if largeCave.MatchString(value) || !existsInSlice(value, path) {
			if value == "end" {
				solution += 1
			} else {
				tmp := make([]string, len(path))
				copy(tmp, path)

				tmp = append(tmp, value)
				solution += firstSolution(data, tmp)
			}
		}
	}

	return solution
}

// secondSolution will return all the possible paths from a given path, allowing to repeat small caves once per path. If
// we try to append a small cave to the path we already visited, we complete the path from that point using the result
// from firstSolution.
func secondSolution(data map[string][]string, path []string) (solution int) {
	for _, value := range data[path[len(path)-1]] {
		if value == "end" {
			solution += 1
		} else {
			tmp := make([]string, len(path))
			copy(tmp, path)

			tmp = append(tmp, value)

			a := smallCave.MatchString(value)
			b := existsInSlice(value, path)

			if a && b {
				solution += firstSolution(data, tmp)
			} else {
				solution += secondSolution(data, tmp)
			}

		}
	}

	return solution
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day12_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	graph := loadInput(bufio.NewScanner(file))

	fmt.Println(graph)

	return firstSolution(graph, []string{"start"}), nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day12_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	graph := loadInput(bufio.NewScanner(file))

	fmt.Println(graph)

	return secondSolution(graph, []string{"start"}), nil
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
