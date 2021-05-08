package strings

import (
	"fmt"
	"os"
	"strings"
)

func LowCase() {
	fmt.Println(strings.ToLower(os.Args[1]))
}
