package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var file1 = flag.String("file1", "", "The first file of words to interleave")
var file2 = flag.String("file2", "", "The second file of words to interleave")

func main() {
	flag.Parse()

	if len(*file1) == 0 || len(*file2) == 0 {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	Interleave(*file1, *file2, bufio.NewWriter(os.Stdout))
	fmt.Println()
}
