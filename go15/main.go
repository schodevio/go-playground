package main

import "fmt"

type Position struct {
	x float64
	y float64
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

type SpecialPosition struct {
	Position // Embedding Position struct
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.Position.Move(x*2, y*2) // Move twice as far
}

type Player struct {
	*Position // Embedding Position struct
}

func NewPlayer() *Player {
	return &Player{
		Position: &Position{},
	}
}

type Enemy struct {
	*Position // Embedding Position struct
}

func NewEnemy() *Enemy {
	return &Enemy{
		Position: &Position{},
	}
}

func main() {
	player := NewPlayer()
	player.Move(5, 10)

	fmt.Println(player.Position)
	fmt.Println(player.x)
}
