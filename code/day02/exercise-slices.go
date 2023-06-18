package main

import "golang.org/x/tour/pic"

/*
Implement Pic

It should return a slice of length dy, each element of which is a slice of dx
8-bit unsigned integers. When you run the program, it will display your
picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y,
and x^y.

You need to use a loop to allocate each []uint8 inside the [][]uint8.
Use uint8(intValue) to convert between types.
*/

func createTwoDimensionalArray(dx, dy int) [][]uint8 {
	array := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		array[i] = make([]uint8, dx)
	}

	return array
}

func Pic(dx, dy int) [][]uint8 {
	var arr = createTwoDimensionalArray(dx, dy)

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			arr[i][j] = uint8((j * 10000) / (i + 2))
		}
	}

	return arr
}

func main() {
	pic.Show(Pic)
}
