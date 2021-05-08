package strings

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func RightTrim() {
	name := "inanç           "

	name = strings.TrimRight(name, " ")
	l := utf8.RuneCountInString(name)

	fmt.Println(l)
}
