package main

import (
	"advent_2021/extra"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// loadInput simply read the input text file and forges two slices of strings.  One containing the unique digits and the
// other their correspondent four secret output digits.
// In case of any error parsing the file, we ship an error to the caller.
func loadInput(scanner *bufio.Scanner) (uniqueDigits [][]string, outputDigits [][]string, err error) {
	for scanner.Scan() {
		splitByType := strings.Split(scanner.Text(), " | ")
		if len(splitByType) != 2 {
			return uniqueDigits, outputDigits, errors.New("something failed when reading the input")
		}

		uniqueDigits = append(uniqueDigits, strings.Split(splitByType[0], " "))
		outputDigits = append(outputDigits, strings.Split(splitByType[1], " "))
	}

	return uniqueDigits, outputDigits, nil
}

// firstExercise just iterates the array of output digits and looks for those which length is known to be one of the
// unique ones for 1, 4, 7 and 8.
func firstExercise() (result int, err error) {
	file, err := os.Open("inputs/day08_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	_, output, err := loadInput(bufio.NewScanner(file))

	for i := 0; i < len(output); i++ {
		for q := 0; q < len(output[i]); q++ {
			digitLength := len(output[i][q])
			if digitLength == 2 || digitLength == 3 || digitLength == 4 || digitLength == 7 {
				result++
			}
		}
	}

	return result, nil
}

// decodeDigits will take an array with ten unique digits, break the code and return another slice of strings with ten
// positions, where each position corresponds to its equivalent "secret" digits.
// The solution is achieved by finding the subtraction of characters composing the digit.  Since we know 4 out of the
// 10 digits, we can find out the remaining ones by combining these subtractions.
func decodeDigits(uniqueDigits []string) []string {
	dictionary := make([]string, 10)
	segments := make([]string, 10) // We will use the position in the slice to determine segments 0->a, 1->b...9->g

	// We want to store the difficult digits by their length, to prevent looping over unwanted ones.
	fiveDigitOnes := make([]string, 0)
	sixDigitOnes := make([]string, 0)

	// We iterate over the unique digits, and we catalog them if they are one of those we know (1, 4, 7, 8) or one of
	// the others, based on the length of the segments string.
	for _, value := range uniqueDigits {
		digitLength := len(value)

		switch digitLength {
		case 5: // 2, 3, 6
			fiveDigitOnes = append(fiveDigitOnes, value)
		case 6: // 0, 5, 9
			sixDigitOnes = append(sixDigitOnes, value)
		case 2: // 1
			dictionary[1] = value
		case 3: // 7
			dictionary[7] = value
		case 4: // 4
			dictionary[4] = value
		case 7: // 8
			dictionary[8] = value
		}
	}

	// Now we can infer the other missing positions.
	// First, we can find the top segment with the subtraction between 1 and 7.
	segments[0] = subtractSegments(dictionary[7], dictionary[1])

	// With this information we can find the 6-segment numbers.
	for _, value := range sixDigitOnes {
		findingNine := subtractSegments(subtractSegments(value, dictionary[4]), dictionary[7]) // (value - 4) - 7
		findingSix := subtractSegments(dictionary[1], subtractSegments(dictionary[8], value))  // 1 - (8 - value)

		if len(findingNine) == 1 {
			// Only 9 satisfies the condition, len((9 - 4) - 7) == 1.
			segments[6] = findingNine
			dictionary[9] = value
		} else if len(findingSix) == 1 {
			// Only six satisfy this condition, len(1 - (8 - 6)) == 1.
			segments[2] = findingSix
			dictionary[6] = value
		} else {
			// Otherwise, the number must be 0.
			dictionary[0] = value
		}
	}
	// Since we know the number 6 and 1.
	segments[5] = subtractSegments(dictionary[1], dictionary[6]) // 1 - 6

	// Now we will proceed to identify the digits formed by five segments.
	for _, value := range fiveDigitOnes {
		// Since we know all segments but one from three, it is easy to identify it.
		findingThree := subtractSegments(value, segments[0]+segments[2]+segments[5]+segments[6])
		// Now, the segments of 9 minus the segments of 5, should be just segment 2.
		findingFive := subtractSegments(dictionary[9], value) // 9 - value

		if len(findingThree) == 1 {
			segments[3] = findingThree
			dictionary[3] = value
		} else if len(findingFive) == 1 {
			dictionary[5] = value
		} else {
			dictionary[2] = value
		}
	}

	return dictionary
}

// translateDigit takes the word defining the segments of a digit and with the help of the dictionary providing the
// translation, returns another string with the number it represents.
func translateDigit(digit string, dictionary []string) (string, error) {
	for i, number := range dictionary {
		if len(digit) == len(number) && len(subtractSegments(digit, number)) == 0 {
			return strconv.Itoa(i), nil
		}
	}

	return "", errors.New("unable to translate")
}

// subtractSegments will find the letters in source that are not present in toCompare, then return them in a single
// string.
func subtractSegments(source string, toCompare string) (result string) {
	buffer := make(map[rune]bool)
	for _, value := range toCompare {
		buffer[value] = true
	}

	for _, value := range source {
		if !buffer[value] {
			result += string(value)
		}
	}

	return result
}

func secondExercise() (int, error) {
	file, err := os.Open("inputs/day08_exercise01.txt")
	if err != nil {
		return -1, err
	}
	defer extra.CloseFile(file)

	unique, output, err := loadInput(bufio.NewScanner(file))
	dictionary := make([]string, 0)
	result := 0

	for i := range unique {
		for p := 0; p < 10; p++ {
			dictionary = decodeDigits(unique[i])
		}

		secret := ""
		for q := 0; q < 4; q++ {
			secretDigit := output[i][q]
			translatedDigit, err := translateDigit(secretDigit, dictionary)
			if err != nil {
				log.Fatal(err)
			}
			secret += translatedDigit
		}
		result += extra.ConvertToInt(secret)
	}

	return result, nil
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
