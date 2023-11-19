package myaochelper

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Takes a file and returns a slice with all the runes
func parse_runes(f *os.File) []rune {
	var chars []rune

	reader := bufio.NewReader(f)
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Rune reading error: ", err)
			os.Exit(1)
		}

		chars = append(chars, char)
	}

	return chars
}

// Takes a file and returns a slice of strings with all the lines
func parse_lines(f *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
