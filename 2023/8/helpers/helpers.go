package helpers

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(fileLocation string) map[int]string {

	var lines = make(map[int]string)
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		count++
		lines[count] = scanner.Text()

	}

	return lines
}
