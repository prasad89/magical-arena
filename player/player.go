package player

import (
    "fmt"
)

// Player represents a game player with basic attributes.
type Player struct {
    Name     string // Player's name
    Health   int    // Player's health points
    Strength int    // Player's strength for defense
    Attack   int    // Player's attack power
}

// PromptAndValidateAttribute prompts the user for an attribute, validates it, and returns it.
func PromptAndValidateAttribute(prompt string) int {
    var attr int
    for {
        fmt.Print(prompt)
        _, err := fmt.Scan(&attr)
        if err != nil {
            fmt.Println("Invalid input. Please enter a numeric value.")
            // Clear the input buffer
            var discard string
            fmt.Scanln(&discard)
            continue
        }
        if attr > 0 && attr <= 100 {
            return attr
        }
        fmt.Println("Invalid value. Please enter a value between 1 and 100.")
    }
}

// CreatePlayer prompts the user to input details for a new player and returns a Player object.
func CreatePlayer(playerNum int) Player {
    // Prompt user for player's name
    var name string
    fmt.Printf("Enter player %d name: ", playerNum)
    fmt.Scan(&name)

    // Get and validate player's health, strength, and attack using the prompt and validate function
    health := PromptAndValidateAttribute(fmt.Sprintf("Enter player %d health (1-100): ", playerNum))
    strength := PromptAndValidateAttribute(fmt.Sprintf("Enter player %d strength (1-100): ", playerNum))
    attack := PromptAndValidateAttribute(fmt.Sprintf("Enter player %d attack (1-100): ", playerNum))
    fmt.Println()

    // Create and return a Player with the validated details
    return Player{Name: name, Health: health, Strength: strength, Attack: attack}
}
