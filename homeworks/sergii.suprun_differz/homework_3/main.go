package main

import "fmt"

func main() {
	height := 4

	if height < 2 {
		fmt.Println("Wrong height")
		return
	}

	for i := 1; i <= height; i++ {
		first := true
		for j := i; j < height; j++ {
			fmt.Print(" ")
		}
		for {
			for j := 0; j < i; j++ {
				fmt.Print("#")
			}
			if first {
				fmt.Print("  ")
				first = false
				continue
			}
			break
		}
		fmt.Println()
	}
}
