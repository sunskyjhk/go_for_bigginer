package main

import (
	"fmt"
)

func main() {
	// arrays
	// declaring an empty array of strings
	var months []string
	months = []string{"January", "February"}
	fmt.Println(months[1])

	// declaring an array with elements
	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	fmt.Println(days[2])

	// slices
	weekdays := days[0:5]
	fmt.Println(weekdays)
	// This returns: [Monday Tuesday Wednesday Thursday Friday]

	// Maps
	youtubeSubscribers := map[string]int{
		"TutorialEdge":     2240,
		"MKBHD":            6580350,
		"Fun Fun Function": 171220,
	}

	fmt.Println(youtubeSubscribers["MKBHD"]) // prints out 6580350

	// structs
	// our Person struct
	type Person struct {
		name string
		age  int
	}

	// declaring a new "Person"
	var myPerson Person
	myPerson.name = "Twice"
	myPerson.age = 35

	fmt.Println(myPerson)

	elliot := Person{name: "Elliot", age: 24}
	fmt.Println("elliot age: ", elliot.age)

	elliot.age = 18
	fmt.Println("corrected elliot age: ", elliot.age)

	// nested structs
	// our tesam struct
	type Team struct {
		name    string
		players [2]Person
	}

	// declaring an empty "Team"
	var myTeam Team
	fmt.Println(myTeam)

	players := [...]Person{{name: "Forest"}, {name: "Gordon"}}
	// declaring a team with players
	celtic := Team{name: "Celtic FC", players: players}
	fmt.Println(celtic)
}
