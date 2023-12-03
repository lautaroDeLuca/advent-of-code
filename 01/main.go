package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	buffer, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	inputArray := strings.Split(string(buffer), "\n")

	finalAnswerAccumulator := 0

	numbersMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	for i := range inputArray {
		finalAnswerAccumulator += (findNumberSubstring(inputArray[i], numbersMap, false))*10 + (findNumberSubstring(reverseString(inputArray[i]), numbersMap, true))
	}
	log.Printf("The sum of all numbers found is: %d", finalAnswerAccumulator)
}

func reverseString(input string) string {
	runes := []rune(input)

	length := len(runes)

	if length == 0 {
		return ""
	}

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	reversedString := string(runes)
	return reversedString
}

func deepCopyMap(originalMap map[string]int) map[string]int {
	copyMap := make(map[string]int)
	for key, value := range originalMap {
		copyMap[key] = value
	}
	return copyMap
}

func deleteEntryFromMap(originalMap map[string]int, key string) map[string]int {
	delete(originalMap, key)
	return originalMap
}

func resetAuxSearchValues(
	numbersMapCopy *map[string]int,
	wordsToDelete *[]string,
	accumulator *string,
	defaultSubStringSet map[string]int,
) {
	*numbersMapCopy = deepCopyMap(defaultSubStringSet)

	*wordsToDelete = []string{}

	*accumulator = ""
}

func findNumberSubstring(
	stringToSearch string,
	subStringSet map[string]int,
	findReversed bool,
) int {
	numbersMapCopy := deepCopyMap(subStringSet)
	wordsToDelete := []string{}
	accumulator := ""

	for i := range stringToSearch {
		resetAuxSearchValues(&numbersMapCopy, &wordsToDelete, &accumulator, subStringSet)
		for _, char := range stringToSearch[i:] {
			if findReversed {
				accumulator = string(char) + accumulator
			} else {
				accumulator += string(char)
			}

			for key := range numbersMapCopy {
				if (!findReversed && !strings.HasPrefix(key, accumulator)) ||
					(findReversed && !strings.HasSuffix(key, accumulator)) {
					wordsToDelete = append(wordsToDelete, key)
				}
			}

			for _, word := range wordsToDelete {
				delete(numbersMapCopy, word)
			}

			if len(numbersMapCopy) == 0 {
				resetAuxSearchValues(&numbersMapCopy, &wordsToDelete, &accumulator, subStringSet)
				break
			} else if len(numbersMapCopy) == 1 {
				_, exists := numbersMapCopy[accumulator]
				if exists {
					return numbersMapCopy[accumulator]
				}

			}
		}
	}

	return 0
}
