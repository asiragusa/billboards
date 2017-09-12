package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Regexp that validates the billboard text. It Accepts strings with only lower-case letters a-z, upper-case
// letters A-Z, digits 0-9 and the space character. The text must not begin or end by space and have a
// length between 1 and 1000 chars
var validInputRE = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9 ]{0,998}[a-zA-Z0-9])?$`)

type BillBoard struct {
	// Width of the BillBoard
	width int

	// Height of the BillBoard
	height int

	// Text of the BillBoard
	text string
}

// Creates a new BillBoard from an input string.
// The string has the form "width height text" where width and height are integers and text is a string.
//
// width and height must be between 1 and 1000 chars long.
// text must match "validInputRe" and cannot contain two adjacent spaces.
// It returns a new BillBoard if successful or an error if line is not valid
func NewBillBoard(line string) (*BillBoard, error) {
	parts := strings.SplitN(line, " ", 3)
	if len(parts) != 3 {
		return nil, fmt.Errorf("missing required fields, line: %s", line)
	}
	width, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	if width <= 0 || width > 1000 {
		return nil, errors.New("width out of range")
	}

	height, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	if height <= 0 || height > 1000 {
		return nil, errors.New("height out of range")
	}

	text := parts[2]
	if !validInputRE.MatchString(text) {
		return nil, fmt.Errorf("bad text: %s", text)
	}

	if strings.Contains(text, "  ") {
		return nil, errors.New("text contains adjacent spaces")
	}

	return &BillBoard{width, height, text}, nil
}

// min returns the smaller of a or b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// isFontTooBig returns true if fontSize + 1 is too big for the BillBoard
func (b BillBoard) isFontTooBig(fontSize int) bool {
	fontSize += 1

	// Split the input text, using space as separator
	words := strings.Split(b.text, " ")

	// Max chars per line, given a BillBoard width and a fontSize
	maxCharsPerLine := b.width / fontSize

	// Max lines in the BillBoard, given a BillBoard height and a fontSize
	maxLines := b.height / fontSize

	currentLines := 1
	currentLineLength := 0

	for _, word := range words {
		wordLen := len(word)

		// Current word is longer than the BillBoard width
		if wordLen > maxCharsPerLine {
			return true
		}

		// Add a leading space only if it's not the beginning of the line
		space := 0
		if currentLineLength > 0 {
			space = 1
		}

		// If the current word does not fit in the current line
		if currentLineLength+space+wordLen > maxCharsPerLine {
			// Increment the line count
			currentLines += 1
			currentLineLength = wordLen
		} else {
			// Increment the current line length
			currentLineLength += space + wordLen
		}

		// Return true if the current lines are more than the max allowed lines
		if currentLines > maxLines {
			return true
		}
	}

	return false
}

// GetFontSize finds the biggest font size allowed for the BillBoard. A binary search is used in order to minimize
// the comparisons. This algorithm runs in O(log(n)) where n is the min between the BillBoard width and height.
//
// It returns an integer representing the maximum allowed font size. If the text does not fit the BillBoard it
// returns 0
func (b BillBoard) GetFontSize() int {
	maxFontSize := min(b.width, b.height)

	return sort.Search(maxFontSize, b.isFontTooBig)
}
