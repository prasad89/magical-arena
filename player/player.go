package player

import (
	"fmt"
)

type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

func CreatePlayer(playerNumber int) Player {
	fmt.Printf("Enter name for player %d: ", playerNumber)
	var name string
	fmt.Scanln(&name)

	fmt.Printf("Enter health for player %d: ", playerNumber)
	var health int
	fmt.Scanln(&health)

	fmt.Printf("Enter strength for player %d: ", playerNumber)
	var strength int
	fmt.Scanln(&strength)

	fmt.Printf("Enter attack for player %d: ", playerNumber)
	var attack int
	fmt.Scanln(&attack)

	return Player{
		Name:     name,
		Health:   health,
		Strength: strength,
		Attack:   attack,
	}
}
