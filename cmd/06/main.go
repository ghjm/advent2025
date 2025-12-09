package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghjm/advent_utils"
)

type data struct {
	lines  []string
	fields [][2]int
}

func (d *data) calcFields() {
	lastEnd := -1
	d.fields = nil
	for x, c := range d.lines[0] {
		if c == ' ' {
			good := true
			for y := 1; y < len(d.lines); y++ {
				if d.lines[y][x] != ' ' {
					good = false
					break
				}
			}
			if good {
				d.fields = append(d.fields, [2]int{lastEnd + 1, x})
				lastEnd = x
			}
		}
	}
	maxLen := 0
	for _, l := range d.lines {
		if len(l) > maxLen {
			maxLen = len(l)
		}
	}
	d.fields = append(d.fields, [2]int{lastEnd + 1, maxLen})
}

func (d *data) getField(y, x int) string {
	field := d.fields[x]
	line := d.lines[y]
	if field[1] > len(line) {
		field[1] = len(line)
	}
	return strings.TrimSpace(line[field[0]:field[1]])
}

func (d *data) doCalc(x int, numbers []int64) (int64, error) {
	var answer int64
	op := d.getField(len(d.lines)-1, x)
	switch op {
	case "+":
		for _, n := range numbers {
			answer += n
		}
	case "*":
		answer = 1
		for _, n := range numbers {
			answer *= n
		}
	default:
		return 0, fmt.Errorf("unknown op %q", op)
	}
	return answer, nil
}

func (d *data) part1() (int64, error) {
	var part1 int64
	for x := range d.fields {
		var numbers []int64
		for y := range len(d.lines) - 1 {
			numbers = append(numbers, utils.MustAtoi64(d.getField(y, x)))
		}
		answer, err := d.doCalc(x, numbers)
		if err != nil {
			return 0, err
		}
		part1 += answer
	}
	return part1, nil
}

func (d *data) part2() (int64, error) {
	var part2 int64
	for x := range d.fields {
		var numbers []int64
		for subX := d.fields[x][0]; subX <= d.fields[x][1]; subX++ {
			numStr := ""
			for y := range len(d.lines) - 1 {
				if subX >= len(d.lines[y]) {
					continue
				}
				ch := d.lines[y][subX]
				if ch != ' ' {
					numStr += string(ch)
				}
			}
			if numStr != "" {
				numbers = append(numbers, utils.MustAtoi64(numStr))
			}
		}
		answer, err := d.doCalc(x, numbers)
		if err != nil {
			return 0, err
		}
		part2 += answer
	}
	return part2, nil
}

func run() error {
	var d data
	err := utils.OpenAndReadLines("input6.txt", func(s string) error {
		d.lines = append(d.lines, s)
		return nil
	})
	if err != nil {
		return err
	}
	d.calcFields()
	part1, err := d.part1()
	if err != nil {
		return err
	}
	fmt.Printf("Part 1: %d\n", part1)
	part2, err := d.part2()
	if err != nil {
		return err
	}
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
