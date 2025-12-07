package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type FreshRange struct {
	Begin int
	End   int
}

type Ingredient struct {
	Id      int
	IsFresh bool
}

type Input struct {
	FreshRanges []*FreshRange
	Ingredients map[int]*Ingredient
}

func (self *FreshRange) IsIngredientIncluded(ingredientId int) bool {
	return ingredientId >= self.Begin && ingredientId <= self.End
}

func main() {
	fmt.Println("Day 5, Part 1")

	input := readInput("input.txt")

	for _, ingredient := range input.Ingredients {
		for _, fr := range input.FreshRanges {
			if fr.IsIngredientIncluded(ingredient.Id) {
				ingredient.IsFresh = true
				break
			}
		}
	}

	freshCount := 0

	for _, ingredient := range input.Ingredients {
		if ingredient.IsFresh {
			freshCount++
		}
	}

	fmt.Printf("FRESH INGREDIENT COUNT: %d\n", freshCount)
}

func readInput(s string) *Input {
	file, err := os.Open(s)
	if err != nil {
		log.Fatalf("Failed to open input file '%s': %v", s, err)
	}

	resultInput := Input{
		FreshRanges: []*FreshRange{},
		Ingredients: map[int]*Ingredient{},
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

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		id, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to read ingredient ID %s", line)
		}

		resultInput.Ingredients[id] = &Ingredient{
			Id: id,
		}
	}

	for _, fr := range resultInput.FreshRanges {
		fmt.Printf("Fresh Range: %d to %d\n", fr.Begin, fr.End)
	}

	for id := range resultInput.Ingredients {
		fmt.Printf("Ingredient: %d\n", id)
	}

	return &resultInput
}

// FRESH INGREDIENT COUNT: 635
