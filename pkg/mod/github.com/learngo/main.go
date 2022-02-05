package main

import "fmt"

import "strings"

func lenAndUpper(name string) (int, string) {
	return len(name), string.ToUpper(name)
}

func main() {
	totalLenght, upperName := lenAndUpper("nico")
	fmt.println(totalLength, upperName)
}
