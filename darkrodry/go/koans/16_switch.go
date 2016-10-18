package main

import (
	"fmt"
	"runtime"
	"time"
)

func goOS() {
        fmt.Print("Go runs on ")
        switch os := runtime.GOOS; os {
        case "darwin":
                fmt.Println("OS X.")
        case "linux":
                fmt.Println("Linux.")
        default:
                // freebsd, openbsd,
                // plan9, windows...
                fmt.Printf("%s.", os)
        }
}

func saturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func greeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	goOS()
	saturday()
	greeting()
}
