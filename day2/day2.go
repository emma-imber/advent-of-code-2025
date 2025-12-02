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

	lines := strings.Split(string(file), ",")

	return lines
}

func main() {
	ranges := readLines()
	//sumOfInvalidIdsPart1 := 0
	sumOfInvalidIdsPart2 := 0

	// iterate through list of ranges
	for i := 0; i < len(ranges); i++ {
		// get start and end of range as ints
		rangeStart, err1 := strconv.Atoi(strings.Split(ranges[i], "-")[0])
		rangeEnd, err2 := strconv.Atoi(strings.Split(ranges[i], "-")[1])

		if err1 != nil || err2 != nil {
			log.Fatal(err1)
		}

		// iterate through range of ids
		for id := rangeStart; id < rangeEnd+1; id++ {
			idAsStr := strconv.Itoa(id)

			// this is for part 1
			// if an id has an odd number of digits it can't be invalid
			// if len(idAsStr)%2 != 0 {
			// 	continue
			// }

			// if idAsStr[0:len(idAsStr)/2] == idAsStr[len(idAsStr)/2:] {
			// 	sumOfInvalidIdsPart1 += id
			// }

			// this is for part 2
			// if an id is made up of repeating substrings, that id will be found 3 times in a string made up of the id repeated twice
			// eg if id is 1212, id twice is 12121212, you find a match for id substring in the middle
			// if we trim first and last characters from 12121212 to get 212121, this contains 1212, our id, once

			doubleId := idAsStr + idAsStr
			trimmedDoubleId := doubleId[1 : len(doubleId)-1]

			if strings.Contains(trimmedDoubleId, idAsStr) {
				sumOfInvalidIdsPart2 += id
			}
		}
	}

	fmt.Println(sumOfInvalidIdsPart2)
}
