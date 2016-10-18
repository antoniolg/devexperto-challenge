package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f_1 := 1
	f_2 := -1
	return func() int {
		f := f_1 + f_2
		f_2 = f_1
		f_1 = f
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

