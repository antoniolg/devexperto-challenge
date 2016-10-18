(ns kata.game)

(defn spare? [roll]
	(= 10 (reduce + (take 2 roll)))
)

(defn strike? [roll]
	(= 10 (first roll))
)

(defn get-score [pins]
  (loop [cnt pins acc 0 round 10]
  	(if (> round 0)
  		(if (strike? cnt)
  			(recur (drop 1 cnt) (+ acc (reduce + (take 3 cnt))) (- round 1))
		  	(if (spare? cnt)
		  		(recur (drop 2 cnt) (+ acc (reduce + (take 3 cnt))) (- round 1))
		  		(recur (drop 2 cnt) (+ acc (reduce + (take 2 cnt))) (- round 1))
		  	)
	  	)
	  	acc
  	)
  )
)