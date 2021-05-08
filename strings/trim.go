package strings

import (
	"fmt"
	"strings"
)

func TrimIt() {
	msg := `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))
}
