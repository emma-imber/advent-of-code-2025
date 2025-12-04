package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readLines() [][]rune {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	linesAsRunes := make([][]rune, len(lines))

	for i, line := range lines {
		linesAsRunes[i] = []rune(line)
	}

	return linesAsRunes
}

func removeEligibleRolls(grid [][]rune) ([][]rune, int) {
	numberToRemove := 0
	var updatedGrid [][]rune

	for y := 0; y < len(grid); y++ {
		rowLength := len(grid[0])
		newRow := make([]rune, rowLength)
		for x := 0; x < rowLength; x++ {
			adjacentChars := ""
			if string(grid[y][x]) == "@" {
				if x > 0 {
					adjacentChars += string(grid[y][x-1])
				}
				if x < rowLength-1 {
					adjacentChars += string(grid[y][x+1])
				}
				if y > 0 {
					adjacentChars += string(grid[y-1][x])
				}
				if y < len(grid)-1 {
					adjacentChars += string(grid[y+1][x])
				}
				if x > 0 && y > 0 {
					adjacentChars += string(grid[y-1][x-1])
				}
				if x < rowLength-1 && y > 0 {
					adjacentChars += string(grid[y-1][x+1])
				}
				if x > 0 && y < len(grid)-1 {
					adjacentChars += string(grid[y+1][x-1])
				}
				if x < rowLength-1 && y < len(grid)-1 {
					adjacentChars += string(grid[y+1][x+1])
				}

				if strings.Count(adjacentChars, "@") < 4 {
					numberToRemove += 1
					newRow[x] = rune('x')
				} else {
					newRow[x] = rune('@')
				}
			} else {
				newRow[x] = rune('.')
			}
		}
		updatedGrid = append(updatedGrid, newRow)
	}

	return updatedGrid, numberToRemove
}

func main() {
	grid := readLines()
	totalNumberRemoved := 0

	for {
		newGrid, numberRemoved := removeEligibleRolls(grid)
		if numberRemoved == 0 {
			break
		} else {
			grid = newGrid
			totalNumberRemoved += numberRemoved
			fmt.Println(totalNumberRemoved)
		}
	}
}
