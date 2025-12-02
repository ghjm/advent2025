package main

import (
	"fmt"
	"os"

	utils "github.com/ghjm/advent_utils"
)

type data struct {
	moves []string
}

func main() {
	d := data{}
	err := utils.OpenAndReadLines("input1.txt", func(s string) error {
		d.moves = append(d.moves, s)
		return nil
	})
	pos := 50
	part1 := 0
	part2 := 0
	for _, m := range d.moves {
		var direction int
		if m[0] == 'L' {
			direction = -1
		} else {
			direction = 1
		}
		distance := utils.MustAtoi(m[1:])
		cur := pos
		for i := 0; i < distance; i++ {
			cur = utils.Mod(cur+direction, 100)
			if cur == 0 {
				part2++
			}
		}
		pos = utils.Mod(pos+distance*direction, 100)
		if pos == 0 {
			part1++
		}
		if cur != pos {
			err = fmt.Errorf("mismatch between pos and cur")
			break
		}
	}
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
