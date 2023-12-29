package main

import (
	"fmt"
	"math/rand"
)

type Jump struct {
	start int
	end   int
}

type Cell struct {
	jump *Jump
}

type Board struct {
	cells [][]Cell
}

func BoardConstractor(boardSize, numberOfSnak, numberOfLadder int) *Board {
	board := initializeCells(boardSize)
	board.addSnakAndLadder(boardSize, numberOfSnak, numberOfLadder)
	board.printBoard(boardSize)
	return board

}
func (board *Board) printBoard(boardSize int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board.cells[i][j].jump != nil {
				fmt.Print(board.cells[i][j].jump.start, "-", board.cells[i][j].jump.end, " ")
			} else {
				fmt.Print(" ---- ")
			}
		}
		fmt.Println()
	}
}
func (board *Board) addSnakAndLadder(boardSize, numberOfSnak, numberOfLadder int) {
	for numberOfSnak > 0 {
		snakHead := rand.Intn(boardSize * boardSize)
		snakTail := rand.Intn(boardSize * boardSize)
		if snakHead <= snakTail {
			temp := snakHead
			snakHead = snakTail
			snakTail = temp
		}
		jump := &Jump{start: snakHead, end: snakTail}
		row, col := getCell(snakHead, boardSize)

		cell := board.cells[row][col]
		cell.jump = jump
		board.cells[row][col] = cell
		numberOfSnak--
	}

	for numberOfLadder > 0 {
		ladderStart := rand.Intn(boardSize * boardSize)
		ladderEnd := rand.Intn(boardSize * boardSize)
		if ladderEnd <= ladderStart {
			temp := ladderEnd
			ladderEnd = ladderStart
			ladderStart = temp
		}
		jump := &Jump{start: ladderStart, end: ladderEnd}
		row, col := getCell(ladderStart, boardSize)

		cell := board.cells[row][col]
		cell.jump = jump
		board.cells[row][col] = cell
		numberOfLadder--
	}
}
func initializeCells(boardSize int) *Board {
	cells := make([][]Cell, boardSize)

	for i := 0; i < boardSize; i++ {
		cells[i] = make([]Cell, boardSize)
	}
	return &Board{cells: cells}
}

func getCell(size, boardSize int) (row, col int) {
	return size / boardSize, size % boardSize
}
