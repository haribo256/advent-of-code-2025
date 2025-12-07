package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type GridCell struct {
	HasPaperRoll bool
	IsAvailable  bool
}

type Grid struct {
	Width  int
	Height int
	Rows   [][]*GridCell
}

type GridReference struct {
	X int
	Y int
}

func NewGridRef(x int, y int) GridReference {
	return GridReference{
		X: x,
		Y: y,
	}
}

func (self *GridReference) Translate(x int, y int) GridReference {
	return NewGridRef(self.X+x, self.Y+y)
}

func (self *Grid) PrintRow(y int) {
	for x := 0; x < self.Width; x++ {
		cell := self.Rows[y][x]
		char := "."
		if cell.HasPaperRoll {
			char = "@"
		}
		if cell.IsAvailable {
			char = "x"
		}
		fmt.Print(char)
	}
	fmt.Println()
}

func (self *Grid) PrintGrid() {
	for y := 0; y < self.Height; y++ {
		self.PrintRow(y)
	}
}

func (self *Grid) HasPaperRollAt(ref GridReference) bool {
	if ref.X < 0 || ref.Y < 0 {
		return false
	}
	if ref.X >= self.Width || ref.Y >= self.Height {
		return false
	}

	return self.Rows[ref.Y][ref.X].HasPaperRoll
}

func (self *Grid) EvaluateAvailability(ref GridReference) bool {
	if !self.HasPaperRollAt(ref) {
		self.Rows[ref.Y][ref.X].IsAvailable = false
		return false
	}

	var gridReferences = []GridReference{
		ref.Translate(-1, -1),
		ref.Translate(0, -1),
		ref.Translate(1, -1),
		ref.Translate(-1, 0),
		ref.Translate(1, 0),
		ref.Translate(-1, 1),
		ref.Translate(0, 1),
		ref.Translate(1, 1),
	}

	count := 0
	for _, gridReference := range gridReferences {
		if self.HasPaperRollAt(gridReference) {
			count++
		}
	}

	isAvailable := count < 4
	self.Rows[ref.Y][ref.X].IsAvailable = isAvailable

	return isAvailable
}

func main() {
	fmt.Println("Day 4, Part 1")

	grid := readGridInput("input.txt")

	availableCount := 0
	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			if grid.EvaluateAvailability(NewGridRef(x, y)) {
				availableCount++
			}
		}
	}

	grid.PrintGrid()

	fmt.Printf("SUM AVAILABLE PAPER-ROLLS %d", availableCount)
}

func readGridInput(inputFileName string) *Grid {
	file, err := os.Open(inputFileName)
	if err != nil {
		panic("File cannot be found or opened")
	}

	width := 0
	rowList := list.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		width = max(width, len(line))

		cells := make([]*GridCell, width)
		for x := 0; x < len(line); x++ {
			char := line[x : x+1]
			cells[x] = &GridCell{
				HasPaperRoll: char == "@",
			}
		}

		rowList.PushBack(cells)
	}

	height := rowList.Len()
	rows := make([][]*GridCell, height)
	row := rowList.Front()
	for y := 0; row != nil; y++ {
		rows[y] = row.Value.([]*GridCell)
		row = row.Next()
	}

	return &Grid{
		Width:  width,
		Height: height,
		Rows:   rows,
	}
}

// SUM AVAILABLE PAPER-ROLLS 1516
