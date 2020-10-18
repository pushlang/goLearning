package main

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("..\\bundling")
	data := box.String("test.txt")
	fmt.Println("Contents of file:", data)
}

//C:\Users\123\go\src\basics\bundling>packr install -v ..\bundling
