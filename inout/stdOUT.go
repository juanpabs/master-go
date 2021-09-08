package inout

import (
	"io" //io package instead of the fmt
	"os" //for read command-line arguments and use os.Stdout
)

func StdOUT() {
	myString := ""
	arguments := os.Args 
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString) // io.WriteString works in the same way as the fmt.Print()
	io.WriteString(os.Stdout, "\n")
}
