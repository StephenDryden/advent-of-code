package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const inputLocation = "input.txt"

var numberMap = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {

	result := 0

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("input is :%s\n", line)

		code := ""
		//read line by line
		for i, text := range line {
			// if rune is decimal digit then add it to the code
			if unicode.IsDigit(text) {
				code = code + string(text)
				continue
			}
			// rune is not decimal therefore loop through map one to nine
			substring := line[i:]
			for s, i := range numberMap {
				// if remaining unchecked string contains one to nine and it exists at the start of the substring
				// add it to the code
				if substring >= s && strings.Contains(substring, s) && substring[0:len(s)] == s {
					code = code + fmt.Sprint(i)
				}
			}

		}
		result = result + calculateCalibrationValue(code)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("final answer is :%v\n", result)
	}
}

func calculateCalibrationValue(code string) int {
	r := []rune(code)
	first := r[0]
	second := r[len(code)-1]
	calibrationValue := int(first-'0')*10 + int(second-'0')
	fmt.Printf("calibration value is : %d\n", calibrationValue)
	return calibrationValue
}
