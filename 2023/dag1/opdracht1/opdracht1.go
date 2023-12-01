package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	sum := 0
	for _, line := range lines {
		sum += parseWritten(line)
		fmt.Println(parseWritten(line))
	}
	fmt.Println("Result:", sum)
}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseNumber(line string) int {
	var first int
	for _, char := range line {
		if (unicode.IsDigit(char)) {
			first = int(char - '0')
			break
		}
	}

	var last int
	for i := len(line)-1; i >= 0; i-- {
		char := rune(line[i])
		if (unicode.IsDigit(char)) {
			last = int(char - '0')
			break
		}
	}

	return first*10 + last
}

func parseWritten(line string) int {
	line = strings.Replace(line, "one", "o1e", -1)
	line = strings.Replace(line, "two", "t2o", -1)
	line = strings.Replace(line, "three", "t3e", -1)
	line = strings.Replace(line, "four", "f4r", -1)
	line = strings.Replace(line, "five", "f5e", -1)
	line = strings.Replace(line, "six", "s6x", -1)
	line = strings.Replace(line, "seven", "s7n", -1)
	line = strings.Replace(line, "eight", "e8t", -1)
	line = strings.Replace(line, "nine", "n9e", -1)

	return parseNumber(line)
}