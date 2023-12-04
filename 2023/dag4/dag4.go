package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cards, err = readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var sum1 = 0
	var sum2 = 0

	var dict = make(map[int] int)

	for index, card := range cards {
		dict[index] += 1
		card = strings.Split(card, ":")[1]
		split := strings.Split(card, "|")
		winningCards := getValues(split[0])
		ownedCards := getValues(split[1])
		var cardsWon = 0
		for _, winningCard := range winningCards {
			for _, ownedCard := range ownedCards {
				if winningCard == ownedCard {
					cardsWon++
				}
			}
		}
		for i := 0; i < cardsWon ; i++{
			dict[index+i+1] += dict[index]
		}
		if cardsWon > 0 {
			points := int(1 * math.Pow(2, float64(cardsWon-1)))
			sum1 += points
		}
	}

	for _, info := range dict {
		sum2 += info
	}
	fmt.Println("Total Points:", sum1)
	fmt.Println("Total Games:", sum2)
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

func getValues(input string) []int {
	var values []int
	for _, value := range strings.Split(input, " ") {
		intValue, _ := strconv.Atoi(strings.Trim(value, " "))
		if intValue != 0 {
			values = append(values, intValue)
		}
	}
	return values
}