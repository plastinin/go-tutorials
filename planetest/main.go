package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"planetest/objects"
	"time"
)

func main() {

	board := objects.Plane{
		Name:              "Boing 747",
		PassengerCapacity: 100,
		Passenger:         true,
		Speed:             uint(rand.Intn(1000)),
	}

	board2 := objects.Plane{
		Name:              "Катамаран",
		PassengerCapacity: 100,
		Passenger:         true,
		Speed:             uint(rand.Intn(1000)),
	}

	flight := objects.NewFlight("DP393", board, "Череповец", "Екатеринбург")
	flight2 := objects.NewFlight("DP808", board2, "Москва", "Анталия")

	pass1 := objects.Passenger{
		Name:    "Юля",
		Surname: "Базюганова",
		Male:    false,
	}
	pass2 := objects.Passenger{
		Name:    "Артем",
		Surname: "Пластинин",
		Male:    true,
	}
	pass3 := objects.Passenger{
		Name:    "Илья",
		Surname: "Боговик",
		Male:    true,
	}
	pass4 := objects.Passenger{
		Name:    "Юля",
		Surname: "Боговик",
		Male:    false,
	}

	flight.AddPassenger(pass1)
	flight.AddPassenger(pass2)
	flight.GetFlightInfo()

	fmt.Println("\n--------------------\n")

	flight2.AddPassenger(pass3)
	flight2.AddPassenger(pass4)
	flight2.GetFlightInfo()

	fmt.Println("\nНажмите, чтобы начать гонку\n")
	scanner := bufio.NewScanner(os.Stdin)
	if ok := scanner.Scan(); !ok {
		fmt.Println("Zopa!")
		return
	}

	Finish := make(chan *objects.Flight)

	go flight.StartFly(Finish)
	go flight2.StartFly(Finish)

loop:
	for {
		select {
		case theWinner := <-Finish:
			fmt.Println(theWinner.Number, theWinner.Board.Name, "финишировал")
			break loop
		default:
			time.Sleep(1 * time.Second)
			flight.PrintCurrentDitance()
			flight2.PrintCurrentDitance()
		}
	}
}
