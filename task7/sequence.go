package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	printNumbers(arguments())
}

func squareRoot(n int) int {
	sRoot := math.Sqrt(float64(n))
	return int(sRoot)
}

func printNumbers(n int) {
	temp := squareRoot(n)
	for i := 1; i <= temp; i++ {
		if i+1 <= temp {
			fmt.Print(i, ", ")
		} else if i <= temp {
			fmt.Print(i, ".")
		}
	}
}

func arguments() int {
	if len(os.Args) == 1 {
		fmt.Println("In order to use program,", "\n",
			"which counting any numbers less than given square root number,", "\n",
			"please, enter 1 parameter like in example", "\n",
			"100")
		os.Exit(1)
	}
	arguments := os.Args
	return int(argsave(arguments))
}

func argsave(argsSlice []string) int {
	var err error = errors.New("An error")
	var userNumber float64
	k := 1
	for err != nil {
		if k >= len(argsSlice) {
			fmt.Println("None of the arguments is number")
			return 0
		}
		userNumber, err = strconv.ParseFloat(argsSlice[k], 64)
		k++
	}
	return int(userNumber)
}
