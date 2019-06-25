package main

import (
	"fmt"
	"strconv"
)

type chessBoard struct{
	width,length int
}

func isValidVal(value *int)bool{
	var str string
	fmt.Scanf("%s", &str)
	n, err:= strconv.Atoi(str)
	if err!=nil||n<=0||n>120 {
		fmt.Printf("Please, enter a correct value. It must be an unsigned integer number, less than 120\n")
		return false
	}  else {
		*value=n
		return true
	}

}
func initOfValues()(chBoard chessBoard){
	fmt.Print("Enter the width of the chessboard:\n ")
	for {
		if isValidVal(&chBoard.width){break}
	}
	fmt.Print("Enter the length of the chessboard:\n ")
	for {
		if isValidVal(&chBoard.length){break}
	}
	return
}
func drawChessBoard(chBoard chessBoard){
		for i := 0; i < chBoard.length; i++ {
		for j := 0; j < chBoard.width; j++ {
			if (i+j)%2 == 0 {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
func main() {
	chBoard:=initOfValues()
	drawChessBoard(chBoard)
}