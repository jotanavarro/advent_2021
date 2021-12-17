package main

import (
	"advent_2021/extra"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func loadInput(scanner *bufio.Scanner) (template []rune, rules map[string]string) {
	scanner.Scan()
	template = []rune(scanner.Text())
	scanner.Scan()

	rules = make(map[string]string)

	for scanner.Scan() {
		parsed := strings.Split(scanner.Text(), " -> ")
		rules[parsed[0]] = parsed[1]
	}

	return template, rules
}

func countElements(pairs map[string]int, template []rune) (result map[string]int) {
	result = make(map[string]int)

	// We only consider the first letter of the string, because the rest are repeating.
	for pair, value := range pairs {
		result[string(pair[0])] += value
	}

	// But because of that, we will be ignoring the last letter of the original template, so we increase the
	// count of that one by one.
	result[string(template[len(template)-1])]++

	return result
}

// calculateResult will return the result of subtracting to the most common element count the one of the least common.
func calculateResult(frequencies map[string]int) (result int) {
	min, max := math.MaxInt, 0

	for _, frequency := range frequencies {
		if frequency > max {
			max = frequency
		}

		if frequency < min {
			min = frequency
		}
	}

	return max - min
}

// countPairs will transform a template into a map containing all the pair of elements it contains.
func countPairs(template []rune) (result map[string]int) {
	result = make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		result[string(template[i])+string(template[i+1])]++
	}

	return result
}

// evaluatePairs will, considering the rules provided, how many pairs appear in the resulting new template.
func evaluatePairs(pairs map[string]int, rules map[string]string) (result map[string]int) {
	result = make(map[string]int)

	for pair, count := range pairs {
		insertion := rules[pair] // This is common to know the following pairs

		// We know that AB -> A<insertion> and <insertion>B.
		result[string(pair[0])+insertion] += count
		result[insertion+string(pair[1])] += count
	}

	return result
}

func firstExercise() (int, error) {
	file, err := os.Open("inputs/day14_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	template, rules := loadInput(bufio.NewScanner(file))

	pairs := countPairs(template)

	for i := 0; i < 10; i++ {
		pairs = evaluatePairs(pairs, rules)
	}

	count := countElements(pairs, template)
	fmt.Printf("Count of elements: %+v\n.", count)

	return calculateResult(count), nil
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day14_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	template, rules := loadInput(bufio.NewScanner(file))

	pairs := countPairs(template)

	for i := 0; i < 40; i++ {
		pairs = evaluatePairs(pairs, rules)
	}

	count := countElements(pairs, template)
	fmt.Printf("Count of elements: %+v\n.", count)

	return calculateResult(count), nil
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
