require 'test/unit'
require_relative 'game' 

class GameTest < Test::Unit::TestCase

  def setup
    @game = Game.new
  end

  def roll_many(n, pins)
    n.times do
      @game.roll(pins)
    end
  end

  def test_gutter_game
    setup
    roll_many(20, 0)
    assert_equal 0, @game.score
  end

  def test_all_ones
    setup
    roll_many(20, 1)
    assert_equal 20, @game.score
  end

  def test_one_spare
    setup
    @game.roll(5)
    @game.roll(5)
    @game.roll(3)
    roll_many(17, 0)
    assert_equal 16, @game.score
  end
  
  def test_one_strike
    setup
    @game.roll(10)
    @game.roll(3)
    @game.roll(4)
    roll_many(16, 0)
    assert_equal 24, @game.score
  end

  def test_perfect_game
    setup
    roll_many(12, 10)
    assert_equal 300, @game.score
  end

end
