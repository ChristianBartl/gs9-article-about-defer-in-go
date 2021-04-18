package main

import "fmt"

type Racehorses []string

func main() {
	racehorses := Racehorses{"Winner", "Dark horse", "Mediocre horse", "Lame duck"}
	executeNormalRace(racehorses)
	executeRaceWithDefer(racehorses)

}

func executeNormalRace(h Racehorses) {
	fmt.Println("################ Start a new horse race. ################")
	for i, horse := range h {
		fmt.Printf("%v: %v \n", i, horse)
	}
}

func executeRaceWithDefer(r Racehorses) {
	fmt.Println("################ Start a new horse race. ################")
	for i, horse := range r {
		defer fmt.Printf("%v: %v \n", i, horse)
	}
}
