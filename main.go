/*
The Interleave CLI utility takes two files as input arguments and produces an
output of the individual words of the two files interleaved together.
Internally, this uses buffered IO readers to read the files (to account for
potentially very large files to prevent memory saturation). The main function
also uses a buffered IO writer to write its output back to stdout for similar
reasons. Additionally, this allows for better unit testing so that the
functionality can be tested without having to read the output from stdout;
instead, it can be written to a string and simply checked.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// Define input flags for the two files to be interleaved.
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
