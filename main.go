package main

import (
	"fmt"

	// "magical-arena/battle"
	"magical-arena/player"
	
)

func main() {
    // Create two players
    player1 := player.CreatePlayer(1)
    player2 := player.CreatePlayer(2)

	fmt.Println(player1)
	fmt.Println(player2)

    // Start the magical arena battle
    // battle.MagicalArena(player1, player2)
}
