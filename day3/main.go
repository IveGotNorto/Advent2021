package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	zeroTree := NewBst()

	for scanner.Scan() {
		buf := scanner.Text()

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
