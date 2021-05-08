package err_handling

import (
	"fmt"
	"os"
	"strconv"
)

func Atoi() {
	// Atoi returns an error value
	// So, you should always check it

	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}
