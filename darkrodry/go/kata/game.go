package kata

type Game struct {
	currentRoll int
	rolls [21]int
}

func(g* Game) roll(pins int) {
	g.rolls[g.currentRoll] = pins
	g.currentRoll += 1
}

func(g* Game) score() int {
	score := 0
	frameIndex := 0
	for frame := 0; frame < 10; frame++ {
		if isStrike(g, frameIndex) {
			score += 10 + countStrikeBonus(g, frameIndex)
			frameIndex += 1
		} else if isSpare(g, frameIndex) {
			score += 10 + countSpareBonus(g, frameIndex)
			frameIndex += 2
		} else {
			score += countPointsInFrame(g, frameIndex)
			frameIndex += 2
		}
	}
	return score
}

func isStrike(g* Game, frameIndex int)  bool {
	return g.rolls[frameIndex] == 10
}

func isSpare(g* Game, frameIndex int) bool {
	return g.rolls[frameIndex] + g.rolls[frameIndex + 1] == 10
}

func countPointsInFrame(g* Game, frameIndex int) int {
	return g.rolls[frameIndex] + g.rolls[frameIndex + 1]
}

func countStrikeBonus(g* Game, frameIndex int) int {
	return g.rolls[frameIndex + 1] + g.rolls[frameIndex + 2]
}

func countSpareBonus(g* Game, frameIndex int) int {
	return g.rolls[frameIndex + 2]
}
