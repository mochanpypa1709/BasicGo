package ifstate

import "fmt"

func Simply() {
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}
