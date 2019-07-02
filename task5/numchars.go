package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numberByChars(arguments())
}

func numberByChars(n int) {
	if n < 10 {
		fmt.Println(tillTen(n))
	} else if n < 20 {
		fmt.Println(tillTwenty(n))
	} else if n < 100 {
		fmt.Println(tillHundred(n))
	} else if n < 1000 {
		fmt.Println(tillThousand(n))
	} else {
		fmt.Println(tillMillion(n))
	}
}
func tillTen(n int) string {
	var s string
	switch n {
	case 0:
		s = ("ноль")
	case 1:
		s = ("один")
	case 2:
		s = ("два")
	case 3:
		s = ("три")
	case 4:
		s = ("четыре")
	case 5:
		s = ("пять")
	case 6:
		s = ("шесть")
	case 7:
		s = ("семь")
	case 8:
		s = ("восемь")
	case 9:
		s = ("девять")
	}
	return s
}
func tillTwenty(n int) string {
	var s string
	if n == 10 {
		s = "десять"
	} else if n == 11 || n == 13 {
		s = tillTen(n-10) + "надцать"
	} else if n == 12 {
		s = "двенадцать"
	} else {
		s = splitLastChar(n-10) + "надцать"
	}
	return s
}
func splitLastChar(n int) string {
	s := tillTen(n)
	return s[0 : len(s)-2]
}
func onlyTens(n int) string {
	var s string
	if 40 <= n && n <= 49 {
		s = "сорок"
	} else if 20 <= n && n <= 39 {
		s = tillTen(n/10) + "дцать"
	} else if 90 <= n && n <= 99 {
		s = "девяносто"
	} else {
		s = tillTen(n/10) + "десят"
	}
	return s
}
func tillHundred(n int) string {
	var s string
	if n%10 > 0 {
		s = onlyTens(n) + " " + tillTen(n%10)
	} else {
		s = onlyTens(n)
	}
	return s
}
func onlyHundreds(n int) string {
	var s string
	if 100 <= n && n <= 199 {
		s = "сто"
	} else if 200 <= n && n <= 299 {
		s = "двести"
	} else if 300 <= n && n <= 499 {
		s = tillTen(n/100) + "ста"
	} else {
		s = tillTen(n/100) + "сот"
	}
	return s
}
func tillThousand(n int) string {
	var s string
	if n%100 > 0 {
		s = onlyHundreds(n) + " " + tillHundred(n%100)
	} else {
		s = onlyHundreds(n)
	}
	return s
}
func onlyThousands(n int) string {
	var s string
	if 1000 <= n && n <= 1999 {
		s = "тысяча"
	} else if 2000 <= n && n <= 2999 {
		s = "две тысячи"
	} else if 3000 <= n && n <= 4999 {
		s = tillTen(n/1000) + " тысячи"
	} else {
		s = checkNumberOfThousands(n/1000) + " тысяч"
	}
	return s
}
func checkNumberOfThousands(n int) string {
	var s string
	if n < 10 {
		s = tillTen(n)
	} else if n < 20 {
		s = tillTwenty(n)
	} else if n < 100 {
		s = tillHundred(n)
	} else {
		s = tillThousand(n)
	}
	return s
}
func tillMillion(n int) string {
	var s string
	if n%1000 > 0 {
		s = onlyThousands(n) + " " + tillThousand(n%1000)
	} else {
		s = onlyThousands(n)
	}
	return s
}

func arguments() int {
	if len(os.Args) == 1 {
		fmt.Println("In order to use program,", "\n",
			"which spelling numbers by chars,", "\n",
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
