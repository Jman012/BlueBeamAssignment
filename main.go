package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Definte input flags for the two files to be interleaved.
var file1 = flag.String("file1", "", "The first file of words to interleave")
var file2 = flag.String("file2", "", "The second file of words to interleave")

func main() {
	// Parse os args into the flags for us automatically.
	flag.Parse()

	// Exit early if neither of the two files were passed into the program.
	if len(*file1) == 0 || len(*file2) == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Send the file paths into the function, and ask it to write to stdout.
	Interleave(*file1, *file2, bufio.NewWriter(os.Stdout))
	fmt.Println() // End with a newline for terminals
}
