package main

import "fmt"

// Rectangle
type Rectangle struct {
	width  int
	height int
}

func main() {
	r1 := &Rectangle{300, 323}

	fmt.Printf("Width: %d\nHeight: %d", r1.width, r1.height)
	fmt.Println(r1)
}
