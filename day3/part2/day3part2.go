package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines() []string {
	file, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	return lines
}

func findNextBattery(batteryBank string, trimEndBy int) int {
	nextBatteryAsInt, err := strconv.Atoi(batteryBank[0:1])
	nextBatteryIndex := 0

	if err != nil {
		log.Fatal(err)
	}

	for j := 1; j < len(batteryBank)-trimEndBy; j++ {

		// can't be more than 9 so save ourselves some loops
		if nextBatteryAsInt == 9 {
			break
		}

		batteryInt, err := strconv.Atoi(batteryBank[j : j+1])

		if err != nil {
			log.Fatal(err)
		}

		if batteryInt > nextBatteryAsInt {
			nextBatteryAsInt = batteryInt
			nextBatteryIndex = j
		}
	}

	return nextBatteryIndex
}

func main() {
	batteryBanks := readLines()
	joltageSum := 0

	// iterate through battery banks
	for i := 0; i < len(batteryBanks); i++ {
		batteryBank := batteryBanks[i]
		joltage := ""

		for batteryIndex := 0; batteryIndex < 12; batteryIndex++ {
			trimEndBy := 11 - batteryIndex
			nextBattery := findNextBattery(batteryBank, trimEndBy)
			joltage += batteryBank[nextBattery : nextBattery+1]
			batteryBank = batteryBank[nextBattery+1:]
		}

		joltageAsInt, err := strconv.Atoi(joltage)
		if err != nil {
			log.Fatal(err)
		}

		joltageSum += joltageAsInt
	}

	fmt.Println(joltageSum)
}
