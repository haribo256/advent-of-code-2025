package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction = uint8

const DIRECTION_LEFT = Direction(0)
const DIRECTION_RIGHT = Direction(1)

type DialTranslation struct {
	Direction      Direction
	Distance       int
	SignedDistance int
}

func (self *DialTranslation) IsLeft() bool {
	return self.Direction == DIRECTION_LEFT
}

func (self *DialTranslation) IsRight() bool {
	return self.Direction == DIRECTION_RIGHT
}

func (self *DialTranslation) String() string {
	if self.IsLeft() {
		return fmt.Sprintf("LEFT:%d", self.Distance)
	} else {
		return fmt.Sprintf("RIGHT:%d", self.Distance)
	}
}

func ParseDialTranslation(input string) (*DialTranslation, error) {
	var direction Direction
	var signedDistance int

	trimmedInput := strings.TrimSpace(input)

	distance, err := strconv.Atoi(input[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to convert %s into a numerical distance", input)
	}

	prefix := string(trimmedInput[0])
	switch prefix {
	case "L":
		direction = DIRECTION_LEFT
		signedDistance = -distance
	case "R":
		direction = DIRECTION_RIGHT
		signedDistance = distance
	default:
		return nil, fmt.Errorf("Failed to convert %s into a dial-translation", input)
	}

	return &DialTranslation{
		Direction:      direction,
		Distance:       distance,
		SignedDistance: signedDistance,
	}, nil
}

type Dial struct {
	AccumulatedPosition int
	Position            int
	ZeroLandings        int
	ZeroRotations       int
}

func NewDial(startingPosition int) Dial {
	return Dial{
		AccumulatedPosition: startingPosition,
		Position:            50,
		ZeroLandings:        0,
		ZeroRotations:       0,
	}
}

func (self *Dial) String() string {
	return fmt.Sprintf("%d", self.AccumulatedPosition)
}

const FULL_ROTATION = 100

func (self *Dial) Translate(translation *DialTranslation) int {
	oldPosition := self.Position

	self.AccumulatedPosition += translation.SignedDistance
	self.Position = self.AccumulatedPosition % FULL_ROTATION
	if self.Position < 0 {
		self.Position += FULL_ROTATION
	}

	if self.Position == 0 {
		self.ZeroLandings++
	}

	signedPositionAfterDistance := oldPosition + translation.SignedDistance
	zeroRotations := int(math.Abs(float64(signedPositionAfterDistance / FULL_ROTATION)))

	if oldPosition != 0 && signedPositionAfterDistance <= 0 {
		zeroRotations++
	}

	self.ZeroRotations += zeroRotations

	return zeroRotations
}

func main() {
	inputFileName := "input.txt"
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatalf("Could not open file %s", inputFileName)
	}

	dial := NewDial(50)
	fmt.Printf("dial=%s\n", &dial)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		dialTranslation, err := ParseDialTranslation(line)
		if err != nil {
			log.Fatalf("failed to parse dial-translation %e", err)
		}

		zeroRotations := dial.Translate(dialTranslation)

		fmt.Printf("translation=%d, dial-position=%d, total-landings=%d, zero-rotations=%d, total-rotations=%d\n", dialTranslation.SignedDistance, dial.Position, dial.ZeroLandings, zeroRotations, dial.ZeroRotations)
	}
}

// Part-1 Answer: 1059
// Part-2 Answer: 6305
