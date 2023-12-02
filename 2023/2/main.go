package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputLocation = "input.txt"

type game struct {
	id     int
	rounds []round
	valid  bool
}

type round struct {
	id    int
	cubes []cube
	valid bool
}

type cube struct {
	colour string
	count  int
}

var games []game

func main() {

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		games = append(games, parseGame(line))
		for _, game := range games {
			printGame(game)
		}

		fmt.Printf("And the final answer is: %v\n", calculateAnswer(games))

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func parseGame(line string) game {

	game := game{}

	line = strings.ReplaceAll(line, "Game ", "")

	parts := strings.SplitAfter(line, ":")
	id, err := strconv.Atoi(strings.ReplaceAll(parts[0], ":", ""))

	if err != nil {
		panic(err)
	}
	game.id = id

	rounds := parseRounds(strings.Split(parts[1], ";"))
	game.rounds = rounds
	game.valid = true
	for _, round := range game.rounds {
		if !round.valid {
			game.valid = false
			break
		}
	}

	return game
}

func parseRounds(rounds []string) []round {
	parsedRounds := []round{}

	for roundId, roundInput := range rounds {

		round := round{}
		cubes := parseCubes(roundInput)
		round.id = roundId + 1
		round.cubes = cubes
		round.valid = validRound(cubes)
		parsedRounds = append(parsedRounds, round)
	}

	return parsedRounds

}

func parseCubes(round string) []cube {

	parsedCubes := []cube{}
	cube := cube{}
	parsedCubeInput := strings.Split(round, ",")

	for _, draw := range parsedCubeInput {
		parsedDraw := strings.Split(strings.TrimSpace(draw), " ")

		count, err := strconv.Atoi(strings.ReplaceAll(parsedDraw[0], " ", ""))

		if err != nil {
			panic(err)
		}

		cube.count = count
		cube.colour = parsedDraw[1]

		parsedCubes = append(parsedCubes, cube)
	}
	return parsedCubes
}

func printGame(game game) {

	for _, round := range game.rounds {
		for _, cube := range round.cubes {
			fmt.Printf("Game %v: Round %v: %v %v dice: Game validity %v\n", game.id, round.id, cube.count, cube.colour, game.valid)
		}
	}

}

func validRound(cubes []cube) bool {

	for _, cube := range cubes {
		switch cube.colour {
		case "green":
			if cube.count > 13 {
				return false
			}
		case "red":
			if cube.count > 12 {
				return false
			}
		case "blue":
			if cube.count > 14 {
				return false
			}
		default:
			return false
		}
	}

	return true
}

func calculateAnswer(games []game) (result int) {

	for _, game := range games {
		if game.valid {
			result = result + game.id
		}
	}

	return result
}
