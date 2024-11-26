package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", scanner.Err())
		os.Exit(1)
	}
	fmt.Printf("\n%-30s Count\n", "Word")
	for w, c := range freq {
		fmt.Printf("%-30s %d\n", w, c)
	}
}
