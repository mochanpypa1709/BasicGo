package variable

import "fmt"

func TypConv() {
	//a, b := 10, 5.5
	//fmt.Println(float64(a) + b)

	//a, b := 10, 5.5
	//a = int(b)
	//fmt.Println(float64(a) + b)

	//age := 2
	//fmt.Println(7.5 + float64(age))

	min := int8(127)
	max := int16(1000)

	fmt.Println(max + int16(min))

}
