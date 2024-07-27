package battle

import (
	"fmt"
	"math/rand"

	"magical-arena/player"
)

func RollDice() int {
	return rand.Intn(6) + 1
}

func AttackDamage(attack int) int {
	return attack * RollDice()
}

func DefendDamage(defend int) int {
	return defend * RollDice()
}

func MagicalArena(player1, player2 player.Player) {
	var attacker, defender *player.Player

	if player1.Health >= player2.Health {
		attacker = &player1
		defender = &player2
	} else {
		attacker = &player2
		defender = &player1
	}

	moveCount := 0

	for player1.Health > 0 && player2.Health > 0 {
		moveCount++
		fmt.Printf("Move %d\n", moveCount)

		attack := AttackDamage(attacker.Attack)
		defend := DefendDamage(defender.Strength)

		damage := attack - defend
		if damage > 0 {
			defender.Health -= damage
		}

		fmt.Printf("%s attacks %s for %d damage.\n", attacker.Name, defender.Name, damage)

		attacker, defender = defender, attacker
	}

	if player1.Health > 0 {
		fmt.Printf("%s wins!\n", player1.Name)
	} else {
		fmt.Printf("%s wins!\n", player2.Name)
	}
}
