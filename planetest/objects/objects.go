package objects

type Plane struct {
	Name              string
	Passenger         bool
	PassengerCapacity uint
	Speed uint
}

type Passenger struct {
	Name    string
	Surname string
	Male    bool
}
