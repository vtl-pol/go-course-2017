// You can edit this code!
package main

import "fmt"

func main() {

	mario := 4 // default value
	X := "#"   // default char to print
	//	var mariomin int = 1		// ! it will be inferred from the right-hand side
	//	var mariomax int = 24		// ! it will be inferred from the right-hand side
	mariomin := 1
	mariomax := 24

	fmt.Print("Enter Mario number: ")
	fmt.Scanln(&mario) // input value

	if mario < mariomin {
		fmt.Println("Mario is too small, set to ", mariomin)
		mario = mariomin
	}
	if mario > mariomax {
		fmt.Println("Mario is too big, set to ", mariomax)
		mario = mariomax
	}

	fmt.Println("")
	// print n-space for "Height" align center
	for n := 1; n <= mario-3; n++ {
		fmt.Print(" ")
	}
	fmt.Println("Height: ", mario)

	// print "input" lines
	for i := 0; i <= mario; i++ {

		fmt.Print(" ")
		// print n-space
		for n := 1; n <= mario-i; n++ {
			fmt.Print(" ")
		}
		// print n-symbol
		for n := 1; n <= i; n++ {
			fmt.Print(X)
		}

		// print 2 space
		fmt.Print("  ")

		// print n-symbol
		for n := 1; n <= i; n++ {
			fmt.Print(X)
		}
		// print n-space
		for n := 1; n <= mario-i; n++ {
			fmt.Print(" ")
		}
		// end line & LF
		fmt.Println(" ")
	}

	//	fmt.Println("Height:  4")
	//	fmt.Println("   *__*   ")
	//	fmt.Println("  **__**  ")
	//	fmt.Println(" ***__*** ")
	//	fmt.Println("****__****")

	fmt.Println("")
	fmt.Println("Hello veets.seer-rndlinker37")

}
