class Game

  def initialize
    @current_roll = 0
    @rolls = Array.new(21) {0}
  end

  def roll(pins)
    @rolls[@current_roll] = pins
    @current_roll += 1
  end

  def score
    score = 0
    frame = 0
    10.times do
      if is_strike?(frame)
        score += 10 + count_strike_bonus(frame)
        frame += 1
      elsif is_spare?(frame)
        score += 10 + count_spare_bonus(frame)
        frame += 2
      else
        score += count_frame_points(frame)
        frame += 2
      end
    end
    return score
  end

  def is_strike?(frame)
    return @rolls[frame] == 10
  end

  def is_spare?(frame)
    return @rolls[frame] + @rolls[frame+1] == 10
  end

  def count_frame_points(frame)
    return @rolls[frame] + @rolls[frame+1]
  end

  def count_spare_bonus(frame)
    return @rolls[frame+2]
  end

  def count_strike_bonus(frame)
    return @rolls[frame+1] + @rolls[frame+2]
  end
end
