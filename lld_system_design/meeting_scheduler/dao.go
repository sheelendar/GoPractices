package main

// ...............................Room struct contains room name and calendar schedule for each day,............................
type Room struct {
	Name     string
	Calendar map[int][]Meeting
}

func (room *Room) getName() string {
	return room.Name
}

func GetRoom(name string) Room {
	return Room{Name: name, Calendar: map[int][]Meeting{}}
}

func (room *Room) book(day, start, end int) bool {
	for _, meeting := range room.Calendar[day] {
		if meeting.Start < end && start < meeting.End {
			return false
		}
	}
	meeting := GetMeeting(start, end, *room)
	room.Calendar[day] = append(room.Calendar[day], meeting)
	return true
}

// .......................Meeting contains start and end time of meeting.........................................
type Meeting struct {
	Start int
	End   int
	Room  Room
}

func GetMeeting(start, end int, room Room) Meeting {
	return Meeting{Start: start, End: end, Room: room}
}

// ...............................................Scheduler conatains all rooms to schedule meeting...............................................
type Scheduler struct {
	Rooms []*Room
}

func GetScheduler(rooms []*Room) Scheduler {
	return Scheduler{Rooms: rooms}
}
func (s *Scheduler) book(day, start, end int) string {
	for _, room := range s.Rooms {
		flag := room.book(day, start, end)
		if flag {
			return room.getName()
		}
	}
	return "No room available"
}
