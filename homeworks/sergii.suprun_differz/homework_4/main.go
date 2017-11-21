package main

import "./linkedlist"
import "fmt"

func main() {

	list := linkedlist.New()

	list.Add("element 0")
	list.Add("element 1")
	list.Add("element 2")
	list.Add("element 3")

	for i := 0; i <= list.Size(); i++ {
		element, error := list.Get(i)
		if error != nil {
			fmt.Println(error)
			break
		}
		fmt.Println(element)
	}

	fmt.Println(list.String())

	list.Remove("element 1")
	fmt.Println(list.String())

	list.Remove(1)
	fmt.Println(list.String())

	list.Remove(0)
	fmt.Println(list.String())
}
