package strings

import "fmt"

func PrintJSON() {
	// this one uses a raw string literal

	// can you see how readable it is?
	// compared to the previous one?

	json := `
{
	"Items": [{
		"Item": {
			"name": "Teddy Bear"
		}
	}]
}`

	fmt.Println(json)
}
