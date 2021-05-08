package variable

import "fmt"

func Swap2() {
	red, blue := "red", "blue"

	red, blue = blue, red

	fmt.Println(red, blue)
}
