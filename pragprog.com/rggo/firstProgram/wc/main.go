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

	// Parse the flags provided by the user
	flag.Parse()

	// Calling the count function to count the number of world
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countLines bool) int {
	// A scannner is used to read text from Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lies flag is not set, we want to count words so we define
	// the scanner split type to words (default is split by lines)
	if !countLines {
		scanner.Split(bufio.ScanWords)
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

