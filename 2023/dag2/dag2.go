package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var games, err = readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var sum1 = 0
	var sum2 = 0
	for i, game := range games {
		game = strings.Split(game, ":")[1]
		result1, result2 := getGameResults(game, i +1)
		sum1 += result1
		sum2 += result2
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

func getGameResults(game string, index int) (int, int) {
	var red = 0
	var green = 0
	var blue = 0
	for _, subgame := range strings.Split(game, ";") {
		for i, char := range subgame {
			if unicode.IsDigit(char) {
				var number = int(char - '0')
				var colorchar = rune(subgame[i+2])
				if unicode.IsDigit(rune(subgame[i+1])) {
					number = number*10 + int(subgame[i+1]-'0')
					colorchar = rune(subgame[i+3])
				}
				switch colorchar {
				case 'r':
					if number > red {
						red = number
					}
				case 'g':
					if number > green {
						green = number
					}
				case 'b':
					if number > blue {
						blue = number
					}
				}

			}
		}
	}
	fmt.Println(red, green, blue)
	if (red <= 12 && green <= 13 && blue <= 14) {
		return index, red * green * blue
	}
	return 0, red * green * blue
}