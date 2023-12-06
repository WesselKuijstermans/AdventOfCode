package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"unicode"
)

func main() {
	var lines, err = readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var sum1 = 0
	var sum2 = 0

	for lineIndex, line := range lines {
		var digits []int = make([]int, 0)
		var adjacent = false
		for charIndex, char := range line {
			if unicode.IsDigit(char) {
				var number = int(char - '0')
				if !adjacent {
					adjacent = isAdjacent(lines, lineIndex, charIndex)
				}
				digits = append(digits, number)
				if charIndex == len(line)-1 || !unicode.IsDigit(rune(line[charIndex+1])) {

					if adjacent {
						number = 0
						for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
							digits[i], digits[j] = digits[j], digits[i]
						}
						for power, digit := range digits {
							number += digit * int(math.Pow10(power))
						}
						sum1 += number
					}
					digits = make([]int, 0)
					adjacent = false
				} else {
					continue
				}
			} else if char == '*' {
				sum2 += findGears(lines, lineIndex, charIndex)
			}
		}
	}
	fmt.Println("Result1:", sum1)
	fmt.Println("Result2:", sum2)
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

func isAdjacent(lines []string, lineIndex int, charIndex int) bool {
	var adjacents []rune = make([]rune, 0)

	// Get the adjacent characters
	for i := lineIndex - 1; i <= lineIndex+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		for j := charIndex - 1; j <= charIndex+1; j++ {
			if j < 0 || j >= len(lines[i]) {
				continue
			}
			adjacents = append(adjacents, rune(lines[i][j]))
		}
	}

	for _, r := range adjacents {
		if !unicode.IsDigit(r) && r != '.' {
			return true
		}
	}
	return false

}

func findGears(lines []string, lineIndex int, charIndex int) int {
	var numbers []int = make([]int, 0)

	for i := lineIndex - 1; i <= lineIndex+1; i++ {
		if i < 0 || i >= len(lines) {
			continue
		}
		for j := charIndex - 1; j <= charIndex+1; j++ {
			if j < 0 || j >= len(lines[i]) {
				continue
			}
			if unicode.IsDigit(rune(lines[i][j])) {
				number := getFullNumber(lines[i], j)
				found := false
				for _, n := range numbers {
					if n == number {
						found = true
						break
					}
				}
				if !found {
					numbers = append(numbers, number)
				}
			}
		}

	}
	if len(numbers) == 2 {
		fmt.Println("Multiplying", numbers[0], "and", numbers[1], "together", numbers[0]*numbers[1])
		return numbers[0] * numbers[1]
	}

	return 0
}

func getFullNumber(line string, charIndex int) int {
	var rightIndex = 0
	var leftIndex = 0
	for i := charIndex; i < len(line); i++ {
		if !unicode.IsDigit(rune(line[i])) {
			rightIndex = i
			break
		}
	}
	for i := rightIndex; i > 0; i-- {
		if !unicode.IsDigit(rune(line[i-1])) {
			leftIndex = i
			break
		}
	}
	number, err := strconv.Atoi(line[leftIndex:rightIndex])
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}
	return number

}
