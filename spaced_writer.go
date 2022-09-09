package main

import "bufio"

// A simple wrapper around a bufio.Writer that writes words
// separated by spaces, and prevents a leading or trailing
// space for the entire output.
type SpacedWordWriter struct {
	writer     *bufio.Writer
	hasWritten bool
}

// Initialize a new SpacedWordWriter with a suppled bufio.Writer.
func NewSpacedWordWriter(writer *bufio.Writer) *SpacedWordWriter {
	return &SpacedWordWriter{
		writer:     writer,
		hasWritten: false,
	}
}

// Write a word to the buffered writer. The first word written will
// have no leading or trailing space. Subsequent words will have a
// space writtn before the word.
func (sw *SpacedWordWriter) WriteWord(word string) {
	if sw.hasWritten {
		sw.writer.WriteString(" ")
	}
	sw.writer.WriteString(word)
	sw.hasWritten = true
}
