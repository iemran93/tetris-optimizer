package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var tetros [][]string
	var data []string

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

	fmt.Println(tetros)
	// check validation
	if Valid(tetros) {
		fmt.Println("valid tetro. lets play")
	} else {
		fmt.Println("not valid tetro")
	}
}
