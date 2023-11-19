package myaochelper

import (
	"fmt"
	"os"
)

// get the filename as a string and open the text file
func read_text_input(filename string) *os.File {
	f, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("File reading error: ", err)
		os.Exit(1)
	}
	defer f.Close()

	return f
}
