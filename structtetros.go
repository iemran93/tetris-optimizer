package main

import "fmt"

// Tetromino structure
type Tetromino struct {
	Chr    string
	Width  int
	Height int
	Shape  []string
}

func StructTetros(tetros [][]string) *[]Tetromino {
	var tetrominoes []Tetromino
	fstChar := 65 // ASCII code for 'A'

	for _, tetro := range tetros {
		chr := string(rune(fstChar))
		width := 0
		height := 0
		shape := []string{}

		colToCut := []int{}
		for col := 0; col < 4; col++ {
			found := false
			for row := 0; row < 4; row++ {
				if tetro[row][col] == '#' {
					found = true
				}
				if row == 3 && !found {
					colToCut = append(colToCut, col)
				}
			}
		}

		for _, row := range tetro {
			// reset width for each new row
			rowWidth := 0
			for _, col := range row {
				if col == '#' {
					rowWidth++
				}
			}
			if rowWidth > 0 {
				// if row containing at least one # count it
				height++
				if rowWidth > width { // take the most number of width
					width = rowWidth
				}
				// add it to the shape list
				shape = append(shape, row)
			}
		}

		// Create a new Tetromino struct and add it to the list
		tetromino := Tetromino{
			Chr:    chr,
			Height: height,
			Width:  width,
			Shape:  shape,
		}

		// Remove empty columns
		for j := range tetromino.Shape {
			str := ""
			for indx, chr := range tetromino.Shape[j] {
				if !contains(colToCut, indx) {
					str += string(chr)
				}
			}
			tetromino.Shape[j] = str
		}

		fmt.Println(tetromino.Shape)
		tetrominoes = append(tetrominoes, tetromino)

		fstChar++
	}
	return &tetrominoes
}

// Helper function to check if a value is present in a slice
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
