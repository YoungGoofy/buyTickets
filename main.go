package main

import (
	"fmt"
	"math/rand"
)

func main() {
	spacelines := []string{"Space Adventures", "SpaceX", "Virgin Galactic"}
	tripType := []string{"Round-trip", "One-way"}
	fmt.Println("Spaceline\tDays\tTrip type\tPrice")
	fmt.Println("==============================================")
	for i := 0; i < 10; i++ {
		spacelineType := rand.Intn(3)
		speed := 16 + rand.Intn(15)
		chooseTrip := rand.Intn(2)
		timeInSpace := (62100000/(speed * 1000))/24
		ticet_cost := 20 + speed
		if chooseTrip == 0 {
			ticet_cost *= 2
		}
		fmt.Printf("%-16v %-6v %-15v $%4v\n", spacelines[spacelineType], timeInSpace, tripType[chooseTrip], ticet_cost)
	}
}
