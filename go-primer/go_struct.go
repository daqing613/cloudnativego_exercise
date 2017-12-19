package main

import (
	"fmt"
	"math"
)

type power struct {
	attack  int
	defense int
}

type location struct {
	x float32
	y float32
	z float32
}

func (loc location) String() string {
	return fmt.Sprintf("(%f, %f, %f)", loc.x, loc.y, loc.z)
}

func (loc location) euclideanDistance(target location) float64 {
	return math.Sqrt((loc.x-target.x)*(loc.x-target.x) + (loc.y-target.y)*(loc.y-target.y) + (loc.z-target.z)*(loc.z-target.z))
}

type nonPlayerCharacter struct {
	name  string
	speed int
	hp    int
	power power
	loc   location
}

func main() {
	fmt.Println("Structs...")
	demon := nonPlayerCharacter{
		name:  "Alfred",
		speed: 21,
		hp:    1000,
		power: power{attack: 75, defense: 50},
		loc:   location{x: 1075.123, y: 521.123, z: 211.231},
	}

	fmt.Println(demon)

	anotherdemon := nonPlayerCharacter{
		name:  "Beelzebub",
		speed: 30,
		hp:    5000,
		power: power{attack: 10, defense: 10},
		loc:   location{x: 32.03, y: 72.54, z: 65.231},
	}

	fmt.Println(anotherdemon)
}
