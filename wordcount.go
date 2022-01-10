package main

import "strings"

func wordcount(src string) (int, error) {
	words := strings.Fields(src)
	return len(words), nil
}
