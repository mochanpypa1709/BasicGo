package print

import "fmt"

func Print() {
	//fmt.Printf("I'm %d years old.\n", 35)
	fmt.Printf("My name is %s and my lastname is %s.\n", "Mohamad", "Chandra")

	// print string type %s
	msg := "My name is %s and my lastname is %s.\n"
	fmt.Printf(msg, "Mohamad", "Chandra")

	//print float type %f
	fmt.Printf("Temperature is %.1f degrees.\n", 29.5)

	//print double quote type %q
	fmt.Printf("%q\n", "hello world")

	//print the type %T
	fmt.Printf("Type of %d is %[1]T\n", 3)
	fmt.Printf("Type of %.2f is %[1]T\n", 3.14)
	fmt.Printf("Type of %s is %[1]T\n", "hello")
	fmt.Printf("Type of %t is %[1]T\n", true)
}
