package main

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	in := `45 350 32
50 100 150`
	input := bufio.NewScanner(strings.NewReader(in))
	for input.Scan() {
		array := strings.Split(input.Text(), " ")
		var values []int
		for _, i := range array {
			j, _ := strconv.Atoi(i)
			values = append(values, j)
		}
		fmt.Printf("%v", GetDivides(values))
	}
}

func GetDivides(values []int) []int {
	minValue := slices.Min(values)
	maxDivider := minValue / 2
	var divides []int

	for i := 2; i <= maxDivider; i++ {
		if CheckDivide(i, values) {
			divides = append(divides, i)
		}
	}

	if CheckDivide(minValue, values) {
		divides = append(divides, minValue)
	}
	return divides
}

func CheckDivide(divide int, values []int) bool {
	checker := true
	for _, value := range values {
		if value%divide != 0 {
			checker = false
		}
	}
	return checker
}
