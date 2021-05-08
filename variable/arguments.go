package variable

import (
	"fmt"
	"os"
)

// STEPS:
//
// Compile it by typing:
//   go build -o myprogram
//
// Then run it by typing:
//   ./myprogram
//
// If you're on Windows, then type:
//   myprogram

func Arguments() {
	//count argument
	//count := len(os.Args) - 1
	//fmt.Printf("There are %d names.\n", count)

	//print the path
	//fmt.Println(os.Args[2])

	//fmt.Println("Hello", os.Args[1])
	//fmt.Println("How are you?")

	var (
		l  = len(os.Args) - 1
		n1 = os.Args[1]
		n2 = os.Args[2]
		n3 = os.Args[3]
	)

	fmt.Println("There are", l, "people !")
	fmt.Println("Hello great", n1, "!")
	fmt.Println("Hello great", n2, "!")
	fmt.Println("Hello great", n3, "!")
	fmt.Println("Nice to meet you all.")

}
