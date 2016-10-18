(ns kata.game-test
  (:use clojure.test
    kata.game))

(deftest test-gutter-game
    (is (= 0 (get-score (repeat 20 0))))
)

(deftest test-all-ones
    (is (= 20 (get-score (repeat 20 1))))
)

(deftest test-one-spare
	(is (= 16 (get-score (concat [5 5 3] (repeat 17 0)))))
)

(deftest test-one-strike
	(is (= 24 (get-score (concat [10 3 4] (repeat 16 0)))))
)

(deftest test-perfect-game
	(is (= 300 (get-score (repeat 12 10))))
)
