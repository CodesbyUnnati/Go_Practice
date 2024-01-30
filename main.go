package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, this is the start of the program")
	// Read from a file
	text, err := os.ReadFile("./text_file/file.txt")
	if err != nil {
		fmt.Printf("Error is:%v", err)

	}
	//Print from the file
	fmt.Println(string(text))
	//Open the file
	ok, err := os.Open("./text_file/file.txt")
	if err != nil {
		fmt.Printf("Error is:%v", err)

	}

	defer ok.Close()

	scanner := bufio.NewScanner(ok)
	//Print text one by one on loop
	i := 0
	//This is for + while condition together
	for scanner.Scan() {

		one := scanner.Text()
		fmt.Println(i, one)
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error is:%v", err)
	}
}
