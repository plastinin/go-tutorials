package objects

import (
	"errors"
	"fmt"
	"time"
)

type Flight struct {
	Number          string
	Board           Plane
	From            string
	To              string
	passengers      []Passenger
	passengersCount uint
	distance        uint
	currentDistance uint
}

func NewFlight(Number string, Board Plane, From string, To string) *Flight {
	return &Flight{
		Number:          Number,
		Board:           Board,
		From:            From,
		To:              To,
		passengers:      make([]Passenger, 0),
		passengersCount: 0,
		distance:        GetDistanceFromTo(From, To),
		currentDistance: 0,
	}
}

func (f *Flight) AddPassenger(Passenger Passenger) error {
	if f.Board.PassengerCapacity <= f.passengersCount {
		return errors.New("Борт заполнен")
	}
	f.passengersCount++
	f.passengers = append(f.passengers, Passenger)

	return nil
}

func (f *Flight) GetFlightInfo() {
	fmt.Println("Flight number:", f.Number)
	fmt.Println("Board:", f.Board.Name)
	fmt.Println("From:", f.From)
	fmt.Println("To:", f.To)
	fmt.Println("Distance:", f.distance, "km")
	fmt.Println("Passengers count:", f.passengersCount)

	if f.passengersCount != 0 {
		fmt.Println("\nPassengers list:")
		for index, passenger := range f.passengers {
			fmt.Println("\t", index+1, ":", passenger.Surname, passenger.Name)
		}
	}
}

func (f *Flight) StartFly(Finishline chan *Flight) {
	for {
		if f.currentDistance >= f.distance {
			Finishline <- f
		} else {
			time.Sleep(1 * time.Second)
			f.currentDistance += f.Board.Speed
		}
	}
}

func (f *Flight) PrintCurrentDitance() {
	fmt.Println(f.Number, f.Board.Name, "-", f.currentDistance)
}