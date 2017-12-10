package main

import (
	"fmt"
	"go-course-2017/homeworks/alexandr.kravchenko_drhihi/homework_5/models"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

type Box struct {
	a, b, height int
	sell         []models.Shape
}

const maxSizeShape = 10

func main() {
	box := initBox()
	fillBox(box)

	box.PrintlnDetails()
	box.Println()

	box.sortCorner(1)
	fmt.Println()
	box.Println()

}

/*
param:
	1 2
	3 4
*/
func (b *Box) sortCorner(corner int) {

	slice := b.getSliceSell()
	sort.Sort(slice)

	_ = corner

	switch corner {
	default:
		panic("the corner is not right")
	case 1:
		// 1
		for k, cy, x, y, fx, fy := 0, 0, 0, 0, 0, 0; k < len(slice); k++ {
			b.sell[y][x] = slice[k]
			x++
			y--
			if y < fy || x > 3 {
				cy++
				y = cy
				x = fx
			}
			if cy > 3 {
				cy--
				fx++
				x = fx
				y = cy
			}
		}
	}

}

func (b *Box) Println() {
	for _, vl := range b.sell {
		fmt.Println(vl)
	}
}

func (b *Box) PrintlnDetails() {
	for _, vl := range b.sell {
		fmt.Println("{", vl)
		for _, vr := range vl {
			fmt.Printf("%T %v\n", vr, vr)
		}
		fmt.Println("}")
	}
}

func initBox() *Box {

	box := new(Box)

	if len(os.Args) != 4 {
		panic("fake parameters")
	}

	for k, v := range os.Args[1:] {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		switch k {
		case 0:
			box.a = i
		case 1:
			box.b = i
		case 2:
			box.height = i
		}
	}

	box.sell = make([]models.Shape, box.height)

	for i := 0; i < box.height; i++ {
		box.sell[i] = make(models.Shape, box.height)
	}

	return box
}

func fillBox(box *Box) {
	types := [...]string{
		"cube",
		"cone",
		"piramid",
	}

	for _, vl := range box.sell {
		for k, _ := range vl {
			switch types[rand.Intn(len(types))] {
			default:
				panic("unknown type")
			case "cube":
				vl[k] = models.NewCube(
					rand.Intn(maxSizeShape), // a
				)
			case "cone":
				vl[k] = models.NewCone(
					rand.Intn(maxSizeShape), // r
					rand.Intn(maxSizeShape), // h
				)
			case "piramid":
				vl[k] = models.NewPiramid(
					rand.Intn(maxSizeShape), // a
					rand.Intn(maxSizeShape), // b
					rand.Intn(maxSizeShape), // h
				)
			}
		}
	}

}

func (b *Box) getSliceSell() models.Shape {
	result := make(models.Shape, b.height*b.height)
	var i int
	for _, vl := range b.sell {
		for _, vr := range vl {
			result[i] = vr
			i++
		}
	}
	return result
}
