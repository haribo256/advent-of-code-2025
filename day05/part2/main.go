package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type FreshRange struct {
	Begin int
	End   int
}

type Input struct {
	FreshRanges []*FreshRange
}

func (self *FreshRange) Length() int {
	return self.End - self.Begin + 1
}

func SubtractRanges(subtraction *FreshRange, sources []*FreshRange) []*FreshRange {
	var result []*FreshRange

	for _, source := range sources {
		iterRange := source.Subtract(subtraction)
		result = append(result, iterRange...)
	}

	return result
}

func (self *FreshRange) Subtract(other *FreshRange) []*FreshRange {
	if self.End < other.Begin || other.End < self.Begin {
		// No intersection
		return []*FreshRange{self}
	}

	if other.Begin <= self.Begin && other.End >= self.End {
		// Complete overlap
		return []*FreshRange{}
	}

	if other.Begin <= self.Begin && other.End < self.End {
		// Cut from start
		return []*FreshRange{
			{
				Begin: other.End + 1,
				End:   self.End,
			},
		}
	}

	if other.Begin > self.Begin && other.End >= self.End {
		// Cut from end
		return []*FreshRange{
			{
				Begin: self.Begin,
				End:   other.Begin - 1,
			},
		}
	}

	if other.Begin > self.Begin && other.End < self.End {
		// Cut from middle, splitting into two
		return []*FreshRange{
			{
				Begin: self.Begin,
				End:   other.Begin - 1,
			},
			{
				Begin: other.End + 1,
				End:   self.End,
			},
		}
	}

	panic("unreachable")
}

func (self *FreshRange) IsOverlapping(other *FreshRange) bool {
	return !(self.End < other.Begin || other.End < self.Begin)
}

func main() {
	fmt.Println("Day 5, Part 2")

	input := readInput("input.txt")

	newRanges := []*FreshRange{}

	for _, sourceRange := range input.FreshRanges {
		addingRanges := []*FreshRange{sourceRange}

		fmt.Printf("source:   %020d-%020d\n", sourceRange.Begin, sourceRange.End)

		for _, existingRange := range newRanges {
			if addingRanges == nil {
				break
			}

			fmt.Printf("  exists: %020d-%020d\n", existingRange.Begin, existingRange.End)

			addingRanges = SubtractRanges(existingRange, addingRanges)
		}

		for _, addingRange := range addingRanges {
			fmt.Printf("     res: %020d-%020d\n", addingRange.Begin, addingRange.End)

		}

		newRanges = append(newRanges, addingRanges...)
	}

	slices.SortStableFunc(newRanges, func(a *FreshRange, b *FreshRange) int {
		if a.Begin < b.Begin {
			return -1
		} else if a.Begin > b.Begin {
			return 1
		} else {
			return 0
		}
	})

	fmt.Println()

	totalCount := 0

	for _, rng := range newRanges {
		totalCount += rng.Length()
		fmt.Printf("in-order: %020d-%020d, length=%d, count=%d\n", rng.Begin, rng.End, rng.Length(), totalCount)
	}

	fmt.Printf("FRESH INGREDIENT COUNT: %d", totalCount)
}

func readInput(s string) *Input {
	file, err := os.Open(s)
	if err != nil {
		log.Fatalf("Failed to open input file '%s': %v", s, err)
	}

	resultInput := Input{
		FreshRanges: []*FreshRange{},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		parts := strings.SplitN(line, "-", 2)

		begin, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Failed to read range %s", line)
		}

		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Failed to read range %s", line)
		}

		resultInput.FreshRanges = append(resultInput.FreshRanges, &FreshRange{
			Begin: begin,
			End:   end,
		})
	}

	return &resultInput
}

// FRESH INGREDIENT COUNT: 369761800782619
