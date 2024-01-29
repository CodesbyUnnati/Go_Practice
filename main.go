package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello, this is the start of the program")
	// text, err := os.ReadFile("./text_file/file.txt")
	//  if err != nil {
	// 	fmt.Printf("Error is:%v", err)

	// }
	//fmt.Println(string(text))
	text, err := os.Open("./text_file/file.txt")
	if err != nil {
		fmt.Printf("Error is:%v", err)

	}

	defer text.Close()

	scanner := bufio.NewScanner(text)

	i := 0
	for scanner.Scan() {

		one := scanner.Text()
		fmt.Println(i, one)
		i++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error is:%v", err)
	}
}
