package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Player struct {
	mu     sync.RWMutex
	health int32
}

func NewPlayer() *Player {
	return &Player{
		health: 100,
	}
}

func (p *Player) GetHealth() int {
	return int(atomic.LoadInt32(&p.health))
}

func (p *Player) TakeDamage(damage int) {
	atomic.AddInt32(&p.health, -int32(damage))
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)

	for {
		fmt.Printf("Player health: %d\r", p.GetHealth())
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Millisecond * 300)

	for {
		p.TakeDamage(rand.Intn(40))
		if p.GetHealth() <= 0 {
			fmt.Println("\nPlayer is dead!")
			break
		}
		<-ticker.C
	}
}

func main() {
	player := NewPlayer()

	go startUILoop(player)
	startGameLoop(player)
}

// ----- Mutexes -----

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"sync"
// 	"time"
// )

// type Player struct {
// 	mu     sync.RWMutex
// 	health int
// }

// func NewPlayer() *Player {
// 	return &Player{
// 		health: 100,
// 	}
// }

// func (p *Player) GetHealth() int {
// 	p.mu.RLock()
// 	defer p.mu.RUnlock()
// 	return p.health
// }

// func (p *Player) TakeDamage(damage int) {
// 	p.mu.Lock()
// 	defer p.mu.Unlock()
// 	p.health -= damage
// }

// func startUILoop(p *Player) {
// 	ticker := time.NewTicker(time.Second)

// 	for {
// 		fmt.Printf("Player health: %d\r", p.GetHealth())
// 		<-ticker.C
// 	}
// }

// func startGameLoop(p *Player) {
// 	ticker := time.NewTicker(time.Millisecond * 300)

// 	for {
// 		p.TakeDamage(rand.Intn(40))
// 		if p.GetHealth() <= 0 {
// 			fmt.Println("\nPlayer is dead!")
// 			break
// 		}
// 		<-ticker.C
// 	}
// }

// func main() {
// 	player := NewPlayer()

// 	go startUILoop(player)
// 	startGameLoop(player)
// }
