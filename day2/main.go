package main

import (
	"fmt"
	"iter"
	"log"
	"strconv"
	"strings"
)

const INPUT_TEXT string = "269351-363914,180-254,79-106,771-1061,4780775-4976839,7568-10237,33329-46781,127083410-127183480,19624-26384,9393862801-9393974421,2144-3002,922397-1093053,39-55,2173488366-2173540399,879765-909760,85099621-85259580,2-16,796214-878478,163241-234234,93853262-94049189,416472-519164,77197-98043,17-27,88534636-88694588,57-76,193139610-193243344,53458904-53583295,4674629752-4674660925,4423378-4482184,570401-735018,280-392,4545446473-4545461510,462-664,5092-7032,26156828-26366132,10296-12941,61640-74898,7171671518-7171766360,3433355031-3433496616"

type InputRange struct {
	First int
	Last  int
}

func readInputRanges(inputText string) iter.Seq[InputRange] {
	inputRanges := strings.SplitSeq(inputText, ",")

	return func(yield func(X InputRange) bool) {
		for inputRange := range inputRanges {
			inputRangeParts := strings.SplitN(inputRange, "-", 2)
			inputRangeFirst, err := strconv.Atoi(inputRangeParts[0])
			if err != nil {
				log.Fatalf("input-range did not have min: %s", inputRange)
			}

			inputRangeLast, err := strconv.Atoi(inputRangeParts[1])
			if err != nil {
				log.Fatalf("input-range did not have max: %s", inputRange)
			}

			yield(InputRange{
				First: inputRangeFirst,
				Last:  inputRangeLast,
			})
		}
	}
}

func main() {
	inputRanges := readInputRanges(INPUT_TEXT)
	part1(inputRanges)

	inputRanges = readInputRanges(INPUT_TEXT)
	part2(inputRanges)
}

func part1(inputRanges iter.Seq[InputRange]) {
	invalidIdSum := 0

	for inputRange := range inputRanges {
		fmt.Println("PART1: RANGE", inputRange.First, "to", inputRange.Last)

		for id := inputRange.First; id <= inputRange.Last; id++ {
			var idString = strconv.Itoa(id)
			idStringLen := len(idString)

			if idStringLen%2 != 0 {
				continue
			}

			idHalfStringLen := idStringLen / 2
			idFirstHalf := idString[0:idHalfStringLen]
			idLastHalf := idString[idHalfStringLen:]

			if idFirstHalf == idLastHalf {
				invalidIdSum += id
				fmt.Println("  PART1: INVALID", idString, "SUM", invalidIdSum)
			}
		}
	}

	fmt.Println("PART1: FINAL SUM", invalidIdSum)
}

func part2(inputRanges iter.Seq[InputRange]) {
	invalidIdSum := 0

	for inputRange := range inputRanges {
		fmt.Println("PART2: RANGE", inputRange.First, "TO", inputRange.Last)

	EACH_ID:
		for id := inputRange.First; id <= inputRange.Last; id++ {
			idString := strconv.Itoa(id)
			idStringLen := len(idString)

		EACH_LENGTH:
			for repeatLen := 1; repeatLen <= idStringLen/2; repeatLen++ {
				if idStringLen%repeatLen != 0 {
					continue
				}

				findString := idString[0:repeatLen]
				expectedOccurenceCount := (idStringLen / repeatLen) - 1
				occurenceCount := 0

				for occurenceIndex := repeatLen; occurenceIndex <= idStringLen-repeatLen; occurenceIndex += repeatLen {
					occurenceString := idString[occurenceIndex : occurenceIndex+repeatLen]
					if occurenceString != findString {
						continue EACH_LENGTH
					}

					occurenceCount++
				}

				if occurenceCount == expectedOccurenceCount {
					invalidIdSum += id
					fmt.Println("  PART2: INVALID", idString)
					continue EACH_ID
				}
			}
		}
	}

	fmt.Println("PART2: FINAL SUM", invalidIdSum)
}

// PART1: FINAL SUM: 32976912643
// PART2: FINAL SUM: 54446379122
