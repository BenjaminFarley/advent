package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputS := openFileAndParse("inputs/Part1Input")
	// Set up results value
	frequencyResults := 0

	// Iterate over slice of string, converting each
	// value into an int and adding value to results
	for _, f := range inputS {
		change, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		frequencyResults += change
	}

	// Print out result
	fmt.Println("Advent of Code - Day 1")
	fmt.Println(frequencyResults)

}

// openFileAndParse takes a filePath to a file that contains
// frequencies and returns a slice of string
func openFileAndParse(filePath string) []string {
	// Load in File
	inputB, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Transform slice of bytes to string and Split into
	// a slice of string
	return strings.Split(string(inputB), "\n")
}
