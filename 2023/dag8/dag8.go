package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	left  string
	right string
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	tree := make(map[string]Node)
	currents := make([]string, 0)

	for _, line := range lines[2:] {
		tree[line[0:3]] = Node{line[7:10], line[12:15]}
		if line[2] == 'A' {
			currents = append(currents, line[0:3])
		}
	}

	steps := 0

	current := "AAA"
	for i := 0; i <= len(lines[0]) ; i++ {
		steps++
		if i == len(lines[0]) {
			i = 0
		}
		if lines[0][i] == 'L' {
			current = tree[current].left
		} else {
			current = tree[current].right
		}
		if current == "ZZZ" {
			break
		}
	}
	fmt.Println(steps)
	steps = 0

	pathsToZ := make(map[string]int)
	for _, current := range currents {
		temp := current

		for i := 0; i <= len(lines[0]) ; i++ {
			fmt.Println("CurrentNode:", temp, tree[temp].left, tree[temp].right)
			steps++
			if i == len(lines[0]) {
				i = 0
			}
			if lines[0][i] == 'L' {
				temp = tree[temp].left
			} else if lines[0][i] == 'R' {
				temp = tree[temp].right
			}
			if temp[2] == 'Z' {
				fmt.Println("Found Z for", current, "in", steps, "steps")
				pathsToZ[current] = steps
				steps = 0
				break
			}
		}
	}
	fmt.Println(pathsToZ)
	fmt.Println("Result2:", lcmMap(pathsToZ))
}
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// Function to find lcm of two numbers
func lcm(a, b int) int {
	return int(math.Abs(float64(a*b))) / gcd(a, b)
}

// Function to find lcm of a map of string to int
func lcmMap(m map[string]int) int {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	if len(keys) < 2 {
		return m[keys[0]]
	}

	result := m[keys[0]]
	for i := 1; i < len(keys); i++ {
		result = lcm(result, m[keys[i]])
	}

	return result
}


func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}
