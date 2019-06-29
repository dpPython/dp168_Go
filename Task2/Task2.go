package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var one1 string
	var one2 string
	fmt.Println("Enter size of the first envelope(float or integer), one by one with Enter")
	//entering first envelope's sides
	fmt.Scan(&one1)
	fmt.Scan(&one2)
	firstEnvelopeSide1,_:= strconv.ParseFloat(one1,64)
	firstEnvelopeSide2,_:= strconv.ParseFloat(one2,64)
	//cyclic check intil it's alright
	for !validInput(firstEnvelopeSide1,firstEnvelopeSide2){
		fmt.Scan(&one1)
		fmt.Scan(&one2)
		firstEnvelopeSide1,_= strconv.ParseFloat(one1,64)
		firstEnvelopeSide2,_= strconv.ParseFloat(one2,64)
	}

	var two1 string
	var two2 string
	fmt.Println("Enter size of the second envelope(float or integer), one by one with Enter")
	//entering second envelope's sides
	fmt.Scan(&two1)
	fmt.Scan(&two2)
	secondEnvelopeSide1,_:= strconv.ParseFloat(two1,64)
	secondEnvelopeSide2,_:= strconv.ParseFloat(two2,64)
	//cyclic check intil it's alright
	for !validInput(secondEnvelopeSide1,secondEnvelopeSide2) {
		fmt.Scan(&two1)
		fmt.Scan(&two2)
		secondEnvelopeSide1, _ = strconv.ParseFloat(one1, 64)
		secondEnvelopeSide2, _ = strconv.ParseFloat(one2, 64)
	}
//computing two diagonals
	firstDiagonal := math.Hypot(firstEnvelopeSide1,firstEnvelopeSide2)
	secondDiagonal := math.Hypot(secondEnvelopeSide1,secondEnvelopeSide2)

	fmt.Println("First envelope diagonal is",firstDiagonal,"cm")
	fmt.Println("Second envelope diagonal is",secondDiagonal,"cm")
//using first type to compare them
	fmt.Println("Check second -> first")
	var check1 bool = comparison(firstEnvelopeSide1,firstEnvelopeSide2,secondEnvelopeSide1,secondEnvelopeSide2,secondDiagonal)
	fmt.Println("Check first -> second")
	var check2 bool = comparison(secondEnvelopeSide1,secondEnvelopeSide2,firstEnvelopeSide1,firstEnvelopeSide2,firstDiagonal)

	if check1{
		fmt.Println("The second envelope can be put inside the first one")
	} else {
		fmt.Println("The second envelope cannot be put inside the first one")
	}
	if check2{
		fmt.Println("The first envelope can be put inside the second one")
	} else {
		fmt.Println("The first envelope cannot be put inside the second one")
	}
//using another type of comparison(that creepy formula)
	fmt.Println("Another type of check:")
	fmt.Println("Check second -> first")
	check1 = secondComparison(firstEnvelopeSide1,firstEnvelopeSide2,secondEnvelopeSide1,secondEnvelopeSide2)
	fmt.Println("Check first -> second")
	check2 = secondComparison(secondEnvelopeSide1,secondEnvelopeSide2,firstEnvelopeSide1,firstEnvelopeSide2)

	if check1{
		fmt.Println("The second envelope can be put inside the first one")
	} else {
		fmt.Println("The second envelope cannot be put inside the first one")
	}
	if check2{
		fmt.Println("The first envelope can be put inside the second one")
	} else {
		fmt.Println("The first envelope cannot be put inside the second one")
	}
}
//comparison func
func comparison(fside1 float64,fside2 float64,secside1 float64, secside2 float64,diagonal float64)bool{
	var check bool = false
	if fside1>secside1 && fside1>secside2 && fside1>diagonal && fside2>secside1 && fside2>secside2 && fside2>diagonal{
		check = true
	}
	if check{
		fmt.Println("Successful")
	} else {
		fmt.Println("Failed to insert")
	}
	return  check
}
//input validation function
func validInput(one,two float64)bool{
	check:=true
	if one==0 || two==0{
		check = false
		fmt.Println("Something went wrong!You should enter either integers or float.Try again")
		fmt.Println("Enter size of the first envelope(float or integer), one by one with Enter")
	}else{
		check = true
	}
	return check
}
//second type comparison function
func secondComparison(fside1 float64,fside2 float64,secside1 float64, secside2 float64)bool{
	check:=false
	scaryFormula:= (2*secside1*secside2*fside1+(secside1*secside1-secside2*secside2)*math.Sqrt(secside1*secside1+secside2*secside2-fside1*fside1))/secside1*secside1+secside2*secside2
	if  secside1<=fside1 && secside2<=fside2 || secside1>fside1 && fside2>=scaryFormula{
		check = true
	}
	return check
}

