package main

import (
	"fmt"
	"math"
	"os"

	"github.com/ghjm/advent_utils"
)

type data struct {
	lines []string
}

func maxDigit(s string) (int64, int) {
	var maxPos int = -1
	var maxValue int64 = -1
	for i, c := range s {
		v := utils.MustAtoi64(string(c))
		if v > maxValue {
			maxValue = v
			maxPos = i
		}
	}
	return maxValue, maxPos
}

func maxJoltage(s string, digits int) int64 {
	digits -= 1
	v, p := maxDigit(s[:len(s)-digits])
	if digits <= 0 {
		return v
	}
	return v*int64(math.Pow10(digits)) + maxJoltage(s[p+1:], digits)
}

func run() error {
	d := data{}
	err := utils.OpenAndReadLines("input3.txt", func(s string) error {
		d.lines = append(d.lines, s)
		return nil
	})
	if err != nil {
		return err
	}
	var part1, part2 int64
	for _, line := range d.lines {
		part1 += maxJoltage(line, 2)
		part2 += maxJoltage(line, 12)
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
