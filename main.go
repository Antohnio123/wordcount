package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	count, err := wordcount(src)
	if count == 0 {
		fmt.Println("Строка не входит в паттерн")
		os.Exit(0)
	}
	fmt.Println(count)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string to count words")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("wordcount:", err)
	os.Exit(1)
}
