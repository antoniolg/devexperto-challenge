package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
	fmt.Println(m["Google"])
	
	m["Github"] = Vertex{37.7824791,-122.3928715}
	fmt.Println(m)
	
	delete(m, "Google")
	fmt.Println(m)
}
