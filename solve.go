package main

// GameBoard structure
type GameBoard struct {
	Width  int
	Height int
	Cells  [][]string
}

func Play(tetrominoes *[]Tetromino) *GameBoard {
	// Starting with board size 2x2
	boardSize := 2
	gameBoard := newGameBoard(boardSize, boardSize)

	for !soltionFound(tetrominoes, gameBoard) {
		boardSize++
		gameBoard = newGameBoard(boardSize, boardSize)
	}
	return gameBoard
}

// Create an empty gameboard (wxh)
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

func soltionFound(tetrominoes *[]Tetromino, board *GameBoard) bool {
	return tryPlaceTetromino(*tetrominoes, board, 0)
}

func tryPlaceTetromino(tetrominoes []Tetromino, board *GameBoard, index int) bool {
	// BaseCase (all tetrominoes placed)
	if index == len(tetrominoes) {
		return true
	}

	for row := 0; row < board.Height; row++ {
		for col := 0; col < board.Width; col++ {
			if board.possible(tetrominoes[index], row, col) {
				board.place(tetrominoes[index], row, col)

				// Try placing the next Tetromino
				if tryPlaceTetromino(tetrominoes, board, index+1) {
					return true // Successfully placed all Tetrominoes
				}

				// Backtrack by removing the Tetromino from the board
				board.remove(tetrominoes[index], row, col)
			}
		}
	}

	// If all positions are tried and no solution is found, backtrack
	return false
}

// Check if possible to place the tetro in the given x, y position
func (board *GameBoard) possible(tetromino Tetromino, x, y int) bool {
	if x+tetromino.Width > board.Width || y+tetromino.Height > board.Height {
		return false
	}

	// Check if the cells are empty
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] == '#' && board.Cells[y+i][x+j] != "." {
				return false
			}
		}
	}
	return true
}

// Plcae the tetro in the given x, y position
func (board *GameBoard) place(tetromino Tetromino, x, y int) {
	for i := 0; i < len(tetromino.Shape); i++ {
		for j := 0; j < len(tetromino.Shape[i]); j++ {
			if tetromino.Shape[i][j] == '#' {
				board.Cells[y+i][x+j] = tetromino.Chr
			}
		}
	}
}

// Remove the tetro in the given x, y position
func (board *GameBoard) remove(tetromino Tetromino, x, y int) {
	for i := 0; i < tetromino.Height; i++ {
		for j := 0; j < tetromino.Width; j++ {
			if tetromino.Shape[i][j] == '#' {
				board.Cells[y+i][x+j] = "."
			}
		}
	}
}
