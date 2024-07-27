# Magical Arena

Welcome to the Magical Arena Game!  
This is a simple command-line game where two players battle in an arena until one wins.
Players input their attributes, and the game simulates attacks and defenses using dice rolls.

## Table of Contents

- [Features](#features)
- [Pre-requisites](#pre-requisites)
- [Installation](#installation)
- [Usage](#usage)
- [File Structure](#file-structure)

## Features

- **Player Creation**: Input player attributes such as name, health, strength, and attack.
- **Dice Rolling**: Simulates dice rolls to determine attack and defense values.
- **Battle Simulation**: Players take turns attacking and defending until one player's health drops to zero or below.
- **Outcome Reporting**: Displays detailed battle logs and the winner of the match.

## Pre-requisites

- **Go**: Make sure you have Go installed on your machine. You can download it from the [official website](https://golang.org/dl/).

## Installation

1. **Build the project:**

    ```bash
    go build -o magical-arena
    ```

## Usage

1. **Start the Game:**

    Run the compiled executable to start the game:

    ```bash
    ./magical-arena
    ```

2. **Input Player Details:**

    Follow the prompts to enter details for two players:
    - Player Name
    - Health
    - Strength
    - Attack

    Example input:

    ```
    Enter player 1 name: Alice
    Enter player 1 health (1-100): 100
    Enter player 1 strength (1-100): 20
    Enter player 1 attack (1-100): 10

    Enter player 2 name: Bob
    Enter player 2 health (1-100): 80
    Enter player 2 strength (1-100): 15
    Enter player 2 attack (1-100): 12
    ```

    Example output:

    ```
    Alice starts the battle against Bob

    Move 1
    Alice attacks Bob with a dice roll of 2
    but Bob defends with a dice roll of 2 and does no damage
    Bob has 80 health remaining

    Move 2
    Bob attacks Alice with a dice roll of 3
    Alice defends with a dice roll of 1, dealing 16 damage
    Alice has 84 health remaining

    Move 3
    Alice attacks Bob with a dice roll of 6
    but Bob defends with a dice roll of 6 and does no damage
    Bob has 80 health remaining

    Move 4
    Bob attacks Alice with a dice roll of 1
    but Alice defends with a dice roll of 5 and does no damage
    Alice has 84 health remaining

    Move 5
    Alice attacks Bob with a dice roll of 6
    Bob defends with a dice roll of 2, dealing 30 damage
    Bob has 50 health remaining

    Move 6
    Bob attacks Alice with a dice roll of 2
    but Alice defends with a dice roll of 6 and does no damage
    Alice has 84 health remaining

    Move 7
    Alice attacks Bob with a dice roll of 2
    Bob defends with a dice roll of 1, dealing 5 damage
    Bob has 45 health remaining

    Move 8
    Bob attacks Alice with a dice roll of 3
    but Alice defends with a dice roll of 5 and does no damage
    Alice has 84 health remaining

    Move 9
    Alice attacks Bob with a dice roll of 3
    Bob defends with a dice roll of 1, dealing 15 damage
    Bob has 30 health remaining

    Move 10
    Bob attacks Alice with a dice roll of 1
    but Alice defends with a dice roll of 6 and does no damage
    Alice has 84 health remaining

    Move 11
    Alice attacks Bob with a dice roll of 2
    but Bob defends with a dice roll of 2 and does no damage
    Bob has 30 health remaining

    Move 12
    Bob attacks Alice with a dice roll of 4
    Alice defends with a dice roll of 1, dealing 28 damage
    Alice has 56 health remaining

    Move 13
    Alice attacks Bob with a dice roll of 5
    Bob defends with a dice roll of 3, dealing 5 damage
    Bob has 25 health remaining

    Move 14
    Bob attacks Alice with a dice roll of 2
    but Alice defends with a dice roll of 4 and does no damage
    Alice has 56 health remaining

    Move 15
    Alice attacks Bob with a dice roll of 1
    but Bob defends with a dice roll of 2 and does no damage
    Bob has 25 health remaining

    Move 16
    Bob attacks Alice with a dice roll of 3
    but Alice defends with a dice roll of 5 and does no damage
    Alice has 56 health remaining

    Move 17
    Alice attacks Bob with a dice roll of 5
    Bob defends with a dice roll of 3, dealing 5 damage
    Bob has 20 health remaining

    Move 18
    Bob attacks Alice with a dice roll of 2
    but Alice defends with a dice roll of 2 and does no damage
    Alice has 56 health remaining

    Move 19
    Alice attacks Bob with a dice roll of 1
    but Bob defends with a dice roll of 2 and does no damage
    Bob has 20 health remaining

    Move 20
    Bob attacks Alice with a dice roll of 4
    but Alice defends with a dice roll of 6 and does no damage
    Alice has 56 health remaining

    Move 21
    Alice attacks Bob with a dice roll of 6
    Bob defends with a dice roll of 2, dealing 30 damage
    Bob has -10 health remaining

    Alice wins!
    ```
    
3. **View Battle Results:**

    The game will simulate the battle and display the results, including attack details, remaining health, and the winner.

## File Structure

Here is the structure of the project:

```
magical-arena/
├── main.go               # Entry point of the application
├── go.mod                # Go module file
├── README.md             # Project documentation
├── player/               # Package for player-related functionality
│   ├── player.go         # Player struct and creation function
│   └── player_test.go    # Tests for player functionality
├── battle/               # Package for battle-related functionality
│   ├── battle.go         # Battle logic and helper functions
│   └── battle_test.go    # Tests for battle functionality
```