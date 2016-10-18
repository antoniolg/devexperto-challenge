(defproject kata "0.1.0-SNAPSHOT"
  :description "Bowling kata in clojure"
  :license {:name "MIT"
            :url "https://opensource.org/licenses/MIT"}
  :dependencies [[org.clojure/clojure "1.8.0"]]
  :main ^:skip-aot kata.game
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})
