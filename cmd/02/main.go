package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/ghjm/advent_utils"
)

type idRange struct {
	start int64
	end   int64
}
type data struct {
	ranges []idRange
}

func hasSameParts(txt string, nParts int) bool {
	if len(txt)%nParts != 0 {
		return false
	}
	initial := txt[0 : len(txt)/nParts]
	subLen := len(initial)
	for i := 1; i < nParts; i++ {
		part := txt[subLen*i : subLen*(i+1)]
		if part != initial {
			return false
		}
	}
	return true
}

func isInvalidP1(number int64) bool {
	return hasSameParts(strconv.FormatInt(number, 10), 2)
}

func isInvalidP2(number int64) bool {
	txt := strconv.FormatInt(number, 10)
	for s := 2; s <= len(txt); s++ {
		if hasSameParts(txt, s) {
			return true
		}
	}
	return false
}

func main() {
	d := data{}
	err := utils.OpenAndReadLines("input2.txt", func(s string) error {
		rs := strings.Split(s, ",")
		for _, r := range rs {
			be := strings.Split(r, "-")
			if len(be) != 2 {
				return errors.New("invalid input")
			}
			nir := idRange{
				start: utils.MustAtoi64(be[0]),
				end:   utils.MustAtoi64(be[1]),
			}
			d.ranges = append(d.ranges, nir)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	var part1 int64
	var part2 int64
	for _, n := range d.ranges {
		for i := n.start; i <= n.end; i++ {
			if isInvalidP1(i) {
				part1 += i
			}
			if isInvalidP2(i) {
				part2 += i
			}
		}
	}
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
