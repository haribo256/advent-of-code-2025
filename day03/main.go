package main

import (
	"bufio"
	"fmt"
	"iter"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	Batteries []*Battery
	AsString  string
}

type Battery struct {
	Value     uint8
	Index     int
	IsJoltage bool
}

func main() {
	part1()
	part2()
}

func readInputBanks() iter.Seq[Bank] {
	return func(yield func(V Bank) bool) {
		file, err := os.Open("input.txt")
		if err != nil {
			log.Fatalln("Failed to open input.txt")
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" {
				continue
			}

			batteries := make([]*Battery, len(line))
			for batteryIndex, batteryValueString := range line {
				batteryValue, err := strconv.Atoi(string(batteryValueString))
				if err != nil {
					log.Fatalln("failed to read voltage from line", line)
				}

				batteries[batteryIndex] = &Battery{
					Value:     uint8(batteryValue),
					Index:     batteryIndex,
					IsJoltage: false,
				}
			}

			if !yield(Bank{
				Batteries: batteries,
				AsString:  line,
			}) {
				return
			}
		}
	}
}

func part1() {
	totalJoltage := 0
	banks := readInputBanks()

	for bank := range banks {
		topHighestValue := uint8(0)
		topHighestIndex := -1
		nextHighestValue := uint8(0)
		nextHighestIndex := -1

		for batteryIndex, batteryValue := range bank.Batteries[:len(bank.Batteries)-1] {
			if batteryValue.Value > topHighestValue {
				topHighestValue = batteryValue.Value
				topHighestIndex = batteryIndex
			}
		}

		for batteryIndex, batteryValue := range bank.Batteries[topHighestIndex+1:] {
			if batteryValue.Value > nextHighestValue {
				nextHighestValue = batteryValue.Value
				nextHighestIndex = batteryIndex + topHighestIndex + 1
			}
		}

		if topHighestIndex == -1 || nextHighestIndex == -1 {
			log.Fatalln("failed to read enough values")
		}

		var bankJoltage uint8
		if topHighestIndex < nextHighestIndex {
			bankJoltage = topHighestValue*10 + nextHighestValue
		} else {
			bankJoltage = nextHighestValue*10 + topHighestValue
		}

		firstIndex := min(topHighestIndex, nextHighestIndex)
		nextIndex := max(topHighestIndex, nextHighestIndex)

		totalJoltage += int(bankJoltage)

		fmt.Printf("PART1: bank=%s[%s]%s[%s]%s, jolt=%d, total=%d\n",
			bank.AsString[:firstIndex], bank.AsString[firstIndex:firstIndex+1], bank.AsString[firstIndex+1:nextIndex], bank.AsString[nextIndex:nextIndex+1], bank.AsString[nextIndex+1:],
			bankJoltage,
			totalJoltage,
		)
	}

	fmt.Println("PART1: SUM JOLTAGE", totalJoltage)
}

func part2() {
	totalJoltage := 0

	banks := readInputBanks()

	for bank := range banks {
		joltageLength := 12
		rangeStart := 0
		for joltageIndex := 0; joltageIndex < joltageLength; joltageIndex++ {
			joltageRemaining := joltageLength - joltageIndex
			rangeEnd := len(bank.Batteries) - (joltageRemaining - 1)

			batterySearch := bank.Batteries[rangeStart:rangeEnd]
			var highestBattery *Battery = batterySearch[0]
			for _, battery := range batterySearch {
				if battery.Value > highestBattery.Value {
					highestBattery = battery
				}
			}

			rangeStart = highestBattery.Index + 1
			highestBattery.IsJoltage = true
		}

		joltageString := ""
		for _, battery := range bank.Batteries {
			if battery.IsJoltage {
				joltageString += strconv.Itoa(int(battery.Value))
			}
		}

		joltage, err := strconv.Atoi(joltageString)
		if err != nil {
			log.Fatalln("failed to convert joltage to int")
		}

		totalJoltage += joltage

		fmt.Print("PART2: bank=")
		for _, battery := range bank.Batteries {
			if battery.IsJoltage {
				fmt.Printf("[%d]", battery.Value)
			} else {
				fmt.Printf("%d", battery.Value)
			}
		}
		fmt.Printf(", ")
		fmt.Printf("jolt=%d, total=%d\n", joltage, totalJoltage)
	}

	fmt.Println("PART2: SUM JOLTAGE", totalJoltage)
}

// PART1: SUM JOLTAGE: 17324
// PART2: SUM JOLTAGE: 171846613143331
