package main

import "testing"

func TestMain(t *testing.T) {
	player := NewPlayer()

	go startUILoop(player)
	startGameLoop(player)
}
