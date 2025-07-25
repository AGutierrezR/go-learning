package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Defining a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Defining a boolean flag -b to count bytes instead of words
	bytes := flag.Bool("b", false, "Count bytes")

	// Parse the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of world
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scannner is used to read text from Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines and bytes flags are not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		// If the count lines flag not set and count bytes is set, we want to count bytes so we define
		scanner.Split(bufio.ScanBytes)
	}

	// Defining a counter
	wc := 0

	for scanner.Scan() {
		// Increment the counter for each word found
		wc++
	}

	// Return the total
	return wc
}
