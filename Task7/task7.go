package main

import (
	"fmt"
	"strconv"
)

func main(){
	var size string
	fmt.Println("Enter the border under you want to get squares of natural numbers")
	newsize := Input(size)

	numbers := make([]int, 0)
	for i:=1;i<=newsize+1;i++{
		numbers = append(numbers,i)
	}

	for i:= range numbers{
		if NaturalSquare(i,newsize) {
			fmt.Print(i)
			fmt.Print(";")
		}
	}
}

func NaturalSquare(nat,border int)bool{
	square:= nat*nat
	if square<border{
		return true
	}else{
		return false
	}
}

func Input(input string) int{
	fmt.Scan(&input)
	newsize,_:= strconv.Atoi(input)
		for newsize == 0 || newsize > 2000100 || newsize< 0 {
			fmt.Println("Incorrect.The size is either invalid or too big")
			fmt.Scan(&input)
			newsize,_= strconv.Atoi(input)
		}
	return newsize
}
