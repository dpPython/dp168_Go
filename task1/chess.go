package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	print(arguments())
}

func print(height, width int) {
	for i := 0; i < height; i += 2 {
		printEven(width)
		if i+1 != height {
			printOdd(width)
		}
	}
}

func printEven(width int) {
	for i := 0; i < width; i++ {
		fmt.Print("*" + " ")
	}
	fmt.Println()
}

func printOdd(width int) {
	for i := 0; i < width; i++ {
		fmt.Print(" " + "*")
	}
	fmt.Println()
}

func arguments() (int, int) {
	if len(os.Args) == 1 {
		fmt.Println("In order to use program, which print chess board,", "\n",
			"please, enter 2 parameters like in example", "\n",
			"5 10 ", "\n")
		os.Exit(1)
	}
	arguments := os.Args
	return argsave(arguments)
}

func argsave(argsSlice []string) (int, int) {
	var err = errors.New("An error")
	var from float64
	var till float64

	if err != nil {
		from, err = strconv.ParseFloat(argsSlice[1], 64)
		till, err = strconv.ParseFloat(argsSlice[2], 64)
	} else {
		fmt.Println(err)
		return 0, 0
	}
	return int(from), int(till)
}
