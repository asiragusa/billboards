package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBillBoardError(t *testing.T) {
	lines := []string{
		"",
		"1",
		"1 1",
		"a 1 text",
		"1 a text",
		"0 1 text",
		"-1 1 text",
		"1 -1 text",
		"1 0 text",
		"1 1 ",
		"1 1  ",
		"1 1 " + strings.Repeat("a", 1001),
		"1 1  text",
		"1 1 text ",
		"1 1 text  text",
		"1 1 :",
		"1 1 ,",
		"1001 1 text",
		"1 1001 text",
	}
	for _, line := range lines {
		billBoard, err := NewBillBoard(line)
		if assert.Error(t, err, `"`+line+`"`) {
			assert.Nil(t, billBoard)
		}
	}
}

func TestNewBillBoardOK(t *testing.T) {
	tx1000 := strings.Repeat("t", 1000)
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	type test struct {
		input  string
		width  int
		height int
		text   string
	}
	lines := []test{
		{"1 1 t", 1, 1, "t"},
		{"1000 1000 t", 1000, 1000, "t"},
		{"1 1 " + tx1000, 1, 1, tx1000},
		{"1 1 " + validChars, 1, 1, validChars},
	}
	for _, line := range lines {
		billBoard, err := NewBillBoard(line.input)
		if assert.NoError(t, err, `"`+line.input+`"`) {
			assert.Equal(t, line.width, billBoard.width)
			assert.Equal(t, line.height, billBoard.height)
			assert.Equal(t, line.text, billBoard.text)
		}
	}
}

func TestIsFontTooBig(t *testing.T) {
	type fontTest struct {
		fontSize int
		retValue bool
	}
	type billBoardTest struct {
		billBoard BillBoard
		tests     []fontTest
	}

	tests := []billBoardTest{
		{
			billBoard: BillBoard{
				width:  10,
				height: 10,
				text:   "t",
			},
			tests: []fontTest{
				{9, false},
				{10, true},
			},
		},
		{
			billBoard: BillBoard{
				width:  10,
				height: 5,
				text:   "tt",
			},
			tests: []fontTest{
				{4, false},
				{5, true},
			},
		},
		{
			billBoard: BillBoard{
				width:  50,
				height: 10,
				text:   "0123456789 test",
			},
			tests: []fontTest{
				{4, false},
				{5, true},
			},
		},
		{
			billBoard: BillBoard{
				width:  50,
				height: 9,
				text:   "0123456789 test",
			},
			tests: []fontTest{
				{3, false},
				{4, true},
			},
		},
	}

	for _, billBoardTest := range tests {
		billBoard := billBoardTest.billBoard
		for _, test := range billBoardTest.tests {
			assert.Equal(t, test.retValue, billBoard.isFontTooBig(test.fontSize),
				fmt.Sprintf("FontSize: %d BillBoard: %#v", test.fontSize, billBoard))
		}
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 2))
	assert.Equal(t, 1, min(2, 1))
}

func TestGetFontSize(t *testing.T) {
	type billBoardTest struct {
		billBoard        BillBoard
		expectedFontSize int
	}
	tests := []billBoardTest{
		{
			billBoard: BillBoard{
				width:  10,
				height: 10,
				text:   "t",
			},
			expectedFontSize: 10,
		},
		{
			billBoard: BillBoard{
				width:  1,
				height: 10,
				text:   "t",
			},
			expectedFontSize: 1,
		},
		{
			billBoard: BillBoard{
				width:  10,
				height: 1,
				text:   "t",
			},
			expectedFontSize: 1,
		},
		{
			billBoard: BillBoard{
				width:  50,
				height: 10,
				text:   "0123456789 test",
			},
			expectedFontSize: 5,
		},
		{
			billBoard: BillBoard{
				width:  9,
				height: 1,
				text:   "0123456789",
			},
			expectedFontSize: 0,
		},
		{
			billBoard: BillBoard{
				width:  10,
				height: 1,
				text:   "0123456789",
			},
			expectedFontSize: 1,
		},
		{
			billBoard: BillBoard{
				width:  20,
				height: 6,
				text:   "012345 012",
			},
			expectedFontSize: 3,
		},
		{
			billBoard: BillBoard{
				width:  100,
				height: 20,
				text:   "0123456 012 0123",
			},
			expectedFontSize: 10,
		},
		{
			billBoard: BillBoard{
				width:  10,
				height: 20,
				text:   "0123 01 0123 01 0123",
			},
			expectedFontSize: 2,
		},
		{
			billBoard: BillBoard{
				width:  55,
				height: 25,
				text:   "012 012 0123",
			},
			expectedFontSize: 8,
		},
		{
			billBoard: BillBoard{
				width:  100,
				height: 20,
				text:   "0123 0123 012 01 012 012",
			},
			expectedFontSize: 7,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedFontSize, test.billBoard.GetFontSize(), fmt.Sprintf("%#v", test.billBoard))
	}

}
