package kata

import "testing"

var game Game

func setUp() {
	game = Game{}
}

func rollMany(n, pins int) {
	for i := 0; i < n; i++ {
		game.roll(pins)
	}
}

func rollSpare() {
	game.roll(5)
	game.roll(5)
}

func assertScore(expected, score int, t *testing.T) {
	if score != expected {
		t.Errorf("Game.score() expected %d, got %d", expected, score)
	}
}

func TestGutterGame(t *testing.T) {
	setUp()
	rollMany(20, 0)
	assertScore(0, game.score(), t)
}

func TestAllOnes(t *testing.T) {
	setUp()
	rollMany(20, 1)
	assertScore(20, game.score(), t)
}

func TestOneSpare(t *testing.T) {
	setUp()
	rollSpare()
	game.roll(3)
	rollMany(17, 0)
	assertScore(16, game.score(), t)
}

func TestOneStrike(t *testing.T) {
	setUp()
	game.roll(10)
	game.roll(4)
	game.roll(3)
	rollMany(16, 0)
	assertScore(24, game.score(), t)
}

func TestPerfectGame(t *testing.T) {
	setUp()
	rollMany(12, 10)
	assertScore(300, game.score(), t)
}
