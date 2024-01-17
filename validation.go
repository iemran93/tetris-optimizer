package main

import (
	"strings"
)

func Valid(tetrominoes [][]string) bool {
	// Check validation:
	// - Characters only . or #
	// - Valid tetromino shape
	if characterValid(tetrominoes) && tetroValid(tetrominoes) && len(tetrominoes) != 0 {
		return true
	}
	return false
}

func characterValid(tetrominoes [][]string) bool {
	// Check validation over characters
	for _, tetromino := range tetrominoes {
		totalHash := 0
		for _, chr := range strings.Join(tetromino, "") {
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

func tetroValid(tetrominoes [][]string) bool {
	// Check if its a valid tetromino shape
	// by counting the total numbers of sides touching
	for _, tetromino := range tetrominoes {
		totalSideTouching := 0
		for rx, row := range tetromino {
			for cx, col := range row {
				if col == '#' {
					sideTouching := 0
					if hasNeighbor(tetromino, rx-1, cx) { // check above
						sideTouching++
					}
					if hasNeighbor(tetromino, rx+1, cx) { // check below
						sideTouching++
					}
					if hasNeighbor(tetromino, rx, cx+1) { // check right
						sideTouching++
					}
					if hasNeighbor(tetromino, rx, cx-1) { // check left
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

func hasNeighbor(tetromino []string, rx, cx int) bool {
	// Check if the position is within bounds and contains '#'
	return rx >= 0 && rx < len(tetromino) && cx >= 0 && cx < len(tetromino[rx]) && tetromino[rx][cx] == '#'
}
