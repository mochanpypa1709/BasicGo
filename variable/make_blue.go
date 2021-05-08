package variable

import "fmt"

func Blue() {
	color := "green"

	color = "blue"

	color = "dark " + color

	fmt.Println(color)
}
