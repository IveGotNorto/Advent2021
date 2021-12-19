package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var slide1 int
	var slide2 int
	var slide3 int
	var slide4 int

	iter1 := 0
	iter2 := 1
	iter3 := 2
	iter4 := 3
	mainIter := 0
	count := 0

	for scanner.Scan() {
		buff, err1 := strconv.Atoi(scanner.Text())
		if err1 != nil {
			fmt.Println(err1)
			break
		}

		// RESET
		if mainIter == iter1+4 {
			iter1 += 4
			slide1 = 0
		} else if mainIter == iter2+4 {
			iter2 += 4
			slide2 = 0
		} else if mainIter == iter3+4 {
			iter3 += 4
			slide3 = 0
		} else if mainIter == iter4+4 {
			iter4 += 4
			slide4 = 0
		}

		// LOAD UP
		if mainIter >= iter1 && mainIter < iter1+3 {
			slide1 += buff
		}

		if mainIter >= iter2 && mainIter < iter2+3 {
			slide2 += buff
		}

		if mainIter >= iter3 && mainIter < iter3+3 {
			slide3 += buff
		}

		if mainIter >= iter4 && mainIter < iter4+3 {
			slide4 += buff
		}

		//fmt.Printf("iter1: %v, iter2: %v, iter3: %v, iter4: %v\n", iter1, iter2, iter3, iter4)
		//fmt.Printf("slide1: %v, slide2: %v, slide3: %v, slide4: %v\n", slide1, slide2, slide3, slide4)

		// RIDE OUT
		if iter1+2 == mainIter {
			//fmt.Printf("comparing %v and %v", slide1, slide4)
			if slide1 > slide4 && iter1 != 0 {
				count += 1
			}
		} else if iter2+2 == mainIter {
			//fmt.Printf("comparing %v and %v", slide2, slide1)
			if slide2 > slide1 {
				count += 1
			}
		} else if iter3+2 == mainIter {
			//fmt.Printf("comparing %v and %v", slide3, slide2)
			if slide3 > slide2 {
				count += 1
			}
		} else if iter4+2 == mainIter {
			//fmt.Printf("comparing %v and %v", slide4, slide3)
			if slide4 > slide3 {
				count += 1
			}
		}

		//fmt.Printf("count: %v\n", count)

		mainIter += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Answer: %v\n", count)
}
