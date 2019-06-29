package main

import (
	"fmt"
	"strconv"
)

func main(){
	var rootBorderInput string
	fmt.Println("Enter the lower border above you want to get Fibonacci numbers")
	//entering lower border
	fmt.Scan(&rootBorderInput)
	rootBorder,_:= strconv.Atoi(rootBorderInput)
	//checking for mistakes and getting rid of them
	if  rootBorder>2000100 || rootBorder<0 {
		for rootBorder > 2000100 || rootBorder < 0 {
			fmt.Println("Incorrect.The size is either invalid or too big")
			fmt.Scan(&rootBorderInput)
			rootBorder,_= strconv.Atoi(rootBorderInput)
		}
	}
	fmt.Println("Lower border is ",rootBorder)

	var rangeBorderInput string
	fmt.Println("Enter the top border under you want to get Fibonacci numbers")
	//entering top border
	rangeBorder := inputCheck(rangeBorderInput)
	fmt.Println("Upper border is ",rangeBorder)

	fmt.Println("Here you go,your fibonacci numbers:")
	if rootBorder==rangeBorder{
		fmt.Println(rangeBorder,":-)")
		fmt.Println("Did you expect to get something different?")
	}
//in case lower border is greater than top we change their places
	if rootBorder>rangeBorder{
		fmt.Println("Well,i think you meant a little bit another thing,so i changed the borders places")
		printFibo(rangeBorder,rootBorder)
	}else{
		printFibo(rootBorder,rangeBorder)
	}
	fmt.Println("")
	fmt.Println("If the program finished working but there were shown no Fibonacci numbers,apparently there are no numbers in the entered range! ")
}
//input check function
func inputCheck(num string) int{
	fmt.Scan(&num)
	numconv,_:= strconv.Atoi(num)
	for numconv == 0 || numconv > 2000100 || numconv < 0 {
		fmt.Println("Incorrect.The size is either invalid or too big")
		fmt.Scan(&num)
		numconv,_= strconv.Atoi(num)
	}
	return numconv
}

//this function displays fibo numbers in a certain range
func printFibo(lowBorder,topBorder int){
	fib1,fib2:=0,1
	fib:=0

	if lowBorder==0{
		fmt.Print(0,";")
	}
	for i:=0;i<=topBorder;i++{
		fib = fib2+fib1
		fib2 = fib1
		fib1=fib

		if fib>topBorder{
			break
		}
		if fib<=topBorder && fib>=lowBorder{
			fmt.Print(fib,";")
		}
	}
}

