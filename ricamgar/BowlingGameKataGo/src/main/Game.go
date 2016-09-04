package main

type Game struct {
	rolls       []int
	currentRoll int
}

func NewGame() *Game {
	game := new(Game)
	game.rolls = make([]int, 20)
	return game
}

func (game *Game) Roll(pins int) {
	game.rolls[game.currentRoll] = pins
	game.currentRoll++
}

func (game *Game) Score() (sum int) {
	for frameIndex, frame := 0, 0; frame < 10; frame++ {
		if game.isStrike(frameIndex) {
			sum += 10 + game.strikeBonus(frameIndex + 1)
			frameIndex ++
		} else if game.isSpare(frameIndex) {
			sum += 10 + game.spareBonus(frameIndex)
			frameIndex += 2
		} else {
			sum += game.sumOfBallsInFrame(frameIndex)
			frameIndex += 2
		}
	}
	return sum
}

func (game *Game) isSpare(frameIndex int) bool {
	return game.sumOfBallsInFrame(frameIndex) == 10
}

func (game *Game) isStrike(frameIndex int) bool {
	return game.rolls[frameIndex] == 10
}

func (game *Game) strikeBonus(frameIndex int) int {
	return game.sumOfBallsInFrame(frameIndex)
}

func (game *Game) spareBonus(frameIndex int) int {
	return game.rolls[frameIndex + 2]
}

func (game *Game) sumOfBallsInFrame(frameIndex int) int {
	return game.rolls[frameIndex] + game.rolls[frameIndex + 1]
}
