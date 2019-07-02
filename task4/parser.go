package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count("task4/JackWelch.txt", "Джек")
	insert("task4/JackWelch.txt", "Бильбо Беггинс", "Джек Уелч")
	bufInsert("task4/JackWelch.txt", "Джек", "Бильбо")
}

func count(filename string, str string) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Unable to open file", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var count int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}
		output2 := strings.Count(line, str)
		if output2 > 0 {
			count++
		}
	}
	fmt.Println(count)
}

func insert(filename string, find string, replaced string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	output := bytes.Replace(input, []byte(find), []byte(replaced), -1)

	if err = ioutil.WriteFile(filename, output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func bufInsert(filename string, find string, replaced string) {
	// open input file
	fi, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	// close fi on exit and check for its returned error
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()
	// make a read buffer
	r := bufio.NewReader(fi)

	// open output file
	fo, err := os.Create("task4/output.txt")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	for {
		n, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len(n) == 0 {
			break
		}

		output := strings.Replace(n, find, replaced, -1)
		if _, err := w.Write([]byte(output)); err != nil {
			panic(err)
		}
	}
	if err = w.Flush(); err != nil {
		panic(err)
	}
}
