package variable

import (
	"fmt"
	"path"
)

func RemFile() {
	dir, _ := path.Split("secret/file.txt")

	fmt.Println(dir)
}
