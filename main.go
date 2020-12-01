package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	numberStrings := strings.Split(string(data), "\r\n")
	numbers := []int{}
	for _, i := range numberStrings {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, j)
	}
	sort.Ints(numbers)
	a, b := findPair(numbers)
	result := a * b
	fmt.Print(result)

}

func findPair(numbers []int) (int, int) {
	for _, num1 := range numbers {
		goal := 2020 - num1
		for j := len(numbers) - 1; j > 0; j-- {
			if numbers[j] == goal {
				return num1, numbers[j]
			}
		}
	}
	return 0, 0
}
