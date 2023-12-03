package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type MinValidGame struct {
	blue     int
	red      int
	green    int
	possible bool
	gameId   int
}

func main() {
	buffer, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parsedInput := strings.Split(string(buffer), "\n")

	rulesMap := make(map[string]int)
	rulesMap["blue"] = 14
	rulesMap["red"] = 12
	rulesMap["green"] = 13

	var results []MinValidGame

	for _, game := range parsedInput {
		if game == "" {
			continue
		}
		separatedByColon := strings.Split(game, ":")
		gameAndNumber := strings.Split(separatedByColon[0], " ")
		gameId, err := strconv.ParseInt(gameAndNumber[1], 10, 32)
		if err != nil {
			panic(err)
		}
		sets := strings.Split(separatedByColon[1], ";")
		results = append(results, checkIfGameIsImpossible(sets, rulesMap, int(gameId)))
	}

	var accumulator int
	var powerAccumulator int

	for _, result := range results {
		power := 0
		colorResults := []int{result.red, result.blue, result.green}
		if result.possible == true {
			accumulator += result.gameId
		}
		for _, value := range colorResults {
			if value > 0 && power != 0 {
				power = power * value
			}
			if value > 0 && power == 0 {
				power += value
			}
		}
		powerAccumulator += power
	}
	log.Println("First Challenge's Answer:", accumulator)
	log.Println("Second Challenge's Answer:", powerAccumulator)
}

func checkIfGameIsImpossible(gameSets []string, rulesMap map[string]int, gameId int) MinValidGame {
	var gameStats MinValidGame
	gameStats.gameId = gameId
	gameStats.possible = true

	for _, set := range gameSets {
		trimmedSet := strings.TrimSpace(set)
		for _, cubeColor := range strings.Split(trimmedSet, ",") {
			trimmedCubeColor := strings.TrimSpace(cubeColor)
			cubeColorSplit := strings.Split(trimmedCubeColor, " ")
			amount, err := strconv.ParseInt(cubeColorSplit[0], 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			if cubeColorSplit[1] == "red" && int(amount) > gameStats.red {
				gameStats.red = int(amount)
			}
			if cubeColorSplit[1] == "blue" && int(amount) > gameStats.blue {
				gameStats.blue = int(amount)
			}
			if cubeColorSplit[1] == "green" && int(amount) > gameStats.green {
				gameStats.green = int(amount)
			}
		}
	}
	if rulesMap["red"] < gameStats.red || rulesMap["blue"] < gameStats.blue ||
		rulesMap["green"] < gameStats.green {
		gameStats.possible = false
	}
	return gameStats
}
