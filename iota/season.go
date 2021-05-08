package iota

import "fmt"

func Seasons() {
	const (
		Spring = (iota + 1) * 3 //0
		Summer                  //6
		Fall                    //9
		Winter                  //12
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
