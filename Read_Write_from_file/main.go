package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("hello, this is the start of the program")

	text, err := os.Open("./text_file/file.txt")
	if err != nil {
		fmt.Printf("Error is: %v", err)

	}

	defer text.Close()

	out, err := os.Create("./text_file/output.txt")
	if err != nil {
		fmt.Printf("Error is:%v", err)

	}
	defer out.Close()

	scanner := bufio.NewScanner(text)

	for scanner.Scan() {

		line := scanner.Text()
		fmt.Printf("The command is: %v\n ", line)
		_, err := out.WriteString(fmt.Sprintf("The command is: %v\n ", line))
		if err != nil {
			fmt.Printf("Error is:%v", err)
		}

		cmd := exec.Command("bash", "-c", line)
		if err != nil {
			fmt.Printf("Error is:%v", err)
		}

		stdout, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error is:%v", err)
		}
		fmt.Printf("The Output is: %v\n ", string(stdout))

		_, err = out.WriteString(fmt.Sprintf("The output is: %v\n ", (string(stdout))))

		if err != nil {
			fmt.Printf("Error is:%v", err)
		}

	}
}
