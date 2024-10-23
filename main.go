package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Character represents the player's character
type Character struct {
	Name      string
	HitPoints int
	Mana      int
	Strength  int
	// Add other attributes as needed
}

// Monster represents a random enemy
type Monster struct {
	Name      string
	HitPoints int
	Strength  int
}

// Spell represents a magical ability
type Spell struct {
	Name        string
	ManaCost    int
	DamageValue int
	EffectValue int // For non-damaging spells like healing or buffs
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create the player's character
	player := Character{
		Name:      "Player",
		HitPoints: 100,
		Mana:      50,
		Strength:  10,
	}

	// Define available spells
	spells := []Spell{
		{Name: "Magic Missile", ManaCost: 10, DamageValue: 20, EffectValue: 0},
		{Name: "Protection from Evil", ManaCost: 15, DamageValue: 0, EffectValue: 10}, // Increase defense
		// Add more spells as needed
	}

	// Game loop
	for {
		// Generate a random monster
		monster := generateMonster()

		// Display the monster's information
		fmt.Printf("A %s appears! (HP: %d, Strength: %d)\n", monster.Name, monster.HitPoints, monster.Strength)

		// Battle loop
		for player.HitPoints > 0 && monster.HitPoints > 0 {
			// Display the current state
			fmt.Printf("Your HP: %d, Mana: %d\n", player.HitPoints, player.Mana)
			fmt.Printf("%s's HP: %d\n", monster.Name, monster.HitPoints)

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
				// Use Mana
				fmt.Println("Choose a spell:")
				for i, spell := range spells {
					fmt.Printf("%d. %s (Mana Cost: %d)", i+1, spell.Name, spell.ManaCost)
					if spell.DamageValue > 0 {
						fmt.Printf(", Damage: %d", spell.DamageValue)
					} else if spell.EffectValue > 0 {
						fmt.Printf(", Effect: %d", spell.EffectValue)
					}
					fmt.Println()
				}
				fmt.Print("Enter your choice: ")

				var spellChoice int
				fmt.Scanln(&spellChoice)

				if spellChoice >= 1 && spellChoice <= len(spells) {
					selectedSpell := spells[spellChoice-1]
					if player.Mana >= selectedSpell.ManaCost {
						player.Mana -= selectedSpell.ManaCost
						if selectedSpell.DamageValue > 0 {
							monster.HitPoints -= selectedSpell.DamageValue
							fmt.Printf("You cast %s, dealing %d damage to the %s.\n", selectedSpell.Name, selectedSpell.DamageValue, monster.Name)
						} else if selectedSpell.EffectValue > 0 {
							// Apply the effect (e.g., increase defense)
							fmt.Printf("You cast %s, increasing your defense by %d.\n", selectedSpell.Name, selectedSpell.EffectValue)
						}
					} else {
						fmt.Println("You don't have enough Mana to cast that spell.")
					}
				} else {
					fmt.Println("Invalid spell choice.")
				}
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
		Name:      name,
		HitPoints: hitPoints,
		Strength:  strength,
	}
}
