package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main(){
	/*reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path: ")
	input, _ := reader.ReadString('\n')*/
	var input string
	fmt.Println("Enter file path: ")
	fmt.Scanf("%s",&input)     //getting our input(file path)
	file,err:= os.Open(input)
	//error check(either file os empty or not)
	if err != nil{
		fmt.Println(err)

		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 10*64)    //stupid as shit, but I can't come up with another solution sorry
	//making a slice of bits read from our file
	for{
		n, err := file.Read(data)
		if err == io.EOF{
			break
		}
		fmt.Print(string(data[:n]))
	}

//D:\task6test.txt
//two counters for two lucky tickets count types - moskow and piter
	ticketsCount1 :=0
	ticketsCount2 :=0
	//file needed inners: "count type""gap""6-number ticket"
	//example: 1 369963 ( lucky Moskow)
	for i:=0;i<len(data);i++{
		numberConv,_:= strconv.Atoi(string(data[i]))
		if numberConv==1 && string(data[i+1])==" "{
			number1,_:= strconv.Atoi(string(data[i+2]))
			number2,_:= strconv.Atoi(string(data[i+3]))
			number3,_:= strconv.Atoi(string(data[i+4]))        //"Moskow" type
			number4,_:= strconv.Atoi(string(data[i+5]))
			number5,_:= strconv.Atoi(string(data[i+6]))
			number6,_:= strconv.Atoi(string(data[i+7]))
			sum1:=number1+number2+number3
			sum2:=number4+number5+number6
			if sum1==sum2{
				ticketsCount1++
			}
		}
		if numberConv==2 && string(data[i+1])==" "{
			number1,_:= strconv.Atoi(string(data[i+2]))
			number2,_:= strconv.Atoi(string(data[i+3]))
			number3,_:= strconv.Atoi(string(data[i+4]))
			number4,_:= strconv.Atoi(string(data[i+5]))        //"Piter" type
			number5,_:= strconv.Atoi(string(data[i+6]))
			number6,_:= strconv.Atoi(string(data[i+7]))
			sum1:=number1+number3+number5
			sum2:=number2+number4+number6
			if sum1==sum2{
				ticketsCount2++
			}
		}
	}
	//result
	fmt.Println("So, in this file there are ",ticketsCount1,"Moskow-type lucky tickets and ",ticketsCount2,"Piter-type lucky tickets")

}

