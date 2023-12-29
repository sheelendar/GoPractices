package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	id              string
	currentPosition int
}

func PalyerConstractor(id string, currentPosition int) *Player {
	return &Player{id: id, currentPosition: currentPosition}
}

type Dice struct {
	diceCount int
	min       int
	max       int
}

func DiceConstractor(numOfDice int) *Dice {
	return &Dice{diceCount: numOfDice, min: 1, max: 6}
}
func (dice *Dice) rollDice() int {
	totalSum := 0
	diceC := 0
	for diceC < dice.diceCount {
		totalSum += rand.Intn(dice.max) + 1
		diceC++
	}
	return totalSum
}

type Game struct {
	board   *Board
	dice    *Dice
	players []*Player
	winner  *Player
}

func GameConstractor() *Game {
	board := BoardConstractor(10, 5, 4)
	dice := DiceConstractor(1)
	players := []*Player{PalyerConstractor("shee", 0), PalyerConstractor("ram", 0)}
	return &Game{board: board, dice: dice, players: players}
}
func (game *Game) startGame() {
	for game.winner == nil {
		playerTurm := game.playerTurn()
		fmt.Println("player turn ", playerTurm.id, " current Position ", playerTurm.currentPosition)
		diceNumber := game.dice.rollDice()
		newPosition := diceNumber + playerTurm.currentPosition
		playerTurm.currentPosition = game.jumpCheck(newPosition)

		fmt.Println("player turn ", playerTurm.id, " new Position ", playerTurm.currentPosition)
		size := len(game.board.cells)
		if playerTurm.currentPosition >= size*size {
			game.winner = playerTurm
		}
	}
	fmt.Println("winnner player ", game.winner.id, "Position ", game.winner.currentPosition)
}

func (game *Game) playerTurn() *Player {
	palyer := game.players[0]
	game.players = game.players[1:]
	game.players = append(game.players, palyer)
	return palyer
}

func (game *Game) jumpCheck(position int) int {
	size := len(game.board.cells)
	if position >= size*size-1 {
		return position
	}
	row, col := getCell(position, size)
	if game.board.cells[row][col].jump != nil {
		jump := game.board.cells[row][col].jump
		return jump.end
	}
	return position
}
