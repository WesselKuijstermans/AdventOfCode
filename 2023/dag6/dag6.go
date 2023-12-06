package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	times := stringArrayToIntArray(strings.Split(lines[0], " ")[1:])
	distances := stringArrayToIntArray(strings.Split(lines[1], " ")[1:])
	sum1 := 1
	fmt.Println("Times:", times, "Distances:", distances)
	for i := 0; i < len(times); i++ {
		distanceToBeat := distances[i]
		timesBeaten := 0
		for speed := 0; speed < times[i]; speed++ {
			distance := speed * (times[i] - speed)
			if distance > distanceToBeat {
				timesBeaten++
			}
		}
		fmt.Println("Times beaten:", timesBeaten)
		sum1 *= timesBeaten
	}

	sum2 := 0
	TotalTime := stringToInt((lines[0]))
	distanceToBeat := stringToInt((lines[1]))
	for speed := 0; speed < TotalTime; speed++ {
		if speed * (TotalTime - speed) > distanceToBeat {
			sum2++
		}
	}
	fmt.Println("Result1:", sum1)
	fmt.Println("Result2:", sum2)
}

func stringArrayToIntArray(input []string) []int {
	var output = make([]int, 0)
	for _, val := range input {
		val = strings.Trim(val, " ")
		intVal, _ := strconv.Atoi(val)
		if intVal != 0 {
			output = append(output, intVal)
		}
	}
	return output
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(strings.Split(strings.Replace(input, " ", "", -1), ":")[1])
	return output
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
	return lines, nil
}

