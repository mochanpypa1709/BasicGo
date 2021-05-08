package strings

import (
	"fmt"
	"os"
)

func RawConcat() {
	name := os.Args[1]

	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)
}
