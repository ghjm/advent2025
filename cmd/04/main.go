package main

import (
	"fmt"
	"os"

	utils "github.com/ghjm/advent_utils"
	"github.com/ghjm/advent_utils/board"
)

type data struct {
	board *board.StdBoard
}

func (d *data) getMovable() []utils.StdPoint {
	var result []utils.StdPoint
	d.board.Iterate(func(p utils.Point[int], v rune) bool {
		count := 0
		for _, cp := range d.board.Diagonals(p, false) {
			if d.board.Get(cp) == '@' {
				count++
			}
		}
		if count < 4 {
			result = append(result, p)
		}
		return true
	})
	return result
}

func run() error {
	d := data{
		board: board.NewStdBoard(),
	}
	d.board.MustFromFile("input4.txt")
	m := d.getMovable()
	part1 := len(m)
	part2 := 0
	for len(m) > 0 {
		part2 += len(m)
		for _, p := range m {
			d.board.Clear(p)
		}
		m = d.getMovable()
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
