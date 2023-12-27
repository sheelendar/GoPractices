package main

import "fmt"

/*
Book room for a meeting and return name of room booked
History of meeting room booked
Wrtie a api that take start and end time and return room booked name
*/
func main() {
	room1 := GetRoom("Atlas")
	room2 := GetRoom("Nexus")
	room3 := GetRoom("Holycow")
	rooms := []*Room{&room1, &room2, &room3}
	scheduler := GetScheduler(rooms)
	fmt.Println(scheduler.book(15, 2, 5))
	fmt.Println(scheduler.book(15, 5, 8))
	fmt.Println(scheduler.book(15, 4, 8))
	fmt.Println(scheduler.book(15, 3, 6))
	fmt.Println(scheduler.book(15, 7, 8))
	fmt.Println(scheduler.book(16, 6, 9))
}
