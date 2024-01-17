package main

// Tetromino structure
type Tetromino struct {
	Chr    string
	Width  int
	Height int
	Shape  []string
}

func StructTetros(tetros [][]string) *[]Tetromino {
	// tetrominoes contains all the tetromino struct
	var tetrominoes []Tetromino

	fstChar := 65 // ASCII code for 'A'

	for _, tetro := range tetros {
		chr := string(rune(fstChar))
		width := 0
		height := 0
		shape := []string{}

		// Columns to not consider in the shape (which has only .'s)
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

		// Assign the height of the tetromino
		for _, row := range tetro {
			for _, col := range row {
				if col == '#' {
					shape = append(shape, row)
					height++
					break
				}
			}
		}

		// Remove empty columns from the tetromino
		for j := range shape {
			str := ""
			for indx, chr := range shape[j] {
				if !contains(colToCut, indx) {
					str += string(chr)
				}
			}
			shape[j] = str
		}

		// Assign the width of the tetromino after removing empty columns
		width = len(shape[0])

		// Create a new Tetromino struct and add it to the tetrominoes list
		tetromino := Tetromino{
			Chr:    chr,
			Height: height,
			Width:  width,
			Shape:  shape,
		}
		tetrominoes = append(tetrominoes, tetromino)

		fstChar++

		// Handle wrap-around characters
		if fstChar > 90 { // ASCII code for 'Z'
			fstChar = 65 // ASCII code for 'A'
		}
	}
	return &tetrominoes
}

// Check if a value is present in a slice
func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
