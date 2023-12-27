package main

import "fmt"

// ..........................Person......................
type Person struct {
	Name string
}

func (person *Person) GetName() string {
	return person.Name
}

func (person *Person) SetName(Name string) *Person {
	person.Name = Name
	return person
}

// ..........................Rider......................
type Rider struct {
	Person
	Rides    map[int]Ride
	CurrRide Ride
}

func RiderConstracter(name string) Rider {
	return Rider{Person: Person{Name: name}, CurrRide: RideConstracter(), Rides: make(map[int]Ride)}
}

func (rider *Rider) createRide(id, source, dest, seat int) {
	if source > dest || seat < 0 {
		fmt.Println("wrong value in source or seats")
		return
	}
	rider.CurrRide = RideConstracter()
	rider.CurrRide.SetSeat(seat)
	rider.CurrRide.SetOrigin(source)
	rider.CurrRide.SetDest(dest)
	rider.CurrRide.SetId(id)
	rider.CurrRide.SetRideStatus(Created)
}

func (rider *Rider) updateRide(id, source, dest, seat int) {
	if rider.CurrRide.GetRideStatus() == WITHDRAWN {
		fmt.Println("Can not update WITHDRAWN ride ")
		return
	}
	if rider.CurrRide.GetRideStatus() == Finished {
		fmt.Println("Can not update finished ride ")
		return
	}
	rider.CurrRide.SetSeat(seat)
	rider.CurrRide.SetOrigin(source)
	rider.CurrRide.SetDest(dest)
	rider.CurrRide.SetId(id)
	rider.Rides[id] = rider.CurrRide
}

func (rider *Rider) withdrowRide(id int) {
	if id != rider.CurrRide.id {
		fmt.Println("Can not update WITHDRAWN diffrent ride ")
		return
	}
	if rider.CurrRide.GetRideStatus() != Created {
		fmt.Println("Can not WITHDRAWN progres ride ")
		return
	}
	rider.CurrRide.SetRideStatus(WITHDRAWN)
}

func (rider *Rider) closeRide() float64 {

	if rider.CurrRide.GetRideStatus() != Created {
		fmt.Println("Ride wasn't in progress . can not close ride")
		return 0
	}
	rider.CurrRide.SetRideStatus(Finished)
	rider.Rides[rider.CurrRide.id] = rider.CurrRide
	isPriorityRider := false
	if len(rider.Rides) > 10 {
		isPriorityRider = true
	}
	return rider.CurrRide.CalculateFare(isPriorityRider)

}

// ..........................Driver......................
type Driver struct {
	Person
}

func DriverConstracter(name string) Driver {
	return Driver{Person: Person{Name: name}}
}
func GetDriver(name string) Driver {
	return Driver{Person: Person{Name: name}}
}
