package main

import (
	"chapter3/words"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "specify a single filename")
		os.Exit(1)
	}
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d words in %s.\n", count, filename)
}
