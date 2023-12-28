package main

const MAX_ROLLS_ALLOWED = 20

type Player struct {
	name        string
	score       int
	framIndex   int
	firstRoll   bool
	rolls       []int
	canPlay     bool
	currentRoll int
}

func GetPlayersConstracter(name string) *Player {
	return &Player{name: name, rolls: make([]int, MAX_ROLLS_ALLOWED),
		firstRoll: true, canPlay: true}
}

func (player *Player) roll(pins int) {
	if player.canPlay == false {
		return
	}
	player.rolls[player.currentRoll] = pins
	player.currentRoll++
	player.updateScore(pins)

}
func (player *Player) updateScore(pins int) {
	if player.isStrike() {
		player.score += 20
		player.rolls[player.currentRoll] = 0
		player.currentRoll++
		player.framIndex += 2
		if player.framIndex == MAX_ROLLS_ALLOWED {
			player.canPlay = false
		}
	} else {
		if player.framIndex >= MAX_ROLLS_ALLOWED {
			player.score += player.rolls[player.framIndex]
			player.canPlay = false
		} else if player.firstRoll {
			player.firstRoll = false
		} else {
			if player.isSPare() {
				player.score += 5
			}
			player.score += player.rolls[player.framIndex] + player.rolls[player.framIndex+1]
			player.framIndex += 2
			player.firstRoll = true
			if player.framIndex >= MAX_ROLLS_ALLOWED {
				player.canPlay = false
			}
		}
	}
}

func (player *Player) isStrike() bool {
	if player.firstRoll == true && player.rolls[player.framIndex] == 10 {
		return true
	}
	return false
}
func (player *Player) isSPare() bool {
	if player.firstRoll == true && player.rolls[player.framIndex] == 5 {
		return true
	}
	return false
}

func (player *Player) GetName() string {
	return player.name
}

func (player *Player) GetScore() int {
	return player.score
}

func (player *Player) SetName(name string) *Player {
	player.name = name
	return player
}

func (player *Player) SetScore(score int) *Player {
	player.score = score
	return player
}
