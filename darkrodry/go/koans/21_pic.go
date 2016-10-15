package main

import "golang.org/x/tour/pic"

func Allocate(dx, dy int) [][]uint8 {
	// Allocate picture array
	array := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		array[i] = make([]uint8, dx)
	}
	return array
}

func Pic(dx, dy int) [][]uint8 {
	// Allocate picture array
	pic := Allocate(dx, dy)
	
	for x:= range pic {
		for y:= range pic {
			pic[x][y] = uint8(x*y)
		}
	}
	
	return pic
}

func main() {
        pic.Show(Pic)
}
