package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := `31 45
50 100`
	input := bufio.NewScanner(strings.NewReader(in))
	for input.Scan() {
		array := strings.Split(input.Text(), " ")
		var values []int
		for _, i := range array {
			j, _ := strconv.Atoi(i)
			values = append(values, j)
		}
		fmt.Printf("%v", GetPrimeNumbers(values[0], values[1]))
	}
}

func GetPrimeNumbers(start, end int) []int {
	naturals := make([]int, 0, end-start+1)
	for i := start; i <= end; i++ {
		isPrime := i > 1 && (i%2 != 0 || i == 2) && (i%3 != 0 || i == 3)
		j := 5
		d := 2

		for isPrime && j*j <= i {
			isPrime = i%j != 0
			j += d
			d = 6 - d
		}
		if isPrime {
			naturals = append(naturals, i)
		}
	}
	return naturals
}
