package switching

import "fmt"

func Fall() {
	i := 142

	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number")
	}

	fmt.Println()
}
