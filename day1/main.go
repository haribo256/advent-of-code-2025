package main

// 4614
// Not 4890 too low
// Not 5175
// Not 5454
// Not 5625 (must be close)
// Not 5766
// Not 6478
// Not 7371 too high
// 8503

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	fmt.Println("File opened")

	dialValue := int(50)
	rotationZeroCount := int(0)
	tickZeroCount := int(0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		fmt.Print(line)
		fmt.Print(": ")

		isLeft, isRight, distance := DecodeLine(line)

		if isLeft {
			fmt.Printf("%d - %d", dialValue, distance)
		} else if isRight {
			fmt.Printf("%d + %d", dialValue, distance)
		}

		newDialValue := dialValue

		if isLeft {
			newDialValue -= distance
		} else {
			newDialValue += distance
		}

		tickZeroTimes := int(0)
		// if dialValue < 0 && newDialValue > 0 {
		// 	tickZeroTimes = int(math.Ceil(math.Abs(float64(newDialValue-dialValue) / 100)))
		// } else if dialValue > 0 && newDialValue < 0 {
		// 	tickZeroTimes = int(math.Ceil(math.Abs(float64(dialValue-newDialValue) / 100)))
		// }
		// tickZeroCount += tickZeroTimes
		min := math.Min(float64(dialValue), float64(newDialValue))
		max := math.Max(float64(dialValue), float64(newDialValue))
		// // // if newDialValue%100 != 0 {
		// // tickZeroTimes += int(math.Floor(max/100) - math.Floor(min/100))

		// if min < 0 && max > 0 {
		// 	tickZeroTimes = int(math.Ceil(math.Abs(max-min) / 100))
		// } else {
		// 	tickZeroTimes = int(math.Floor(max/100) - math.Floor(min/100))
		// }

		tickZeroTimes += int(math.Floor(float64(distance) / float64(100)))

		// L814: 14 - 814 = 0
		// }

		if (newDialValue % 100) == 0 {
			rotationZeroCount++

			// if distance%100 != 0 {
			// 	tickZeroTimes++
			// }
			// tickZeroTimes++
		}

		if ((min < 0 && max > 0) || (newDialValue%100) == 0) && tickZeroTimes == 0 {
			tickZeroTimes++
		}

		tickZeroCount += tickZeroTimes

		fmt.Printf(" = %d", newDialValue%100)
		fmt.Printf(", %d to %d", dialValue, newDialValue)
		// fmt.Printf(", %d to %d", int(math.Abs(float64(dialValue%100))), int(math.Abs(float64(newDialValue%100))))
		fmt.Printf(", ticking %d zeros", tickZeroTimes)
		fmt.Printf(", totalling %d rotational zeros", rotationZeroCount)
		fmt.Printf(", totalling %d tick zeros", tickZeroCount)
		fmt.Println()

		dialValue = newDialValue % 100
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// func runTests() {
// 	if MoveDial(50, true, false, 68) != 82 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(82, true, false, 30) != 52 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(52, false, true, 48) != 0 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(0, true, false, 5) != 95 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(95, false, true, 60) != 55 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(55, true, false, 55) != 0 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(0, true, false, 1) != 99 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(99, true, false, 99) != 0 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(0, false, true, 14) != 14 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(14, true, false, 82) != 32 {
// 		log.Fatal("test failed")
// 	}

// 	if MoveDial(99, true, false, 100) != 99 {
// 		log.Fatal("test failed")
// 	}
// 	if MoveDial(99, false, true, 100) != 99 {
// 		log.Fatal("test failed")
// 	}

// 	if MoveDial(99, false, true, 552) != 51 {
// 		log.Fatal("test failed")
// 	}

// 	if MoveDial(0, true, false, 245) != 55 {
// 		log.Fatal("test failed")
// 	}

// 	if MoveDial(0, false, true, 245) != 45 {
// 		log.Fatal("test failed")
// 	}
// }

// func MoveDial(currentDial int, isLeft bool, isRight bool, amount int) int {
// 	newValue := 0

// 	if isRight {
// 		newValue = currentDial + amount
// 	} else if isLeft {
// 		newValue = currentDial - amount
// 	} else {
// 		log.Fatal(errors.New("number is neither left or right"))
// 	}

// 	if newValue < 0 {
// 		newValue = int(math.Abs(float64(newValue)) - 100)
// 	}

// 	remainder := math.Mod(float64(newValue), 100)

// 	return int(math.Abs(remainder))
// }

// func MoveDial(currentDial int, isLeft bool, isRight bool, distance int) int {
// 	newValue := 0

// 	if isRight {
// 		newValue = currentDial + distance
// 	} else if isLeft {
// 		newValue = currentDial - distance
// 	} else {
// 		log.Fatal(errors.New("number is neither left or right"))
// 	}

// 	if newValue < 0 {
// 		return int(100 - math.Abs(math.Mod(float64(newValue), 100)))
// 		// newValue = int(math.Abs(float64(newValue)) - 100)
// 	} else {
// 		remainder := math.Mod(float64(newValue), 100)
// 		return int(math.Abs(remainder))
// 	}
// }

func DecodeLine(input string) (isLeft bool, isRight bool, amount int) {
	amountString, isLeft := strings.CutPrefix(input, "L")
	if !isLeft {
		amountString, isRight = strings.CutPrefix(input, "R")
	}

	if !isLeft && !isRight {
		log.Fatal(errors.New("failed to decode"))
	}

	amount, err := strconv.Atoi(amountString)
	if err != nil {
		log.Fatal("Failed")
	}

	return
}
