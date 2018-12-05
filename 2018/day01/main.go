package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Load in Example File
	inputB, err := ioutil.ReadFile("inputs/puzzleInput")
	if err != nil {
		log.Fatal(err)
	}

	// Transform slice of bytes to string and Split into
	// a slice of string
	inputS := strings.Split(string(inputB), "\n")

	frequencyResults := 0

	for _, f := range inputS {
		change, err := strconv.Atoi(f)
		if err != nil {
			log.Fatal(err)
		}
		frequencyResults += change
	}

	fmt.Println("Advent of Code - Day 1")
	fmt.Println(frequencyResults)

}
