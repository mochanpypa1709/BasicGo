package strings

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func CountChar() {
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)
}
