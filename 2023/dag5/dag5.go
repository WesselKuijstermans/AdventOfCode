package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var lines, err = readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	part1(lines)
	part2(lines)
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

func stringArrayToIntArray(input []string) []int {
	var output = make([]int, 0)
	for _, val := range input {
		intVal, _ := strconv.Atoi(val)
		output = append(output, intVal)
	}
	return output
}

func part1 (input []string) {
	source := stringArrayToIntArray(strings.Split(input[0], " ")[1:])
		
	var destination = make([]int, len(source))
	copy(destination, source)

	
	for _, line := range input[1:] {
		fmt.Println("Line:", line)
		if (len(line) == 0 || !unicode.IsDigit(rune(line[0]))) {
			copy(source, destination)
		} else {
			mappings := stringArrayToIntArray(strings.Split(line, " "))
			fmt.Println("Source:", source)
			fmt.Println("Mappings: ", mappings)
			for i := range source {
				rangeIndex := source[i] - mappings[1]
				fmt.Println("Range Index:", rangeIndex)
				if 0 <= rangeIndex && rangeIndex < mappings[2] {
					destination[i] = mappings[0] + rangeIndex
				}
			}
			fmt.Println("Destination:", destination)
		}
	}
	sort.Slice(destination, func(i, j int) bool {
		return destination[i] < destination[j]
	})
	fmt.Println("Lowest Location:", destination[0])

}

func part2(input []string) {
	source := stringArrayToIntArray(strings.Split(input[0], " ")[1:])
	var destination = make([]int, len(source))
	copy(destination, source)
	
	fmt.Println(source)
	for i := 0 ; i < len(source) ; i++ {
		
	}
	
}