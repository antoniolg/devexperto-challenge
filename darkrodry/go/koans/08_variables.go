package main

import "fmt"

func createVars() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}

func initVars() {
	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func shortVars() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func main() {
	createVars()
	initVars()
	shortVars()
}