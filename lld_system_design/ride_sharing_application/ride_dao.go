package main

type RideStatus int

const (
	Active = iota + 1
	Created
	Finished
	Cancelled
	WITHDRAWN
	Failed

	AMT_PER_KM = 20
)

// ..........................Ride......................
type Ride struct {
	id         int
	origin     int
	dest       int
	seat       int
	rideStatus RideStatus
}

func RideConstracter() Ride {
	return Ride{}
}
func (ride *Ride) CalculateFare(isPriorityRider bool) float64 {
	dist := ride.dest - ride.origin
	priority := float64(1)
	if ride.seat < 2 {
		if isPriorityRider {
			priority = .75
		}
		return float64(dist) * AMT_PER_KM * priority
	}
	priority = .75
	if isPriorityRider {
		priority = .5
	}
	return float64(dist) * AMT_PER_KM * priority * float64(ride.seat)
}

// Getter and Setter function
func (ride *Ride) GetId() int {
	return ride.id
}

func (ride *Ride) GetOrigin() int {
	return ride.origin
}

func (ride *Ride) GetDest() int {
	return ride.dest
}

func (ride *Ride) GetSeat() int {
	return ride.seat
}

func (ride *Ride) GetRideStatus() RideStatus {
	return ride.rideStatus
}

func (ride *Ride) SetId(id int) *Ride {
	ride.id = id
	return ride
}

func (ride *Ride) SetOrigin(origin int) *Ride {
	ride.origin = origin
	return ride
}

func (ride *Ride) SetDest(dest int) *Ride {
	ride.dest = dest
	return ride
}

func (ride *Ride) SetSeat(seat int) *Ride {
	ride.seat = seat
	return ride
}

func (ride *Ride) SetRideStatus(rideStatus RideStatus) *Ride {
	ride.rideStatus = rideStatus
	return ride
}
