package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size string
	fmt.Println("Enter the needed size of the board(integer only)")
	fmt.Scan(&size)
	convsize,_:= strconv.Atoi(size)
	//getting correct size in case of any mistake made by a user
	for convsize==0{
		fmt.Println("Incorrect.The size is either invalid or too big")
		fmt.Scan(&size)
		convsize,_ = strconv.Atoi(size)
	}
//check
	if convsize>120 || convsize<0{
		fmt.Println("I think this chessboard is too large to play,100 is max")
		convsize = 8
		fmt.Println("But we made it standard")
	}

	var str1 string = "*"
	var str2 string = " "
//displaying the board
	for i:=0; i<convsize; i++{
		for j:=0;j<convsize/2;j++{
			if i%2==0{
				fmt.Print(str1,str2)
			} else {
				fmt.Print(str2,str1)
			}
		}
		fmt.Println()
	}
}