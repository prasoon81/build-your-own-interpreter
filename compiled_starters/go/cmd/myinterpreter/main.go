package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// But make sure to write the output to stderr.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Error: Missing file argument")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting file info: %v\n", err)
		os.Exit(1)
	}

	fileContents := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContents)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage
	// if len(fileContents) > 0 {
	// 	panic("Scanner not implemented")
	// } else {
	// 	fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	// }
}