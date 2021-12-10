package main

import (
	"advent_2021/dataStructures"
	"advent_2021/extra"
	"bufio"
	"container/heap"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

// Defining the interface of a priority queue.

func loadInput(scanner *bufio.Scanner) (result []string) {
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

// invertCharacter will return the closing character for a given (, {, [ or <.
func invertCharacter(c rune) (result rune, err error) {
	switch c {
	case '(':
		result = ')'
	case '[':
		result = ']'
	case '{':
		result = '}'
	case '<':
		result = '>'
	default:
		return '0', errors.New("not matching character")
	}
	return result, nil
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day10_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	var points = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	input := loadInput(bufio.NewScanner(file))
	scoreTable := map[rune]int{
		')': 0,
		']': 0,
		'}': 0,
		'>': 0,
	}
	score := 0

	for line := range input {
		rpq := make(dataStructures.RunePriorityQueue, 0)
		heap.Init(&rpq)

		for i, r := range input[line] {
			if r == '(' || r == '[' || r == '{' || r == '<' {
				inverted, err := invertCharacter(r)
				if err != nil {
					log.Fatal(err)
				}
				item := &dataStructures.Item{
					Value:    inverted,
					Priority: i,
				}
				heap.Push(&rpq, item)
			} else {
				required := heap.Pop(&rpq).(*dataStructures.Item)
				// If the enclosing character does NOT match the previous one, we consider this line corrupted, annotate
				// the score and keep going.
				if r != required.Value {
					scoreTable[r]++
					break
				}
			}
		}
	}

	for i := range scoreTable {
		score += scoreTable[i] * points[i]
	}

	return score, nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day10_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	var points = map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	input := loadInput(bufio.NewScanner(file))
	scoreTable := make([]int, 0)

	for line := range input {
		rpq := make(dataStructures.RunePriorityQueue, 0)
		lineRunes := make([]rune, 0)
		lineScore := 0
		corrupted := false
		heap.Init(&rpq)

		for i, r := range input[line] {
			// First we check the line, accumulate the symbols required and discard the corrupted ones.
			if r == '(' || r == '[' || r == '{' || r == '<' {
				inverted, err := invertCharacter(r)
				if err != nil {
					log.Fatal(err)
				}
				item := &dataStructures.Item{
					Value:    inverted,
					Priority: i,
				}
				heap.Push(&rpq, item)
			} else {
				required := heap.Pop(&rpq).(*dataStructures.Item)
				// The line is corrupted if the required symbol does not match.
				if r != required.Value {
					corrupted = true
					break
				}
			}
		}
		if !corrupted {
			// Now we traverse the required symbol queue and store those we need to make valid lines.
			for rpq.Len() > 0 {
				item := heap.Pop(&rpq).(*dataStructures.Item)
				lineRunes = append(lineRunes, item.Value)
			}

			// We calculate the score for this line using the second exercise rules.
			for i := range lineRunes {
				lineScore = lineScore*5 + points[lineRunes[i]]
			}

			scoreTable = append(scoreTable, lineScore)
		}
	}

	sort.Ints(scoreTable)

	return scoreTable[len(scoreTable)/2], nil
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
