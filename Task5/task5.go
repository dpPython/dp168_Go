package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var size int
	fmt.Println("Enter the number you want to convert to string(integers only)")
	//integer input
	//if a user inputs ababba he/she gets 0
	//if a user inputs 1h2bab23 gets 1(example)
	fmt.Scan(&size)
	var minus bool
	if size<0{
		minus=true
		size = size*(-1)
	}
	fmt.Print("Your number is: ",size," - ")

	stringSize:= strconv.Itoa(size)
	numberLength:=len(stringSize)
	numberArray:= strings.Split(stringSize,"")

	if minus{
		fmt.Print("minus ")
	}
	conversion(numberLength,numberArray,size)
}

//up to 9
func oneDigit(sl[]string,counter int){
		n,_:=strconv.Atoi(sl[counter-1])
		switch n{
		case 0:
			fmt.Print("zero")
		case 1:
			fmt.Print("one")
		case 2:
			fmt.Print("two")
		case 3:
			fmt.Print("three")
		case 4:
			fmt.Print("four")
		case 5:
			fmt.Print("five")
		case 6:
			fmt.Print("six")
		case 7:
			fmt.Print("seven")
		case 8:
			fmt.Print("eight")
		case 9:
			fmt.Print("nine")
		default:
			fmt.Println("Not able to convert this number")
		}
}
// from 10 to 19
func upToTwenty(sl[]string,counter int){
		secNum, _ := strconv.Atoi(sl[counter-1])
		switch secNum{
		case 0:
			fmt.Print("ten")
		case 1:
			fmt.Print("eleven")
		case 2:
			fmt.Print("twelve")
		case 3:
			fmt.Print("thirteen")
		case 4:
			fmt.Print("fourteen")
		case 5:
			fmt.Print("fifteen")
		case 6:
			fmt.Print("sixteen")
		case 7:
			fmt.Print("seventeen")
		case 8:
			fmt.Print("eighteen")
		case 9:
			fmt.Print("nineteen")
		default:
			fmt.Println("")
		}
}
//from 20 to 99
func twoDigit(sl[]string,counter int){
	secNum, _ := strconv.Atoi(sl[counter-2])
	switch secNum {
	case 2:
		fmt.Print("twenty ")
		oneDigit(sl, counter)
	case 3:
		fmt.Print("thirty ")
		oneDigit(sl, counter)
	case 4:
		fmt.Print("fourty ")
		oneDigit(sl, counter)
	case 5:
		fmt.Print("fifty ")
		oneDigit(sl, counter)
	case 6:
		fmt.Print("sixty ")
		oneDigit(sl, counter)
	case 7:
		fmt.Print("seventy ")
		oneDigit(sl, counter)
	case 8:
		fmt.Print("eighty ")
		oneDigit(sl, counter)
	case 9:
		fmt.Print("ninty ")
		oneDigit(sl, counter)
	default:
		fmt.Println("")
	}
}
//for hundreds up to 999
func threeDigits(sl[]string,counter int){
	thirdNum, _ := strconv.Atoi(sl[counter-3])
	switch thirdNum {
	case 1:
		fmt.Print("one hundred ")
		hardCheck(sl,counter)
	case 2:
		fmt.Print("two hundred ")
		hardCheck(sl,counter)
	case 3:
		fmt.Print("three hundred ")
		hardCheck(sl,counter)
	case 4:
		fmt.Print("four hundred ")
		hardCheck(sl,counter)
	case 5:
		fmt.Print("five hundred ")
		hardCheck(sl,counter)
	case 6:
		fmt.Print("six hundred ")
		hardCheck(sl,counter)
	case 7:
		fmt.Print("seven hundred ")
		hardCheck(sl,counter)
	case 8:
		fmt.Print("eight hundred ")
		hardCheck(sl,counter)
	case 9:
		fmt.Print("nine hundred ")
		hardCheck(sl,counter)
	default:
		fmt.Println("")
	}
}
//for thousands up to 10k
func fourDigits(sl[]string,counter int){
	one:=1
	oneDigit(sl,one)
	fmt.Print(" thousand ")
	threeDigits(sl,counter)
}
//from 10k to 20k
func fiveDigitsUpToTwenty(sl[]string,counter int){
	two:=2
	upToTwenty(sl,two)
	fmt.Print(" thousand ")
	threeDigits(sl,counter)
}
//thousands till hundreds
func fiveDigits(sl[]string,counter int){
	two:=2
	twoDigit(sl,two)
	fmt.Print(" thousand ")
	threeDigits(sl,counter)
}
// for thousand hundreds
func sixDigits(sl[]string,counter int,fnumber int){
	threeDigits(sl,fnumber)
	fmt.Print(" thousand ")
	threeDigits(sl,counter)
}
//checking decades number in hundreds
func hardCheck(sl[]string,counter int){
	secNum, _ := strconv.Atoi(sl[counter-2])
	if secNum==0{
		oneDigit(sl,counter)
	}
	if secNum==1{
		upToTwenty(sl,counter)
	}else{
		twoDigit(sl,counter)
	}
}
//conversion function
func conversion(counter int,sl[]string,number int){
	switch {
	case counter == 1:
		oneDigit(sl, counter)
	case counter == 2 && number < 20:
		upToTwenty(sl, counter)
	case counter == 2 && number >= 20 && number < 100:
		twoDigit(sl, counter)
	case counter == 3 && number >= 100 && number < 1000:
		threeDigits(sl, counter)
	case counter == 4 && number >= 1000 && number < 10000:
		fourDigits(sl, counter)
	case counter == 5 && number >= 10000 && number < 20000:
		fiveDigitsUpToTwenty(sl, counter)
	case counter == 5 && number >= 20000 && number < 100000:
		fiveDigits(sl, counter)
	case counter == 6 && number >= 100000 && number < 1000000:
		sixDigits(sl, counter, 3)
	case counter == 7 && number >= 1000000 && number < 10000000:
		oneDigit(sl, 1)
		fmt.Print(" million ")
		sixDigits(sl, counter, 4)
	case counter == 8 && number >= 10000000 && number < 20000000:
		upToTwenty(sl, 2)
		fmt.Print(" million ")
		sixDigits(sl, counter, 5)
	case counter == 8 && number >= 20000000 && number < 100000000:
		twoDigit(sl, 2)
		fmt.Print(" million ")
		sixDigits(sl, counter, 5)
	case counter == 9 && number >= 100000000 && number < 1000000000:
		threeDigits(sl,3)
		fmt.Print(" million ")
		sixDigits(sl, counter, 5)
	}
}


