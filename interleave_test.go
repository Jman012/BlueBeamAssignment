package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestInterleaveScanners(t *testing.T) {
	buf := bytes.NewBufferString("")
	writer := bufio.NewWriter(buf)

	str1 := "cat apple zebra"
	str2 := "fox dog orange pluto"
	expected := "cat fox apple dog zebra orange pluto"

	scanner1 := bufio.NewScanner(strings.NewReader(str1))
	scanner1.Split(bufio.ScanWords)
	scanner2 := bufio.NewScanner(strings.NewReader(str2))
	scanner2.Split(bufio.ScanWords)

	interleaveScanners(scanner1, scanner2, writer)

	got := buf.String()
	if got != expected {
		t.Errorf("Got '%v', expected '%v'", got, expected)
	}
}
