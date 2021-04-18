package main

import "fmt"

type Racehorses []string

func main() {
	racehorses := Racehorses{"Winner", "Dark horse", "Mediocre horse", "Lame duck"}
	executeNormalRace(racehorses)
	executeRaceWithDefer(racehorses)

}

func executeNormalRace(rh Racehorses) {
	fmt.Println("Start a normal horse race: 🏇🏇🏇🏇")
	var resultTemplate string
	for i, horse := range rh {
		if i == 0 {
			resultTemplate = "%v: %v 🥇\n"
		} else {
			resultTemplate = "%v: %v \n"
		}
		fmt.Printf(resultTemplate, i+1, horse)
	}
}

func executeRaceWithDefer(rh Racehorses) {
	fmt.Println("Start a horse race including defer: 🏇🏇🏇🏇")
	var resultTemplate string
	for i, horse := range rh {
		if i == 0 {
			resultTemplate = "%v: %v 🥇\n"
		} else {
			resultTemplate = "%v: %v \n"
		}
		defer fmt.Printf(resultTemplate, i+1, horse)
	}
}
