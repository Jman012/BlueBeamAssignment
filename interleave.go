package main

import (
	"bufio"
	"fmt"
	"os"
)

// Interleave the words of the two given files into one stream of words,
// written to the buffered io writer.
func Interleave(file1Path string, file2Path string, writer *bufio.Writer) {

	// Open the files and get a word scanner.
	file1, scanner1 := openForWords(file1Path)
	defer file1.Close()

	file2, scanner2 := openForWords(file2Path)
	defer file2.Close()

	// Pass the word scanners for each file into the dedicated
	// function to perform the interleaving.
	interleaveScanners(scanner1, scanner2, writer)
}

// Helper function for opening and preparing os files/scanners.
func openForWords(filePath string) (*os.File, *bufio.Scanner) {
	// Open the file and fatally end the program if there's an error,
	// such as missing files or read permissions.
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file '%v': %v\n", filePath, err)
		os.Exit(1)
	}

	// Setup a buffered scanner for words from the file (as opposed
	// to reading all contents of the file in initially).
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	return file, scanner
}

// Implementation of interleaving two streams of words into the resulting writer.
func interleaveScanners(scanner1 *bufio.Scanner, scanner2 *bufio.Scanner, writer *bufio.Writer) {
	// Handle the logic of spaces between the words for us, to reduce copied/pasted code.
	spacedWordWriter := NewSpacedWordWriter(writer)
	// Flags to keep track of when we should break the loop.
	wrote1, wrote2 := false, false
	for {
		wrote1 = false
		wrote2 = false

		if scanner1.Scan() {
			spacedWordWriter.WriteWord(scanner1.Text())
			wrote1 = true
		}

		if scanner2.Scan() {
			spacedWordWriter.WriteWord(scanner2.Text())
			wrote2 = true
		}

		if !wrote1 && !wrote2 {
			break
		}
	}

	writer.Flush()
}
