package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   	bool		= false
	String 	string 		= "awesome"
	Int 	int 		= 1
	MaxInt 	uint64     	= 1<<64 - 1
	Byte 	byte 		= 8
	Float 	float32		= 32.2
	Complex complex128 	= cmplx.Sqrt(-5 + 12i)
	Rune 	rune 		= 'a'
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, String, String)
	fmt.Printf(f, Int, Int)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, Byte, Byte)
	fmt.Printf(f, Float, Float)
	fmt.Printf(f, Complex, Complex)
	fmt.Printf(f, Rune, Rune)
}