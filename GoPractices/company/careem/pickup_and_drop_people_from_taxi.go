package main

import "fmt"

/*
There is a car with capacity empty seats. The vehicle only drives east (i.e., it cannot turn around and drive west).

You are given the integer capacity and an array trips where trips[i] = [numPassengersi, fromi, toi]
indicates that the ith trip has numPassengersi passengers and the locations to pick them up and drop them
off are fromi and toi respectively. The locations are given as the number of kilometers due east from the
car's initial location.

Return min capacity of car so all passengers can be served once.



Example 1:

Input: trips = [[2,1,5],[3,3,7]]
Output: 5 :  capacity 5 to serve all passenger

Example 2:

Input: trips = [[2,1,5],[3,3,7],[3,6,7]]
Output: 6 :capacity 6 to serve all passenger
*/

func main() {
	trips := [][]int{{2, 1, 5}, {3, 3, 7}, {3, 6, 7}}
	fmt.Println(GetMinCapacityOfCar(trips))
}
func GetMinCapacityOfCar(trips [][]int) int {
	return 0
}
