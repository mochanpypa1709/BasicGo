package variable

import "fmt"

func MultiAssign() {
	var (
		lang    string
		version int
	)

	lang, version = "go", 2

	fmt.Println(lang, "version", version)
}
