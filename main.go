package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var tetrominoes [][]string
	var input []string

	// Record the start time
	startTime := time.Now()

	// Read the args
	args := os.Args
	if len(args) != 2 {
		log.Fatal("ERROR: Check Arguments length")
	}
	filename := args[1]

	// Open the file
	inputFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("ERROR: While opening the file\n", err)
	}

	// Iterate over each line in the file
	scanner := bufio.NewScanner(inputFile)
	lineIndex := 0
	for scanner.Scan() {
		lineIndex++
		line := scanner.Bytes()
		// Check formatting
		if (len(line) != 4 && lineIndex != 5) || (lineIndex == 5 && len(line) != 0) {
			log.Fatal("ERROR: Formatting error")
		}

		// Append the data to the tetrominoes & reset lineIndex
		input = append(input, string(line))
		if lineIndex == 5 {
			tetrominoes = append(tetrominoes, input[:4])
			lineIndex = 0
			input = nil
		}
	}
	if len(input) != 0 { // Check if last newline exist
		log.Fatal("ERROR: Formatting error")
	}

	// Check validation & start solving
	if Valid(tetrominoes) {
		// Create a structure for eac tetromino [width, height, shape, char]
		tetrominoes := *StructTetros(tetrominoes)

		// Begin solving
		solution := *Play(&tetrominoes)

		// Print the solution
		for row := range solution.Cells {
			fmt.Println(strings.Join(solution.Cells[row], ""))
		}

		// Calculate the elapsed time
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		fmt.Printf("Program took %s to finish.\n", elapsedTime)

	} else {
		log.Fatal("ERROR: Tetrominoes not valid")
	}
}
