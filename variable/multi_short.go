package variable

import (
	"fmt"
)

func MultiShort() {
	a, _ := multi()

	fmt.Println(a)
}

func multi() (int, int) {
	return 5, 4
}
