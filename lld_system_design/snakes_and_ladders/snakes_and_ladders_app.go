package main

/*
.............................Snakes And Ladders...........................

Rules of the Game
• It’s a multiplayer game which has a 10x10 board(numbered 1 to 100), a unique colored
pin(something to indicate the players’ position on the board) for each player and a
dice(values 1 to 6) to be rolled which tells the next place for the current player’s pin.
• All the pins are initially kept outside of the board. The players participate in a round-robin
manner i.e., they initially agree upon a sequence to throw the dice and take subsequent turns
• The dice value indicates the next position the pin should move to. If the new position has a
snake head, the player ends up at the square indicated by the snake tail. If it’s a ladder head,
the player pin ends up at the other ladder head
• Ladders are meant to be “climbed up”. It helps a player to reach the number 100 faster.
Snakes are meant to do the opposite. For instance, if a player is at position 61, and gets the
number 3 as the dice value, s/he moves the pin to the number 64. If there’s a snake head at
64, the snake tail would always point to a number lesser than 64. The player should move
the pin to the corresponding value. The situation is vice-versa for ladders, in which the player
always moves the pin to a higher value.
• The player who(whose pin) reaches the number 100 first, wins the game. If the dice rolled
points to a number greater than 100, the move is not counted and the turn is passed.


Assumptions
• There’s at least one possible way to win the game.
• A board block/cell can either have one of the types: snake head, snake tail, any of the ladder
ends or nothing

Application requirements
• Snakes, Ladders and number of players are decided before the game begins. Once the game
begins, these can’t be changed.
• A rolled dice should have a random value.
• The game ends when a player reaches the number 100.
• After the game ends, the system should display the final values of each player and also
declare the winner
*/

func main() {
	game := GameConstractor()
	game.startGame()
}

/*
// output


 ----  ----  ---- 3-86  ----  ----  ----  ---- 8-80  ----
 ----  ----  ----  ----  ---- 15-8  ----  ----  ----  ----
 ----  ----  ----  ----  ----  ---- 26-85  ----  ----  ----
 ----  ----  ----  ----  ----  ----  ----  ----  ----  ----
 ----  ----  ----  ----  ----  ----  ----  ----  ---- 49-27
 ----  ----  ----  ----  ---- 55-33  ----  ----  ----  ----
 ----  ---- 62-90  ----  ----  ----  ----  ----  ----  ----
 ----  ----  ----  ----  ----  ----  ----  ----  ----  ----
 ----  ----  ----  ----  ----  ----  ----  ----  ---- 89-86
 ----  ---- 92-28  ----  ----  ----  ----  ----  ----  ----
player turn  shee  current Position  0
player turn  shee  new Position  1
player turn  ram  current Position  0
player turn  ram  new Position  6
player turn  shee  current Position  1
player turn  shee  new Position  5
player turn  ram  current Position  6
player turn  ram  new Position  12
player turn  shee  current Position  5
player turn  shee  new Position  7
player turn  ram  current Position  12
player turn  ram  new Position  8
player turn  shee  current Position  7
player turn  shee  new Position  10
player turn  ram  current Position  8
player turn  ram  new Position  14
player turn  shee  current Position  10
player turn  shee  new Position  16
player turn  ram  current Position  14
player turn  ram  new Position  18
player turn  shee  current Position  16
player turn  shee  new Position  17
player turn  ram  current Position  18
player turn  ram  new Position  20
player turn  shee  current Position  17
player turn  shee  new Position  23
player turn  ram  current Position  20
player turn  ram  new Position  85
player turn  shee  current Position  23
player turn  shee  new Position  85
player turn  ram  current Position  85
player turn  ram  new Position  86
player turn  shee  current Position  85
player turn  shee  new Position  86
player turn  ram  current Position  86
player turn  ram  new Position  87
player turn  shee  current Position  86
player turn  shee  new Position  91
player turn  ram  current Position  87
player turn  ram  new Position  91
player turn  shee  current Position  91
player turn  shee  new Position  93
player turn  ram  current Position  91
player turn  ram  new Position  94
player turn  shee  current Position  93
player turn  shee  new Position  98
player turn  ram  current Position  94
player turn  ram  new Position  95
player turn  shee  current Position  98
player turn  shee  new Position  103
winnner player  shee Position  103

*/
