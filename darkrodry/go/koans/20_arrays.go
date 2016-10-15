package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	
	var slice []int = primes[1:4]
	fmt.Println(slice)

	primes[2] = 17
	fmt.Println(slice)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	makeSlice := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d %v\n", len(makeSlice), cap(makeSlice), makeSlice)

	for i := range primes {
		fmt.Print(i)
	}
}
