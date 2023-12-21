/*
You are working on a video game where the player has to go through a level without falling into any obstacles.

The player starts at position zero and can move in three ways:
- L (left)  => one position to the left
- R (right) => one position to the right
- J (jump)  => move two positions, in the direction of the previous move

The player starts at position 0 and the exit will always be at position 10.

The instructions never lead the player outside the level boundaries, and the first move is always right.

Write a function that given the instructions and the positions of the obstacles, returns True if the instructions lead to the exit position, and False otherwise.

For example:

Obstacles 1: [4,6]

--------------------------------------------------------

	->                X         X                   Exit

--------------------------------------------------------
0    1    2    3    4    5    6    7    8    9    10

Instructions 1: "RRRJJRRR" -> True.

	                                                        0,1,2
	JUMP      JUMP                                  RRRLJ = i+1, i+1

-----------------/-----\---/-----\----------------------

	->   ->   ->      X         X      ->   ->   -> Exit

--------------------------------------------------------
0    1    2    3    4    5    6    7    8    9    10

Instructions 2: "RRRLJ" -> False, it would just lead back to the start.

Instructions 3: "RRRJJRRRL" -> True, extra instructions can be ignored.

Instructions 4: "RRRLRJJRRR" -> True, the player is allowed to move back and forth.

Instructions 5: "RRRRRRRRRR" -> False, due to falling into an obstacle.

Instructions 6: "RRJJJJ" -> False, as the jump would land on an obstacle.

Instructions 7: "RLRRRJJRRLLJJJLRRRJJRRR" -> True, even after many instructions, exit is reached.

Instructions 8: "RRRJJRLJJJRRR" -> False, as directions of jumps must be observed.

Instructions 9: "R" -> False, as the exit position isn't reached.

Instructions 10: "RJJJJR" -> True, as it gets to the exit after avoiding the obstacles.

Example -2:
Obstacles 2: [2,9,4], Instructions 11: "RJJRRRJ" -> True, as it gets to the exit after avoiding the obstacles.

	0,1,3,5 6,7,8,10

Obstacles 3: [], Instructions 9: -> False

All test cases:

obstacles_1 = [4,6]
obstacles_2 = [2,9,4]
obstacles_3 = []

level(obstacles_1, instructions_1) # True
level(obstacles_1, instructions_2) # False
level(obstacles_1, instructions_3) # True
level(obstacles_1, instructions_4) # True
level(obstacles_1, instructions_5) # False
level(obstacles_1, instructions_6) # False
level(obstacles_1, instructions_7) # True
level(obstacles_1, instructions_8) # False
level(obstacles_1, instructions_9) # False
level(obstacles_1, instructions_10) # True
level(obstacles_2, instructions_11) # True
level(obstacles_3, instructions_9)  # False

Complexity variables:

N - number of instructions.
*/
package main

import "fmt"

func main() {
	var obstacles_1 []int = []int{4, 6}
	/*var obstacles_2 []int = []int {2, 9, 4}
	  var obstacles_3 []int = []int {} */

	//var instructions_1 string = "RRRJJRRR" //true
	//var instructions_1 string = "RRRLJ" //false
	//var instructions_1 string = "RRRJJRRRL" // true
	//var instructions_1 string = "RRRLRJJRRR" // true
	//var instructions_1 string = "RRRRRRRRRR" //false
	//var instructions_1 string = "RRJJJJ"  //false
	//var instructions_1 string = "RLRRRJJRRLLJJJLRRRJJRRR" //true
	//var instructions_1 string = "RRRJJRLJJJRRR" //false
	//var instructions_1 string = "R" //false
	//var instructions_1 string = "RJJJJR" // true
	var instructions_1 string = "RJJRRRJ" // true

	fmt.Println(findPath(obstacles_1, instructions_1))

}

func findPath(obs []int, ins string) bool {

	obstacles := make(map[int]bool)
	n := len(ins)
	for i := 0; i < len(obs); i++ {
		obstacles[obs[i]] = true
	}

	step := 0
	isRight := true
	for i := 0; i < n; i++ {
		fmt.Println(step, " - ", i)
		if obstacles[step] {
			return false
		}
		if ins[i] == 'L' {
			step--
			isRight = false
		} else if ins[i] == 'R' {
			step++
			isRight = true
		} else if ins[i] == 'J' { // 3 5
			if isRight {
				step = step + 2
			} else {
				step = step - 2
			}
		}
		if step < 0 {
			return false
		}

		if step >= 10 {
			return true
		}
	}
	return false
}
