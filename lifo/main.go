package main

import "fmt"

type Racehorses []string

func main() {
	racehorses := Racehorses{"Winner", "Dark horse", "Mediocre horse", "Lame duck"}
	executeNormalRace(racehorses)
	executeRaceWithDefer(racehorses)

}

func executeNormalRace(h Racehorses) {
	fmt.Println("Start a normal horse race: 🏇🏇🏇🏇")
	for i, horse := range h {
		if i == 0 {
			fmt.Printf("%v: %v 🥇\n", i+1, horse)

		} else {
			fmt.Printf("%v: %v \n", i+1, horse)
		}
	}
}

func executeRaceWithDefer(r Racehorses) {
	fmt.Println("Start a horse race including defer: 🏇🏇🏇🏇")
	for i, horse := range r {
		if i == 0 {
			defer fmt.Printf("%v: %v 🥇\n", i+1, horse)

		} else {
			defer fmt.Printf("%v: %v \n", i+1, horse)
		}
	}
}
