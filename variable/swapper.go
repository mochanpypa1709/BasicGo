package variable

import "fmt"

func Swap() {
	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color, color2)
}
