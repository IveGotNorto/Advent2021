package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	depth := 0
	horizontal := 0
	aim := 0
	for scanner.Scan() {
		buf := scanner.Text()
		result := strings.Split(buf, " ")
		command := result[0]
		value, err := strconv.Atoi(result[1])

		if err != nil {
			fmt.Println(err)
			break
		}

		if command == "up" {
			aim -= value
		} else if command == "down" {
			aim += value
		} else if command == "forward" {
			horizontal += value
			depth += (aim * value)
		}
	}

	fmt.Println("Answer", horizontal*depth)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
