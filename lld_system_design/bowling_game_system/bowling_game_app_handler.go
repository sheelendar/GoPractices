package main

import (
	"fmt"
	"math/rand"
)

/*
						DESIGN A BOWLING GAME SYSTEM

RULES OF THE GAME
• The game can be played by multiple players
• Primary goal is to knock off all the pins to gain scores. The game consist of ten frames.
In each frame the bowler (player) has two chances to knock off the pins.

• If the bowler is able to knock down all the pins in one strike, it is called a strike.
If the bowler does it with two it is called a spare.

• If there is a spare, the player gets 5 bonus points. For strike, the player gets 10 bonus points.

• In the final set, a player who rolls a spare or a strike is allowed to roll the extra balls to complete the set. However, only a maximum of three balls can be
rolled in the final set

• During the game, players and their scores will be maintained and shown by the system and winner will be declared at the end of the game.
• Likewise, multiple games can be played in parallel on multiple free lanes.
*/

func main() {
	p1 := GetPlayersConstracter("shee")
	p2 := GetPlayersConstracter("ram")
	p3 := GetPlayersConstracter("shayam")
	p4 := GetPlayersConstracter("mohan")
	players := []*Player{p1, p2, p3, p4}

	game := GetGameConstracter()
	index := game.createSession(players)
	var s1 []int
	var s2 []int
	var s3 []int
	var s4 []int

	for i := 0; i < 20; i++ {
		score := rand.Intn(11)
		s1 = append(s1, score)
		game.roll(p1, index, score)

		score = rand.Intn(11)
		s2 = append(s2, score)
		game.roll(p2, index, score)

		score = rand.Intn(11)
		s3 = append(s3, score)
		game.roll(p3, index, score)

		score = rand.Intn(11)
		s4 = append(s4, score)
		game.roll(p4, index, score)
	}
	fmt.Println("Player-1 ", p1.GetName(), " score ", p1.GetScore())
	for _, s := range s1 {
		fmt.Print(s, ", ")
	}
	fmt.Println()

	fmt.Println("Player-2 ", p2.GetName(), " score ", p2.GetScore())
	for _, s := range s2 {
		fmt.Print(s, ", ")
	}
	fmt.Println()

	fmt.Println("Player-3 ", p3.GetName(), " score ", p3.GetScore())
	for _, s := range s3 {
		fmt.Print(s, ", ")
	}
	fmt.Println()

	fmt.Println("Player-4 ", p4.GetName(), " score ", p4.GetScore())
	for _, s := range s4 {
		fmt.Print(s, ", ")
	}
	fmt.Println()
	game.declearWinner(index)
	game.createSession(players)
	game.createSession(players)
	game.createSession(players)
	game.createSession(players)
}
