package main

// GameBoard structure
type GameBoard struct {
	Width  int
	Height int
	Cells  [][]string
}

func Solve(tetrominoes *[]Tetromino) *GameBoard {
	// initialize a gameboard (starting with 2x2)
	gameBoadrd := newGameBoard(2, 2)

	// is possible to place tetromino on this board
	for x := 0; x < gameBoadrd.Width; x++ {
		for y := 0; y < gameBoadrd.Height; y++ {
			for _, tetromino := range *tetrominoes {
				if gameBoadrd.possible(tetromino, x, y) {
					// if possible place it
					gameBoadrd.place(tetromino, x, y)
					solve()
				} else {
				}
			}
		}
	}
}

func newGameBoard(width, height int) *GameBoard {
	board := make([][]string, height)
	for i := range board {
		board[i] = make([]string, width)
		for j := range board[i] {
			board[i][j] = "."
		}
	}
	return &GameBoard{
		Width:  width,
		Height: height,
		Cells:  board,
	}
}

func (board *GameBoard) possible(tetromino Tetromino, x, y int) bool {
	if x+tetromino.Width > board.Width || y+tetromino.Height > board.Height {
		return false
	}
	return true
}

func (board *GameBoard) place(tetromino Tetromino, x, y int) {
	// iterate over the tetromino to place it in the board
	// if its # place it in the x, y position
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] == '#' {
				board.Cells[y+i][x+j] = tetromino.Chr
			}
		}
	}
}
