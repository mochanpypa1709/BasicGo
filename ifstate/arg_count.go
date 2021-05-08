package ifstate

import (
	"fmt"
	"os"
)

func ArgCount() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	fmt.Printf("total args: %d\n", l)

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf(
			`There are two: "%s %s"`+"\n",
			args[1], args[2],
		)
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}
}
