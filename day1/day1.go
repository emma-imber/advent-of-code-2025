package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// create function to read input file into an array as we'll be doing this every day
func readLines() []string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

func main() {
	rotations := readLines()

	// initialise dial position and counters
	dialPosition := 50
	// part 1
	numberOfZerosAtEndOfRotation := 0
	// part 2
	numberOfZerosDuringRotation := 0

	// iterate through rotation instructions
	for step := 0; step < len(rotations); step++ {
		// check number to move dial by
		rotationNumber, err := strconv.Atoi(rotations[step][1:])
		if err != nil {
			log.Fatal(err)
		}
		// increment clicks one by one
		for click := 0; click < rotationNumber; click++ {
			// check direction of rotation
			if rotations[step][0] == 'L' {
				dialPosition -= 1
			} else {
				dialPosition += 1
			}
			// if dial position is on a zero after the click, increment counter
			if (dialPosition % 100) == 0 {
				numberOfZerosDuringRotation += 1
			}
		}
		// if dial position is on a zero at the end of rotation, increment counter
		if (dialPosition % 100) == 0 {
			numberOfZerosAtEndOfRotation += 1
		}
	}

	fmt.Println(numberOfZerosAtEndOfRotation)
	fmt.Println(numberOfZerosDuringRotation)
}
