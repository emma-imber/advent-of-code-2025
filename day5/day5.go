package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines() ([]string, []string) {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileParts := strings.Split(string(file), "\n\n")
	ranges := strings.Split(fileParts[0], "\n")
	ids := strings.Split(fileParts[1], "\n")

	return ranges, ids
}

func mergeRanges(rangeStartAndEnds [][]int) [][]int {
	combinedRanges := [][]int{}

	for l := 0; l < len(rangeStartAndEnds); l++ {
		added := false
		for m := 0; m < len(combinedRanges); m++ {
			// check if start of range is within an existing range
			if rangeStartAndEnds[l][0] >= combinedRanges[m][0] && rangeStartAndEnds[l][0] <= combinedRanges[m][1] {
				// if end of range is bigger, extend existing range
				if rangeStartAndEnds[l][1] > combinedRanges[m][1] {
					combinedRanges[m][1] = rangeStartAndEnds[l][1]
				}
				added = true
				break
			// check if end of range is within an existing range
			} else if rangeStartAndEnds[l][1] >= combinedRanges[m][0] && rangeStartAndEnds[l][1] <= combinedRanges[m][1] {
				// if start of range is smaller, extend existing range
				if rangeStartAndEnds[l][0] < combinedRanges[m][0] {
					combinedRanges[m][0] = rangeStartAndEnds[l][0]
				}
				added = true
				break
			// check if range completely contains existing range
			} else if rangeStartAndEnds[l][0] <= combinedRanges[m][0] && rangeStartAndEnds[l][1] >= combinedRanges[m][1] {
				combinedRanges[m] = rangeStartAndEnds[l]
				added = true
				break
			}
		}
		if !added {
			combinedRanges = append(combinedRanges, rangeStartAndEnds[l])
		}
	}

	return combinedRanges
}

func main() {
	ranges, ids := readLines()
	rangeStartAndEnds := [][]int{}

	// iterate through ranges to build map of fresh ids
	for i := 0; i < len(ranges); i++ {
		// get start and end of range as ints
		rangeStart, err1 := strconv.Atoi(strings.Split(ranges[i], "-")[0])
		rangeEnd, err2 := strconv.Atoi(strings.Split(ranges[i], "-")[1])

		if err1 != nil || err2 != nil {
			log.Fatal(err1)
		}

		rangeStartAndEnds = append(rangeStartAndEnds, []int{rangeStart, rangeEnd})
	}

	freshIdCounter := 0

	// part 1
	// iterate through ids to check which are fresh
	for j := 0; j < len(ids); j++ {
		idAsInt, err := strconv.Atoi(ids[j])

		if err != nil {
			log.Fatal(err)
		}

		for k := 0; k < len(rangeStartAndEnds); k++ {
			if idAsInt >= rangeStartAndEnds[k][0] && idAsInt <= rangeStartAndEnds[k][1] {
				freshIdCounter += 1
				break
			}
		}
	}

	fmt.Println(freshIdCounter)

	// part 2
	// iterate through ranges to combine any that overlap
	combinedRanges := [][]int{}

	for {
		combinedRanges = mergeRanges(rangeStartAndEnds)

		// keep merging until no more changes
		if len(combinedRanges) == len(rangeStartAndEnds) {
			break
		} else {
			rangeStartAndEnds = combinedRanges
		}
	}

	freshIdsFromRangesCounter := 0

	// iterate through merged ranges to count ids
	for n := 0; n < len(combinedRanges); n++ {
		freshIdsFromRangesCounter += combinedRanges[n][1] - combinedRanges[n][0] + 1
	}

	fmt.Println((freshIdsFromRangesCounter))
}
