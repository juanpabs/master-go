package inout

import (
	"bufio" // for file input output
	"fmt"
	"os" // functions idependent of the OS (LINUX or windows)
)

func StdIN() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f) // create a scanner to read command line (Stdin)
	for scanner.Scan() {           // scanner.Scan() to read line by line
		fmt.Println(">", scanner.Text())
	}
}
