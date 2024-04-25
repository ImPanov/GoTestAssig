package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testValue := `1
2
3
7
11
101
100001
100005`

	in := bufio.NewScanner(strings.NewReader(testValue))
	for in.Scan() {
		value, _ := strconv.Atoi(in.Text())
		fmt.Print(GetComputerWithSuffix(value))
	}
}

func GetComputerWithSuffix(value int) string {
	switch {
	case value%100 > 4 && value < 20:
		return fmt.Sprintf("%d Компьютеров ", value)
	case value%10 == 1:
		return fmt.Sprintf("%d Компьютер ", value)
	case value%10 < 5:
		return fmt.Sprintf("%d Компьютера ", value)
	default:
		return ""
	}
}
