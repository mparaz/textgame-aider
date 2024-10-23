package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Character represents the player's character
type Character struct {
	Name     string
	HitPoints int
	Mana     int
	Strength int
	// Add other attributes as needed
}

// Monster represents a random enemy
type Monster struct {
	Name     string
	HitPoints int
	Strength int
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create the player's character
	player := Character{
		Name:     "Player",
		HitPoints: 100,
		Mana:     50,
		Strength: 10,
	}

	// Game loop
	for {
		// Generate a random monster
		monster := generateMonster()

		// Display the monster's information
		fmt.Printf("A %s appears! (HP: %d, Strength: %d)\n", monster.Name, monster.HitPoints, monster.Strength)

		// Battle loop
		for player.HitPoints > 0 && monster.HitPoints > 0 {
			// Player's turn
			fmt.Println("What do you want to do?")
			fmt.Println("1. Attack")
			fmt.Println("2. Use Mana")
			fmt.Print("Enter your choice: ")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				// Attack the monster
				damage := player.Strength
				monster.HitPoints -= damage
				fmt.Printf("You deal %d damage to the %s.\n", damage, monster.Name)
			case 2:
				// Use Mana (you can implement your own logic here)
				fmt.Println("You use some of your Mana.")
			default:
				fmt.Println("Invalid choice.")
			}

			// Monster's turn
			if monster.HitPoints > 0 {
				damage := monster.Strength
				player.HitPoints -= damage
				fmt.Printf("The %s deals %d damage to you.\n", monster.Name, damage)
			}
		}

		// Check the outcome of the battle
		if player.HitPoints > 0 {
			fmt.Println("You defeated the monster!")
		} else {
			fmt.Println("You were defeated by the monster.")
			break
		}
	}
}

// generateMonster generates a random monster
func generateMonster() Monster {
	// Generate random monster attributes
	name := fmt.Sprintf("Monster %d", rand.Intn(100))
	hitPoints := rand.Intn(50) + 10
	strength := rand.Intn(10) + 1

	return Monster{
		Name:     name,
		HitPoints: hitPoints,
		Strength: strength,
	}
}
