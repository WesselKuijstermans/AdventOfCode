package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandInfo struct {
	cards string
	bid   int
	value int
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	hands := make(map[string]HandInfo)
	for _, line := range lines {
		sep := strings.Split(line, " ")
		cards := sep[0]
		bid, _ := strconv.Atoi(sep[1])
		hands[cards] = HandInfo{cards, bid, getHandValue(cards)}
	}

	keys := make([]string, 0, len(hands))
	for k := range hands {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Keys:", keys)

	rankedList := make([]HandInfo, 0, len(hands))
	for _, k := range keys {
		rankedList = append(rankedList, hands[k])
	}

	sort.Slice(rankedList, func(i, j int) bool {
		if rankedList[i].value == rankedList[j].value {
			fmt.Println("Comparing", rankedList[i].cards, "to", rankedList[j].cards)
			return compareCards(rankedList[i].cards, rankedList[j].cards) > 0
		}
		return rankedList[i].value < rankedList[j].value
	})

	for _, hand := range rankedList {
		fmt.Println(hand)
	}

	sum1 := 0
	for index, hand := range rankedList {
		sum1 += hand.bid * (index + 1)
	}
	fmt.Println("Result1:", sum1)
}

func compareCards(card1 string, card2 string) int {
	for i := 0; i < len(card1); i++ {
		if card1[i] != card2[i] {
			result := getCardIndex(string(card1[i])) - getCardIndex(string(card2[i]))
			return result
		}
	}
	return 0
}

func getCardIndex(card string) int {
	cardHierarchy := []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

	for i := 0; i <= len(cardHierarchy); i++ {
		if cardHierarchy[i] == card {
			return i
		}
	}
	return -1
}

func getHandValue(hand string) int {
	occurrences := map[string]int{"A": 0, "K": 0, "Q": 0, "J": 0, "T": 0, "9": 0, "8": 0, "7": 0, "6": 0, "5": 0, "4": 0, "3": 0, "2": 0}
	for _, card := range hand {
		occurrences[string(card)]++
	}

	// Five of a kind
	for card, count := range occurrences {
		if card != "J" {
			if count == 5 || count+occurrences["J"] == 5 {
				return 6
			}
		}
		if count == 5 {
			return 6
		}
	}

	// Four of a kind
	for card, count := range occurrences {
		if card != "J" {
			if count == 4 || count+occurrences["J"] == 4 {
				return 5
			}
		}
		if count == 4 {
			return 5
		}
	}

	// Full house
	hasThree := false
	hasTwo := false
	for card, count := range occurrences {
		if card != "J" {
			if count == 3 || count+occurrences["J"] == 3 && !hasThree {
				hasThree = true
			} else if count == 2 {
				hasTwo = true
			}
		} else {
			if count == 3 {
				hasThree = true
			} else if count == 2 {
				hasTwo = true
			}
		}
	}
	if hasThree && hasTwo {
		return 4
	}

	// Three of a kind
	for card, count := range occurrences {
		if card != "J" {
			if count == 3 || count+occurrences["J"] == 3 {
				return 3
			}
		}
		if count == 3 {
			return 3
		}

	}

	// Two pair
	pairCount := 0
	for card, count := range occurrences {
		if card != "J" {
			if count == 2 || count+occurrences["J"] == 2 {
				pairCount++
			}
		} else {
			if count == 2 {
				pairCount++
			}
		}
	}
	return pairCount
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not open file %s", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.Trim(scanner.Text(), " "))
	}
	return lines, nil
}
