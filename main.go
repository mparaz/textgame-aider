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
	MaxHitPoints int
	Mana      int
	Strength  int
	Weapon    *Weapon // Pointer to the character's current weapon
}

// Monster represents a random enemy
type Monster struct {
	Name        string
	Description string
	HitPoints   int
	Strength    int
	CanCastSpells bool
	Spells      []Spell
}

// Spell represents a magical ability
type Spell struct {
	Name        string
	ManaCost    int
	DamageValue int
	EffectValue int // For non-damaging spells like healing or buffs
}

// Weapon represents a weapon that can be equipped
type Weapon struct {
	Name     string
	Damage   int
	Accuracy int // Chance to hit the target
}

// Room represents a single room in the dungeon
type Room struct {
	Description string
	Monster     *Monster // Pointer to a monster, if present
	Weapon      *Weapon  // Pointer to a weapon, if present
}

// Dungeon represents the entire dungeon
type Dungeon struct {
	Rooms       []*Room
	CurrentRoom int
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create the player's character
	player := Character{
		Name:      "Player",
		HitPoints: 100,
		MaxHitPoints: 100,
		Mana:      50,
		Strength:  10,
		Weapon:    &Weapon{Name: "Rusty Sword", Damage: 5, Accuracy: 80}, // Starting weapon
	}

	// Define available spells
	spells := []Spell{
		{Name: "Magic Missile", ManaCost: 10, DamageValue: 20, EffectValue: 0},
		{Name: "Protection from Evil", ManaCost: 15, DamageValue: 0, EffectValue: 10}, // Increase defense
		// Add more spells as needed
	}

	// Generate a random dungeon
	dungeon := generateDungeon()

	// Game loop
	for {
		// Regenerate hit points over time
		if player.HitPoints < player.MaxHitPoints {
			player.HitPoints += 1
			if player.HitPoints > player.MaxHitPoints {
				player.HitPoints = player.MaxHitPoints
			}
		}

		// Check if the player is dead
		if player.HitPoints <= 0 {
			fmt.Println("You have died a glorious death!")
			printDeathAsciiArt()
			break
		}

		// Display the current room's description
		currentRoom := dungeon.Rooms[dungeon.CurrentRoom]
		fmt.Println(currentRoom.Description)

		// Check if there's a weapon in the room
		if currentRoom.Weapon != nil {
			fmt.Printf("You found a %s (Damage: %d, Accuracy: %d%%)!\n", currentRoom.Weapon.Name, currentRoom.Weapon.Damage, currentRoom.Weapon.Accuracy)
			fmt.Printf("Your current weapon: %s (Damage: %d, Accuracy: %d%%)\n", player.Weapon.Name, player.Weapon.Damage, player.Weapon.Accuracy)
			fmt.Println("Enter 'equip' to equip the new weapon, or any other input to continue.")
			var choice string
			fmt.Scanln(&choice)
			if choice == "equip" {
				player.Weapon = currentRoom.Weapon
				currentRoom.Weapon = nil // Remove the weapon from the room
				fmt.Printf("You equipped the %s.\n", player.Weapon.Name)
			}
		}

		// Check if there's a monster in the room
		if currentRoom.Monster != nil {
			monster := currentRoom.Monster

			// Display the monster's information
			fmt.Printf("%s\n", monster.Description)
			fmt.Printf("(HP: %d, Strength: %d)\n", monster.HitPoints, monster.Strength)

			// Battle loop
			for player.HitPoints > 0 && monster.HitPoints > 0 {
				// Display the current state
				fmt.Printf("Your HP: %d, Mana: %d\n", player.HitPoints, player.Mana)
				fmt.Printf("%s's HP: %d\n", monster.Name, monster.HitPoints)
				fmt.Printf("Equipped Weapon: %s (Damage: %d, Accuracy: %d%%)\n", player.Weapon.Name, player.Weapon.Damage, player.Weapon.Accuracy)

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
					hit := rand.Intn(100) < player.Weapon.Accuracy
					if hit {
						damage := player.Weapon.Damage
						monster.HitPoints -= damage
						fmt.Printf("You deal %d damage to the %s with your %s.\n", damage, monster.Name, player.Weapon.Name)
					} else {
						fmt.Println("Your attack missed!")
					}
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
					if monster.CanCastSpells && len(monster.Spells) > 0 {
						// Monster casts a spell
						spell := monster.Spells[rand.Intn(len(monster.Spells))]
						if spell.DamageValue > 0 {
							player.HitPoints -= spell.DamageValue
							fmt.Printf("The %s casts %s, dealing %d damage to you.\n", monster.Name, spell.Name, spell.DamageValue)
						} else if spell.EffectValue > 0 {
							// Apply the effect (e.g., increase defense)
							fmt.Printf("The %s casts %s, increasing its defense by %d.\n", monster.Name, spell.Name, spell.EffectValue)
						}
					} else {
						// Monster attacks with its strength
						damage := monster.Strength
						player.HitPoints -= damage
						fmt.Printf("The %s deals %d damage to you.\n", monster.Name, damage)
					}
				}
			}

			// Check the outcome of the battle
			if player.HitPoints > 0 {
				fmt.Println("You defeated the monster!")
				currentRoom.Monster = nil // Remove the monster from the room
			} else {
				fmt.Println("You were defeated by the monster.")
				break
			}
		}

		// Ask the player for their next move
		fmt.Println("What do you want to do next?")
		fmt.Println("1. Go forward")
		fmt.Println("2. Go backward")
		fmt.Println("3. Turn left")
		fmt.Println("4. Turn right")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Go forward
			dungeon.CurrentRoom = (dungeon.CurrentRoom + 1) % len(dungeon.Rooms)
		case 2:
			// Go backward
			dungeon.CurrentRoom = (dungeon.CurrentRoom - 1 + len(dungeon.Rooms)) % len(dungeon.Rooms)
		case 3:
			// Turn left
			// Implement your logic for turning left
		case 4:
			// Turn right
			// Implement your logic for turning right
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// generateMonster generates a random monster
func generateMonster() Monster {
	// Define monster types and descriptions
	monsterTypes := []struct {
		Name        string
		Description string
		BaseHitPoints int
		CanCastSpells bool
		Spells      []Spell
	}{
		{"Goblin", "A small, green-skinned creature with a mischievous grin and a sharp dagger.", 20, false, nil},
		{"Orc", "A towering, muscular humanoid with a fierce expression and a massive club.", 40, false, nil},
		{"Skeleton", "A reanimated skeleton, its bones clattering as it moves, wielding a rusty sword.", 30, false, nil},
		{"Slime", "A gelatinous blob that oozes across the ground, leaving a trail of slime behind.", 25, false, nil},
		{"Dragon", "A majestic, winged beast with scales that glisten like gems and a maw that breathes fire.", 80, true, []Spell{{Name: "Fire Breath", ManaCost: 0, DamageValue: 20, EffectValue: 0}}},
		// Add more monster types as needed
	}

	// Select a random monster type
	monsterType := monsterTypes[rand.Intn(len(monsterTypes))]

	// Generate random monster attributes
	strength := rand.Intn(10) + 1

	return Monster{
		Name:        monsterType.Name,
		Description: monsterType.Description,
		HitPoints:   monsterType.BaseHitPoints,
		Strength:    strength,
		CanCastSpells: monsterType.CanCastSpells,
		Spells:      monsterType.Spells,
	}
}

// generateDungeon generates a random dungeon
func generateDungeon() Dungeon {
	numRooms := rand.Intn(10) + 5 // Generate between 5 and 14 rooms

	rooms := make([]*Room, numRooms)
	for i := range rooms {
		rooms[i] = &Room{
			Description: fmt.Sprintf("You are in room %d.", i+1),
			Monster:     generateMonsterForRoom(),
			Weapon:      generateWeaponForRoom(),
		}
	}

	return Dungeon{
		Rooms:       rooms,
		CurrentRoom: 0,
	}
}

// generateMonsterForRoom generates a monster for a room, or nil if no monster
func generateMonsterForRoom() *Monster {
	if rand.Intn(2) == 0 {
		return nil // No monster in this room
	}
	monster := generateMonster()
	return &monster
}

// generateWeaponForRoom generates a weapon for a room, or nil if no weapon
func generateWeaponForRoom() *Weapon {
	if rand.Intn(3) == 0 {
		return nil // No weapon in this room
	}
	weapons := []Weapon{
		{Name: "Rusty Dagger", Damage: 3, Accuracy: 90},
		{Name: "Iron Mace", Damage: 6, Accuracy: 80},
		{Name: "Elven Bow", Damage: 5, Accuracy: 85},
		{Name: "Enchanted Staff", Damage: 4, Accuracy: 75},
		// Add more weapons as needed
	}
	weapon := weapons[rand.Intn(len(weapons))]
	return &weapon
}

// printDeathAsciiArt prints an ASCII art representation of the player's death
func printDeathAsciiArt() {
	deathArt := `
    _____
   /     \
  | () () |
   \  ^  /
    |||||
    |||||
`
	fmt.Println(deathArt)
}
