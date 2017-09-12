package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Reads the input and returns an array of BillBoards.
//
// The first line of the input is the BillBoard count,
// the following lines contain width, height and the text of the BillBoard.
func readInput(scanner *bufio.Scanner) ([]*BillBoard, error) {
	// Eof after the first line
	if !scanner.Scan() {
		return nil, errors.New("bad input file")
	}
	// Convert the first line to int
	lines, err := strconv.Atoi(scanner.Text())

	if err != nil {
		return nil, fmt.Errorf("impossible to parse line count: %s", err.Error())
	}

	if lines <= 0 || lines > 20 {
		return nil, fmt.Errorf("bad line count: %d", lines)
	}

	billBoards := make([]*BillBoard, lines)
	i := 0

	// Loop on every line
	for scanner.Scan() {
		text := scanner.Text()
		// Accept a blank line as last
		if text == "" && i == lines {
			break
		}

		// Return an error if the input has more lines than the lines parameter
		if i >= lines {
			return nil, errors.New("input has too many lines")
		}

		// Create a new BillBoard
		billBoard, err := NewBillBoard(text)
		if err != nil {
			return nil, fmt.Errorf("bad input, line: %d, error: %s", i+1, err.Error())
		}

		billBoards[i] = billBoard
		i += 1
	}

	// IO Error
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Input has less lines than the lines parameter
	if i < lines {
		return nil, fmt.Errorf("expected %d lines but found %d", lines, i)
	}

	return billBoards, nil
}

func main() {
	// Parse the input from Stdin
	billBoards, err := readInput(bufio.NewScanner(os.Stdin))
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	// Loop over the BillBoards and get the font size
	for i, billboard := range billBoards {
		fmt.Printf("Case #%d: %d\n", i, billboard.GetFontSize())
	}
}
