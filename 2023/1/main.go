package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

const inputLocation = "input.txt"

func main() {

	result := 0

	file, err := os.Open(inputLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("input is :%s\n", scanner.Text())

		code := ""
		for _, text := range scanner.Text() {
			if unicode.IsDigit(text) {
				code = code + string(text)
			}
		}
		fmt.Printf("code is : %v\n", code)
		result = result + obtainCalibrationValue(code)
		fmt.Printf("running total is : %v\n", code)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("final answer is :%v\n", result)

	}

}

func obtainCalibrationValue(code string) int {
	r := []rune(code)
	first := r[0]
	second := r[len(code)-1]
	calibrationValue := int(first-'0')*10 + int(second-'0')
	fmt.Printf("calibration value is : %d\n", calibrationValue)
	return calibrationValue
}
