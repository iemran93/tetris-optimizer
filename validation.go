package main

import (
	"strings"
)

func Valid(tetros [][]string) bool {
	if characterValid(tetros) && tetroValid(tetros) {
		return true
	}
	return false
}

func characterValid(tetros [][]string) bool {
	// check validation over character
	for _, tetro := range tetros {
		totalHash := 0
		for _, chr := range strings.Join(tetro, "") {
			if chr != '.' && chr != '#' {
				return false
			}
			if chr == '#' {
				totalHash++
			}
		}
		if totalHash != 4 {
			return false
		}
	}
	return true
}

func tetroValid(tetros [][]string) bool {
	// check if its a valid tetro
	// count the total number of sides touching
	for _, tetro := range tetros {
		totalSideTouching := 0
		for rx, row := range tetro {
			for cx, col := range row {
				if col == '#' {
					sideTouching := 0
					if hasNeighbor(tetro, rx-1, cx) { // check above
						sideTouching++
					}
					if hasNeighbor(tetro, rx+1, cx) { // check below
						sideTouching++
					}
					if hasNeighbor(tetro, rx, cx+1) { // check right
						sideTouching++
					}
					if hasNeighbor(tetro, rx, cx-1) { // check left
						sideTouching++
					}
					if sideTouching == 0 || sideTouching > 3 {
						return false
					}
					totalSideTouching += sideTouching
				}
			}
		}
		if totalSideTouching != 6 && totalSideTouching != 8 {
			return false
		}
	}
	return true
}

func hasNeighbor(tetro []string, rx, cx int) bool {
	// Check if the position is within bounds and contains '#'
	return rx >= 0 && rx < len(tetro) && cx >= 0 && cx < len(tetro[rx]) && tetro[rx][cx] == '#'
}
