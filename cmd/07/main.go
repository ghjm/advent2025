package main

import (
	"fmt"
	"os"

	utils "github.com/ghjm/advent_utils"
	"github.com/ghjm/advent_utils/board"
)

type data struct {
	b *board.StdBoard
}

func (d *data) calcs() (int64, int64, error) {
	var part1 int64
	b := board.NewRunePlusBoard[int, int64](board.WithBounds[int, board.RunePlusData[int64]](d.b.Bounds()))
	d.b.Iterate(func(p utils.Point[int], v rune) bool {
		b.Set(p, board.RunePlusData[int64]{Value: v})
		return true
	})
	for y := b.Bounds().P1.Y + 1; y <= b.Bounds().P2.Y; y++ {
		for x := b.Bounds().P1.X; x <= b.Bounds().P2.X; x++ {
			above := b.Get(utils.StdPoint{X: x, Y: y - 1})
			if above.Value == 'S' || above.Value == '|' {
				if above.Value == 'S' {
					above.Extra = 1
				}
				cv := b.Get(utils.StdPoint{X: x, Y: y})
				if cv.Value == '^' {
					part1++
					for _, dx := range []int{-1, +1} {
						p := utils.StdPoint{X: x + dx, Y: y}
						if b.Contains(p) {
							v := b.Get(p)
							b.Set(p, board.RunePlusData[int64]{Value: '|', Extra: v.Extra + above.Extra})
						}
					}
				} else {
					p := utils.StdPoint{X: x, Y: y}
					v := b.Get(p)
					b.Set(p, board.RunePlusData[int64]{Value: '|', Extra: v.Extra + above.Extra})
				}
			}
		}
	}
	var part2 int64
	for x := b.Bounds().P1.X; x <= b.Bounds().P2.X; x++ {
		v := b.Get(utils.StdPoint{X: x, Y: b.Bounds().P2.Y})
		part2 += v.Extra
	}
	return part1, part2, nil
}

func run() error {
	d := data{
		b: board.NewStdBoard(),
	}
	err := d.b.FromFile("input7.txt")
	if err != nil {
		return err
	}
	part1, part2, err := d.calcs()
	if err != nil {
		return err
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
