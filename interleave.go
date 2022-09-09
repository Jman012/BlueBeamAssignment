package main

import (
	"bufio"
	"fmt"
	"os"
)

func Interleave(file1Path string, file2Path string, writer *bufio.Writer) {

	file1, err := os.Open(file1Path)
	if err != nil {
		fmt.Printf("Error opening file '%v': %v\n", file1Path, err)
		os.Exit(1)
	}
	defer file1.Close()

	file2, err := os.Open(file2Path)
	if err != nil {
		fmt.Printf("Error opening file '%v': %v\n", file2Path, err)
		os.Exit(1)
	}
	defer file2.Close()

	scanner1 := bufio.NewScanner(file1)
	scanner1.Split(bufio.ScanWords)

	scanner2 := bufio.NewScanner(file2)
	scanner2.Split(bufio.ScanWords)

	interleaveScanners(scanner1, scanner2, writer)
}

func interleaveScanners(scanner1 *bufio.Scanner, scanner2 *bufio.Scanner, writer *bufio.Writer) {
	wrote1, wrote2 := false, false
	for {
		wrote1 = false
		wrote2 = false

		if scanner1.Scan() {
			writer.WriteString(scanner1.Text())
			writer.WriteString(" ")
			wrote1 = true
		}

		if scanner2.Scan() {
			writer.WriteString(scanner2.Text())
			writer.WriteString(" ")
			wrote2 = true
		}

		if !wrote1 && !wrote2 {
			break
		}
	}

	writer.Flush()
}
