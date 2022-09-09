package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestInterleave(t *testing.T) {
	// Prepare a string buffer and a buffered writer for it,
	// as opposed to testing with stdout.
	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)

	expected := "cat fox apple dog zebra orange pluto"

	Interleave("a.txt", "b.txt", writer)

	got := buf.String()
	if got != expected {
		t.Errorf("Got '%v', expected '%v'", got, expected)
	}
}

func TestInterleaveScanners(t *testing.T) {
	testInternalInterleave(t, "cat apple zebra", "fox dog orange pluto", "cat fox apple dog zebra orange pluto")
}

func TestEmpty(t *testing.T) {
	testInternalInterleave(t, "", "", "")
}

func TestMixed(t *testing.T) {
	testInternalInterleave(t, "", "cat dog", "cat dog")
	testInternalInterleave(t, "cat dog", "", "cat dog")
}

func TestSpaces(t *testing.T) {
	testInternalInterleave(t, " cat     apple ", " dog ", "cat dog apple")
}

// Helper function to convert the inputs into scanners and perform the testing.
func testInternalInterleave(t *testing.T, str1 string, str2 string, expected string) {
	// Prepare a string buffer and a buffered writer for it,
	// as opposed to testing with stdout.
	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)

	// Setup word-separated scanners for the input strings
	scanner1 := bufio.NewScanner(strings.NewReader(str1))
	scanner1.Split(bufio.ScanWords)
	scanner2 := bufio.NewScanner(strings.NewReader(str2))
	scanner2.Split(bufio.ScanWords)

	// Perform the interleaving with the inputs
	interleaveScanners(scanner1, scanner2, writer)

	// Check
	got := buf.String()
	if got != expected {
		t.Errorf("Got '%v', expected '%v'", got, expected)
	}
}
