package loop

import (
	"fmt"
)

func Break() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}

		sum += i

		fmt.Println(i, "-->", sum)

		i++
	}

	fmt.Println(sum)
}
