package switching

import (
	"fmt"
	"os"
)

func City() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
		// break // unnecessary in Go

		// vip := true
		// fmt.Println("VIP trip?", vip)
	case "Tokyo":
		fmt.Println("Japan")
		// fmt.Println("VIP trip?", vip)
	default:
		fmt.Println("Where?")
	}

	// if city == "Paris" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}
