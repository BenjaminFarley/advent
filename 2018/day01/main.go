package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Part 1
	fmt.Println("Starting Part 2")
	input1S := openFileAndParse("inputs/Part2Input")
	// Set up results value
	frequencyResults := 0

	// Iterate over slice of string, converting each
	// value into an int and adding value to results
	for _, f := range input1S {
		frequencyResults += f
	}

	// Print out result
	fmt.Println("Advent of Code - Day 1 - Part 1")
	fmt.Println(frequencyResults)

	// Part 2
	fmt.Println("Starting Part 2")
	input2S := openFileAndParse("inputs/Part2Input")
	fmt.Println("Advent of Code - Day 1 - Part 2")
	findDuplicate(input2S)

}

// openFileAndParse takes a filePath to a file that contains
// frequencies and converts them into a slice of int
func openFileAndParse(filePath string) []int {
	// Load in File
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Transform slice of bytes to string and Split into
	// a slice of string
	listOfStrings := strings.Split(string(file), "\n")

	var numbers []int
	for _, n := range listOfStrings {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}
	return numbers
}

// findDuplicate takes and parses a slice of frequencies
// to find the first duplicate value
func findDuplicate(input []int) int {

	// Set up results value
	frequencyResult := 0

	foundDuplicate := map[int]bool{0: true}

	// Iterate over input frequencies and check to see if
	// the value is already stored in frequencyResults. If
	// not, then append the result in frequencyResults.
	for {
		for _, n := range input {
			frequencyResult += n
			if foundDuplicate[frequencyResult] {
				fmt.Println("Found a duplicate")
				fmt.Println(frequencyResult)
				return frequencyResult
			}
			foundDuplicate[frequencyResult] = true
		}
	}
}
