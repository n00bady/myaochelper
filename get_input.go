package myaochelper

import (
	"fmt"
	"os"
)

// get the filename as a string and open the text file
func Read_Text_Input(filename string) *os.File {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error: ", err)
		os.Exit(1)
	}
	return f
}
