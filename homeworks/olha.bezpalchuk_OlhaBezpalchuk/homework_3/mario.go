package main

import (
	"fmt"
)

func build_line(material string, brick_length int, line_length int) (line string) {
	line = ""

	piece := ""
	spaces := ""

	// use bricks
	for i := 0; i < brick_length; i++ {
		piece += material
	}

	// add spaces to the line
	for i := brick_length; i < line_length; i++ {
		spaces += " "
	}

	line = spaces + piece + "  " + piece + spaces 
	return
}

func build_pyramid(material string, height int) () {
	fmt.Printf("Height: %d\n", height)
	for i := 1; i <= height; i++ {
		fmt.Println(build_line(material, i, height))
	}
}

func main() {

	brick := "#"
	height := 5

	build_pyramid(brick, height)
}
