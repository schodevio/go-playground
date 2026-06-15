package main

import "fmt"

// Pointer is 8 bytes in size on a 64-bit architecture.
// Pointer is 4 bytes in size on a 32-bit architecture.

type Player struct {
	health int
}

func (p *Player) takeExplosionDamage(damage int) {
	fmt.Println("BOOOOM!")
	p.health -= damage
}

func takeExplosionDamage(player *Player, damage int) {
	fmt.Println("BOOOOM!")
	player.health -= damage
}

func main() {
	player := &Player{
		health: 100,
	}

	fmt.Println("Player health before taking damage:", player.health)

	// player = nil                    // somewhere else in the code, the player object is set to nil, which means it no longer points to a valid memory address
	// takeExplosionDamage(player, 10) // pass the address of the player object to the function
	player.takeExplosionDamage(10)

	fmt.Println("Player health after taking damage:", player.health)
}

// ==========

// type Player struct {
// 	health int
// }

// func takeExplosionDamage(player Player) {
// 	fmt.Println("BOOOOM!")

// 	explosionDamage := 10
// 	player.health -= explosionDamage
// }

// func main() {
// 	player := Player{
// 		health: 100,
// 	}

// 	fmt.Println("Player health before taking damage:", player.health)
// 	takeExplosionDamage(player) // player object is copied and passed to the function, so the original player object remains unchanged
// 	fmt.Println("Player health after taking damage:", player.health)
// }
