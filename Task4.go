package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	in := `5
8`
	input := bufio.NewScanner(strings.NewReader(in))
	for input.Scan() {
		value, _ := strconv.Atoi(input.Text())
		MultipleTable(value)
	}
}

func MultipleTable(value int) {
	var array [][]int
	for i := 1; i <= value; i++ {
		var row []int
		for j := 1; j <= value; j++ {
			row = append(row, i*j)
		}
		array = append(array, row)
	}

	fmt.Printf("    ")
	for _, i2 := range array[0] {
		fmt.Printf("%3v ", i2)
	}
	fmt.Printf("\n")

	for i, ints := range array {
		fmt.Printf("%3v ", array[0][i])
		for _, i2 := range ints {
			fmt.Printf("%3v ", i2)
		}
		fmt.Printf("\n")
	}
}
