package main

// Tetromino structure
type Tetromino struct {
	Chr    string
	Width  int
	Height int
	Shape  []string
}

func StructTetros(tetros [][]string) *[]Tetromino {
	var terominoes []Tetromino
	fstChar := 65 // ASCII code for 'A'

	for _, tetro := range tetros {
		chr := string(rune(fstChar))
		width := 0
		height := 0
		shape := []string{}

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
		teromino := Tetromino{
			Chr:    chr,
			Height: height,
			Width:  width,
			Shape:  shape,
		}
		terominoes = append(terominoes, teromino)

		fstChar++
	}
	return &terominoes
}
