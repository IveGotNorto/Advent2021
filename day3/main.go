package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	arr := strings.Split(strings.ReplaceAll(string(data), " ", ""), "\n")

	i := 0
	zeroCount := 0
	oneCount := 0

	var mostChar byte = '0'
	var leastChar byte = '0'

	mostFilter := make([]string, len(arr))
	leastFilter := make([]string, len(arr))
	copy(mostFilter, arr)
	copy(leastFilter, arr)

	for i < len(arr[0]) {
		for _, bin := range mostFilter {
			if bin[i] == '0' {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}

		if oneCount >= zeroCount {
			mostChar = '1'
		} else {
			mostChar = '0'
		}

		oneCount = 0
		zeroCount = 0

		for _, bin := range leastFilter {
			if bin[i] == '0' {
				zeroCount += 1
			} else {
				oneCount += 1
			}
		}

		if zeroCount <= oneCount {
			leastChar = '0'
		} else {
			leastChar = '1'
		}

		if len(mostFilter) > 1 {
			mostFilter = filter(mostFilter, func(x string) bool { return x[i] == mostChar })
		}

		if len(leastFilter) > 1 {
			leastFilter = filter(leastFilter, func(x string) bool { return x[i] == leastChar })
		}

		zeroCount = 0
		oneCount = 0
		i += 1
	}

	fmt.Printf("most: %v, least: %v\n", mostFilter, leastFilter)
	most, err := strconv.ParseInt(mostFilter[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	least, err := strconv.ParseInt(leastFilter[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("answer: %v\n", most*least)

}

func filter(arr []string, test func(string) bool) (ret []string) {
	for _, item := range arr {
		if test(item) {
			ret = append(ret, item)
		}
	}
	return
}
