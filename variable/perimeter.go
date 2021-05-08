package variable

import "fmt"

func Perim() {
	var (
		perimeter     int
		width, height = 5, 6
	)

	// first calculates: (width + height)
	// then            :  multiplies it with 2

	// just like in math

	perimeter = 2 * (width + height)

	fmt.Println(perimeter)
}
