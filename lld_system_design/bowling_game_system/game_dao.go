package main

import "fmt"

var GameSessionID = 1

type Game struct {
	gameSession map[int]*GameSession
	alley       []int
}

func GetGameConstracter() *Game {
	return &Game{alley: []int{1, 2, 3, 4}, gameSession: make(map[int]*GameSession)}
}
func (game *Game) makeActive(alley int) {

}
func (game *Game) createSession(players []*Player) int {
	if len(game.alley) == 0 {
		fmt.Println("All alley are cooupied")
		return -1
	}
	gameSession := GetGameSessionConstracter(players)
	gameSession.SetPlayers(players)
	gameSession.SetAlley(game.alley[len(game.alley)-1])
	game.alley = game.alley[:len(game.alley)-1]
	game.gameSession[gameSession.GetId()] = gameSession
	return gameSession.GetId()

}
func (game *Game) roll(player *Player, index, pins int) {
	gameSession := game.gameSession[index]
	gameSession.makeRoll(player, pins)
}
func (game *Game) declearWinner(index int) {
	winnerFlag := game.gameSession[index].declearWinner()
	if !winnerFlag {
		fmt.Println("No winner for this Game")
	}
}

// ..........................GameSession.............................
type GameSession struct {
	alley   int
	id      int
	players []*Player
}

func GetGameSessionConstracter(list []*Player) *GameSession {
	return &GameSession{alley: -1, id: GetUniqueID(), players: list}
}
func (gamesession *GameSession) makeRoll(player *Player, pins int) int {
	for _, p := range gamesession.players {
		if p.GetName() == player.GetName() {
			player.roll(pins)
		}
	}
	return 0
}
func (gamesession *GameSession) declearWinner() bool {
	maxScore := 0
	var winner *Player
	for _, p := range gamesession.players {
		if p.canPlay {
			fmt.Println("player hasn't completed yet. the current score")
			fmt.Println("match is in progress")
			return false
		}
		if p.GetScore() > maxScore {
			maxScore = p.GetScore()
			winner = p
		}
	}
	if winner != nil {
		fmt.Println("this winner is ", winner.GetName(), "winner socre ", winner.GetScore())
	}
	return true
}

func GetUniqueID() int {
	id := GameSessionID
	GameSessionID++
	return id
}

func (gamesession *GameSession) GetAlley() int {
	return gamesession.alley
}

func (gamesession *GameSession) GetId() int {
	return gamesession.id
}

func (gamesession *GameSession) GetPlayers() []*Player {
	return gamesession.players
}

func (gamesession *GameSession) SetAlley(alley int) *GameSession {
	gamesession.alley = alley
	return gamesession
}

func (gamesession *GameSession) SetId(id int) *GameSession {
	gamesession.id = id
	return gamesession
}

func (gamesession *GameSession) SetPlayers(players []*Player) *GameSession {
	gamesession.players = players
	return gamesession
}
