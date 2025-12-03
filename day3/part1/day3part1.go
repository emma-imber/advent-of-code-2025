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

func main() {
	batteryBanks := readLines()
	joltageSum := 0

	// iterate through battery banks
	for i := 0; i < len(batteryBanks); i++ {
		batteryBank := batteryBanks[i]
		firstBatteryAsInt, err := strconv.Atoi(batteryBank[0:1])
		firstBatteryIndex := 0

		if err != nil {
			log.Fatal(err)
		}

		// find first battery - will never be the battery at the end so slice off end
		for j := 1; j < len(batteryBank)-1; j++ {
			// can't be more than 9 so save ourselves some loops
			if firstBatteryAsInt == 9 {
				break
			}

			batteryInt, err := strconv.Atoi(batteryBank[j : j+1])

			if err != nil {
				log.Fatal(err)
			}

			if batteryInt > firstBatteryAsInt {
				firstBatteryAsInt = batteryInt
				firstBatteryIndex = j
			}

		}

		// set second battery as one after first
		secondBatteryAsInt, err := strconv.Atoi(batteryBank[firstBatteryIndex+1 : firstBatteryIndex+2])
		secondBatteryIndex := firstBatteryIndex + 1

		if err != nil {
			log.Fatal(err)
		}

		// find second battery - iterate from first battery to end
		for k := firstBatteryIndex + 1; k < len(batteryBank); k++ {
			// can't be more than 9 so save ourselves some loops
			if secondBatteryAsInt == 9 {
				break
			}

			batteryInt, err := strconv.Atoi(batteryBank[k : k+1])

			if err != nil {
				log.Fatal(err)
			}

			if batteryInt > secondBatteryAsInt {
				secondBatteryAsInt = batteryInt
				secondBatteryIndex = k
			}

		}

		joltage, err := strconv.Atoi(batteryBank[firstBatteryIndex:firstBatteryIndex+1] + batteryBank[secondBatteryIndex:secondBatteryIndex+1])

		if err != nil {
			log.Fatal(err)
		}

		joltageSum += joltage
	}

	fmt.Println(joltageSum)
}
