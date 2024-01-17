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
	var tetros [][]string
	var data []string

	// Record the start time
	startTime := time.Now()

	// read the args
	args := os.Args
	if len(args) != 2 {
		log.Fatal("Args length error")
	}
	filename := args[1]

	// open the file
	dataFile, err := os.Open(filename)
	if err != nil {
		log.Fatal("error while opening the file\n", err)
	}

	// iterate over each line in the file
	scanner := bufio.NewScanner(dataFile)
	lineIndex := 0
	for scanner.Scan() {
		lineIndex++
		line := scanner.Bytes()
		// check formatting
		if len(line) != 4 && lineIndex != 5 || lineIndex == 5 && len(line) != 0 {
			log.Fatal("Formatting error")
		}
		// append the data to the tetros
		// reset line index
		data = append(data, string(line))
		if lineIndex == 5 {
			tetros = append(tetros, data[:4])
			lineIndex = 0
			data = nil
		}
	}

	// fmt.Println(tetros)

	// Check validation
	if Valid(tetros) {
		fmt.Println("Valid tetro, lets play")

		// create a structure for the input tetros width, height, character
		tetrominoes := *StructTetros(tetros)

		// start to solve it
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
		log.Fatal("Tetro not valid")
	}
}
