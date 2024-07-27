package battle

import (
	"fmt"
	"math/rand"

	"magical-arena/player"
)

// RollDice simulates rolling a six-sided die and returns a number between 1 and 6.
func RollDice() int {
	return rand.Intn(6) + 1
}

// CalculateAttackDamage calculates the damage dealt based on attack power and a dice roll.
func CalculateAttackDamage(attack int) (int, int) {
	diceRoll := RollDice()
	return attack * diceRoll, diceRoll
}

// CalculateDefenseDamage calculates the defense value based on strength and a dice roll.
func CalculateDefenseDamage(strength int) (int, int) {
	diceRoll := RollDice()
	return strength * diceRoll, diceRoll
}

// MatchStatus displays the outcome of an attack, including the dice rolls and the defender's remaining health.
func MatchStatus(attacker, defender player.Player, damage, attackDiceRoll, defenseDiceRoll int) {
	if damage > 0 {
		fmt.Printf(
			"%s attacks %s with a dice roll of %d\n"+
				"%s defends with a dice roll of %d, dealing %d damage\n"+
				"%s has %d health remaining\n",
			attacker.Name,
			defender.Name,
			attackDiceRoll,
			defender.Name,
			defenseDiceRoll,
			damage,
			defender.Name,
			defender.Health,
		)
	} else {
		fmt.Printf(
			"%s attacks %s with a dice roll of %d\n"+
				"but %s defends with a dice roll of %d and does no damage\n"+
				"%s has %d health remaining\n",
			attacker.Name,
			defender.Name,
			attackDiceRoll,
			defender.Name,
			defenseDiceRoll,
			defender.Name,
			defender.Health,
		)
	}
	fmt.Println()
}

// MagicalArena simulates a battle between two players until one wins.
func MagicalArena(player1, player2 player.Player) {
	var attacker, defender *player.Player

	// Determine which player starts as the attacker based on health.
	if player1.Health >= player2.Health {
		attacker = &player1
		defender = &player2
	} else {
		attacker = &player2
		defender = &player1
	}

	// Initialize move count.
	moveCount := 0

	// Display the start of the battle.
	fmt.Printf("%s starts the battle against %s\n", attacker.Name, defender.Name)
	fmt.Println()

	// Continue the battle until one player's health drops to 0 or below.
	for player1.Health > 0 && player2.Health > 0 {
		// Increment and display the move count.
		moveCount++
		fmt.Printf("Move %d\n", moveCount)

		// Calculate damage dealt and apply it to the defender.
		attackDamage, attackDiceRoll := CalculateAttackDamage(attacker.Attack)
		defenseDamage, defenseDiceRoll := CalculateDefenseDamage(defender.Strength)
		damage := attackDamage - defenseDamage
		if damage > 0 {
			defender.Health -= damage
		}

		// Display the result of the attack.
		MatchStatus(*attacker, *defender, damage, attackDiceRoll, defenseDiceRoll)

		// Swap roles: attacker becomes defender and vice versa.
		attacker, defender = defender, attacker
	}

	// Determine and display the winner.
	if player1.Health > 0 {
		fmt.Printf("%s wins!\n", player1.Name)
	} else {
		fmt.Printf("%s wins!\n", player2.Name)
	}
}
