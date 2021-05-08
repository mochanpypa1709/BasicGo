package iota

import "fmt"

func Month2() {
	const (
		_   = iota // 0
		Jan        // 1
		Feb        // 2
		Mar        // 3
	)

	fmt.Println(Jan, Feb, Mar)
}
