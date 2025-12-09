package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/ghjm/advent_utils"
)

type data struct {
	ranges      [][2]int64
	ingredients []int64
}

func (d *data) fresh(ingredient int64) bool {
	for _, r := range d.ranges {
		if ingredient >= r[0] && ingredient <= r[1] {
			return true
		}
	}
	return false
}

func (d *data) optimize() {
	var newRanges [][2]int64
	for _, r := range d.ranges {
		newRanges = append(newRanges, r)
	}
	i := 0
	for i < len(newRanges) {
		overlaps := false
		ir := newRanges[i]
		for j, jr := range newRanges {
			if j == i {
				continue
			}
			if (ir[0] >= jr[0] && ir[0] <= jr[1]) || (ir[1] >= jr[0] && ir[1] <= jr[1]) {
				overlaps = true
				jr[0] = min(ir[0], jr[0])
				jr[1] = max(ir[1], jr[1])
				newRanges[i] = jr
				newRanges = slices.Delete(newRanges, j, j+1)
				break
			}
		}
		if overlaps {
			i = 0
		} else {
			i = i + 1
		}
	}
	d.ranges = newRanges
}

func run() error {
	d := data{}
	in_ingredients := false
	err := utils.OpenAndReadLines("input5.txt", func(s string) error {
		if s == "" {
			in_ingredients = true
			return nil
		}
		if in_ingredients {
			d.ingredients = append(d.ingredients, utils.MustAtoi64(s))
		} else {
			values := strings.Split(s, "-")
			if len(values) != 2 {
				return fmt.Errorf("invalid input")
			}
			newRange := [2]int64{utils.MustAtoi64(values[0]), utils.MustAtoi64(values[1])}
			if newRange[0] > newRange[1] {
				return fmt.Errorf("invalid input")
			}
			d.ranges = append(d.ranges, newRange)
		}
		return nil
	})
	if err != nil {
		return err
	}
	d.optimize()
	part1 := 0
	for _, ing := range d.ingredients {
		if d.fresh(ing) {
			part1++
		}
	}
	var part2 int64
	for _, r := range d.ranges {
		part2 += r[1] - r[0] + 1
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
