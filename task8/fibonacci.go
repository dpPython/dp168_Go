package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	printRange(arguments())
}

func fibSlice(till int) []int {
	var fibSlice = []int{}
	fibSlice = append(fibSlice, 0)
	x, y := 0, 1
	for i := 0; i < till; i++ {
		x, y = y, x+y
		fibSlice = append(fibSlice, x)
	}
	return fibSlice
}

func printRangeByIndex(from int, till int) {
	slice := fibSlice(till)
	for i, num := range slice {
		if i >= from && i+1 <= till {
			fmt.Print(num, ", ")
		} else if i+1 > till && i <= till {
			fmt.Print(num, ".")
		}
	}
}

func printRangeByValue(from int, till int) {
	slice := fibSlice(till)
	for _, num := range slice {
		if num >= from && num+num/2 < till {
			fmt.Print(num, ", ")
		} else if num+num/2 >= till && num <= till {
			fmt.Print(num, ".")
		}
	}
}

func printRange(from int, till int, rangeBy bool) {
	if rangeBy {
		printRangeByValue(from, till)
	} else {
		printRangeByIndex(from, till)
	}
}

func arguments() (int, int, bool) {
	if len(os.Args) == 1 {
		fmt.Println("In order to use program, which counting fibonacci numbers,", "\n",
			"please, enter 3 parameters like in example", "\n",
			"5 10 false", "\n",
			"Where true or false is type of range (by value/index)")
		os.Exit(1)
	}
	arguments := os.Args
	return argsave(arguments)
}

func argsave(argsSlice []string) (int, int, bool) {
	var err = errors.New("An error")
	var from float64
	var till float64
	var rangeBy bool

	if err != nil {
		from, err = strconv.ParseFloat(argsSlice[1], 64)
		till, err = strconv.ParseFloat(argsSlice[2], 64)
		rangeBy, err = strconv.ParseBool(argsSlice[3])
	} else {
		fmt.Println(err)
		return 0, 0, false
	}
	return int(from), int(till), rangeBy
}
