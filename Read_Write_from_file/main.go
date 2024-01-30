package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("hello, this is the start of the program")
	// Open a file
	text, err := os.Open("./text_file/file.txt")
	if err != nil {
		fmt.Printf("Error is: %v", err)

	}
	// Close a file after all operations are completed;{Cleanup}
	defer text.Close()
	// Create a new file
	out, err := os.Create("./text_file/output.txt")
	if err != nil {
		fmt.Printf("Error is:%v", err)

	}
	// Close a file after all operations are completed;{Cleanup}
	defer out.Close()

	scanner := bufio.NewScanner(text)

	for scanner.Scan() {
		// Scan the whole file line by line
		line := scanner.Text()
		fmt.Printf("The command is: %v\n ", line)
		// Write the the above print statement in the output.txt file
		_, err := out.WriteString(fmt.Sprintf("The command is: %v\n ", line))
		if err != nil {
			fmt.Printf("Error is:%v", err)
		}
		// Execute the command in bash terminal
		cmd := exec.Command("bash", "-c", line)
		if err != nil {
			fmt.Printf("Error is:%v", err)
		}
		// Store o/p in stdout
		stdout, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error is:%v", err)
		}
		// Print the stdout in string format
		fmt.Printf("The Output is: %v\n ", string(stdout))
		// Write the the above print statement in the output.txt file
		_, err = out.WriteString(fmt.Sprintf("The output is: %v\n ", (string(stdout))))

		if err != nil {
			fmt.Printf("Error is:%v", err)
		}

	}
}
