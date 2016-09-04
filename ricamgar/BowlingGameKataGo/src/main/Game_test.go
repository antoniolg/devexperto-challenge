package main

import (
	"testing"
)

func TestGutterGame(t *testing.T) {
	game := NewGame()
	game.rollMany(20, 0)
	assertScore(t, 0, game.Score())
}

func TestAllOnes(t *testing.T) {
	game := NewGame()
	game.rollMany(20, 1)
	assertScore(t, 20, game.Score())
}

func TestOneSpare(t *testing.T) {
	game := NewGame()
	game.rollSpare()
	game.Roll(3)
	game.rollMany(17, 0)
	assertScore(t, 16, game.Score())

}

func TestStrike(t *testing.T) {
	game := NewGame()
	game.Roll(10)
	game.Roll(3)
	game.Roll(4)
	game.rollMany(16, 0)
	assertScore(t, 24, game.Score())
}

func (game *Game) rollMany(times int, pins int) {
	for i := 0; i < times; i++ {
		game.Roll(pins)
	}
}

func (game *Game) rollSpare()  {
	game.Roll(5)
	game.Roll(5)
}

func assertScore(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected %d but was %d", expected, actual)
	}
}


