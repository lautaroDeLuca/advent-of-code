package main

import (
	"log"
	"os"
	"strings"
)

type FoundIndexes struct {
	firstFound   string
	firstIndex   int
	first        bool
	lastFound    string
	last         bool
	lastIndex    int
	parsedOutput int
}

func main() {
	buffer, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	inputArray := strings.Split(string(buffer), "\n")

	accumulator := 0
	for i := range inputArray {
		accumulator += (findFirstNumber(inputArray[i], false))*10 + (findFirstNumber(reverseString(inputArray[i]), true))
	}
	log.Printf("The sum of all numbers found is: %d", accumulator)
}

func deleteEntryFromSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func deepCopySlice(source []string) []string {
	target := make([]string, len(source))
	copy(target, source)
	return target
}

func findStringInSlice(target string, slice []string) bool {
	for _, str := range slice {
		if str == target {
			return true // Return true if the target string is found
		}
	}
	return false // Return false if the target string is not found
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

func findFirstNumber(stringToSearch string, reversed bool) int {
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

	numbersMapCopy := deepCopyMap(numbersMap)

	wordsToDelete := []string{}
	accumulator := ""

	for i := range stringToSearch {
		for _, char := range stringToSearch[i:] {
			if reversed {
				accumulator = string(char) + accumulator
			} else {
				accumulator += string(char)
			}
			for key := range numbersMapCopy {
				if reversed {
					if !strings.HasSuffix(key, accumulator) {
						wordsToDelete = append(wordsToDelete, key)
					}
				} else {
					if !strings.HasPrefix(key, accumulator) {
						wordsToDelete = append(wordsToDelete, key)
					}
				}
			}

			for _, word := range wordsToDelete {
				delete(numbersMapCopy, word)
			}

			if len(numbersMapCopy) == 0 {
				accumulator = ""
				numbersMapCopy = deepCopyMap(numbersMap)
				wordsToDelete = []string{}
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
