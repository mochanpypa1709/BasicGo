package strings

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func Banger() {
	msg := os.Args[1]
	l := utf8.RuneCountInString(msg)

	s := msg + strings.Repeat("!", l)

	fmt.Println(s)
}
