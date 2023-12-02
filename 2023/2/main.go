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
	id           int
	rounds       []round
	isValid      bool
	minimumCubes []cube
}

type round struct {
	id      int
	cubes   []cube
	isValid bool
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

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("The answer to part 1 is: %v\n", calculateAnswer(games))
	fmt.Printf("The answer to part 2 is: %v\n", calculateTotalCubePowers(games))
}

func parseGame(line string) game {

	game := game{}

	line = strings.ReplaceAll(line, "Game ", "")

	splitLine := strings.Split(line, ":")
	id, err := strconv.Atoi(splitLine[0])

	if err != nil {
		panic(err)
	}
	game.id = id

	game.rounds = parseRounds(strings.Split(splitLine[1], ";"))
	game.isValid = true
	game.minimumCubes = calculateMinimumCubes(game)
	for _, round := range game.rounds {
		if !round.isValid {
			game.isValid = false
			break
		}
	}

	return game
}

func parseRounds(rounds []string) []round {
	parsedRounds := []round{}

	for roundId, roundRaw := range rounds {
		round := round{}
		round.id = roundId + 1
		round.cubes = parseCubes(roundRaw)
		round.isValid = validRound(round.cubes)
		parsedRounds = append(parsedRounds, round)
	}

	return parsedRounds

}

func parseCubes(round string) []cube {

	parsedCubes := []cube{}
	cube := cube{}

	for _, cubeRaw := range strings.Split(round, ",") {
		parsedCube := strings.Split(strings.TrimSpace(cubeRaw), " ")

		count, err := strconv.Atoi(strings.ReplaceAll(parsedCube[0], " ", ""))

		if err != nil {
			panic(err)
		}

		cube.count = count
		cube.colour = parsedCube[1]

		parsedCubes = append(parsedCubes, cube)
	}
	return parsedCubes
}

func printGame(game game) {

	for _, round := range game.rounds {
		for _, cube := range round.cubes {
			fmt.Printf("Game %v: Round %v: %v %v dice: Game validity %v: Minimum %v %v, minimum %v %v, minimum %v %v.\n", game.id, round.id, cube.count, cube.colour, game.isValid, game.minimumCubes[0].colour, game.minimumCubes[0].count, game.minimumCubes[1].colour, game.minimumCubes[1].count, game.minimumCubes[2].colour, game.minimumCubes[2].count)
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
		if game.isValid {
			result = result + game.id
		}
	}

	return result
}

func calculateMinimumCubes(game game) []cube {

	minimumCubes := []cube{}

	blueCube := cube{
		colour: "blue",
		count:  0,
	}
	greenCube := cube{
		colour: "green",
		count:  0,
	}
	redCube := cube{
		colour: "red",
		count:  0,
	}

	for _, round := range game.rounds {
		for _, cube := range round.cubes {
			switch cube.colour {
			case "blue":
				if cube.count > blueCube.count {
					blueCube.count = cube.count
				}
			case "red":
				if cube.count > redCube.count {
					redCube.count = cube.count
				}
			case "green":
				if cube.count > greenCube.count {
					greenCube.count = cube.count
				}
			}
		}
	}

	minimumCubes = append(minimumCubes, blueCube, greenCube, redCube)

	return minimumCubes
}

func calculateTotalCubePowers(games []game) int {
	answer := 0
	for _, game := range games {
		printGame(game)
		answer = answer + (game.minimumCubes[0].count * game.minimumCubes[1].count * game.minimumCubes[2].count)
	}

	return answer
}
